package controllers

import (
	"fmt"
	"lem-in/models"
	"os"
)

func FindPaths(lemin models.LemInData) models.Paths {
	_, ok := lemin.Links[lemin.Start]
	if !ok || lemin.Start == "" {
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(1)
	}
	_, ok = lemin.Links[lemin.End]
	if !ok || lemin.End == "" {
		fmt.Println("ERROR: invalid data format, no end room found")
		os.Exit(1)
	}

	visited := make(map[string]bool)
	var path []string
	var paths [][]string

	find(lemin.Start, lemin.End, lemin.Links, visited, path, &paths)
	return models.Paths{AllPaths: paths}
}

func find(start, end string, edges map[string][]string, visited map[string]bool, path []string, paths *[][]string) {
	visited[start] = true
	path = append(path, start)

	if start == end {
		temp := make([]string, len(path[1:]))
		copy(temp, path[1:])
		*paths = append(*paths, temp)
	} else {
		for _, vertex := range edges[start] {
			if !visited[vertex] {
				find(vertex, end, edges, visited, path, paths)
			}
		}
	}
	visited[start] = false
}

