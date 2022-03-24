package problem

import (
	"fmt"
	"math"
	"math/rand"
	// "sort"
)

// type RandomAlgorithm struct {
// }

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
