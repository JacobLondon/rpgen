import matplotlib.pyplot as plt
import random
import os
from collections import defaultdict

random.seed(os.urandom(128))
def randint(a, b, seed_bytes=128):
    random.seed(os.urandom(seed_bytes))
    return random.randint(a,b)

def d20():
    return randint(1, 20)

count = 0
bins = defaultdict(int)

GOES = 1000
FGOES = float(GOES)

for _ in range(GOES):
    bins[d20()] += 1

for key, val in bins.items():
    print(f"{key}\t{val / FGOES}%")

avg = sum(bins.values()) / 20.0 / FGOES
print(f"Avg: {avg}%")
