package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)
	
type card struct {
    cardNumber string	
    numbers []int
	winningNumbers []int
	pieces int
}


func main() {	
	inputFile := "input.txt"
	part1(inputFile)
	part2(inputFile)
}

func part1(filename string) {
	cards := make([]card, 0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		card := card {}
		s := string(line)
		parts := strings.Split(s, ":")
		numbers := strings.Split(parts[1], "|")[0]
		card.cardNumber = parts[0]
		
		for _, number := range strings.Split(numbers, " ") {
			if number != " " && number != "" {
				result, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println("Parse error", err)
					return
				}
				
				card.winningNumbers = append(card.winningNumbers, result)
			}
		}
		numbers = strings.Split(parts[1], "|")[1]

		for _, number := range strings.Split(numbers, " ") {
			if number != " " && number != "" {
				result, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println("Parse error", err)
					return
				}
				card.numbers = append(card.numbers, result)
			}
		}
		cards = append(cards, card)
		line, isPrefix, err = r.ReadLine()
	}

	sum := int64(0)
	for _, card := range cards {
		count := 0
		for _, number := range card.numbers {
			if slices.Index(card.winningNumbers, number) > -1 {
				count++
			}
		}
		fmt.Println(card.cardNumber, "Number count: ", count)
		if count > 0 {
			sum += int64(math.Pow(2, float64(count-1)))
		}
	}
	fmt.Println("Sum of cards: ", sum)
}

func part2(filename string) {
	cards := make([]card, 0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		card := card {}
		s := string(line)
		parts := strings.Split(s, ":")
		numbers := strings.Split(parts[1], "|")[0]
		card.cardNumber = parts[0]
		
		for _, number := range strings.Split(numbers, " ") {
			if number != " " && number != "" {
				result, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println("Parse error", err)
					return
				}
				
				card.winningNumbers = append(card.winningNumbers, result)
			}
		}
		numbers = strings.Split(parts[1], "|")[1]

		for _, number := range strings.Split(numbers, " ") {
			if number != " " && number != "" {
				result, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println("Parse error", err)
					return
				}
				card.numbers = append(card.numbers, result)
			}
		}
		card.pieces = 1
		cards = append(cards, card)
		line, isPrefix, err = r.ReadLine()
	}

	sum := int64(0)
	for index, card := range cards {
		count := 0
		for _, number := range card.numbers {
			if slices.Index(card.winningNumbers, number) > -1 {
				count++
			}
		}

		if count > 0 {
			cardIndex := index+1
			for i := 0; i < count; i++ {			
				cards[cardIndex].pieces += card.pieces
				cardIndex++
				if cardIndex > len(cards) {
					break
				}
			}
		}
		fmt.Println(card.cardNumber, "Number of cards count: ", card.pieces)
		sum += int64(card.pieces)
	}
	fmt.Println("Pieces of cards: ", sum)
}