guardPosition = []
initial_visited_positions = set()
direction = 'up'
total_cols = 0
total_rows = 0
p2 = 0

def read_file(file_path):
    global total_rows, total_cols
    with open(file_path, "r") as file:
        array = [list(line.strip()) for line in file if line.strip()]
        total_rows = len(array)
        total_cols = len(array[0]) if array else 0
        return array

def find_guard(array):
    for row_num, row in enumerate(array):
        for col_num, element in enumerate(row):
            if element == "^":
                guardPosition.append((row_num, col_num))

def turn_right(direction):
    if direction == 'up':
        return 'right'
    elif direction == 'right':
        return 'down'
    elif direction == 'down':
        return 'left'
    elif direction == 'left':
        return 'up'

def move(delta, row_num, col_num):
    return row_num + delta[0], col_num + delta[1]

direction_deltas = {
    'up': (-1, 0),
    'right': (0, 1),
    'down': (1, 0),
    'left': (0, -1),
}

def go_over_initial(array):
    global p2
    for visit in initial_visited_positions:
        temp_array = [row[:] for row in array]
        temp_array[visit[0]][visit[1]] = "#"

        direction = 'up'
        row_num, col_num, direction = part2_move_guard(temp_array, guardPosition[0][0], guardPosition[0][1], direction)

    return p2

def part2_move_guard(array, row_num, col_num, direction):
    global p2
    visited_positions = set()
    while True:
        delta = direction_deltas[direction]
        new_row, new_col = move(delta, row_num, col_num)
        new_pos = (new_row, new_col, direction)

        if new_row < 0 or new_row >= total_rows or new_col < 0 or new_col >= total_cols:
            break
        if new_pos in visited_positions:
            p2 += 1
            break

        if array[new_row][new_col] == '#':
            direction = turn_right(direction)
        else:
            visited_positions.add((new_row, new_col, direction))
            row_num, col_num = new_row, new_col 

    return row_num, col_num, direction

def move_guard(array, row_num, col_num, direction):
    initial_visited_positions.add((row_num, col_num))

    while True:
        delta = direction_deltas[direction]
        new_row, new_col = move(delta, row_num, col_num)

        if new_row < 0 or new_row >= total_rows or new_col < 0 or new_col >= total_cols:
            break

        if array[new_row][new_col] == '#':
            direction = turn_right(direction)
        else:
            initial_visited_positions.add((new_row, new_col))
            row_num, col_num = new_row, new_col 

    return row_num, col_num, direction

array = read_file("py/inputs/day6.txt")
find_guard(array)

for row_num, col_num in guardPosition:
    direction = 'up'
    row_num, col_num, direction = move_guard(array, row_num, col_num, direction)

print('Part 1: ', len(initial_visited_positions))
print('Part 2: ', go_over_initial(array))
