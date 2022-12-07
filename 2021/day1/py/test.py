from solution import *

test_input = """199
200
208
210
200
207

240
269
260
263
"""

numbers = [int(line) for line in test_input.splitlines() if len(line) > 0]


def test_part1():
    assert 7 == part1(numbers)


def test_part1a():
    assert 7 == part1a(numbers)


def test_part1b():
    assert 7 == part1b(numbers)


def test_part2():
    assert 5 == part2(numbers)


def test_part2a():
    assert 5 == part2a(numbers)


def test_part2b():
    assert 5 == part2b(numbers)
