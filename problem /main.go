package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	City string
	Temp float64
}

func readlines(filename string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		file, err := os.Open(filename)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out
}

func parse(input <-chan string) <-chan Measurement {
	out := make(chan Measurement)

	go func() {
		defer close(out)

		for line := range input {
			parts := strings.Split(line, ";")
			if len(parts) != 2 {
				continue
			}

			city := parts[0]
			temp, err := strconv.ParseFloat(parts[1], 64)

			if err != nil {
				continue
			}

			out <- Measurement{
				City: city,
				Temp: temp,
			}
		}
	}()

	return out
}
func main() {
	lines := readlines("stations.csv")

	measurements := parse(lines)

	count := 0
	sum := 0

	for m := range measurements {
		count++
		if count == 10 {
			break
		}
		sum += int(m.Temp)

		fmt.Printf("%s %.1f\n", m.City, m.Temp)
	}
	fmt.Println(sum)
}
