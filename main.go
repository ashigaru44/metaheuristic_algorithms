package main

import (
	//"math/rand"
	"fmt"
	"math"
	"meta-heur/tsp/problem"
	"time"
	//"time"
	//"sort"
)

//var r rand.Rand

type Result struct {
	path     *[]int
	distance int
	duration int64
}

type Results struct {
	name    string
	results []Result

	best_dist  int
	avg_dist   int
	worst_dist int

	best_time  int64
	avg_time   int64
	worst_time int64
}

type TestAlgorithm struct {
	p1           *problem.Problem
	duration     int64
	k            int
	initial_path *[]int
}

type Algorithm int

const (
	Nearest    Algorithm = 0
	RandomTime           = 1
	Opt2                 = 2
	RandomK              = 3
)

func printResults(res *Results) {

}

func calculateDistances(p1 *problem.Problem, res *Results) {
	for i := range res.results {
		res.results[i].distance = p1.EvaluateSolution2(res.results[i].path)
	}
}

func calculateResults(p1 *problem.Problem, res *Results) {
	best_dist := math.MaxInt32
	worst_dist := 0
	avg_dist := 0

	best_time := int64(math.MaxInt64)
	worst_time := int64(0)
	avg_time := int64(0)

	calculateDistances(p1, res)

	// CALCULATE DISTANCE VALUES

	for i := range res.results {
		new_dist := res.results[i].distance
		avg_dist += new_dist
		if worst_dist < new_dist {
			worst_dist = new_dist
		}
		if best_dist > new_dist {
			best_dist = new_dist
		}
	}

	res.best_dist = best_dist
	res.worst_dist = worst_dist
	res.avg_dist = avg_dist / len(res.results)

	// CALCULATE DURATION VALUES

	for i := range res.results {
		new_time := res.results[i].duration
		avg_time += new_time
		if worst_time < new_time {
			worst_time = new_time
		}
		if best_time > new_time {
			best_time = new_time
		}
	}

	res.best_time = best_time
	res.worst_time = worst_time
	res.avg_time = avg_time / int64(len(res.results))

}

func test_Algorithm(alg Algorithm, ta TestAlgorithm, k int) Results {
	var results = Results{results: make([]Result, k)}
	for i := 0; i < k; i++ {
		switch alg {
		case Nearest:
			results.results[i] = testNearest(ta.p1)
		case RandomTime:
			results.results[i] = testRandomTime(ta.p1, ta.duration)
		case Opt2:
			results.results[i] = test_2opt(ta.p1, ta.initial_path)
		case RandomK:
			results.results[i] = testRandomK(ta.p1, ta.k)
		}

	}
	calculateResults(ta.p1, &results)
	return results
}

// func test_Algorithm(algorithm func(*problem.Problem, int64, *[]int) (Result), p1 *problem.Problem, duration int64, initial_path *[]int) (Results) {
// 	k := 100
// 	var results = Results{results: make([]Result, k) }
// 	for i:=0; i<k; i++ {
// 		results.results[i] = algorithm(p1, duration, initial_path)
// 	}
// 	return results
// }

// func test_Algorithm3(algorithm func(params... interface{}) (Result), params... interface{}) (Results) {
// 	k := 100
// 	var results = Results{results: make([]Result, k) }
// 	for i:=0; i<k; i++ {
// 		results.results[i] = algorithm(params)
// 	}
// 	return results
// }

// func test_Algorithm2(algorithm func(params... interface{}) (Result), params... interface{}) (Results) {
// 	k := 100
// 	var results = Results{results: make([]Result, k) }
// 	for i:=0; i<k; i++ {
// 		results.results[i] = algorithm(params)
// 	}
// 	return results
// }

// func test_Algorithm_Nearest(params... interface{}) (Result) {
// 	var p1 *problem.Problem = (*problem.Problem)(params[0])
// 	start := time.Now()
// 	path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
// 	duration := time.Since(start).Nanoseconds()
// 	var elapsed float64 = float64(duration) / 1000000
//     fmt.Println("test Nearest took ", elapsed, " ms")
// 	fmt.Println("Distance = ", p1.EvaluateSolution2(path))
// 	fmt.Println("")
// 	dist := p1.EvaluateSolution2(path)
// 	return Result{path, dist, duration}
// }

func testNearest(p1 *problem.Problem) Result {
	start := time.Now()
	path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	duration := time.Since(start).Nanoseconds()
	// var elapsed float64 = float64(duration) / 1000000
	// fmt.Println("test Nearest took ", elapsed, " ms")
	// fmt.Println("Distance = ", p1.EvaluateSolution2(path))
	// fmt.Println("")
	dist := p1.EvaluateSolution2(path)
	return Result{path, dist, duration}
}

