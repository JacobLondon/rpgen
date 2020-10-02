import random
import signal
import subprocess
import sys
import re
import os

cancel = False
def handle(sig, frame):
    global cancel
    cancel = True
signal.signal(signal.SIGINT, handle)

def roll_one(n: int) -> int:
    return random.randint(1, n)

def roll_add(n: int, v: int) -> int:
    return roll_one(n) + v

def roll_many(xs: list):
    try:
        die = int(next(xs))
    except StopIteration:
        return

    try:
        yield roll_add(die, int(next(xs)))
    except StopIteration:
        yield roll_one(die)

    yield from roll_many(xs)

print(">> ", end='', flush=True)
for line in sys.stdin:
    args = line.split()

    if not args:
        continue

    if re.match(r'[0-9]+', args[0]):
        try:
            getlist = lambda: list(roll for roll in roll_many(iter(args)))
            manylist = list(getlist() for _ in range(10))
            print(manylist)
        except ValueError as e:
            print(e)
            continue

    elif args[0] == 'char':
        try:
            with open('out.txt', 'r') as fp:
                lines = fp.readlines()
                while not cancel:
                    choice = random.choice(lines)
                    if choice[0] == '\t' and all(arg in choice for arg in args[1:]):
                        print(choice.lstrip().rstrip())
                        break
        except:
            pass

    elif args[0] == 'gen':
        proc = subprocess.Popen("./rpgen", stdout=subprocess.PIPE)
        proc.communicate()

    elif args[0] == 'c' or args[0] == 'clear':
        print(chr(27) + "[2J")

    cancel = False
    print(">> ", end='', flush=True)
