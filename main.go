package main

import (
	"fmt"
	"lem-in/controllers"
	"lem-in/readfile"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: missing filename")
		fmt.Println("USAGE: go run main.go example00.txt")
		return
	}

	filePath := os.Args[1]
	data := readfile.ReadFile(filePath)
	graph, ants := readfile.ParseData(data)

	paths := controllers.FindPaths(graph)
	ValidPaths := controllers.ValidPaths(paths)
	sortedCombPaths := controllers.SortCombPaths(ValidPaths)
	bestCombPaths := controllers.BestCombPaths(ants, sortedCombPaths)
	sendAnts := controllers.SendAnts(ants, data, bestCombPaths)

	turns := 0
	for _, v := range sendAnts {
		fmt.Println(v)
		turns++
	}
	fmt.Println()
	fmt.Printf("Turns number: %v\n", turns)
}
