package main

import (
	//"math/rand"

	"meta-heur/tsp/problem"
	//"time"
	//"sort"
)

//var r rand.Rand

func main() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	problem_path := "../data/berlin52.tsp"
	p1 := problem.InitProblem(problem_path)
	// path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	path, _ := problem.Random(*p1)
	problem.Tabu_search(*p1, path, 10, 3, 5)
	// utils.CompareAlgorithms(p1, utils.Opt2, utils.Nearest, 40)

	// ta := TestAlgorithm{p1, int64(0), 50, path}
	// test_Algorithm(Opt2, ta, 1)

	// result_dist := res.avg_dist
	// fmt.Println(result_dist)
	// rest(p1)
	// ta.duration = res.avg_time
	// res = test_Algorithm(RandomTime, ta, 1)
	// res = test_Algorithm(Opt2, ta, 1)
	//test_Algorithm(RandomTime, ta, 100)
	//test_Algorithm(Opt2, ta, 100)

	// p1.PrintProblem()
	// pr := problem.GenerateProblem(10)
	// fmt.Println(saved_problem_path)

	// res := testNearest(p1)
	// testRandomTime(p1, res.duration)
	// test_2opt(p1, res.duration, res.path)

	// pr.PrintProblem()
}
