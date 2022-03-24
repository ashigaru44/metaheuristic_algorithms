package main

import (
	//"math/rand"
	"fmt"
	"meta-heur/tsp/problem"
	"os/exec"
	//"time"
)

//var r rand.Rand

func main() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	problem_path := "./berlin52.tsp"
	p1 := problem.InitProblem(problem_path)
	// p1.PrintProblem()
	// pr := problem.GenerateProblem(10)
	// fmt.Println(saved_problem_path)

	// pr.PrintProblem()
	path, distance := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	p1.Path = *path
	saved_problem_path := p1.SaveProblemToFile()
	fmt.Println("Path:", *path)
	fmt.Println("Distance = ", distance)
	err := exec.Command("python", "./visualize.py", saved_problem_path).Run()
	if err != nil {
		panic(err)
	}
}
