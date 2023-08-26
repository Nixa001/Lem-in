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
	SortComb map[int][][]string
	BestComb    [][]string
}
