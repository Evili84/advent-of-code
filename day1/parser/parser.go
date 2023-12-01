package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(input string) int {
	numbers := ReadLine(input)

	sum := 0

	for index := range numbers {
		sum += numbers[index]
	}
	return sum
}

func ReadLine(filename string) []int {
	
	numberStrings := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	result := make([]int,0)
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer f.Close()
    r := bufio.NewReaderSize(f, 4*1024)
    line, isPrefix, err := r.ReadLine()
    for err == nil && !isPrefix {
        s := string(line)
		//fmt.Println(s)
        re := regexp.MustCompile("[^0-9]+")
		numbers := re.ReplaceAllString(s,"")
		
		first := numbers[0:1]
		lastIndex := strings.Index(s,first)
		for key, value := range numberStrings {
			wordIndex := strings.Index(s, key)
			
			if wordIndex > -1 && wordIndex < lastIndex {
				first = strconv.Itoa(value)
				if(lastIndex > -1 && wordIndex < lastIndex) {
					lastIndex = wordIndex
					first = strconv.Itoa(value)
				}
			}
		}

		last := numbers[len(numbers)-1:]

		lastIndex = strings.LastIndex(s, last)
		for key, value := range numberStrings {
			wordIndex := strings.LastIndex(s, key)

			if wordIndex > -1 && wordIndex > lastIndex {
				last = strconv.Itoa(value)
				lastIndex = wordIndex			
			}
		}		

		i, parseErr := strconv.ParseInt(first+last, 10, 64 )

		if parseErr != nil {
			return nil
		}

		fmt.Println(s, first, last, i)

		result = append(result, int(i))

        line, isPrefix, err = r.ReadLine()
    }
    if isPrefix {
        fmt.Println("buffer size to small")
        return nil
    }
    if err != io.EOF {
        fmt.Println(err)
        return nil
    }

	return result
}