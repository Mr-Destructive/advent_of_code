import os
dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, '../inp/inp1.txt')

def get_text(input_file):
    with open(input_file, 'r') as f:
        lines = f.readlines()
    return lines

with open(input_file, 'r') as f:
    lines = f.readlines()

entries = len(lines)
count = 0

lines[0] = int(lines[0].strip())

for i in range(1, entries):
    lines[i] = int(lines[i].strip())
    if lines[i] > lines[i-1]:
        count += 1

# print(count)

count2 = 0

w1 = 0
for i in range(0, entries-2):
    w2 = lines[i] + lines[i+1] + lines[i+2]
    if i > 2:
        if w2 > w1:
            count2 += 1
    w1 = lines[i] + lines[i+1] + lines[i+2]
    i += 3

# print(count2)
