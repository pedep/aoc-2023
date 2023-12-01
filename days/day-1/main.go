// golang hello world

package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
)

func main() {
  file, err := os.Open("days/day-1/input.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var part1 []int
  var part2 []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    part1 = append(part1, part1Line(scanner.Text()))
    part2 = append(part2, part2Line(scanner.Text()))
  }

  part1sum := 0
  for _, num := range part1 {
    part1sum += num
  }

  part2sum := 0
  for _, num := range part2 {
    part2sum += num
  }

  fmt.Println(fmt.Sprintf("Part 1: %d", part1sum))
  fmt.Println(fmt.Sprintf("Part 2: %d", part2sum))
}

func part1Line(line string) int {
  var first string
  var last string
  for i := 0; i < len(line); i++ {
    if line[i] >= byte('0') && line[i] <= byte('9') {
      if first == "" {
        first = string(line[i])
      }

      last = string(line[i])
    }
  }

  n, err := strconv.Atoi(first + last)

  if err != nil {
    panic(err)
  }

  return n
}

func part2Line(line string) int {
  var first struct {char string; index int}
  var last struct {char string; index int}

  strings := []struct {char string; strings []string}{
    {"0", []string{"0"}},
    {"1", []string{"1", "one"}},
    {"2", []string{"2", "two"}},
    {"3", []string{"3", "three"}},
    {"4", []string{"4", "four"}},
    {"5", []string{"5", "five"}},
    {"6", []string{"6", "six"}},
    {"7", []string{"7", "seven"}},
    {"8", []string{"8", "eight"}},
    {"9", []string{"9", "nine"}},
  }

  for _, str := range strings {
    for _, substr := range str.strings {
      min, max := findSubstring(line, substr)

      if min == -1 || max == -1 {
        continue
      }

      if first.char == "" || min < first.index {
        first.char = str.char
        first.index = min
      }

      if last.char == "" || max > last.index {
        last.char = str.char
        last.index = max
      }
    }
  }

  n, err := strconv.Atoi(first.char + last.char)

  if err != nil {
    panic(err)
  }

  return n
}

func findSubstring(s string, substr string) (int, int) {
  min := -1
  max := -1

  for i := 0; i < len(s); i++ {
    if s[i] == substr[0] {
      if i+len(substr) > len(s) {
        continue
      }

      if s[i:i+len(substr)] == substr {
        if min == -1 || i < min {
          min = i
        }

        if max == -1 || i > max {
          max = i
        }
      }
    }
  }

  return min, max
}
