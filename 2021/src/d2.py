import os
from d1 import get_text

dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, '../inp/inp2.txt')

lines = get_text(input_file)

horizontal, depth = 0, 0

for i in lines:
    func, dist = i.split(' ')
    dist = int(dist)
    if func == 'forward':
        horizontal += dist
    elif func == 'up':
        depth -= dist
    elif func == 'down':
        depth += dist

# print(horizontal*depth)
# 1924923

horizontal, aim, depth = 0, 0, 0

for i in lines:
    func, dist = i.split(' ')
    dist = int(dist)
    if func == 'forward':
        horizontal += dist
        depth += aim*dist
    elif func == 'up':
        aim -= dist
    elif func == 'down':
        aim += dist

# print(horizontal*depth)
# 1982495697



