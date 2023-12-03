package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile := "input.txt"
	//part1(inputFile)
	part2(inputFile)
}

func part1(filename string) {
	parts := make([]string, 0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {

		s := string(line)
		parts = append(parts, s)

		line, isPrefix, err = r.ReadLine()
	}

	sum := int64(0)
	regRow := regexp.MustCompile("[\\d]+")
	regx := regexp.MustCompile("[^(.)]")

	for rowIndex, row := range parts {
		values := ""
		numberIndexes := regRow.FindAllIndex([]byte(row), -1)

		for _, numberIndex := range numberIndexes {
			value, err := strconv.ParseInt(string(row[numberIndex[0]:numberIndex[1]]), 10, 64)

			if numberIndex[0] > 0 {
				numberIndex[0] -= 1
			}

			if err != nil {
				fmt.Println("Error parsing", err)
				return
			}

			if numberIndex[0] == 0 {
				if row[numberIndex[1]] != 46 {
					if err != nil {
						fmt.Println("Error parsing", err)
					}
					values += " " + strconv.FormatInt(value, 10)
					sum += value
					continue
				}
			}

			if numberIndex[1] == len(row) {
				if row[numberIndex[0]] != 46 {
					if err != nil {
						fmt.Println("Error parsing", err)
					}
					values += " " + strconv.FormatInt(value, 10)
					sum += value
					continue
				}
			}

			if numberIndex[0] > 0 && numberIndex[1] < len(row) {
				fmt.Print(string(row[numberIndex[0]]), string(row[numberIndex[1]]))
				if row[numberIndex[0]] != 46 || row[numberIndex[1]] != 46 {
					if err != nil {
						fmt.Println("Error parsing", err)
					}
					values += " " + strconv.FormatInt(value, 10)
					sum += value
					continue
				}
			}

			if numberIndex[1] < len(row)-1 {
				numberIndex[1] += 1
			}

			resultsUp := make([][]int, 0)
			resultsDown := make([][]int, 0)
			if rowIndex < len(parts)-1 {
				resultsDown = regx.FindAllIndex([]byte(parts[rowIndex+1][numberIndex[0]:numberIndex[1]]), -1)
			}

			if rowIndex > 0 {
				resultsUp = regx.FindAllIndex([]byte(parts[rowIndex-1][numberIndex[0]:numberIndex[1]]), -1)
			}
			if len(resultsUp) > 0 || len(resultsDown) > 0 {
				values += " " + strconv.FormatInt(value, 10)
				sum += value
				continue
			}

		}
		fmt.Println(row, values)
	}
	fmt.Println(sum)
}

func part2(filename string) {
	parts := make([]string, 0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {

		s := string(line)
		parts = append(parts, s)

		line, isPrefix, err = r.ReadLine()
	}
	values1 := ""
	values2 := ""
	valuesRow := ""	
	sum := int64(0)
	regRow := regexp.MustCompile("[\\d]+")
	regxGear := regexp.MustCompile("[(*)]")

	for rowIndex, row := range parts {
		
		gearIndexes := regxGear.FindAllIndex([]byte(row), -1)
		valueRow := int64(0)
		if len(gearIndexes) > 0 {
			for _, gearValue := range gearIndexes {
				value1 := int64(0)
				value2 := int64(0)
				numberIndexes := regRow.FindAllIndex([]byte(row), -1)

				for numberIndex, indexValue := range numberIndexes {
					columnIndex := 0
					if numberIndex < len(numberIndexes) -1 {
						columnIndex = numberIndex +1
					}

					if indexValue[1] == gearValue[0] || indexValue[0] == gearValue[1] {
						value1, err = strconv.ParseInt(row[indexValue[0]:indexValue[1]], 10, 64)
					} 
					if numberIndexes[columnIndex][0] == gearValue[1] {
						value2, err = strconv.ParseInt(row[numberIndexes[columnIndex][0]:numberIndexes[columnIndex][1]], 10, 64)
					}					
					
					if value1 > 0 && value2 > 0 && value1 != value2 {						
						fmt.Println(row, value1, value2)
						fmt.Println("")
						sum += value1*value2
						value1 = 0
						value2 = 0
						break
					}

					if (value1 > 0 && value1 == value2) {
						valueRow = value1						
						value2 = 0
						value1 = 0
						break
					}

				}
				if value1 > 0 || value2 > 0 {
					valueRow = value1 + value2					
					value2 = 0
					value1 = 0
				}				

				if rowIndex > 0  && rowIndex < len(parts) -1 {
					numberIndexesDown := regRow.FindAllIndex([]byte(parts[rowIndex+1]), -1)
					numberIndexesUp := regRow.FindAllIndex([]byte(parts[rowIndex-1]), -1)
					for _, indexValue := range numberIndexesUp {
	
						if gearValue[0] <= indexValue[1] && (gearValue[0] >= indexValue[0]-1) {
							if value1 == 0 {
								value1, err = strconv.ParseInt(parts[rowIndex-1][indexValue[0]:indexValue[1]], 10, 64)
							} else {
								value2, err = strconv.ParseInt(parts[rowIndex-1][indexValue[0]:indexValue[1]], 10, 64)
							}							
						}
					}
					
					for _, indexValue := range numberIndexesDown {											
	
						if gearValue[0] <= indexValue[1] && (gearValue[0] >= indexValue[0]-1) {
							if value1 == 0 {
								value1, err = strconv.ParseInt(parts[rowIndex+1][indexValue[0]:indexValue[1]], 10, 64)	
							} else {
								value2, err = strconv.ParseInt(parts[rowIndex+1][indexValue[0]:indexValue[1]], 10, 64)
							}
						}
					}
					if value1 > 0 && value2 > 0 {
						values1 += " " + strconv.FormatInt(value1, 10)
						values2 += " " + strconv.FormatInt(value2, 10)
						sum += value1*value2
					} else if (value1 > 0 && valueRow > 0) || (value2 > 0 && valueRow > 0) {
						valuesRow += " " + strconv.FormatInt(valueRow, 10)
						if value1 > 0 {
							values1 += " " + strconv.FormatInt(value1, 10)
						}
						
						if value2 > 0 {
							values2 += " " + strconv.FormatInt(value2, 10)
						}

						sum += (value1*valueRow) + (value2*valueRow)
					} 
				}
				value1 = 0
				value2 = 0
				valueRow= 0
			}
			if values1 != "" && values2 != "" {
				fmt.Println(parts[rowIndex-1], values1)
				fmt.Println(row, valuesRow)
				fmt.Println(parts[rowIndex+1], values2)
				fmt.Println("")
			}
			values1 = ""
			values2 = ""
			valuesRow = ""
		}
	}
	fmt.Println(sum)
}
