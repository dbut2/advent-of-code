package harness

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	strings2 "strings"
	"time"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/strings"
)

// Harness manages running each days' challenge code, including parsing the
// input before calling solve, and managing test expectations from examples.
type Harness[T any, U comparable] struct {
	preProcessor PreProcessor[T]
	solve        func(T) U
	silent       bool
	dontSubmit   bool
	metadata     metadata
}

// HarnessOpt modifies the Harness when initialising.
type HarnessOpt[T any, U comparable] func(harness *Harness[T, U])

// WithPreProcessor explicitly sets the PreProcessor to be used.
func WithPreProcessor[T any, U comparable](preProcessor PreProcessor[T]) HarnessOpt[T, U] {
	return func(h *Harness[T, U]) {
		h.preProcessor = preProcessor
	}
}

func WithSilence[T any, U comparable]() HarnessOpt[T, U] {
	return func(h *Harness[T, U]) {
		h.silent = true
	}
}

func WithNoSubmit[T any, U comparable]() HarnessOpt[T, U] {
	return func(h *Harness[T, U]) {
		h.dontSubmit = true
	}
}

// New returns a new Harness. At minimum a solve function and some inputs are
// required.
func New[T any, U comparable](solve func(T) U, opts ...HarnessOpt[T, U]) *Harness[T, U] {
	h := Harness[T, U]{
		preProcessor: defaultPreProcessor[T](),
		solve:        solve,
		metadata:     getMetadata(),
	}

	for _, opt := range opts {
		opt(&h)
	}

	return &h
}

type metadata struct {
	workdir         string
	year, day, part int
}

func getMetadata() metadata {
	_, file, _, _ := runtime.Caller(2)

	dir := filepath.Dir(file)

	r := regexp.MustCompile(`(\d{4})/(\d{2})/(\d{1})\.go`)
	parts := r.FindStringSubmatch(file)

	return metadata{
		workdir: dir,
		year:    sti.Int(parts[1]),
		day:     sti.Int(parts[2]),
		part:    sti.Int(parts[3]),
	}
}

// PreProcessor is a function that process the input data before passing to the
// solve function.
type PreProcessor[T any] func(string) T

// SplitSequence trims and splits the input on the seq string sequence
func SplitSequence(seq string) PreProcessor[[]string] {
	return func(s string) []string {
		return strings2.Split(strings2.TrimSpace(s), seq)
	}
}

// SplitNewlines is the default processing that splits the input on newlines
func SplitNewlines() PreProcessor[[]string] {
	return SplitSequence("\n")
}

// Ints process the input to a slice of ints for each line
func Ints() PreProcessor[[][]int] {
	return func(s string) [][]int {
		return lists.Map(SplitNewlines()(s), func(l string) []int {
			return strings.Ints(l)
		})
	}
}

// Grid processes the input as a grid of bytes
func Grid() PreProcessor[space.Grid[byte]] {
	return func(s string) space.Grid[byte] {
		return space.NewGridFromInput(SplitNewlines()(s))
	}
}

func DoubleSection() PreProcessor[[2]string] {
	return func(s string) [2]string {
		return [2]string(SplitSequence("\n\n")(s))
	}
}

func DoubleSectionLines() PreProcessor[[2][]string] {
	return func(s string) [2][]string {
		parts := DoubleSection()(s)
		return [2][]string{SplitNewlines()(parts[0]), SplitNewlines()(parts[1])}
	}
}

func defaultPreProcessor[T any]() PreProcessor[T] {
	switch any(*new(T)).(type) {
	case string:
		return any(PreProcessor[string](strings2.TrimSpace)).(PreProcessor[T])
	case []string:
		return any(SplitNewlines()).(PreProcessor[T])
	case [][]string:
		return any(PreProcessor[[][]string](func(input string) [][]string {
			var groups [][]string
			for _, group := range strings2.Split(strings2.TrimSpace(input), "\n\n") {
				groups = append(groups, strings2.Split(group, "\n"))
			}
			return groups
		})).(PreProcessor[T])
	case [2]string:
		return any(DoubleSection()).(PreProcessor[T])
	case [2][]string:
		return any(DoubleSectionLines()).(PreProcessor[T])
	case []int:
		return any(PreProcessor[[]int](strings.Ints)).(PreProcessor[T])
	case [][]int:
		return any(Ints()).(PreProcessor[T])
	case space.Grid[byte]:
		return any(Grid()).(PreProcessor[T])
	default:
		panic("no supported preprocessor for type")
	}
}

func (h *Harness[T, U]) run(s string) U {
	return h.solve(h.preProcessor(s))
}

