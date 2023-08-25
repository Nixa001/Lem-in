package controllers

import (
	"fmt"
	"lem-in/models"
	"strconv"
)

func SendAnts(ants models.Ants, data []string, bestCombination models.Paths) []string {
	inputData(data)
	queue := assignAnts(ants, bestCombination)
	determineOrder(queue)
	return calculateSteps(queue, bestCombination)
}

func determineOrder(queue [][]string) []int {
	order := []int{}
	longest := len(queue[0])
	for i := 0; i < len(queue); i++ {
		if len(queue[i]) > longest {
			longest = len(queue[i])
		}
	}
	for j := 0; j < longest; j++ {
		for i := 0; i < len(queue); i++ {
			if j < len(queue[i]) {
				x, _ := strconv.Atoi(queue[i][j])
				order = append(order, x)
			}
		}
	}
	return order
}

func assignAnts(ants models.Ants, bestCombination models.Paths) [][]string {
	queue := make([][]string, len(bestCombination.BestCombination))

	for i := 1; i <= ants.NbrAnts; i++ {
		ant := strconv.Itoa(i)
		minSteps := len(bestCombination.BestCombination[0]) + len(queue[0])
		minIndex := 0

		for j, path := range bestCombination.BestCombination {
			steps := len(path) + len(queue[j])
			if steps < minSteps {
				minSteps = steps
				minIndex = j
			}
		}
		queue[minIndex] = append(queue[minIndex], ant)
	}
	return queue
}

func inputData(data []string) {
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println()
}

func calculateSteps(queue [][]string, bestCombination models.Paths) []string {
	container := make([][][]string, len(queue))
	for i, path := range queue {
		for _, ant := range path {
			adder := []string{}
			for _, vertex := range bestCombination.BestCombination[i] {
				str := "L" + ant + "-" + vertex
				adder = append(adder, str)
			}
			container[i] = append(container[i], adder)
		}
	}

	finalMoves := []string{}
	for _, paths := range container {
		for j, moves := range paths {
			for k, vertex := range moves {
				if j+k > len(finalMoves)-1 {
					finalMoves = append(finalMoves, vertex+" ")
				} else {
					finalMoves[j+k] = finalMoves[j+k] + vertex + " "
				}
			}
		}
	}
	return finalMoves
}
