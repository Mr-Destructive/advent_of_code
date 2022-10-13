import os
from d1 import get_text

dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, '../inp/6.input')

lines = get_text(input_file)
fishes = lines[0].strip()

fishes = fishes.split(',')
fishes = [int(f) for f in fishes]

def count_fishes(days):
    day_count = {k:0 for k in range(0, 9)}
    for i in fishes:
        day_count[i] = day_count[i]+1

    for i in range(days):
        day = i % len(day_count)
        day_count[(day + 7) % len(day_count)] += day_count[day]
    return day_count

# PART 1
# print(sum(count_fishes(80).values()))
# 391671

# PART 2
# print(sum(count_fishes(256).values()))
# 1754000560399
