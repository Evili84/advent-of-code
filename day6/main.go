package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

func part1(filename string) {
	times := make([]int64, 0)
	distances := make([]int64, 0)
	regx := regexp.MustCompile("(\\d+)")

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	row := 1
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()	
	for err == nil && !isPrefix {		
		s := string(line)
		if s != "" {
			values := regx.FindAllString(strings.Split(s, ":")[1], -1)
			for _, v := range values {
				val, _ := strconv.ParseInt(v, 10, 64) 
				if(row == 1) {
					times = append(times, val)
				} else {
					distances = append(distances, val)
				}
			}
		}		
		row++
		line, isPrefix, err = r.ReadLine()
	}
	
	fmt.Println(times,distances)
	result := int64(1)
	for i := 0; i < len(times); i++ {
		
		min := int64(math.Ceil(float64(distances[i]) / float64(times[i]))) 
		for (times[i] - min) * min <= distances[i] {
			min++
		}
		
		max := min
		for (times[i] - max) * max > distances[i]{
			//fmt.Println((times[i] - max) * max)
			max++
		}
		max--
		//fmt.Println(times[i], distances[i], min, max)
		result *= (max - min + 1)
	}
	fmt.Println(result)
}

func part2(filename string) {
	times := make([]int64, 0)
	distances := make([]int64, 0)
	regx := regexp.MustCompile("(\\d+)")

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	row := 1
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()	
	for err == nil && !isPrefix {		
		s := string(line)
		if s != "" {
			values := regx.FindAllString(strings.ReplaceAll(strings.Split(s, ":")[1], " ", ""), -1)
			for _, v := range values {
				val, _ := strconv.ParseInt(v, 10, 64) 
				if(row == 1) {
					times = append(times, val)
				} else {
					distances = append(distances, val)
				}
			}
		}		
		row++
		line, isPrefix, err = r.ReadLine()
	}
	
	fmt.Println(times,distances)
	result := int64(1)
	for i := 0; i < len(times); i++ {
		
		min := int64(math.Ceil(float64(distances[i]) / float64(times[i]))) 
		for (times[i] - min) * min <= distances[i] {
			min++
		}
		
		max := min
		for (times[i] - max) * max > distances[i]{
			max++
		}
		max--
		result *= (max - min + 1)
	}
	fmt.Println(result)
}