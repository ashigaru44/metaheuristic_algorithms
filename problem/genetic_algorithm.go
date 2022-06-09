package problem

import (
	"math/rand"
	"time"
)

var probability_mutation_individual = 0.05

// "sort"
type generation struct {
	individuals [][]int
	distances   []int
}

func (p generation) path(id int) *[]int {
	return &p.individuals[id]
}

func (p generation) size() int {
	return len(p.distances)
}

func generate_random_generation(p Problem, size int) generation {
	var individuals = make([][]int, size)
	var distances = make([]int, size)
	for i := range individuals {
		sol, dis := Random(p)
		individuals[i] = *sol
		distances[i] = dis
	}
	return generation{individuals: individuals, distances: distances}
}

func empty_generation(p Problem, size int) generation {
	var individuals = make([][]int, size)
	var distances = make([]int, size)
	return generation{individuals: individuals, distances: distances}
}

// ----------------------------------------------------------   SELECTIONS

func tournament_selection(g generation, tournament_size int) *[]int {
	fighters := make([]int, 0)
	for len(fighters) < tournament_size {
		random_fighter := rand.Intn(g.size())
		id, _ := Find(fighters, random_fighter)
		if id == -1 {
			fighters = append(fighters, random_fighter)
		}
	}
	best_distance := g.distances[0]
	best_fighter := g.path(0)
	for _, fighter := range fighters {
		if g.distances[fighter] < best_distance {
			best_distance = g.distances[fighter]
			best_fighter = g.path(fighter)
		}
	}
	return best_fighter
}

// ----------------------------------------------------------   CROSSOVERS

func ordered_crossover(parent_one *[]int, parent_two *[]int) *[]int {

	child := make([]int, len(*parent_one))
	for i := range child {
		child[i] = -1
	}

	size := len(*parent_one)
	start_point := rand.Intn(size)
	end_point := rand.Intn(size)

	for start_point == end_point {
		end_point = rand.Intn(size)
	}

	if start_point > end_point {
		temp := end_point
		end_point = start_point
		start_point = temp
	}

	for i := start_point; i <= end_point; i++ {
		child[i] = (*parent_one)[i]
	}

	parent_index := 0
	for i := 0; i < size; i++ {
		if child[i] == -1 {
			for j := 0; j < size; j++ {
				if (*parent_two)[parent_index] == child[j] {
					parent_index++
					j = -1
				}
			}
			child[i] = (*parent_two)[parent_index]
		}
	}
	return &child
}

// ----------------------------------------------------------   MUTATIONS

func mutation_swap(individual *[]int) {
	dimension := len(*individual)
	first_index := rand.Intn(dimension)
	second_index := rand.Intn(dimension)

	//Make sure that indexes arent equal
	for second_index == first_index {
		second_index = rand.Intn(dimension)
	}

	first_value := (*individual)[first_index]
	second_value := (*individual)[second_index]

	(*individual)[first_index] = second_value
	(*individual)[second_index] = first_value
}

func mutation_swap_linear(individual *[]int) {
	dimension := len(*individual)
	for i := 0; i < len(*individual)-1; i++ {
		if rand.Float64() <= probability_mutation_individual {
			second_index := rand.Intn(dimension-i-1) + i + 1

			first_value := (*individual)[i]
			second_value := (*individual)[second_index]

			(*individual)[i] = second_value
			(*individual)[second_index] = first_value
		}
	}

}

func mutation_invert(individual *[]int) {
	dimension := len(*individual)
	start := rand.Intn(dimension)
	end := rand.Intn(dimension)

	//Make sure that indexes arent equal
	for start == end {
		end = rand.Intn(dimension)
	}

	if start > end {
		temp := start
		start = end
		end = temp
	}

	for i := 0; i+start < end; i++ {
		temp := (*individual)[start+i]
		(*individual)[start+i] = (*individual)[end-i]
		(*individual)[end-i] = temp
	}
}

func Genetic_generate_solution(p Problem,
	probability_cross float32,
	probability_mutate float32,
	population_size int,
	iterations int,
	tournament_size int,
) *[]int {

	rand.Seed(time.Now().UnixNano())
	old_population := generate_random_generation(p, population_size)
	new_population := empty_generation(p, population_size)

	for i := 0; i < iterations; i++ {
		for j := 0; j < population_size; j++ {

			parent_1 := *tournament_selection(old_population, tournament_size)
			var child []int
			if rand.Float32() <= probability_cross {
				parent_2 := *tournament_selection(old_population, tournament_size)
				child = *ordered_crossover(&parent_1, &parent_2)
			} else {
				child = parent_1
			}
			if rand.Float32() <= probability_mutate {
				mutation_swap_linear(&child)
			}

			new_population.individuals[j] = child
			new_population.distances[j] = p.EvaluateSolution2(&child)
		}
		old_population = new_population
		new_population = generation{}
		new_population = empty_generation(p, population_size)

	}

	best_dist := old_population.distances[0]
	best_ind := old_population.individuals[0]
	for i := range old_population.distances {
		if old_population.distances[i] < best_dist {
			best_dist = old_population.distances[i]
			best_ind = old_population.individuals[i]
		}
	}

	return &best_ind
}

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
