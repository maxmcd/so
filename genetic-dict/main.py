# dict1 = {"GGA": 64231, "GAT": 66582}
# dict2 = {"TCC": 64231, "ATC": 66582}

# dict3 = {
#     "%s:%s" % (k, l): v
#     for (k, v), (l, b) in zip(
#         sorted(dict1.items(), key=lambda x: x[1]),
#         sorted(dict2.items(), key=lambda x: x[1]),
#     )
# }
# print(dict3)

def  lcw(u,v):
    for r in range(len(u)+1):
        lcw[r][len(v)+1]=0
    for c in range(len(v)+1):
        lcw[len(u)+1][c]=0
    maxlcw=0
    for c in range(len(v)+1,-1,-1):
        for r in range(len(u)+1,-1,-1):
            if u[r]==v[c]:
                lcw[r][c]=1+lcw[r+1][c+1]
            else:
                lcw[r][c]=0
            if  lcw[r][c]>maxlcw:
                maxlcw=lcw[r][c]
     return(maxlcw)
lcw(['fgshf'],['vdbhf'])  
