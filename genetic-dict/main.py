dict1 = {"GGA": 64231, "GKAT": 66582, "GAT": 66582, "GGAT": 6682}
dict2 = {"TCC": 64231, "ATC": 682, "GAT": 66582}

dict3 = {
    "%s:%s" % (k, l): v
    for (k, v), (l, b) in zip(
        sorted(dict1.items(), key=lambda x: x[1]),
        sorted(dict2.items(), key=lambda x: x[1]),
    )
}
print(dict3)

print(zip(
        sorted(dict1.items(), key=lambda x: x[1]),
        sorted(dict2.items(), key=lambda x: x[1]),
    ))

list(set(["A","A","A","B","C","C","D","D","D","B","B"])) # => ['D', 'C', 'B', 'A']
