from collections import deque
import os
import requests
from time import perf_counter


def get_input(year, day):
    session = os.environ["AOC_SESSION"]
    ret = requests.get(
        f"https://adventofcode.com/{year}/day/{day}/input",
        cookies=dict(session=session)
    )

    return ret.text


def parse_input(text):
    heights = []
    zeros = []
    start = None
    end = None
    for row, line in enumerate(text.strip().split()):
        heights.append([])
        for col, c in enumerate(line):
            if c == 'S':
                heights[-1].append(0)
                start = (row, col)
            elif c == 'E':
                heights[-1].append(25)
                end = (row, col)
            else:
                heights[-1].append(ord(c) - ord('a'))

            if heights[-1][-1] == 0:
                zeros.append((row, col))
            
    return heights, start, end, zeros

def shortest_path(height, start, end):
    n_rows = len(height)
    n_cols = len(height[0])

    distance = [[-1 for _ in range(n_cols)] for _ in range(n_rows)]

    row, col = start
    distance[row][col] = 0
    to_check = deque()
    to_check.append(start)
    directions = ((0, 1), (0, -1), (1, 0), (-1, 0))

    while len(to_check) > 0:
        row, col = to_check.popleft()

        for drow, dcol in directions:
            nrow, ncol = row + drow, col + dcol
            # outside the grid
            if nrow < 0 or nrow == n_rows or ncol < 0 or ncol == n_cols:
                continue

            # already visited
            if distance[nrow][ncol] != -1:
                continue

            # too high to reach
            if (height[nrow][ncol] - height[row][col]) > 1:
                continue

            distance[nrow][ncol] = distance[row][col] + 1

            # found target position
            if (row, col) == (end):
                return distance[row][col]

            to_check.append((nrow, ncol))

    return 2**64



if __name__ == "__main__":
    text = get_input(2022, 12)
    height, start, end, zeros = parse_input(text)
    t0 = perf_counter()
    answer = shortest_path(height, start, end)
    print(f'Part 1: {(perf_counter() - t0) * 1e3:.3f} ms')
    print(answer)

    t0 = perf_counter()
    shortest = 2**64
    n_rows = len(height)
    n_cols = len(height[0])
    for pos in zeros:
        path_length = shortest_path(height, pos, end)
        shortest = min(shortest, path_length)
    print(f'Part 2: {(perf_counter() - t0) * 1e3:.3f} ms')
    print(shortest)
