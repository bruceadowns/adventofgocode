def part1(lines):
    count = 0
    for i in range(0, len(lines) - 1):
        if lines[i] < lines[i + 1]:
            count += 1
    return count


def part1a(lines):
    pairs = zip(lines, lines[1:])
    increases = [b - a for a, b in pairs if a < b]
    return len(increases)


def part1b(lines):
    return len([True for a, b in zip(lines, lines[1:]) if a < b])


def part2(lines):
    count = 0
    for i in range(0, len(lines) - 3):
        if (
            lines[i] + lines[i + 1] + lines[i + 2]
            < lines[i + 1] + lines[i + 2] + lines[i + 3]
        ):
            count += 1
    return count


def part2a(lines):
    windows = [sum(w) for w in zip(lines, lines[1:], lines[2:])]
    pairs = zip(windows, windows[1:])
    increases = [b - a for a, b in pairs if a < b]
    return len(increases)


def part2b(lines):
    windows = [sum(w) for w in zip(lines, lines[1:], lines[2:])]
    return len([True for a, b in zip(windows, windows[1:]) if a < b])
