var int: x0;
var int: y0;
var int: z0;

var int: dx0;
var int: dy0;
var int: dz0;


var int: t0;
constraint x0 + dx0 * t0 = 225415405941969 + 23 * t0;
constraint y0 + dy0 * t0 = 400648127977931 + -204 * t0;
constraint z0 + dz0 * t0 = 79201130433258 + 617 * t0;

var int: t1;
constraint x0 + dx0 * t1 = 353783687623292 + -80 * t1;
constraint y0 + dy0 * t1 = 138575899489956 + 156 * t1;
constraint z0 + dz0 * t1 = 318416438572569 + 21 * t1;

var int: t2;
constraint x0 + dx0 * t2 = 215751176267772 + -120 * t2;
constraint y0 + dy0 * t2 = 376619563956940 + 126 * t2;
constraint z0 + dz0 * t2 = 230133299986253 + -352 * t2;

var int: t3;
constraint x0 + dx0 * t3 = 323665001411559 + 21 * t3;
constraint y0 + dy0 * t3 = 234169007288514 + -57 * t3;
constraint z0 + dz0 * t3 = 312785643963696 + 87 * t3;

var int: t4;
constraint x0 + dx0 * t4 = 254448034578908 + -86 * t4;
constraint y0 + dy0 * t4 = 463170080025786 + -444 * t4;
constraint z0 + dz0 * t4 = 244090582025017 + -22 * t4;


solve satisfy;

output [show(x0), show(y0), show(z0)]
