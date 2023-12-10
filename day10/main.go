package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type pipe struct {
	connection []int
	connected bool
}

type pipeMaze struct {
	pipes [][]pipe
	start []int
}

func main() {
	inputFile := "input.txt"
	pipeDream := buildPipeDream(inputFile)
	part1(pipeDream)
	part2(pipeDream)
}

func part1(pipeDream pipeMaze) {
	fmt.Println("********* PART 1 START *********")

	pos := make([]int,4)
	pos2 := pos

	if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]+1].connection[3] == 1 {
		pos = []int{pipeDream.start[0],pipeDream.start[1]+1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]+1][pipeDream.start[1]].connection[0] == 1 {
		pos = []int{pipeDream.start[0]+1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]-1].connection[2] == 1 {
		pos = []int{pipeDream.start[0],pipeDream.start[1]-1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]-1][pipeDream.start[1]].connection[1] == 1 { 
		pos = []int{pipeDream.start[0]-1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	}

	if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]-1].connection[1] == 1 {
		pos2 = []int{pipeDream.start[0],pipeDream.start[1]-1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]+1][pipeDream.start[1]].connection[0] == 1 {
		pos2 = []int{pipeDream.start[0]+1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]-1].connection[1] == 1 {
		pos2 = []int{pipeDream.start[0],pipeDream.start[1]-1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]-1][pipeDream.start[1]].connection[3] == 1 { 
		pos2 = []int{pipeDream.start[0]-1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	}

	pipeDream.pipes[pos[0]][pos[1]].connected = true
	pipeDream.pipes[pos2[0]][pos2[1]].connected = true

	steps := 1
	for true {
		steps++
		pipeDream.pipes[pos[0]][pos[1]].connected = true
		pipe := pipeDream.pipes[pos[0]][pos[1]]
		
		if pipe.connection[0] == 1 && (pos[0]-1 != pos[2] || pos[1] != pos[3]) {
			pos = []int{pos[0]-1,pos[1],pos[0],pos[1]}
		} else if pipe.connection[1] == 1 && (pos[0] != pos[2] || pos[1] +1 != pos[3]) {
			pos = []int{pos[0],pos[1]+1,pos[0],pos[1]}
		} else if pipe.connection[2] == 1 && (pos[0]+1 != pos[2] || pos[1] != pos[3]) {
			pos = []int{pos[0]+1,pos[1],pos[0],pos[1]}
		} else if pipe.connection[3] == 1 && (pos[0] != pos[2] || pos[1] - 1 != pos[3]) {
			pos = []int{pos[0],pos[1]-1,pos[0],pos[1]}		
		} 
		
		pipeDream.pipes[pos2[0]][pos2[1]].connected = true
		pipe = pipeDream.pipes[pos2[0]][pos2[1]]

		if pipe.connection[3] == 1 && (pos2[0] != pos2[2] || pos2[1] - 1 != pos2[3]) {
			pos2 = []int{pos2[0],pos2[1]-1,pos2[0],pos2[1]}
		} else if pipe.connection[2] == 1 && (pos2[0]+1 != pos2[2] || pos2[1] != pos2[3]) {
			pos2 = []int{pos2[0]+1,pos2[1],pos2[0],pos2[1]}			
		} else if pipe.connection[1] == 1 && (pos2[0] != pos2[2] || pos2[1] +1 != pos2[3]) {
			pos2 = []int{pos2[0],pos2[1]+1,pos2[0],pos2[1]}
		} else if pipe.connection[0] == 1 && (pos2[0]-1 != pos2[2] || pos2[1] != pos2[3]) {
			pos2 = []int{pos2[0]-1,pos2[1],pos2[0],pos2[1]}
		} 

		if pos[0] == pos2[0] && pos[1] == pos2[1] {
			break
		}
		
	}
	fmt.Println(steps)
}

func buildPipeDream(inputFile string) pipeMaze {
	pipeDream := pipeMaze{pipes: make([][]pipe, 0), start: make([]int,2)}
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return pipeDream
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()

	for err == nil && !isPrefix {		
		row := make([]pipe,0)
		s := string(line)
		
		for index, char := range s {
			if char != 46 {
				connection := make([]int, 0)
				pipe := pipe{}
				switch char {
					case 124:
						connection = append(connection, 1,0,1,0)
					case 45:
						connection = append(connection, 0,1,0,1)
					case 76:
						connection = append(connection, 1,1,0,0)				
					case 74:
						connection = append(connection, 1,0,0,1)					
					case 55:
						connection = append(connection, 0,0,1,1)										
					case 70:
						connection = append(connection, 0,1,1,0)
					case 83:
						pipeDream.start[0] = int(len(pipeDream.pipes))
						pipeDream.start[1] = int(index)
						connection = append(connection, 1,1,1,1)
				}
 				pipe.connection = connection
				row = append(row, pipe)
			} else {
				row = append(row, pipe{connection: []int{0,0,0,0}})
			}
		}

		pipeDream.pipes = append(pipeDream.pipes, row)
		line, isPrefix, err = r.ReadLine()
	}
	return pipeDream
}

