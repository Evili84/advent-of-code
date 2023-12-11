package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type visited struct {
	start int
	end int
}

type galaxyMap struct {
	imageData [][]string
	galaxies [][]int
	expansionRows []int
	expansionCols []int
}

func main() {
	inputFile := "input.txt"
	part1(inputFile, 2)
	part1(inputFile, 1000000)
}

func part1(inputFile string, expansionRate int) {
	_, mapData := makeMapData(inputFile)
	mapData = expandMap(mapData)
	visitedData := make([]visited,0)
	point := 0
	sum := 0
	max := len(mapData.galaxies)-1
	for max > point {		
		for i := point; i < max; i++ {
			v := visited{start: i, end: max}
			if !slices.Contains(visitedData, v) {
				visitedData = append(visitedData, v)
			} else {
				continue
			}
			cols := 0
			rows := (mapData.galaxies[max][0]-mapData.galaxies[i][0])
			for _, v := range mapData.expansionRows {
				if v < mapData.galaxies[max][0] && v > mapData.galaxies[i][0] {
					rows+= expansionRate-1
				}
			}
		

			if mapData.galaxies[i][1] > mapData.galaxies[max][1] {
				cols = (mapData.galaxies[i][1]-mapData.galaxies[max][1])
				for _, v := range mapData.expansionCols {
					if v < mapData.galaxies[i][1] && v > mapData.galaxies[max][1] {
						cols+= expansionRate-1
					}
				}
			} else {
				cols = (mapData.galaxies[max][1]-mapData.galaxies[i][1])
				for _, v := range mapData.expansionCols {
					if v < mapData.galaxies[max][1] && v > mapData.galaxies[i][1] {
						cols+= expansionRate-1
					}
				}
			}

			sum += rows + cols
		}		
		max--
	}

	fmt.Println(sum)
	
}

func makeMapData(inputFile string) (error, galaxyMap) {

	mapData := galaxyMap{imageData: make([][]string, 0)}
	mapData.galaxies = make([][]int,0)
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return err, mapData
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		points := make([]string,0)
		for _, point := range s {
			points = append(points, string(point))
		}
		mapData.imageData = append(mapData.imageData, points)
		line, isPrefix, err = r.ReadLine()
	}
	return err, mapData
}

func expandMap(mapData galaxyMap) galaxyMap {

	mapData.expansionCols = make([]int,0)
	mapData.expansionRows = make([]int,0)
	
	for i := 0; i < len(mapData.imageData); i++ {
		found := false
		for j := 0; j < len(mapData.imageData[i]); j++ {
			if mapData.imageData[i][j] == "#" {
				found = true
				break
			}
		}		
		if !found {
			mapData.expansionRows = append(mapData.expansionRows, i)
		}
	}
	
	for i := 0; i < len(mapData.imageData[0]); i++ {
		found := false
		for j := 0; j < len(mapData.imageData); j++ {
			if mapData.imageData[j][i] == "#" {
				found =true;
				break
			}
		}
	
		if !found {
			mapData.expansionCols = append(mapData.expansionCols, i)
		}
	}

	for i, d := range mapData.imageData {
		for j, v := range d {
			if v == "#" {
				mapData.galaxies = append(mapData.galaxies, []int{i,j})
			}
		}
	}

	return mapData
}
