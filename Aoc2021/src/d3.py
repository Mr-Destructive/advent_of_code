import os
from d1 import get_text

dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, '../inp/inp3.txt')

lines = get_text(input_file)
lines = list(map(lambda s: s.strip(), lines))

entries = len(lines)
length = len(lines[0])

gamma_str, epsilon_str = "", ""

for i in range(length):
    zero, one = 0, 0
    bit = 0

    for j in range(0, entries):
        if lines[j][i] == '0':
            zero += 1
        else:
            one += 1
    
    if zero > one:
        gamma_str += "0"
        epsilon_str += "1"
    else:
        gamma_str += "1"
        epsilon_str += "0"

gamma_rate = int(gamma_str,2)
epsilon_rate = int(epsilon_str,2)

# print(gamma_rate * epsilon_rate)
