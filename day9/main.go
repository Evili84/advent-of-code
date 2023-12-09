package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type sequence struct {
	values [][]int
}

func main() {
	inputFile := "input.txt"
	part1(inputFile)
}

func part1(inputFile string) {
	fmt.Println("********* PART 1 START *********")

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	sequences := make([]sequence, 0)
	for err == nil && !isPrefix {
		seq := sequence{ values: make([][]int, 1)}
		s := string(line)
		for _, num := range strings.Split(s, " ") {
			value, _ := strconv.Atoi(num)
			seq.values[0] = append(seq.values[0], value)
		}
		sequences = append(sequences, seq)

		line, isPrefix, err = r.ReadLine()
	}

	sum, sum2 := 0, 0
	for index, seq := range sequences {

		calculcateSequence(&seq)
		sum += seq.values[0][len(seq.values[0])-1]
		sum2 += seq.values[0][0]
		sequences[index] = seq
	}
	fmt.Println(sum, sum2)
}

func calculcateSequence(seq *sequence) {
	steps := 0	
	seq.values = append(seq.values, make([]int, 0))
	for true {
		sums := 0 
		for i := 0; i < len(seq.values[steps])-1; i++ {			
			seq.values[steps+1] = append(seq.values[steps+1], seq.values[steps][i+1] - seq.values[steps][i])
			sums += seq.values[steps+1][len(seq.values[steps+1])-1]
		}		
		if slices.Max(seq.values[steps+1]) == slices.Min(seq.values[steps+1]) {
			break
		}
		steps++		

		seq.values = append(seq.values, make([]int, 0))
	}	
	steps = len(seq.values)-1
	for steps > 0 {
		seq.values[steps-1] = append(seq.values[steps-1], seq.values[steps][len(seq.values[steps])-1]+ seq.values[steps-1][len(seq.values[steps-1])-1])
		seq.values[steps-1] = append([]int{seq.values[steps-1][0] - seq.values[steps][0]}, seq.values[steps-1]...)
		steps--
	}
}

