import itertools
import re

def evaluate_left_to_right(expression):
    tokens = re.split(r'(\D)', expression)
    
    result = int(tokens[0])
    
    for i in range(1, len(tokens) - 1, 2):
        operator = tokens[i]
        next_number = int(tokens[i + 1])
        
        if operator == '+':
            result += next_number
        elif operator == '*':
            result *= next_number

    return result

def generate_all_combinations(arr):
    operators = ['+', '*']
    num_slots = len(arr) - 1
    all_combinations = []

    for operator_combination in itertools.product(operators, repeat=num_slots):
        expression = str(arr[0])
        for i, op in enumerate(operator_combination):
            expression += op + str(arr[i + 1])
        all_combinations.append(expression)
    
    return all_combinations

def parse_input_from_file():
    result = []
    with open("py/inputs/day7.txt", 'r') as file:
        lines = file.readlines()
        for line in lines:
            x_part, y_part = line.split(':')
            x = int(x_part.strip())
            y = list(map(int, y_part.strip().split()))
            result.append((x, y))
    return result

count = 0
for input in parse_input_from_file():
    combinations = generate_all_combinations(input[1])
    for c in combinations:
        if input[0] == evaluate_left_to_right(c):
            count += input[0]
            break
print(count)
