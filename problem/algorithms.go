package problem

// import "fmt"

// type RandomAlgorithm struct {
// }

func RandomAlgorithm(p Problem) *[][2]int {
	temp_nodes := make([][2]int, len(p.nodes)-1)
	copy(temp_nodes, p.nodes[1:])
	// fmt.Println("p.nodes:")
	// for i := range p.nodes {
	// 	fmt.Println(p.nodes[i])
	// }
	// fmt.Println("Temp_nodes:")
	// for i := range temp_nodes {
	// 	fmt.Println(temp_nodes[i])
	// }
	// fmt.Println("Result_path:")
	result_path := make([][2]int, len(p.nodes))
	result_path[0] = p.nodes[0]
	// for i := range result_path {
	// fmt.Println(result_path[i])
	// }

	return &result_path
}
