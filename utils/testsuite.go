package utils

import (
	//"math/rand"
	"fmt"
	"meta-heur/tsp/problem"
	"os"
)

const TEST_REPEAT = 2

func Test_run_GA() {
	problems := generate_testing_problems()
	f, _ := os.Create("./testsuite_output.txt")

	crossing_testing(problems, *f)
	mutation_testing(problems, *f)
	population_size_testing(problems, *f)
	tournament_size_testing(problems, *f)
	elitism_size_testing(problems, *f)
	f.Sync()
	defer f.Close()
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

func crossing_testing(p [16]problem.Problem, f os.File) {
	crossing_params := [5]float32{0.1, 0.25, 0.5, 0.75, 0.9}
	MSG := "crossing_ordered"

	fmt.Println(MSG + " ENTERED")
	for i := 0; i < len(crossing_params); i++ {
		f.WriteString(fmt.Sprintf("%s;%f;", MSG, crossing_params[i]))
		for m := 0; m < len(p); m++ {
			best_dist := 0
			for j := 0; j < TEST_REPEAT; j++ {
				_, dist := problem.Genetic_generate_solution(p[m], crossing_params[i], 0.2, 1000, 10000, 8, 0.02, i+j)
				best_dist += dist
			}
			best_dist = int(best_dist / TEST_REPEAT)
			f.WriteString(fmt.Sprintf("%d;", best_dist))
		}
		f.WriteString(fmt.Sprintf("\n"))
	}
	fmt.Println(MSG + " LEFT")
}

func mutation_testing(p [16]problem.Problem, f os.File) {
	mutation_params := [5]float32{0.1, 0.25, 0.5, 0.75, 0.9}
	MSG := "mutation_inverted"

	fmt.Println(MSG + " ENTERED")
	for i := 0; i < len(mutation_params); i++ {
		f.WriteString(fmt.Sprintf("%s;%f;", MSG, mutation_params[i]))
		for m := 0; m < len(p); m++ {
			best_dist := 0
			for j := 0; j < TEST_REPEAT; j++ {
				_, dist := problem.Genetic_generate_solution(p[m], 0.7, mutation_params[i], 1000, 10000, 8, 0.02, i+j)
				best_dist += dist
			}
			best_dist = int(best_dist / TEST_REPEAT)
			f.WriteString(fmt.Sprintf("%d;", best_dist))
		}
		f.WriteString(fmt.Sprintf("\n"))
	}
	fmt.Println(MSG + " LEFT")
}

func population_size_testing(p [16]problem.Problem, f os.File) {
	pop_size_params := [3]int{100, 1000, 5000}
	MSG := "population"

	fmt.Println(MSG + " ENTERED")
	for i := 0; i < len(pop_size_params); i++ {
		f.WriteString(fmt.Sprintf("%s;%d;", MSG, pop_size_params[i]))
		for m := 0; m < len(p); m++ {
			best_dist := 0
			for j := 0; j < TEST_REPEAT; j++ {
				_, dist := problem.Genetic_generate_solution(p[m], 0.7, 0.2, pop_size_params[i], 10000, 8, 0.02, i+j)
				best_dist += dist
			}
			best_dist = int(best_dist / TEST_REPEAT)
			f.WriteString(fmt.Sprintf("%d;", best_dist))
		}
		f.WriteString(fmt.Sprintf("\n"))
	}
	fmt.Println(MSG + " LEFT")
}

func tournament_size_testing(p [16]problem.Problem, f os.File) {
	tournament_size_params := [3]float32{0.01, 0.05, 0.2}
	MSG := "tournament"

	fmt.Println(MSG + " ENTERED")
	for i := 0; i < len(tournament_size_params); i++ {
		f.WriteString(fmt.Sprintf("%s;%f;", MSG, tournament_size_params[i]))
		for m := 0; m < len(p); m++ {
			best_dist := 0
			for j := 0; j < TEST_REPEAT; j++ {
				_, dist := problem.Genetic_generate_solution(p[m], 0.7, 0.2, 1000, 10000, int(tournament_size_params[i]*float32(p[m].GetDim())), 0.02, i+j)
				best_dist += dist
			}
			best_dist = int(best_dist / TEST_REPEAT)
			f.WriteString(fmt.Sprintf("%d;", best_dist))
		}
		f.WriteString(fmt.Sprintf("\n"))
	}
	fmt.Println(MSG + " LEFT")
}

func elitism_size_testing(p [16]problem.Problem, f os.File) {
	elitism_size_params := [3]float32{0.01, 0.05, 0.2}
	MSG := "elitism"

	fmt.Println(MSG + " ENTERED")
	for i := 0; i < len(elitism_size_params); i++ {
		f.WriteString(fmt.Sprintf("%s;%f;", MSG, elitism_size_params[i]))
		for m := 0; m < len(p); m++ {
			best_dist := 0
			for j := 0; j < TEST_REPEAT; j++ {
				_, dist := problem.Genetic_generate_solution(p[m], 0.7, 0.2, 1000, 10000, 8, elitism_size_params[i], i+j)
				best_dist += dist
			}
			best_dist = int(best_dist / TEST_REPEAT)
			f.WriteString(fmt.Sprintf("%d;", best_dist))
		}
		f.WriteString(fmt.Sprintf("\n"))
	}
	fmt.Println(MSG + " LEFT")
}
