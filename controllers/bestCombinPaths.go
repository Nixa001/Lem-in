package controllers

import (
	"lem-in/models"
	"math"
)

func BestCombPaths(ants models.Ants, pathGroups models.Paths) models.Paths {
	minLevel := math.MaxInt32
	var bestCombination [][]string

	for _, combination := range pathGroups.SortedCombinations {
		spaceInPath := 0
		totalPathLength := 0
		levelOfAnts := 0
		longestPath := len(combination[len(combination)-1])
		for i := 1; i < len(combination); i++ {
			spaceInPath = spaceInPath + longestPath - len(combination[i])
			totalPathLength = totalPathLength + len(combination[i])
		}
		levelOfAnts = (ants.NbrAnts-spaceInPath)/len(combination) + longestPath

		if levelOfAnts < minLevel {
			minLevel = levelOfAnts
			bestCombination = combination
		}
	}

	return models.Paths{BestCombination: bestCombination}
}

func SortCombPaths(combinations models.Paths) models.Paths {
	pathGroups := make(map[int][][]string)
	for _, combination := range combinations.ValidPaths {
		category := len(combination)
		currentCombLength := combLength(combination)
		if _, ok := pathGroups[category]; ok {
			valueInMap := pathGroups[category]
			if currentCombLength < combLength(valueInMap) {
				pathGroups[category] = combination
			}
		} else {
			pathGroups[category] = combination
		}
	}
	return models.Paths{SortedCombinations: pathGroups}
}

func combLength(combination [][]string) int {
	length := 0
	for _, path := range combination {
		length = length + len(path)
	}
	return length
}
