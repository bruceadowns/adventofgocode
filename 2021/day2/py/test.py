from solution import *

test_input = """forward 5
down 5
forward 8
up 3
down 8
forward 2
"""

lines = test_input.splitlines()


def test_part1():
    assert 150 == part1(lines)


def test_part1a():
    assert 150 == part1a(lines)


def test_part2():
    assert 900 == part2(lines)