func testRandom(p1 *problem.Problem) {
	start := time.Now()
	path, _ := problem.Random_k(*p1, 2000)
	var elapsed float64 = float64(time.Since(start).Nanoseconds()) / 1000000
	fmt.Println("test Random took ", elapsed, " ms")
	fmt.Println("Distance = ", p1.EvaluateSolution2(path))
	fmt.Println("")
}

func testRandomTime(p1 *problem.Problem, old_duration int64) Result {
	start := time.Now()
	path, dist := problem.Random_time(*p1, old_duration)
	duration := time.Since(start).Nanoseconds()
	//dist := p1.EvaluateSolution2(path)
	var elapsed float64 = float64(duration) / 1000000
	fmt.Println("test RandomTime took ", elapsed, " ms")
	fmt.Println("Distance = ", p1.EvaluateSolution2(path))
	fmt.Println("")
	return Result{path, dist, duration}
}

func testRandomK(p1 *problem.Problem, k int) Result {
	start := time.Now()
	path, dist := problem.Random_k(*p1, k)
	duration := time.Since(start).Nanoseconds()
	//dist := p1.EvaluateSolution2(path)
	// var elapsed float64 = float64(duration) / 1000000
	// fmt.Println("test RandomTime took ", elapsed, " ms")
	// fmt.Println("Distance = ", p1.EvaluateSolution2(path))
	// fmt.Println("")
	return Result{path, dist, duration}
}

func test_2opt(p1 *problem.Problem, initial_path *[]int) Result {
	start := time.Now()
	path, dist := problem.Opt2_PickBest(*p1, p1.Adj_matrix, initial_path)
	duration := time.Since(start).Nanoseconds()
	// var elapsed float64 = float64(duration) / 1000000
	//dist := p1.EvaluateSolution2(path)
	// fmt.Println("\ntest 2Opt took ", elapsed, " ms")
	// fmt.Println("Distance = ", p1.EvaluateSolution2(path))
	// fmt.Println("")
	return Result{path, dist, duration}

}

func rest(p1 *problem.Problem) {
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

func compareAlgorithms(p *problem.Problem, alg1 Algorithm, alg2 Algorithm, repeats int) {
	alg1_distances := make([]int, repeats)
	alg2_distances := make([]int, repeats)
	alg3_distances := make([]int, repeats)

	for i := 0; i < repeats; i++ {
		initPath, _ := problem.NearestNeighbourAllPoints(*p, p.Adj_matrix)
		// initPath, _ := problem.Random(*p)
		ta := TestAlgorithm{p, int64(0), 100, initPath}
		fmt.Println()
		alg1_result := test_Algorithm(alg1, ta, 1)
		alg2_result := test_Algorithm(alg2, ta, 1)
		alg3_result := test_Algorithm(RandomK, ta, 1)
		// alg1_dist := alg1_result.avg_dist
		// alg2_dist := alg2_result.avg_dist
		alg1_distances[i] = alg1_result.results[0].distance
		alg2_distances[i] = alg2_result.results[0].distance
		alg3_distances[i] = alg3_result.results[0].distance

		// fmt.Println("Results_1: ")
		// fmt.Println(alg1_result.results)
		// fmt.Println("Results_2: ")
		// fmt.Println(alg2_result.results)
	}
	for i := range alg1_distances {
		fmt.Println(alg1_distances[i])
	}

	for i := range alg2_distances {
		fmt.Println(alg2_distances[i])
	}

	for i := range alg3_distances {
		fmt.Println(alg3_distances[i])
	}

	// initPath, _ := problem.Random(*p)
	// ta := TestAlgorithm{p, int64(0), 100, initPath}
	// fmt.Println()
	// alg1_result := test_Algorithm(alg1, ta, repeats)
	// alg2_result := test_Algorithm(alg2, ta, repeats)
	// alg3_result := test_Algorithm(RandomK, ta, repeats)
	// // alg1_dist := alg1_result.avg_dist
	// // alg2_dist := alg2_result.avg_dist
	// for i := range alg1_result.results {
	// 	fmt.Println(alg1_result.results[i].distance)
	// }
	// for i := range alg2_result.results {
	// 	fmt.Println(alg2_result.results[i].distance)
	// }
	// for i := range alg3_result.results {
	// 	fmt.Println(alg3_result.results[i].distance)
	// }

}

func main() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	problem_path := "./berlin52.tsp"
	p1 := problem.InitProblem(problem_path)
	// path, _ := problem.NearestNeighbourAllPoints(*p1, p1.Adj_matrix)
	// path, _ := problem.Random(*p1)

	compareAlgorithms(p1, Opt2, Nearest, 40)
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
