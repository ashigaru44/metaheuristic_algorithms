package problem

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"os/exec"
)

type Problem struct {
	name       string
	dim        int
	nodes      [][2]int
	Adj_matrix [][]int
	Path       []int
	// Edge_weight_type string
	// Type             string
	// Comment          string
}

func (p Problem) PrintProblem() {
	fmt.Printf("\nName: %s", p.name)
	fmt.Printf("\nDimension: %v\nNodes:\n", p.dim)
	for i, val := range p.nodes {
		fmt.Printf("%v.\t%v\n", i, val)
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
	for i := 1; i < len(solution.path); i++ {
		dist += p.Adj_matrix[solution.path[i]][solution.path[i-1]]
	}
	dist += p.Adj_matrix[solution.path[len(solution.path)]][solution.path[0]]
	return dist
}

func (p Problem) EvaluateSolution2(solution *[]int) int {
	dist := 0
	for i := 1; i < len(*solution); i++ {
		dist += p.Adj_matrix[(*solution)[i]][(*solution)[i-1]]
	}
	dist += p.Adj_matrix[(*solution)[len(*solution)-1]][(*solution)[0]]
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
			} else if strings.Contains(text, "DIMENSION") {
				splitted := strings.Split(text, " ")
				problem.dim, err = strconv.Atoi(splitted[len(splitted)-1])
				check(err)
			}
		} else {
			splitted := strings.Split(text, " ")
			x, err := strconv.ParseFloat(splitted[1], 32)
			check(err)
			y, err := strconv.ParseFloat(splitted[2], 32)
			check(err)
			problem.nodes = append(problem.nodes, [2]int{int(x), int(y)})
		}
	}

	problem.PrintProblem()
	problem.Adj_matrix = *problem.adjacency_matrix()
	// for i := range problem.adj_matrix {
	// fmt.Println(problem.adj_matrix[i])
	// }
	file.Close()
	return &problem
}

func (p Problem) adjacency_matrix() *[][]int {
	adj_matrix := make([][]int, p.dim)
	for i := range adj_matrix {
		adj_matrix[i] = make([]int, p.dim)
	}

	for i := range adj_matrix {
		for j := range adj_matrix[i] {
			xd := p.nodes[i][0] - p.nodes[j][0]
			yd := p.nodes[i][1] - p.nodes[j][1]
			adj_matrix[i][j] = int(math.Sqrt(float64(xd*xd + yd*yd)))
		}
	}
	return &adj_matrix
}

func (p Problem) GetDistance(i1 int, i2 int) int {
	return p.Adj_matrix[i1][i2]
}

func (p Problem) SaveProblemToFile() string {
	file_path := "\\data\\problem_data.txt"
	f, err := os.Create("." + file_path)
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Print("len: ")
	fmt.Println(len(p.nodes))
	for i := range p.nodes {
		_, err := w.WriteString(fmt.Sprintln(p.nodes[i]))
		check(err)
	}
	if len(p.Path) > 0 {
		w.WriteString(fmt.Sprintln())
		for i := range p.Path {
			_, err := w.WriteString(fmt.Sprintln(p.Path[i]))
			check(err)
		}
	}
	w.Flush()
	ex, err := os.Executable()
	check(err)
	return filepath.Dir(ex) + file_path
	// return "C:\\Users\\mielcare\\Documents\\repos\\metaheuristic_algorithms\\data\\problem_data.txt"
}

func ShowGraph(p *Problem, path *[]int) {
	p.Path = *path
	saved_problem_path := p.SaveProblemToFile()
	fmt.Println("Path:", *path)
	fmt.Println("Distance = ", p.EvaluateSolution2(path))
	err := exec.Command("python", "./visualize.py", saved_problem_path).Run()
	if err != nil {
		panic(err)
	}
}
