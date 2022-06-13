package problem

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

var probability_mutation_individual = 0.02

// "sort"
// type individual struct {
// 	path []int
// 	distance int
// }

type generation struct {
	individuals [][]int
	distances   []int
}

type island struct {
	old_population generation
	new_population generation
}

func (p generation) get_results() (int, int, int) {
	best_dist := math.MaxInt32
	worst_dist := 0
	avg_dist := 0

	for i := range p.distances {
		new_dist := p.distances[i]
		avg_dist += new_dist
		if worst_dist < new_dist {
			worst_dist = new_dist
		}
		if best_dist > new_dist {
			best_dist = new_dist
		}
	}
	return worst_dist, avg_dist / len(p.distances), best_dist
}

type ByDistance generation

func (a ByDistance) Len() int           { return len(a.individuals) }
func (a ByDistance) Less(i, j int) bool { return a.distances[i] < a.distances[j] }
func (a ByDistance) Swap(i, j int) {
	a.distances[i], a.distances[j] = a.distances[j], a.distances[i]
	a.individuals[i], a.individuals[j] = a.individuals[j], a.individuals[i]
}

func (p generation) path(id int) *[]int {
	return &p.individuals[id]
}

func (p generation) size() int {
	return len(p.distances)
}

func (p generation) elitism(elite_individuals int, new_pop generation) generation {
	sort.Sort(ByDistance(p))

	for i := 0; i < elite_individuals; i++ {
		new_pop.individuals[i] = p.individuals[i]
		new_pop.distances[i] = p.distances[i]
	}
	// new_pop.individuals[:elite_individuals] = p.individuals[:elite_individuals]
	return new_pop
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
	best_distance := math.MaxInt32
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

func crossover_pm(p Problem, parent_one *[]int, parent_two *[]int) *[]int {
	child_1 := crossover_pm_helper(parent_one, parent_two)
	child_2 := crossover_pm_helper(parent_one, parent_two)
	child_1_dist := p.EvaluateSolution2(child_1)
	child_2_dist := p.EvaluateSolution2(child_2)
	if child_1_dist > child_2_dist {
		return child_2
	} else {
		return child_1
	}
}

func crossover_pm_helper(parent_one *[]int, parent_two *[]int) *[]int {
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

	for i := start_point; i <= end_point; i++ {
		if !Contains((*parent_two)[i], &child) {
			searched_val := (*parent_one)[i]
			searched_val_index := Where(searched_val, parent_two)

			for child[searched_val_index] != -1 {
				searched_val := child[searched_val_index]
				searched_val_index = Where(searched_val, parent_two)
			}
			child[searched_val_index] = (*parent_two)[i]
		}
	}

	for i := range child {
		if child[i] == -1 {
			child[i] = (*parent_two)[i]
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
	elitism_size float32,
	id int,
	crossing_variant string,
	mutation_variant string,
) (*[]int, int) {

	// file_name := "./test_output_" + strconv.Itoa(id) + ".txt"
	// f, _ := os.Create(file_name)
	// f.WriteString(fmt.Sprintf("Worst;Avg;Best\n"))
	//check(err)
	// defer f.Close()
	rand.Seed(time.Now().UnixNano())
	elite_individuals := int(elitism_size * float32(population_size))
	old_population := generate_random_generation(p, population_size)
	new_population := empty_generation(p, population_size)

	for i := 0; i < iterations; i++ {
		mean := 0
		for j := 0; j < population_size; j++ {
			if new_population.distances[j] == 0 {
				parent_1 := *tournament_selection(old_population, tournament_size)
				var child []int
				if rand.Float32() <= probability_cross {
					parent_2 := *tournament_selection(old_population, tournament_size)
					switch crossing_variant {
					case "pm":
						child = *crossover_pm(p, &parent_1, &parent_2)
					case "ordered":
						child = *ordered_crossover(&parent_1, &parent_2)
					}
				} else {
					child = parent_1
				}
				if rand.Float32() <= probability_mutate {
					switch mutation_variant {
					case "invert":
						mutation_invert(&child)
					case "linear":
						mutation_swap_linear(&child)
					}
				}

				new_population.individuals[j] = child
				new_population.distances[j] = p.EvaluateSolution2(&child)
			}
			mean += new_population.distances[j]
		}
		mean /= population_size
		old_population = new_population

		// for j := 0; j < population_size; j++ {
		// 	fmt.Println(j, "=", new_population.distances[j])
		// }
		new_population = generation{}
		new_population = empty_generation(p, population_size)
		new_population = old_population.elitism(elite_individuals, new_population)
		//fmt.Print("\nIteration: ", i)
		//fmt.Print("\t Worst individual: ", old_population.distances[population_size-1], "\t Average individual: ", mean, "\t Best individual: ", old_population.distances[0])
		/*f.WriteString(fmt.Sprintf("%d;%d;%d\n",
		old_population.distances[population_size-1],
		mean,
		old_population.distances[0]))
		*/

	}
	best_dist := old_population.distances[0]
	best_ind := old_population.individuals[0]
	for i := range old_population.distances {
		if old_population.distances[i] < best_dist {
			best_dist = old_population.distances[i]
			best_ind = old_population.individuals[i]
		}
	}

	return &best_ind, best_dist
}

func GA_Islands_generate_solution(p Problem,
	probability_cross float32,
	probability_mutate float32,
	population_size int,
	iterations int,
	tournament_size int,
	islands_amount int,
	migration_rate float32,
	migration_interval int,
	id int,
	migration_start_ratio int,
) (*[]int, int) {

	file_name := "./test_island_output_" + strconv.Itoa(id) + ".txt"
	f, _ := os.Create(file_name)
	f.WriteString(fmt.Sprintf("Worst;Avg;Best\n"))
	//check(err)
	defer f.Close()

	file_name2 := "./test_islands_best_output_" + strconv.Itoa(id) + ".txt"
	f2, _ := os.Create(file_name2)
	names := ""
	for i := 1; i < islands_amount; i++ {
		names += "i_" + strconv.Itoa(i) + ";"
	}
	names += "i_" + strconv.Itoa(islands_amount) + "\n"
	f2.WriteString(names)
	//check(err)
	defer f2.Close()

	rand.Seed(time.Now().UnixNano())

	// prepare islands populations
	island_shift := 1
	island_population_size := population_size / islands_amount
	migration_size := int(migration_rate * float32(island_population_size))
	island_shrinked_population_size := island_population_size - migration_size
	islands := make([]island, islands_amount)
	for i := 0; i < islands_amount; i++ {
		islands[i].old_population = generate_random_generation(p, island_population_size)
		islands[i].new_population = empty_generation(p, island_population_size)
	}

	for i := 0; i < iterations; i++ {
		for l := 0; l < islands_amount; l++ {
			for j := 0; j < island_population_size; j++ {
				// handle new subpopulation on island
				parent_1 := *tournament_selection(islands[l].old_population, tournament_size)
				var child []int
				if rand.Float32() <= probability_cross {
					parent_2 := *tournament_selection(islands[l].old_population, tournament_size)
					child = *ordered_crossover(&parent_1, &parent_2)
				} else {
					child = parent_1
				}
				if rand.Float32() <= probability_mutate {
					mutation_invert(&child)
				}

				islands[l].new_population.individuals[j] = child
				islands[l].new_population.distances[j] = p.EvaluateSolution2(&child)
			}

			// clear new_population
			islands[l].old_population = islands[l].new_population
			islands[l].new_population = generation{}
			islands[l].new_population = empty_generation(p, island_population_size)

			/*_, _, best := islands[l].old_population.get_results()

			if l < islands_amount-1 {
				f2.WriteString(fmt.Sprintf("%d;", best))
			} else {
				f2.WriteString(fmt.Sprintf("%d", best))
			}*/

		}
		/*worst, avg, best := islands[0].old_population.get_results()
		f.WriteString(fmt.Sprintf("%d;%d;%d\n", worst, avg, best))
		f2.WriteString(fmt.Sprintf("\n"))*/

		// handle migration
		if i%migration_interval == 0 && i >= iterations/migration_start_ratio {
			//fmt.Println("Migration ", i/migration_interval+1, " / ", iterations/migration_interval)
			migrants := make([]generation, islands_amount)

			for l := 0; l < islands_amount; l++ {
				sort.Sort(ByDistance(islands[l].old_population))

				// remove worst members from island's population
				islands[l].old_population.individuals = islands[l].old_population.individuals[:island_shrinked_population_size]
				islands[l].old_population.distances = islands[l].old_population.distances[:island_shrinked_population_size]

				// add island's best members' copies to ships
				target_island := (l + island_shift) % islands_amount
				migrants[target_island] = empty_generation(p, migration_size)
				migrants[target_island].individuals = islands[l].old_population.individuals[:migration_size]
				migrants[target_island].distances = islands[l].old_population.distances[:migration_size]
			}

			for l := 0; l < islands_amount; l++ {

				// simultaneously transfer migrants to destination islands
				for m := 0; m < migration_size; m++ {
					islands[l].old_population.individuals = append(islands[l].old_population.individuals, migrants[l].individuals[m])
					islands[l].old_population.distances = append(islands[l].old_population.distances, migrants[l].distances[m])
				}

			}
			// shift destination island to the right by one.
			island_shift = island_shift + 1
			// islands are formed as circle, so adjustment is needed as follows:
			if island_shift%islands_amount == 0 {
				island_shift = 1
			}
		}

	}
	f.Sync()
	f2.Sync()

	best_dist := islands[0].old_population.distances[0]
	best_ind := islands[0].old_population.individuals[0]
	for l := 0; l < islands_amount; l++ {
		for i := range islands[l].old_population.distances {
			if islands[l].old_population.distances[i] < best_dist {
				best_dist = islands[l].old_population.distances[i]
				best_ind = islands[l].old_population.individuals[i]
			}

		}
	}

	return &best_ind, best_dist
}

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
