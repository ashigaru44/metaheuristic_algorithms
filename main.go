package main

import (
	//"math/rand"
	"fmt"
	"io/ioutil"
	"log"
	"meta-heur/tsp/problem"
	//"sort"
)

const FOLDER_PATH = "C:\\Users\\ashig\\Documents\\GoLang\\meta-heur\\input_data\\"

//var r rand.Rand

func main() {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	files, err := ioutil.ReadDir(FOLDER_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var problems_files []string
	for _, f := range files {
		problems_files = append(problems_files, f.Name())
	}
	for _, file_path := range problems_files {
		problem_path := FOLDER_PATH + file_path
		// problem_path := "./input_data/berlin52.tsp"
		fmt.Println(problem_path)
		p1 := problem.InitProblem(problem_path)
		path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
		fmt.Println("distance = ", p1.EvaluateSolution2(path))
		// problem.ShowGraph(p1, path)
	}
	// problem_path := FOLDER_PATH + problems_files[0]
	// // problem_path := "./input_data/berlin52.tsp"
	// fmt.Println(problem_path)
	// p1 := problem.InitProblem(problem_path)
	// path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	// // path, _ := problem.Random(*p1)
	// // var distance int
	// // path, distance = problem.Tabu_search_concurrent(*p1, ":", 200, 0.985, 30, 8, 1)

	// _, distance_2OPT := problem.Opt2(*p1, p1.Adj_matrix, path)

	// utils.Test_run_GA_2()
	// // utils.Random_sampling()

	// //path = problem.Genetic_generate_solution(*p1, 0.7, 0.2, 1000, 10000, 10, 0.02, 0)
	// // path = problem.GA_Islands_generate_solution(*p1, 0.75, 0.5, 1000, 7000, 4, 10, 0.05, 20, 0, 2)
	// // path, distance = problem.Tabu_search(*p1, path, 1000, 0.985, 30)
	// fmt.Println("distance 2OPT = ", distance_2OPT)
	// fmt.Println("distance = ", p1.EvaluateSolution2(path))
	// problem.ShowGraph(p1, path)
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
	// pr := problem.GenerateProblem(50, 1, 20, false)
	// fmt.Println(saved_problem_path)

	// res := testNearest(p1)
	// testRandomTime(p1, res.duration)
	// test_2opt(p1, res.duration, res.path)

	// pr.PrintProblem()
}
