package utils

import (
	//"math/rand"
	"meta-heur/tsp/problem"
)

const TEST_REPEAT = 5

func test_run_GA() {
	problems := generate_testing_problems()

	for _, problem := range problems {
		crossing_testing(problem)
	}

	for _, problem := range problems {
		mutation_testing(problem)
	}

	for _, problem := range problems {
		population_size_testing(problem)
	}

	for _, problem := range problems {
		tournament_size_testing(problem)
	}

	for _, problem := range problems {
		elitism_size_testing(problem)
	}
}

func generate_testing_problems() [16]problem.Problem {
	var problems [16]problem.Problem
	problem1_path := "./input_data/berlin52.tsp"
	problem2_path := "./input_data/pr76.tsp"

	problems[0] = *problem.InitProblem(problem1_path)
	problems[1] = *problem.InitProblem(problem2_path)
	for i := 2; i < len(problems); i++ {
		if i < 8 {
			problems[i] = *problem.GenerateProblem(200, 40, 150, false)
		} else {
			problems[i] = *problem.GenerateProblem(200, 40, 150, true)
		}
	}
	return problems
}

func crossing_testing(p problem.Problem) {
	crossing_params := [5]float32{0.1, 0.25, 0.5, 0.75, 0.9}
	for i := 0; i < len(crossing_params); i++ {
		for j := 0; j < TEST_REPEAT; j++ {
			problem.Genetic_generate_solution(p, crossing_params[i], 0.2, 1000, 10000, 8, 0.02, i+j)
		}
	}
}

func mutation_testing(p problem.Problem) {
	mutation_params := [5]float32{0.1, 0.25, 0.5, 0.75, 0.9}
	for i := 0; i < len(mutation_params); i++ {
		for j := 0; j < TEST_REPEAT; j++ {
			problem.Genetic_generate_solution(p, 0.7, mutation_params[i], 1000, 10000, 8, 0.02, i+j)
		}
	}
}

func population_size_testing(p problem.Problem) {
	pop_size_params := [3]int{100, 1000, 5000}
	for i := 0; i < len(pop_size_params); i++ {
		for j := 0; j < TEST_REPEAT; j++ {
			problem.Genetic_generate_solution(p, 0.7, 0.2, pop_size_params[i], 10000, 8, 0.02, i+j)
		}
	}
}

func tournament_size_testing(p problem.Problem) {
	tournament_size_params := [3]float32{0.01, 0.05, 0.2}
	for i := 0; i < len(tournament_size_params); i++ {
		for j := 0; j < TEST_REPEAT; j++ {
			problem.Genetic_generate_solution(p, 0.7, 0.2, 1000, 10000, int(tournament_size_params[i]*float32(p.GetDim())), 0.02, i+j)
		}
	}
}

func elitism_size_testing(p problem.Problem) {
	elitism_size_params := [3]float32{0.01, 0.05, 0.2}
	for i := 0; i < len(elitism_size_params); i++ {
		for j := 0; j < TEST_REPEAT; j++ {
			problem.Genetic_generate_solution(p, 0.7, 0.2, 1000, 10000, 8, elitism_size_params[i], i+j)
		}
	}
}
