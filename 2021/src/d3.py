import os
from d1 import get_text

dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, '../inp/3.input')

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

# PART 2


def find_most_zerone(binary_list, bit, ox_co2):

    zero, one = 0, 0
    z_list, o_list = [], []
    bin_list = []
    entries = len(binary_list)
    
    for j in range(0, entries):
        if binary_list[j][bit] == '0':
            zero += 1
            z_list.append(binary_list[j])
        else:
            one += 1
            o_list.append(binary_list[j])

    if zero > one:
        if ox_co2 == '1':
            common = 0
            bin_list = z_list
        else:
            common = 1
            bin_list = o_list
    else:
        if ox_co2 == '1':
            common = 1
            bin_list = o_list
        else:
            common = 0
            bin_list = z_list


    if common is None:
        for j in bin_list:
            if j[bit] == ox_co2:
                bin_list = j

    return bin_list
            
    
bin_list = lines
co_gen = lines

for i in range(0, length):
    
    if len(bin_list) > 1:
        bin_list= find_most_zerone(bin_list, i, '1')
    if len(co_gen) > 1:
        co_gen = find_most_zerone(co_gen, i, '0')

    
oxy_gen = int("".join(bin_list), 2)
co_gen = int("".join(co_gen), 2)

# print(oxy_gen*co_gen)
# 1370737
