package main

import (
	//"math/rand"
	"fmt"
	"meta-heur/tsp/problem"

	//"time"
	//"sort"
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

	path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	problem.ShowGraph(p1, path)
	// END_OF_PATH_VISUALIZATION


	//path, _ = problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	//sort.Ints(*path)
	//fmt.Println("Distance = ", distance)
	fmt.Println("Path:", *path)
	fmt.Println("Distance = ", p1.EvaluateSolution2(path))


	path, _ = problem.Opt2(*p1, p1.Adj_matrix, path)
	problem.ShowGraph(p1, path)

	path, _ = problem.Opt2_PickBest(*p1, p1.Adj_matrix, path)
	problem.ShowGraph(p1, path)


	problem.Random_k(*p1, 100)
}
