package main

import "github.com/Ite-2022-pwr/sem5-pea-proj1-ak/benchmark"

func main() {
	//G, err := utils.ReadGraphFromFile("data/input/tsp_10_test1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(G.ToString())
	//
	//var tsp atsp.ATSP
	////tsp = atsp.NewBruteForceSolver(G)
	////benchmark.MeasureSolveTime(tsp, "Brute Force")
	//
	//tsp = atsp.NewBranchAndBoundSolver(G)
	//benchmark.MeasureSolveTime(tsp, "Branch and Bound")
	//
	//tsp = atsp.NewDynamicProgrammingSolver(G)
	//benchmark.MeasureSolveTime(tsp, "Dynamic Programming")

	//benchmark.BranchAndBound()
	//benchmark.DynamicProgramming()
	//benchmark.BruteForce()

	//menu.RunMenu()
	benchmark.All()
}
