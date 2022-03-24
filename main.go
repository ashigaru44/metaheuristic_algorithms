package main

import (
	//"math/rand"
	"fmt"
	"meta-heur/tsp/problem"
	"os/exec"
	//"time"
	"sort"
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
  // PATH_VISUALIZATION
	path, distance := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	p1.Path = *path
	saved_problem_path := p1.SaveProblemToFile()
	fmt.Println("Path:", *path)
	fmt.Println("Distance = ", distance)
	err := exec.Command("python", "./visualize.py", saved_problem_path).Run()
	if err != nil {
		panic(err)
	}
  // END_OF_PATH_VISUALIZATION
  path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	sort.Ints(*path)
	//fmt.Println("Distance = ", distance)
	fmt.Println("Path:", *path)
	fmt.Println("Distance = ", p1.EvaluateSolution2(path))

	problem.Opt2(*p1, p1.Adj_matrix, path)
}
