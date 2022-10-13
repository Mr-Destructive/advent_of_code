import os
from d1 import get_text

dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, '../inp/7.input')

lines = get_text(input_file)
positions = lines[0].strip()

positions = positions.split(',')
positions = [int(f) for f in positions]
length = len(positions)

# PART 1
fuel_cost = []
for i in range(length):
    pos = [abs(p-i) for p in positions]
    fuel_cost.append(sum(pos))

# print(min(fuel_cost))
# 347011


# PART 2
fuel_cost = []
for i in range(length):
    pos = [int(abs(p-i)*(abs(p-i)+1)/2) for p in positions]
    fuel_cost.append(sum(pos))

# print(min(fuel_cost))
# 98363777
