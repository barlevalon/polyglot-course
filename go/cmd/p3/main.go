package main

import (
	"fmt"
	"strings"
)

type Thing = int

const (
	Tree Thing = iota
	Snow
)

func getInput() string {
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`;
}

func main() {
  treeCount := 0
  for r, line := range strings.Split(getInput() , "\n") {
    if line[r * 3 % len(line)] == '#' {
      treeCount += 1
    }
  }

  fmt.Printf("Tree count: %v\n", treeCount)
}
