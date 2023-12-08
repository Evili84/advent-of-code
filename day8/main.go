package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	inputFile := "input.txt"
	part1(inputFile)
	part2(inputFile)
}

type order struct {
	input []rune
	locations []string
	instructions [][]string
}

type pathMap struct {
	input []rune
	locations map[string][]string
}

func part2(inputFile string) {

	fmt.Println("********* PART 2 START *********")

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	
	inputValue := make(map[string]int)
	inputValue["R"] = 1
	inputValue["L"] = 0
	
	path := pathMap{locations: make(map[string][]string, 0)}

	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	lineIndex := 0
	for err == nil && !isPrefix {
		s := string(line)
		
		if lineIndex == 0 {
			path.input = []rune(s)
			
		} else if s != ""  {
			parts := strings.Split(s, " = ")			
			value := strings.Split(strings.ReplaceAll(strings.ReplaceAll(parts[1], ")", ""), "(", ""), ", ")
			path.locations[parts[0]] = value
		}
		lineIndex++
		line, isPrefix, err = r.ReadLine()
	}
	

	steps := make([]string, 0)	
	command := 0
	keys := maps.Keys(path.locations)	
	stepCount := uint64(1)

	for i := 0; i < len(keys); i++ {
		if strings.LastIndex(keys[i],"A") == 2 {
			steps = append(steps, keys[i])
		}
	}

	ends := make([]int, len(steps))
	endCount := 0

	for true {
		for i := 0; i < len(steps); i++ {			
			steps[i] = path.locations[steps[i]][inputValue[string(path.input[command])]]	
					
			if strings.LastIndex(steps[i], "Z") > -1 {
				ends[i] = int(stepCount)
				endCount++
			}
		}

		if command < len(path.input) -1 {
			command++
		} else {
			command = 0
		}

		if endCount == len(steps)  {
			break
		}
		stepCount++
	}
	sort.Sort(sort.IntSlice(ends))
	one, two := ends[0], ends[1]
	ends = slices.Delete(ends, 0,2)
	fmt.Println(LCM(one, two, ends...))

}

func part1(inputFile string) {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	
	inputValue := make(map[string]int)
	inputValue["R"] = 1
	inputValue["L"] = 0
	
	path := order{instructions: make([][]string, 0), locations: make([]string, 0)}

	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	lineIndex := 0
	for err == nil && !isPrefix {
		s := string(line)
		
		if lineIndex == 0 {
			path.input = []rune(s)
			
		} else if s != ""  {
			parts := strings.Split(s, " = ")
			path.locations = append(path.locations, parts[0])
			value := strings.Split(strings.ReplaceAll(strings.ReplaceAll(parts[1], ")", ""), "(", ""), ", ")
			path.instructions = append(path.instructions, value)

		}
		lineIndex++
		line, isPrefix, err = r.ReadLine()
	}

	location := path.locations[0]
	stepCount := 1
	command := 0
	index := slices.Index(path.locations, "AAA")
	indexEnd := slices.Index(path.locations, "ZZZ")
	for 1 < 2 {
		location = path.locations[index]
		next := path.instructions[index][inputValue[string(path.input[command])]]
		index = slices.Index(path.locations, next)

		if command < len(path.input) -1 {
			command++
		} else {
			command = 0
		}
		fmt.Println(location, path.instructions[index], string(path.input[command]), next, index)
		if index == indexEnd {
			break
		}		
		stepCount++		
	}

	fmt.Println(stepCount)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}