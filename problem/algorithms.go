package problem

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"
	// "sort"
)

type TabuNeigbourAlgorithm int

const (
	SwapTabu TabuNeigbourAlgorithm = 0
	Opt2Tabu                       = 1
)

// UTILS

var wg sync.WaitGroup

func initTempAdjMatrix(adj_matrix [][]int, size int) [][]int {
	tmp_adj_matrix := make([][]int, size)

	for i := range tmp_adj_matrix {
		tmp_adj_matrix[i] = make([]int, size)
	}

	for i := range adj_matrix {
		for j := range adj_matrix[i] {
			tmp_adj_matrix[i][j] = adj_matrix[i][j]
		}
	}
	return tmp_adj_matrix
}

// Random ////////////////////////////////////////////////////////////////////////////

func Random(p Problem) (*[]int, int) {
	var distance = 0
	temp_nodes := make([]int, len(p.nodes)-1)
	for i := range temp_nodes {
		temp_nodes[i] = i + 1
	}

	result_path := make([]int, len(p.nodes))
	result_path[0] = 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(result_path)-1; i++ {
		next_index := rand.Intn(len(temp_nodes))
		result_path[i+1] = temp_nodes[next_index]
		temp_nodes = append(temp_nodes[:next_index], temp_nodes[next_index+1:]...)
		distance += p.GetDistance(result_path[i], result_path[i+1])
	}
	return &result_path, distance
}

// Random repeat k times ////////////////////////////////////////////////////////////////////////////

func Random_k(p Problem, k int) (*[]int, int) {
	// fmt.Println("Random_k, k = ", k)
	var best_route *[]int
	for i := 0; i < k; i++ {
		new_route, _ := Random(p)
		if best_route == nil || p.EvaluateSolution2(new_route) < p.EvaluateSolution2(best_route) {
			best_route = new_route
		}
	}
	//fmt.Println("Path:", *best_route)
	//fmt.Println("Distance = ", p.EvaluateSolution2(best_route))
	return best_route, p.EvaluateSolution2(best_route)
}

// Random repeat for duration ////////////////////////////////////////////////////////////////////////////

func Random_time(p Problem, duration int64) (*[]int, int) {
	fmt.Println("Random_time, time = ", duration)
	var best_route *[]int
	start := time.Now()
	i := 0
	for time.Since(start).Nanoseconds() < duration {
		i++
		new_route, _ := Random(p)
		if best_route == nil || p.EvaluateSolution2(new_route) < p.EvaluateSolution2(best_route) {
			best_route = new_route
		}
	}
	//fmt.Println("Path:", *best_route)
	//fmt.Println("Distance = ", p.EvaluateSolution2(best_route))
	fmt.Println("Random_time, loops = ", i)
	return best_route, p.EvaluateSolution2(best_route)
}

// 2OPT BASIC ////////////////////////////////////////////////////////////////////////////

func opt2Swap(solution *[]int, a int, b int) *[]int {
	new_route := make([]int, len(*solution))
	for i := 0; i <= a-1; i++ {
		new_route[i] = (*solution)[i]
	}
	dec := 0
	for i := a; i <= b; i++ {
		new_route[i] = (*solution)[b-dec]
		dec++
	}
	for i := b + 1; i < len(*solution); i++ {
		new_route[i] = (*solution)[i]
	}

	return &new_route
}

func opt2Rec(p Problem, adj_matrix [][]int, solution *[]int, best_distance int) (*[]int, int) {
	for i := 0; i < p.dim-1; i++ {
		for j := i + 1; j < p.dim; j++ {
			var new_route *[]int = opt2Swap(solution, i, j)
			//fmt.Println("Current path: ", *new_route)

			var new_distance int = p.EvaluateSolution2(new_route)
			if new_distance < best_distance {
				//fmt.Println("new_distance: ", new_distance)
				return opt2Rec(p, adj_matrix, new_route, new_distance)
			}
		}
	}
	return solution, best_distance
}

func Opt2(p Problem, adj_matrix [][]int, solution *[]int) (*[]int, int) {
	//var best_path *[]int
	fmt.Println("Opt2")
	var distance = p.EvaluateSolution2(solution)
	var new_route, new_distance = opt2Rec(p, adj_matrix, solution, distance)
	return new_route, new_distance
}

// 2OPT PICK BEST ////////////////////////////////////////////////////////////////////////////

