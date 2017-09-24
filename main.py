import random

f = open("./artist_list.csv").read().split("\n")

random.shuffle(f)

print f[0]