func part2(pipeDream pipeMaze) {

	fmt.Println("********* PART 2 START *********")
	allPositions := make(map[int][]int, 0)
	vectices := make([][]int, 0)
	pipeMap := make([][]string, len(pipeDream.pipes))

	for i := 0; i < len(pipeMap); i++ {
		pipeMap[i] = make([]string, len(pipeDream.pipes[0]))
		for j := 0; j < len(pipeMap[i]); j++ {
			pipeMap[i][j] = "X"
		}
	}

	pos := make([]int,4)
	pos2 := pos

	pipeMap[pipeDream.start[0]][pipeDream.start[1]] = "o"

	if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]+1].connection[3] == 1 {
		pos = []int{pipeDream.start[0],pipeDream.start[1]+1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]+1][pipeDream.start[1]].connection[0] == 1 {
		pos = []int{pipeDream.start[0]+1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]-1].connection[2] == 1 {
		pos = []int{pipeDream.start[0],pipeDream.start[1]-1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]-1][pipeDream.start[1]].connection[1] == 1 { 
		pos = []int{pipeDream.start[0]-1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	}

	if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]-1].connection[1] == 1 {
		pos2 = []int{pipeDream.start[0],pipeDream.start[1]-1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]+1][pipeDream.start[1]].connection[0] == 1 {
		pos2 = []int{pipeDream.start[0]+1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]][pipeDream.start[1]-1].connection[1] == 1 {
		pos2 = []int{pipeDream.start[0],pipeDream.start[1]-1,pipeDream.start[0],pipeDream.start[1]}	
	} else if pipeDream.pipes[pipeDream.start[0]-1][pipeDream.start[1]].connection[3] == 1 { 
		pos2 = []int{pipeDream.start[0]-1,pipeDream.start[1],pipeDream.start[0],pipeDream.start[1]}	
	}

	pipeMap[pos[0]][pos[1]] = "o" 
	pipeMap[pos2[0]][pos2[1]] = "o"

	steps := 1
	for true {
		steps++
		allPositions[pos[0]] = append(allPositions[pos[0]],pos[1] )
		allPositions[pos2[0]] = append(allPositions[pos2[0]],pos2[1] )
		vectices = append(vectices,[]int{pos[0],pos[1]}, []int{pos2[0],pos2[1]})
		pipe := pipeDream.pipes[pos[0]][pos[1]]
		
		pipeMap[pos[0]][pos[1]] = "o"
		pipeMap[pos2[0]][pos2[1]] = "o"
		
		if pipe.connection[0] == 1 && (pos[0]-1 != pos[2] || pos[1] != pos[3]) {			
			pos = []int{pos[0]-1,pos[1],pos[0],pos[1]}			
		} else if pipe.connection[1] == 1 && (pos[0] != pos[2] || pos[1] +1 != pos[3]) {
			pos = []int{pos[0],pos[1]+1,pos[0],pos[1]}
		} else if pipe.connection[2] == 1 && (pos[0]+1 != pos[2] || pos[1] != pos[3]) {
			pos = []int{pos[0]+1,pos[1],pos[0],pos[1]}			
		} else if pipe.connection[3] == 1 && (pos[0] != pos[2] || pos[1] - 1 != pos[3]) {
			pos = []int{pos[0],pos[1]-1,pos[0],pos[1]}		
		} 

		pipe = pipeDream.pipes[pos2[0]][pos2[1]]

		if pipe.connection[3] == 1 && (pos2[0] != pos2[2] || pos2[1] - 1 != pos2[3]) {
			pos2 = []int{pos2[0],pos2[1]-1,pos2[0],pos2[1]}
		} else if pipe.connection[2] == 1 && (pos2[0]+1 != pos2[2] || pos2[1] != pos2[3]) {			
			pos2 = []int{pos2[0]+1,pos2[1],pos2[0],pos2[1]}					
		} else if pipe.connection[1] == 1 && (pos2[0] != pos2[2] || pos2[1] +1 != pos2[3]) {
			pos2 = []int{pos2[0],pos2[1]+1,pos2[0],pos2[1]}
		} else if pipe.connection[0] == 1 && (pos2[0]-1 != pos2[2] || pos2[1] != pos2[3]) {
			pos2 = []int{pos2[0]-1,pos2[1],pos2[0],pos2[1]}	
		} 

		if pos[0] == pos2[0] && pos[1] == pos2[1] {
			pipeMap[pos[0]][pos[1]] = "o"
			allPositions[pos[0]] = append(allPositions[pos[0]], pos[1])
			vectices = append(vectices,[]int{pos[0],pos[1]})
			break
		}		
	}
	
	for i,j := 0,len(pipeMap)-1 ; i < len(pipeMap); i++ {
		for k,l := 0, len(pipeDream.pipes[0])-1; k < len(pipeMap[i]); k++ {
			if pipeMap[i][l] == "X" {
				if (len(allPositions[i]) > 0 && (l < slices.Min(allPositions[i]) || l > slices.Max(allPositions[i]))) ||
					i == len(pipeDream.pipes)-1 || 
					i == 0 || pipeMap[i-1][l] == " " || 
					pipeMap[i][l-1] == " " || pipeMap[i][l+1] == " "  {
					pipeMap[i][l] = " "
				}
			}
			if pipeMap[j][k] == "X" {
				if (len(allPositions[j]) > 0 && (k < slices.Min(allPositions[j]) || k > slices.Max(allPositions[j]))) ||
					j == len(pipeDream.pipes)-1 || 
					j == 0 || pipeMap[j+1][k] == " " || 
					pipeMap[j][k-1] == " " || pipeMap[j][k+1] == " " || pipeMap[j-1][k] == " "  {
					pipeMap[j][k] = " "
				}
			}
			l--
		}
		j--
	}

	count,parts,sum  := 0, 0, 0
	for i, row := range pipeMap {				
		for j, cell := range row {
			if i == pipeDream.start[0] && j == pipeDream.start[1] {
				continue
			}
			if cell == "o" {
				if  pipeDream.pipes[i][j].connection[0] == 1 {
					count++
				}				
				continue
			}
			if cell == " " {
				continue
			}

			if cell == "X" && count%2 != 0 {
				parts++
			}
		}
		fmt.Println(row,parts,i)
		sum += parts
		parts = 0
		count = 0
	}
	fmt.Println(sum)
}
