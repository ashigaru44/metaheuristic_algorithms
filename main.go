package main

import (
	//"math/rand"
	"meta-heur/tsp/problem"
	//"time"
)

//var r rand.Rand

func main() {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	problem_path := "C:/Users/ashig/Documents/meta-heur_data/berlin52.tsp"
	problem.InitProblem(problem_path)
}
