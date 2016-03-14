import random
import sys

print(1)
for i in range(0,100000):
    if random.randint(0,1) == 0:
        sys.stdout.write('A')
    else:
        sys.stdout.write('B')
sys.stdout.write('\n')


