package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)



type seed struct {
	start int64
	end int64
}

type conversion struct {
	sourceStart int64
	destinationStart int64
	sourceEnd int64
	destinationEnd int64
}

func main() {
	inputFile := "input.txt"
	part1(inputFile)
	part2(inputFile)
}

func part1(filename string) {

	seeds := make([]int64, 0)
	maps := make(map[string][]conversion)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	firstLine := true
	lastString := ""
	for err == nil && !isPrefix {
		s := string(line)
		if s != "" {
			if firstLine {
				firstLine = false
				parts := strings.Split(s, ":")
				numbers := strings.Split(parts[1], " ")

				for _, v := range numbers {
					if v != "" {
						number, _ := strconv.ParseInt(v, 10, 64)
						seeds = append(seeds, number)
					}
				}
			} else {
				if strings.Index(s, ":") > -1 {
					lastString = strings.Split(s, ":")[0]
					maps[lastString] = make([]conversion, 0)
				} else {
					numbers := strings.Split(s, " ")					
					startSource := int64(0)
					startDestination := int64(0)
					for index, v := range numbers {						
						if v != "" {
							number, _ := strconv.ParseInt(v, 10, 64)
							switch index {
							case 0:
								startSource = number
							case 1:
								startDestination = number							
							case 2:
								conv := conversion {}
								conv.destinationStart = startSource
								conv.destinationEnd = startSource + 1*number
								conv.sourceStart = startDestination
								conv.sourceEnd = startDestination + 1*number
								maps[lastString] = append(maps[lastString], conv)
							}
						}
					}					
				}
			}
		}		
		line, isPrefix, err = r.ReadLine()
	}	
	slices.Sort(seeds)
	soil := int64(0)
	fertilizer := int64(0)
	water := int64(0)
	light := int64(0)
	temp := int64(0)
	humid := int64(0)
	loc := int64(0)
	smallestLoc := int64(math.MaxInt64)
	for _, seed := range seeds {
		soil = findValue(maps["seed-to-soil map"], seed)
		fertilizer = findValue(maps["soil-to-fertilizer map"], soil)
		water = findValue(maps["fertilizer-to-water map"], fertilizer)
		light = findValue(maps["water-to-light map"], water)
		temp = findValue(maps["light-to-temperature map"], light)
		humid = findValue(maps["temperature-to-humidity map"], temp)
		loc = findValue(maps["humidity-to-location map"], humid)

		if loc < int64(smallestLoc) {
			smallestLoc = loc
		}			
	}
	fmt.Println("Part 1 smallest location: ", smallestLoc)
}

func part2(filename string) {
	seedtosoil := "seed-to-soil map"
	soiltofertilizer := "soil-to-fertilizer map"
	fertilizertowater := "fertilizer-to-water map"
	watertolight := "water-to-light map"
    lighttotemperature := "light-to-temperature map"
	temperaturetohumidity := "temperature-to-humidity map"
	humiditytolocation := "humidity-to-location map"


	seeds := make([]seed, 0)
	maps := make(map[string][]conversion)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	firstLine := true
	lastString := ""
	for err == nil && !isPrefix {
		s := string(line)
		if s != "" {
			if firstLine {
				firstLine = false
				parts := strings.Split(s, ": ")
				numbers := strings.Split(parts[1], " ")

				for i := 1; i < len(numbers); i = i+2 {
					seed := seed {}
					seed.start, _ = strconv.ParseInt(numbers[i-1], 10, 64)
					end, _ := strconv.ParseInt(numbers[i], 10, 64)
					seed.end = seed.start + 1*end-1
					seeds = append(seeds, seed)
				}

			} else {
				if strings.Index(s, ":") > -1 {
					lastString = strings.Split(s, ":")[0]
					maps[lastString] = make([]conversion, 0)
				} else {
					numbers := strings.Split(s, " ")					
					startSource := int64(0)
					startDestination := int64(0)
					for index, v := range numbers {						
						if v != "" {
							number, _ := strconv.ParseInt(v, 10, 64)
							switch index {
							case 0:
								startSource = number
							case 1:
								startDestination = number							
							case 2:
									conv := conversion {}
									conv.destinationStart = startSource
									conv.destinationEnd = startSource + number -1
									conv.sourceStart = startDestination
									conv.sourceEnd = startDestination + number -1
									maps[lastString] = append(maps[lastString], conv)
							}
						}
					}					
				}
			}
		}		
		line, isPrefix, err = r.ReadLine()
	}	

	smallestLoc := int64(math.MaxInt64)

	soil := int64(0)
	fertilizer := int64(0)
	water := int64(0)
	light := int64(0)
	temp := int64(0)
	humid := int64(0)
	found := false

	// Calculate the location backwards

	for i := int64(0); i < smallestLoc; i++ {

		humid = findValueReverse(maps[humiditytolocation], i)
		temp = findValueReverse(maps[temperaturetohumidity], humid)
		light = findValueReverse(maps[lighttotemperature], temp)
		water = findValueReverse(maps[watertolight], light)
		fertilizer = findValueReverse(maps[fertilizertowater], water)
		soil = findValueReverse(maps[soiltofertilizer], fertilizer)
		seed := findValueReverse(maps[seedtosoil], soil)

		for _, seedRange := range seeds {
			if seed >= seedRange.start && seed <= seedRange.end {
				fmt.Println("Part 2 smallest location: ", i, seed)
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	
	soil = findValue(maps[seedtosoil], 4002147451)
	fertilizer = findValue(maps[soiltofertilizer], soil)
	water = findValue(maps[fertilizertowater], fertilizer)
	light = findValue(maps[watertolight], water)
	temp = findValue(maps[lighttotemperature], light)
	humid = findValue(maps[temperaturetohumidity], temp)
	loc := findValue(maps[humiditytolocation], humid)

	fmt.Println(loc)
}

func findValueReverse(conversions []conversion, value int64) int64 {
	result := value
	for _, conv := range conversions  {
		if value >= conv.destinationStart && value <= conv.destinationEnd {
			result = conv.sourceStart + (value - conv.destinationStart)
			return  result
		}		
	}
	return result
}


func findValue(conversions []conversion, value int64) int64 {
	result := value
	for _, conv := range conversions  {
		if value >= conv.sourceStart && value <= conv.sourceEnd {
			result = conv.destinationStart + (value - conv.sourceStart)
			return result			
		}		
	}
	return result
}