// Run will execute the harness with the main input.
func (h *Harness[T, U]) Run() {
	out := h.run(h.getInput())
	if !h.silent {
		fmt.Println(out)
	}
	if !h.dontSubmit && !h.hasCompletedToday() {
		fmt.Println(h.submitAnswer(out))
	}
}

func (h *Harness[T, U]) readInput(s string) string {
	data, _ := os.ReadFile(filepath.Join(h.metadata.workdir, fmt.Sprintf("%s.txt", s)))
	return string(data)
}

func (h *Harness[T, U]) writeInput(s string, input string) {
	err := os.WriteFile(filepath.Join(h.metadata.workdir, fmt.Sprintf("%s.txt", s)), []byte(input), 0644)
	if err != nil {
		panic(err.Error())
	}
}

func (h *Harness[T, U]) getInput() string {
	input := h.readInput("input")
	if input == "" {
		input = h.fetchInput()
		h.writeInput("input", input)
	}
	return input
}

func (h *Harness[T, U]) getAnswers() []U {
	data, _ := os.ReadFile(filepath.Join(h.metadata.workdir, "answers.txt"))

	var answers []U
	switch any(*new(U)).(type) {
	case string:
		for _, line := range bytes.Split(data, []byte("\n")) {
			answers = append(answers, any(string(line)).(U))
		}
	case int:
		for _, line := range bytes.Split(data, []byte("\n")) {
			answers = append(answers, any(sti.Int(string(line))).(U))
		}
	}
	return answers
}

func (h *Harness[T, U]) addAnswer(v U) {
	answers := h.getAnswers()
	answers = append(answers, v)

	var data []byte
	switch any(*new(U)).(type) {
	case string:
		for _, a := range answers {
			data = append(data, []byte(fmt.Sprintf("%v\n", a))...)
		}
	case int:
		for _, a := range answers {
			data = append(data, []byte(fmt.Sprintf("%d\n", a))...)
		}
	}

	err := os.WriteFile(filepath.Join(h.metadata.workdir, "answers.txt"), data, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func (h *Harness[T, U]) hasAnswered(v U) bool {
	return slices.Contains(h.getAnswers(), v)
}

func (h *Harness[T, U]) fetchInput() string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", h.metadata.year, h.metadata.day)
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		panic("no AOC_SESSION env")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	return string(bytes)
}

func (h *Harness[T, U]) getTestInput(n int) string {
	return h.readInput(fmt.Sprintf("test%d", n))
}

func (h *Harness[T, U]) submitAnswer(answer U) string {
	if h.hasAnswered(answer) {
		panic("already tried")
	}
	h.addAnswer(answer)
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", h.metadata.year, h.metadata.day)
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		panic("no AOC_SESSION env")
	}

	form := url2.Values{
		"level":  {fmt.Sprintf("%d", h.metadata.part)},
		"answer": {fmt.Sprintf("%v", answer)},
	}

	req, err := http.NewRequest("POST", url, strings2.NewReader(form.Encode()))
	if err != nil {
		panic(err.Error())
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	r := regexp.MustCompile(`<article><p>(.+)</p></article>`)
	return r.FindStringSubmatch(string(bytes))[1]
}

func (h *Harness[T, U]) hasCompletedToday() bool {
	url := fmt.Sprintf("https://adventofcode.com/%d/leaderboard/private/view/1573050.json", h.metadata.year)
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		panic("no AOC_SESSION env")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var v any
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		panic(err.Error())
	}

	completions := v.(map[string]any)["members"].(map[string]any)["1573050"].(map[string]any)["completion_day_level"]

	parts, ok := completions.(map[string]any)[fmt.Sprintf("%d", h.metadata.day)]
	if !ok {
		return false
	}

	_, ok = parts.(map[string]any)[fmt.Sprintf("%d", h.metadata.part)]
	if !ok {
		return false
	}

	return ok
}

// Expect will set a required assertion on the output of solve for a test input.
func (h *Harness[T, U]) Expect(n int, expected U) {
	input := h.getTestInput(n)
	out := h.run(input)
	if out != expected {
		panic(fmt.Sprintf("test failed!\n expected: %v\ngot: %v\n", expected, out))
	}
}

// Benchmark is a utility to run a benchmark on solve using the main input.
func (h *Harness[T, U]) Benchmark(cond any) {
	var c benchmark.Condition
	switch v := cond.(type) {
	case benchmark.Condition:
		c = v
	case int:
		c = benchmark.Count(v)
	case time.Duration:
		c = benchmark.Time(v)
	default:
		panic("unknown condition")
	}

	input := h.preProcessor(h.getInput())
	benchmark.Run(func() {
		h.solve(input)
	}, c)
}
