package problem

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	// "sort"
)

// UTILS

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
	fmt.Println("Random_k, k = ", k)
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

func opt2Swap(solution *[]int, a int, b int) (*[]int) {
	new_route := make([]int, len(*solution))
	for i := 0; i <= a - 1; i++ {
		new_route[i] = (*solution)[i]
	}
	dec := 0
	for i := a; i <= b ; i++ {
		new_route[i] = (*solution)[b - dec]
		dec++
	}
	for i := b + 1; i < len(*solution); i++ {
		new_route[i] = (*solution)[i]
	}

	return &new_route
}

func opt2Rec(p Problem, adj_matrix [][]int, solution *[]int, best_distance int) (*[]int, int) {
	for i := 0; i < p.dim - 1; i++ {
		for j := i + 1; j < p.dim; j++ {
			var new_route*[] int = opt2Swap(solution, i, j)
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
	var best_new_route*[] int
	for i := 0; i < p.dim - 1; i++ {
		for j := i + 1; j < p.dim; j++ {
			var new_route*[] int = opt2Swap(solution, i, j)

			var new_distance = p.EvaluateSolution2(new_route)
			if new_distance < best_distance {
				best_new_distance = new_distance
				best_new_route = new_route

			}
		}
	}
	if best_new_distance < best_distance {
		//fmt.Println("best_new_distance: ", best_new_distance)
		return opt2_PickBestRec(p, adj_matrix, best_new_route, best_new_distance)
	}
	return solution, best_distance
}


func Opt2_PickBest(p Problem, adj_matrix [][]int, solution *[]int) (*[]int, int) {
	var distance = p.EvaluateSolution2(solution)
	var new_route, new_distance = opt2_PickBestRec(p, adj_matrix, solution, distance)
	return new_route, new_distance
}

func Accelerated_opt2Swap(solution *[]int, a int, b int) (*[]int) {
	new_route := make([]int, len(*solution))
	for i := 0; i <= a - 1; i++ {
		new_route[i] = (*solution)[i]
	}
	dec := 0
	for i := a; i <= b ; i++ {
		new_route[i] = (*solution)[b - dec]
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
	var best_new_route*[] int
	for i := 0; i < p.dim - 1; i++ {
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
	distance := (*next_distances)[p.dim - 1]
	var new_route, new_distance = Accelerated_opt2_PickBestRec(p, solution, distance, next_distances)
	return new_route, new_distance
}

// GREEDY / NEAREST ////////////////////////////////////////////////////////////////////////////

func NearestNeighbourAllPoints(p Problem, adj_matrix [][]int) (*[]int, int) {
	var best_path *[]int
	best_dist := math.MaxInt32
	for i := 0; i < p.dim; i++ {
		tmp_path, tmp_dist := NearestNeighbour(p, adj_matrix, i)
		fmt.Println("Current path: ", *tmp_path)
		fmt.Println("Current dist = ", tmp_dist)
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

// func slice_contain(slice1 *[]int, slice2 *[]int) {
// for i:
// }
