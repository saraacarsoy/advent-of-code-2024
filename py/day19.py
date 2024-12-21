def can_form_string(towels, designs):
    if designs == "":
        return True

    for comb in towels:
        if designs.startswith(comb):
            remaining_string = designs[len(comb):]
            if can_form_string(towels, remaining_string): # I despise recursion #
                return True

    return False

def check_all_combinations(towels, designs):
    count = 0
    for target in designs:
        if can_form_string(towels, target):
            count += 1
    print(count)

def parse_input_from_file():
    with open("py/inputs/day19.txt", 'r') as file:
        lines = file.readlines()
    
    towel_list = [x.strip() for x in lines[0].strip().split(',')]
    designs = [line.strip() for line in lines[2:]]
    
    return towel_list, designs

towel_list, designs = parse_input_from_file()
check_all_combinations(towel_list, designs)
