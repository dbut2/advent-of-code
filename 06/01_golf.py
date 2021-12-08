f=[0]*9
for g in open("i").read().split(","):f[int(g)]+=1
for i in range(80):f[(i+7)%9]+=f[i%9]
print(sum(f))