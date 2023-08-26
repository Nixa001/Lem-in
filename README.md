<h1 align="center">LEM-IN</h1>

<table align="center" border="0">
  <tr>
    <td><a href="https://imgbb.com/"><img src="https://i.ibb.co/tL6SpfD/zone01.png" alt="zone01" border="0" width="100px"></a></td>
    <td><a href="https://imgbb.com/"><img src="https://i.ibb.co/W2vM8Ws/golang.png" alt="golang" border="0" width="120px"></a></td>
   </tr>
</table>

The goal of lem-in is to navigate a group of ants through a maze from the start room to the end room in the least number of turns, while avoiding collisions and following specific rules.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Algorithm](#algorithm)
- [Contributing](#contributing)

## Getting Started

### Prerequisites

Before running Lem-In, make sure you have Go installed on your system. You can download and install it from the [official Go website](https://golang.org/).

### Installation

1. Clone the repository:

   ```bash
   git clone https://learn.zone01dakar.sn/git/nifaye/lem-in
   ```

2. Navigate to the project directory:

   ```bash
   cd lem-in
   ```

3. Build the executable:

   ```go
   go build
   ```

## Usage

- Run Lem-In Go with the following command:

  ```bash
  ./lem-in <input.txt>
  ```

- Replace input.txt with your own input file containing the Lem-In data.

## Algorithm

Lem-In Go uses a combination of graph traversal algorithms and heuristics to find the optimal path for the ants. The main steps of the algorithm include:

1.  Parsing the input data to create a graph representation of the maze.
2.  Applying graph traversal algorithms (e.g., [Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)) to find the shortest paths from the start room to the end room.
[See more](https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893)
3.  Distributing ants along the paths while avoiding collisions.

## Contributing

Contributions to Lem-In are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

## Autors

- [@igueye](https://learn.zone01dakar.sn/git/igueye)
- [@nifaye](https://learn.zone01dakar.sn/git/nifaye)
