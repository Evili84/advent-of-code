package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cube struct {
    color string
    amount  int64
}
	
type game struct {
    gameName string
    sets map[int][]cube
}

func main() {
	inputFile := os.Args[0]
	inputFile = "input.txt"

	if(inputFile == "") {
		fmt.Println("No input file given, exiting")
		return 
	}

	games := parseFile(inputFile)

	part1(games)

	part2(games)
}

func part1(games []game) {

	cubeCount := map[string]int64{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	
	sum := 0
	invalidGame := false
	for _, game := range games {		
		invalidGame = false
		for _, cubes := range game.sets {
			for _, cube := range cubes {
				if cubeCount[cube.color] < cube.amount {
					invalidGame = true
					break
				}
			}
			if invalidGame {
				break
			}
		}
		if !invalidGame {
			re := regexp.MustCompile("[\\d]+")
			value, err := strconv.ParseInt(re.FindString(game.gameName), 10, 64)

			if err != nil {
				fmt.Print(err)
				return
			}

			sum += int(value)
		}
	}

	fmt.Print(sum)
}

func part2(games []game) {

	cubeCount := map[string]int64{
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	sum := 0
	for _, game := range games {		
		for _, cubes := range game.sets {
			for _, cube := range cubes {
				if cubeCount[cube.color] < cube.amount {
					cubeCount[cube.color] = cube.amount
				}
			}
		}
		fmt.Println(cubeCount)
		gamesum := 0
		for key := range cubeCount {			 
			if cubeCount[key] > 0 {
				if gamesum == 0 {
					gamesum += int(cubeCount[key])
				} else {
					gamesum *= int(cubeCount[key])
				}
				
			}
			cubeCount[key] = 0
		}
		sum += gamesum
	}

	fmt.Println(sum)
}

func parseFile(filename string) []game {
	games := make([]game,0)
	f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer f.Close()
    r := bufio.NewReaderSize(f, 4*1024)
    line, isPrefix, err := r.ReadLine()
    for err == nil && !isPrefix {
		game := game{ gameName: "", sets:  make(map[int][]cube)}
        s := string(line)
		parts := strings.Split(s, ":")
		game.gameName = parts[0]

		subset := strings.Split(parts[1], ";")

		sets := make(map[int][]cube)
		for i := 0; i < len(subset); i++ {
			set := make([]cube,0)
			cubes := strings.Split(subset[i], ",")
			for j := 0; j < len(cubes); j++ {
				cube := cube{ color: "", amount: 0}
				re := regexp.MustCompile("[\\d]+")
				amount := re.FindString(cubes[j])
				cube.amount, err = strconv.ParseInt(amount, 10 , 64)
				re = regexp.MustCompile("[^\\d^\\s]+")
				cube.color = re.FindString(cubes[j])
				set = append(set, cube)
			}
			sets[i] = set
		}
		game.sets = sets
		games = append(games, game)
		line, isPrefix, err = r.ReadLine()
	}
	return games
}