func opt2_PickBestRec(p Problem, adj_matrix [][]int, solution *[]int, best_distance int) (*[]int, int) {
	var best_new_distance int = math.MaxInt32
	var best_new_route *[]int
	for i := 0; i < p.dim-1; i++ {
		for j := i + 1; j < p.dim; j++ {
			var new_route *[]int = opt2Swap(solution, i, j)

			var new_distance = p.EvaluateSolution2(new_route)
			if new_distance < best_distance {
				best_new_distance = new_distance
				best_new_route = new_route

			}
		}
	}
	if best_new_distance < best_distance {
		// fmt.Println("best_new_distance: ", best_new_distance)
		// fmt.Println(best_new_distance)
		return opt2_PickBestRec(p, adj_matrix, best_new_route, best_new_distance)
	}
	return solution, best_distance
}

func Opt2_PickBest(p Problem, adj_matrix [][]int, solution *[]int) (*[]int, int) {
	var distance = p.EvaluateSolution2(solution)
	var new_route, new_distance = opt2_PickBestRec(p, adj_matrix, solution, distance)
	return new_route, new_distance
}

func Accelerated_opt2Swap(solution *[]int, a int, b int) *[]int {
	new_route := make([]int, len(*solution))
	for i := 0; i <= a-1; i++ {
		new_route[i] = (*solution)[i]
	}
	dec := 0
	for i := a; i <= b; i++ {
		new_route[i] = (*solution)[b-dec]
		dec++
	}
	for i := b + 1; i < len(*solution); i++ {
		new_route[i] = (*solution)[i]
	}

	return &new_route
}

// ACCELERATED OPT 2 PICK BEST ////////////////////////////////////////////////////////////////////////////

func Accelerated_opt2_PickBestRec(p Problem, solution *[]int, best_distance int, distances *[]int) (*[]int, int) {
	var best_new_distance int = math.MaxInt32
	var best_new_route *[]int
	for i := 0; i < p.dim-1; i++ {
		for j := i + 1; j < p.dim; j++ {
			//new_distance := (*distances)[p.dim-1] - p.Adj_matrix[(*solution)[i]][(*solution)[i+1]] - p.Adj_matrix[(*solution)[j]][(*solution)[j-1]]
			new_distance := best_distance - p.Adj_matrix[(*solution)[i]][(*solution)[i+1]] - p.Adj_matrix[(*solution)[j]][(*solution)[j-1]]
			new_distance += p.Adj_matrix[(*solution)[i]][(*solution)[j-1]] + p.Adj_matrix[(*solution)[j]][(*solution)[i+1]]
			if new_distance < best_distance {
				best_new_route = Accelerated_opt2Swap(solution, i, j)
				best_new_distance = p.EvaluateSolution2(best_new_route)
				fmt.Println("BEST NEW DISTANCE  01  :", best_new_distance)
				fmt.Println("BEST NEW DISTANCE  02  :", p.EvaluateSolution2(best_new_route))
				fmt.Println("")
				//best_new_route = new_route
			}
			//var new_distance = p.EvaluateSolution2(new_route)

		}
	}
	if best_new_distance < best_distance {
		new_distances := p.EvaluateSolutionIncrement(best_new_route)
		return Accelerated_opt2_PickBestRec(p, best_new_route, best_new_distance, new_distances)
	}
	return solution, best_distance
}

func Accelerated_Opt2_PickBest(p Problem, solution *[]int) (*[]int, int) {
	next_distances := p.EvaluateSolutionIncrement(solution)
	distance := (*next_distances)[p.dim-1]
	var new_route, new_distance = Accelerated_opt2_PickBestRec(p, solution, distance, next_distances)
	return new_route, new_distance
}

// GREEDY / NEAREST ////////////////////////////////////////////////////////////////////////////

func NearestNeighbourAllPoints(p Problem, adj_matrix [][]int) (*[]int, int) {
	var best_path *[]int
	best_dist := math.MaxInt32
	for i := 0; i < p.dim; i++ {
		tmp_path, tmp_dist := NearestNeighbour(p, adj_matrix, i)
		// fmt.Println("Current path: ", *tmp_path)
		// fmt.Println("Current dist = ", tmp_dist)
		if tmp_dist < best_dist {
			best_path = tmp_path
			best_dist = tmp_dist
		}
	}
	return best_path, best_dist
}

