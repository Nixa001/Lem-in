package readfile

import (
	"bufio"
	"fmt"
	"lem-in/models"
	"os"
	"strconv"
	"strings"
)

func ParseData(data []string) (models.LemInData, models.Ants) {
	var lemIn models.LemInData
	var ant models.Ants

	lemIn.Links = make(map[string][]string)

	nbrAnts, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Printf("Error converting number of ants to integer: %v", err)
		os.Exit(1)
	}
	ant.NbrAnts = nbrAnts
	if ant.NbrAnts < 1 { //|| ant.NbrAnts > 1000_000_000
		fmt.Printf("ERROR: invalid data format, invalid number of Ants\n")
		os.Exit(1)
	}

	for i := 1; i < len(data); i++ {
		if data[i] == "##start" {
			fields := strings.Fields(data[i+1])
			start := fields[0]
			lemIn.Start = string(start)
		} else if data[i] == "##end" {
			fields := strings.Fields(data[i+1])
			end := fields[0]
			lemIn.End = string(end)
		} else if strings.Contains(data[i], "-") {
			link := strings.Split(data[i], "-")
			from := link[0]
			to := link[1]
			if from == "" || to == "" {
				fmt.Printf("Invalid link format")
				os.Exit(1)
			}
			lemIn.Links[from] = append(lemIn.Links[from], to)
			lemIn.Links[to] = append(lemIn.Links[to], from)
		} else if strings.Contains(data[i], " ") {
			room := strings.Split(data[i], " ")
			v := room[0]
			if strings.HasPrefix(v, "#") {
				fmt.Printf("Invalid room name: %s", v)
				os.Exit(1)
			}
			lemIn.Rooms = append(lemIn.Rooms, v)
		}
	}
	return lemIn, ant
}

func ReadFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	if len(fileLines) == 0 {
		fmt.Println("The file is empty.")
		os.Exit(1)
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
		os.Exit(1)
	}
	return fileLines
}
