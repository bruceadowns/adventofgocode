from solution import part1
from solution import part2

lines = [line for line in open("../input.txt").readlines() if line != ""]

print(part1(lines))
print(part2(lines))