func NearestNeighbour(p Problem, adj_matrix [][]int, start_index int) (*[]int, int) {
	var distance = 0
	tmp_adj_matrix := initTempAdjMatrix(adj_matrix, p.dim)

	temp_nodes := make([]int, len(p.nodes)-1)
	for i := range temp_nodes {
		temp_nodes[i] = i + 1
	}

	result_path := make([]int, len(p.nodes))
	result_path[0] = start_index
	mark_column(&tmp_adj_matrix, start_index)

	var next_index int
	var dist_to_point int
	// for i := range tmp_adj_matrix {
	// 	fmt.Println(tmp_adj_matrix[i])
	// }
	for i := 0; i < len(result_path)-1; i++ {
		next_index, dist_to_point = getNearestPoint(&tmp_adj_matrix, result_path[i])
		result_path[i+1] = next_index
		// fmt.Println("next_index = ", next_index)
		mark_column(&tmp_adj_matrix, next_index)
		// for i := range tmp_adj_matrix {
		// 	fmt.Println(tmp_adj_matrix[i])
		// }
		distance += dist_to_point
	}

	return &result_path, distance
}

func getNearestPoint(adj_matrix *[][]int, point_index int) (int, int) {
	min := math.MaxInt32
	min_index := -1
	for i, val := range (*adj_matrix)[point_index] {
		if i == 0 || (val < min && val > 0) {
			min = val
			min_index = i
		}
	}
	return min_index, min
}

func mark_column(adj_matrix *[][]int, index int) {
	for i := range *adj_matrix {
		(*adj_matrix)[i][index] = math.MaxInt32
	}
}

// TABU SEARCH ////////////////////////////////////////////////////////////////////////////

type neighbouring_solution struct {
	swap_elements [2]int
	distance      int
}

func (s neighbouring_solution) getValue() int {
	return s.distance
}

func (s1 neighbouring_solution) isBiggerThan(s2 neighbouring_solution) bool {
	return s1.distance > s2.distance
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func Tabu_search_concurrent(p Problem, initial_method string, terminate_criteria int, aspiration_criteria float32, tabuTenure int, num_of_routines int, alg TabuNeigbourAlgorithm) (*[]int, int) {
	initial_paths := make([]*[]int, num_of_routines)
	channels_paths := make([]chan []int, num_of_routines)
	channels_dists := make([]chan int, num_of_routines)
	best_distances := make([]int, num_of_routines)
	var best_path []int
	best_distance := math.MaxInt32

	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		initial_paths[i], _ = Random_k(p, 100)
	}
	for i := 0; i < num_of_routines; i++ {
		channels_paths[i] = make(chan []int, 1)
		channels_dists[i] = make(chan int, 1)
		go Tabu_search(p, initial_paths[i], terminate_criteria, aspiration_criteria, tabuTenure, channels_paths[i], channels_dists[i], alg)
	}
	wg.Wait()

	var best_dist_index int

	for i, dist_channel := range channels_dists {
		best_distances[i] = <-dist_channel
	}

	for i, dist := range best_distances {
		if dist < best_distance {
			best_distance = dist
			best_dist_index = i
		}
	}

	best_path = <-channels_paths[best_dist_index]

	return &best_path, best_distance
}

