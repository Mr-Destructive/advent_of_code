import os
import copy
from pathlib import Path


def get_board(s):
    return [int(n) for n in s.split()]


def is_bingo(board, current_board):
    for i in range(5):
        for j in range(5):
            if board[i * 5 + j] in current_board:
                break
        else:
            return True

        for j in range(5):
            if board[i + 5 * j] in current_board:
                break
        else:
            return True
    else:
        return False


dirname = os.path.dirname(__file__)
input_file = os.path.join(dirname, "../inp/4.test")

content = Path(input_file).read_text().split("\n\n")

inp = [n for n in content[0].split(",")]
content = content[1:]
boards = []

for i, b in enumerate(content):
    boards.append(get_board(content[i]))

boards_copy = copy.deepcopy(boards)
over, win_over = False, False
num_of_boards = len(boards)
last_call = inp[0]
winner, unmarked = 0, 0
marked_boards = set()

for i, call in enumerate(inp):
    for b in range(num_of_boards):
        if int(call) in boards[b]:
            boards[b].remove(int(call))

    for b in range(num_of_boards):
        unmarked = 0
        if is_bingo(boards_copy[b], boards[b]):
            for c in boards[b]:
                unmarked += c
                last_call = int(call)
                if not win_over:
                    print(f"Board {b+1} won!")
                    win_over = True
                # over = True
            if not b in marked_boards:
                marked_boards.add(b)
            if len(marked_boards) == num_of_boards:
                print(f"Board {b+1} won!")
                over = True
                break
    if win_over and not over:
        print("First Winner = ", last_call * unmarked)
    if len(marked_boards) == num_of_boards:
        over = True
        break
print("Last Winner", unmarked * last_call)
