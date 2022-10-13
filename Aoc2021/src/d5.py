import os
from d1 import get_text

dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, '../inp/5.test')

lines = get_text(input_file)
for line in lines:
    x, y = line.strip().split(' -> ')
    x1, y1 = x.split(',')
    x2, y2 = y.split(',')
    print(x1, y1, x2, y2)
