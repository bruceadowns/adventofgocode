from parse import parse


def part1(lines):
    horizontal = 0
    depth = 0

    for line in lines:
        result = parse("{} {}", line)
        direction = result[0]
        scalar = int(result[1])

        if direction == "forward":
            horizontal += scalar
        elif direction == "down":
            depth += scalar
        elif direction == "up":
            depth -= scalar

    return horizontal * depth


def part1a(lines):
    horizontal = 0
    depth = 0

    steps = [parse("{} {}", line) for line in lines]
    steps = [(direction, int(scalar)) for direction, scalar in steps]
    for direction, scalar in steps:
        if direction == "forward":
            horizontal += scalar
        elif direction == "down":
            depth += scalar
        elif direction == "up":
            depth -= scalar

    return horizontal * depth


def part2(lines):
    horizontal = 0
    depth = 0
    aim = 0

    for line in lines:
        result = parse("{} {}", line)
        direction = result[0]
        scalar = int(result[1])

        if direction == "forward":
            horizontal += scalar
            depth += aim * scalar
        elif direction == "down":
            aim += scalar
        elif direction == "up":
            aim -= scalar

    return horizontal * depth
