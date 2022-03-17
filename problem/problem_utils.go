package problem

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
	name       string
	dim        int
	nodes      [][2]int
	adj_matrix [][]int
	// Edge_weight_type string
	// Type             string
	// Comment          string
}

func (p Problem) PrintProblem() {
	fmt.Printf("\nName: %s", p.name)
	fmt.Printf("\nDimension: %v", p.dim)
	for i, val := range p.nodes {
		fmt.Printf("\n%v.\t%v", i, val)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GenerateProblem(size int, min int, max int) *Problem {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var dim int
	if min == 0 && max == 0 {
		dim = r.Int() % 50
	} else {
		dim = r.Int()%(max-min) + min
	}
	nodes := make([][2]int, dim)

	for i := 0; i < dim; i++ {
		nodes[i][0] = r.Int()%(2*size) - size
		nodes[i][1] = r.Int()%(2*size) - size
	}

	fmt.Println(dim)

	var problem = Problem{name: "generated", dim: dim, nodes: nodes}
	return &problem
}

func (p Problem) EvaluateSolution(solution *Solution) int {
	dist := 0
	for i := 1; i <= len(solution.path); i++ {
		dist += p.adj_matrix[solution.path[i]][solution.path[i-1]]
	}
	return dist
}

func InitProblem(path string) *Problem {
	file, err := os.Open(path)
	check(err)
	var problem = Problem{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")

		_, err := strconv.Atoi(splitted[0])
		if err != nil {
			if strings.Contains(text, "NAME") {
				splitted := strings.Split(text, " ")
				problem.name = splitted[len(splitted)-1]
				fmt.Println("name: " + problem.name)
			} else if strings.Contains(text, "DIMENSION") {
				splitted := strings.Split(text, " ")
				problem.dim, err = strconv.Atoi(splitted[len(splitted)-1])
				check(err)
				fmt.Println(problem.dim)
			}
		} else {
			splitted := strings.Split(text, " ")
			fmt.Println(splitted)
			x, err := strconv.ParseFloat(splitted[1], 32)
			check(err)
			y, err := strconv.ParseFloat(splitted[2], 32)
			check(err)
			problem.nodes = append(problem.nodes, [2]int{int(x), int(y)})
		}
	}


	file.Close()
	return &problem
}