func Tabu_search(p Problem,
	initial_path *[]int,
	terminate_criteria int,
	aspiration_criteria float32,
	tabuTenure int,
	ch_best_path chan<- []int,
	// ch_best_new_path <-chan []int,
	ch_best_dist chan<- int,
	alg TabuNeigbourAlgorithm) (*[]int, int) {

	f, _ := os.Create("./output.txt")
	//check(err)
	defer f.Close()

	defer timeTrack(time.Now(), "Tabu_search")
	defer func() {
		close(ch_best_dist)
		close(ch_best_path)
	}()
	defer wg.Done()
	// tabu list contains previous swap movements e.g. [[2;5],[1;7],[9:15]]
	var tabu_list [][2]int
	best_path := initial_path
	best_distance := math.MaxInt32
	current_path := initial_path
	var current_distance int
	var neighbouring_solutions []neighbouring_solution

	iter := 0
	terminate := 0
	for terminate < terminate_criteria {
		// fmt.Print("\n###iter ", iter, " Current distance: ", current_distance, " Best_distance: ", best_distance)
		current_distance = math.MaxInt32
		switch alg {
		case SwapTabu:
			neighbouring_solutions = generate_solutions(p, current_path, tabuTenure)
		case Opt2Tabu:
			neighbouring_solutions = generate_solutions_2opt(p, current_path, tabuTenure)
		}

		if len(neighbouring_solutions) == 0 {
			fmt.Println("lmao")
		}
		//neighbouring_solutions := generate_solutions(p, current_path, tabuTenure)
		for {
			var best_move [2]int
			best_swap_dist := math.MaxInt32
			best_sol_index := -1
			for i, s := range neighbouring_solutions {
				if s.getValue() < best_swap_dist {
					best_move = s.swap_elements
					best_swap_dist = s.distance
					best_sol_index = i
				}
			}

			// Check if move is in tabu list
			var isTabu bool = false
			for _, move := range tabu_list {
				if move == best_move {
					isTabu = true
				}
			}
			// Case with move NOT IN tabu list
			if !isTabu {
				//current_path = swap(current_path, best_move[0], best_move[1])
				// current_path = 2opt
				switch alg {
				case SwapTabu:
					current_path = swap(current_path, best_move[0], best_move[1])
				case Opt2Tabu:
					current_path = opt2Swap(current_path, best_move[0], best_move[1])
				}
				current_distance = p.EvaluateSolution2(current_path)
				//fmt.Println(current_distance)
				if best_swap_dist < best_distance {
					best_path = current_path
					best_distance = current_distance
					terminate = 0

				} else {
					// Move is not better then previously found, incrementing iterator
					terminate++
				}
				if len(tabu_list) >= tabuTenure {
					tabu_list = tabu_list[1:]
				}
				tabu_list = append(tabu_list, best_move)

				break
				// Case with move IN tabu list
			} else {
				if float32(best_swap_dist) < float32(best_distance)*aspiration_criteria {
					// Ignore tabu list if found distance is better then aspiration criteria
					// fmt.Print("\nAspiration criteria has been acheived\nBest distance: ", best_distance)
					current_path = swap(current_path, best_move[0], best_move[1])
					current_distance = p.EvaluateSolution2(current_path)
					best_path = current_path
					best_distance = current_distance
					terminate = 0
					break
				} else {
					// break
					neighbouring_solutions[best_sol_index].distance = math.MaxInt32
					continue
				}
			}
		}
		iter++
		//fmt.Println(best_distance)
		f.WriteString(fmt.Sprintf("%d\n", best_distance))
	}
	f.Sync()
	fmt.Println("END")
	ch_best_dist <- best_distance
	ch_best_path <- *best_path
	return best_path, best_distance
}

func swap(path *[]int, i int, j int) *[]int {
	path_copy := make([]int, len(*path))
	copy(path_copy, *path)
	i = find_index_by_element(&path_copy, i)
	j = find_index_by_element(&path_copy, j)
	path_copy[i], path_copy[j] = path_copy[j], path_copy[i]
	return &path_copy
}

func find_index_by_element(arr *[]int, element int) int {
	for i := range *arr {
		if (*arr)[i] == element {
			return i
		}
	}
	return -1
}

func generate_solutions(p Problem, path *[]int, tabuTenure int) []neighbouring_solution {
	var neighbouring_solutions []neighbouring_solution
	base_score := p.EvaluateSolution2(path)
	best_score := base_score
	for i := 0; i < len(*path); i++ {
		for j := i + 1; j <= len(*path)-1; j++ {
			new_path := swap(path, i, j)
			new_path_score := p.EvaluateSolution2(new_path)
			if new_path_score < best_score {
				if len(neighbouring_solutions) > tabuTenure {
					neighbouring_solutions = neighbouring_solutions[1:]
				}
				neighbouring_solutions = append(neighbouring_solutions, neighbouring_solution{[2]int{i, j}, new_path_score})
			} else if len(neighbouring_solutions) <= tabuTenure {

				neighbouring_solutions = append(neighbouring_solutions, neighbouring_solution{[2]int{i, j}, new_path_score})
			}
		}
	}

	return neighbouring_solutions
}

func generate_solutions_2opt(p Problem, path *[]int, tabuTenure int) []neighbouring_solution {
	var neighbouring_solutions []neighbouring_solution
	base_score := p.EvaluateSolution2(path)
	best_score := base_score
	for i := 0; i < len(*path)-1; i++ {
		for j := i + 1; j < len(*path); j++ {
			new_path := opt2Swap(path, i, j)
			new_path_score := p.EvaluateSolution2(new_path)
			if new_path_score < best_score {
				if len(neighbouring_solutions) > tabuTenure {
					neighbouring_solutions = neighbouring_solutions[1:]
				}
				neighbouring_solutions = append(neighbouring_solutions, neighbouring_solution{[2]int{i, j}, new_path_score})
			} else if len(neighbouring_solutions) <= tabuTenure {

				neighbouring_solutions = append(neighbouring_solutions, neighbouring_solution{[2]int{i, j}, new_path_score})
			}
		}
	}

	return neighbouring_solutions
}
