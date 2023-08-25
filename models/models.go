package models

type LemInData struct {
	Start    string
	End      string
	Links    map[string][]string
	Rooms []string
}

type Ants struct {
	NbrAnts int
}

type Paths struct {
	AllPaths           [][]string
	ValidPaths         [][][]string
	SortedCombinations map[int][][]string
	BestCombination    [][]string
}
