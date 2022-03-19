package main

import (
	//"math/rand"
	"meta-heur/tsp/problem"
	//"time"
)



//var r rand.Rand

func main() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
<<<<<<< Updated upstream
	problem_path := "C:/Users/wolski/Documents/!meta/metaheuristic_algorithms/berlin52.tsp"
	p1 := problem.InitProblem(problem_path)
	p1.PrintProblem()
	pr := problem.GenerateProblem(10)
	pr.PrintProblem()
=======
	problem_path := "./berlin52.tsp"
	problem.InitProblem(problem_path)
>>>>>>> Stashed changes
}
