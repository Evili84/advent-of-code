package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type cardT struct {
	value int
	pieces int
}

type hand struct {
	Cardstring string
	Cards    []cardT
	Bid      int
}

func main() {
	inputFile := "input.txt"

	//part1(inputFile)
	part2(inputFile)
}

func part1(inputFile string) {
	cards := make(map[string]int, 14)

	cards["A"] = 14
	cards["K"] = 13
	cards["Q"] = 12
	cards["J"] = 11
	cards["T"] = 10
	cards["9"] = 9
	cards["8"] = 8
	cards["7"] = 7
	cards["6"] = 6
	cards["5"] = 5
	cards["4"] = 4
	cards["3"] = 3
	cards["2"] = 2

	hands := make([]hand, 0)

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		hand := hand{}
		parts := strings.Split(s, " ")

		hand.Cards = make([]cardT, 0)
		hand.Bid, _ = strconv.Atoi(parts[1])

		line, isPrefix, err = r.ReadLine()
		card := cardT{}
		hand.Cardstring = parts[0]

		for _, v := range parts[0] {
						
			cardCount := strings.Count(parts[0], string(v))
			
			card.pieces = cardCount
			card.value = cards[string(v)]
			if card.pieces > 0 {
				if slices.Index(hand.Cards, card) == -1 {
					hand.Cards = append(hand.Cards, card)
				}				
			}
			card = cardT{}		
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		handValue := cardCount(hands[i], false)
		handValue2 := cardCount(hands[j], false)

		if handValue != handValue2 {
			return handValue < handValue2
		}
		cardValue1 := cardValue(hands[i], cards, false)
		cardValue2 := cardValue(hands[j], cards, false)
		

		for i := 0; i < len(cardValue1); i++ {
			if cardValue1[i] != cardValue2[i] {
				return cardValue1[i] < cardValue2[i]
			}
		}
		return 1 < 2
	})

	sum := 0
	for index, hand := range hands {
		fmt.Println(hand)
		sum += (index + 1) * hand.Bid
	}
	fmt.Println("Total value: ", sum)	
}

func part2(inputFile string) {
	cards := make(map[string]int, 14)

	cards["A"] = 14
	cards["K"] = 13
	cards["Q"] = 12	
	cards["T"] = 10
	cards["9"] = 9
	cards["8"] = 8
	cards["7"] = 7
	cards["6"] = 6
	cards["5"] = 5
	cards["4"] = 4
	cards["3"] = 3
	cards["2"] = 2
	cards["J"] = 1
	
	hands := make([]hand, 0)

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		hand := hand{}
		parts := strings.Split(s, " ")

		hand.Cards = make([]cardT, 0)
		hand.Bid, _ = strconv.Atoi(parts[1])

		line, isPrefix, err = r.ReadLine()
		card := cardT{}
		hand.Cardstring = parts[0]

		for _, v := range parts[0] {
						
			cardCount := strings.Count(parts[0], string(v))
			
			card.pieces = cardCount
			card.value = cards[string(v)]
			if card.pieces > 0 {
				if slices.Index(hand.Cards, card) == -1 {
					hand.Cards = append(hand.Cards, card)
				}				
			}
			card = cardT{}		
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		handValue := cardCount(hands[i], true)
		handValue2 := cardCount(hands[j], true)

		if handValue != handValue2 {
			return handValue < handValue2
		}
		cardValue1 := cardValue(hands[i], cards, true)
		cardValue2 := cardValue(hands[j], cards, true)

		for i := 0; i < len(cardValue1); i++ {
			if cardValue1[i] != cardValue2[i] {
				return cardValue1[i] < cardValue2[i]
			}
		}
		return 1 < 2
	})

	sum := 0
	for index, hand := range hands {
		fmt.Println(hand, cardCount(hand, true))
		sum += (index + 1) * hand.Bid
	}
	fmt.Println("Total value: ", sum)	
}


func cardValue(hand hand, cards map[string]int, jokers bool) []int {
	values := make([]int, 0)
	for _, v := range hand.Cardstring {
		if jokers && string(v) == "J" {
			values = append(values, cards[string(v)])
		} else {
			values = append(values, cards[string(v)])
		}		
	}
	return values
}

func cardCount(hand hand, joker bool) int {
	counts := make([]int, 0)
	jokers := 0
	for _, v := range hand.Cards {
		if joker && v.value == 1 {
			jokers = v.pieces
		} else {
			counts = append(counts, v.pieces)
		}
	}

	if hand.Cardstring == "JJJ8J" {
		fmt.Println("tsek")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	
	if len(counts) == 0 && jokers == 5 {
		return 7
	}

	for i := 0; i < len(counts); i++ {
		if counts[i] == 5 || jokers == 5 || (joker && jokers + counts[i] == 5) {
			return 7
		} 

		if counts[i] == 4 || (joker && jokers + counts[i] == 4)  {
			return 6
		} 

		if counts[i] == 3 && counts[i+1] == 2 || (joker && jokers + counts[i] == 3 && counts[i+1] == 2)  {
			return 5
		} 

		if counts[i] == 3  || (joker && jokers + counts[i] == 3) {
			return 4
		} 

		if counts[i] == 2 && counts[i+1] == 2  {
			return 3
		} 		
		if counts[i] == 2 || (joker && jokers + counts[i] == 2) {
			return 2
		}
		return 1	
	}
	return 0
}