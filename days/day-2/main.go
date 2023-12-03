
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "regexp"
  "strconv"
)

func main() {
  part1()
  part2()
}

func part1() {
  file, err := os.Open("days/day-2/input.txt")

  if err != nil {
    panic(err)
  }

  defer file.Close()

  sum := 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if gamePossible(line) {
      sum += getGamenum(line)
    }
  }

  fmt.Println("Part 1: ", sum)
}

func part2() {
  file, err := os.Open("days/day-2/input.txt")

  if err != nil {
    panic(err)
  }

  defer file.Close()

  sum := 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    sum += getGamePower(line)
  }

  fmt.Println("Part 2: ", sum)
}

var limits = map[string]int {
  "red": 12,
  "green": 13,
  "blue": 14,
}

func gamePossible(line string) bool {
  pulls := getPulls(line)

  for _, pull := range pulls {
    n, p := getPull(pull)

    if limits[p] < n {
      return false
    }
  }

  return true
}

func getGamenum(line string) int {
  chunks := strings.Split(line, ":")

  game := strings.Split(chunks[0], " ")[1]

  gameNum, err := strconv.Atoi(game)

  if err != nil {
    panic(err)
  }

  return gameNum
}

func getPulls(line string) []string {
  chunks := strings.Split(line, ":")

  re := regexp.MustCompile("(,|;)")
  pulls := re.Split(chunks[1], -1)

  return pulls
}

func getPull(pull string) (int, string) {
  p := strings.Split(strings.TrimSpace(pull), " ")

  n, err := strconv.Atoi(p[0])

  if err != nil {
    panic(err)
  }

  return n, p[1]
}

func getGamePower(line string) int {
  pulls := getPulls(line)

  minPulls := map[string]int {
    "red": 0,
    "green": 0,
    "blue": 0,
  }

  for _, pull := range pulls {
    n, p := getPull(pull)

    if minPulls[p] < n {
      minPulls[p] = n
    }
  }

  return minPulls["red"] * minPulls["green"] * minPulls["blue"]
}
