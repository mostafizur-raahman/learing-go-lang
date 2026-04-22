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

type AggregateData struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

type CityStats struct {
	City  string
	Min   float64
	Max   float64
	Mean  float64
	Count int
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

func filterAbove30(input <-chan Measurement) <-chan Measurement {
	out := make(chan Measurement)

	go func() {
		defer close(out)

		for m := range input {
			if m.Temp > 30 {
				out <- m
			}
		}
	}()
	return out
}

func aggregate(input <-chan Measurement) <-chan CityStats {
	out := make(chan CityStats)
	go func() {
		defer close(out)
		stats := make(map[string]*AggregateData)

		for m := range input {
			data, exist := stats[m.City]

			if !exist {
				stats[m.City] = &AggregateData{
					Min:   m.Temp,
					Max:   m.Temp,
					Sum:   m.Temp,
					Count: 1,
				}
			} else {
				if m.Temp < data.Min {
					data.Min = m.Temp
				}
				if m.Temp > data.Max {
					data.Max = m.Temp
				}
				data.Sum += m.Temp
				data.Count++
			}
		}

		for city, data := range stats {
			out <- CityStats{
				City:  city,
				Min:   data.Min,
				Max:   data.Max,
				Mean:  data.Sum / float64(data.Count),
				Count: data.Count,
			}
		}
	}()

	return out
}
func main() {
	lines := readlines("stations.csv")

	measurements := parse(lines)

	hotReadings := filterAbove30(measurements)

	results := aggregate(hotReadings)

	for result := range results {
		fmt.Println(result.City)
	}
}
