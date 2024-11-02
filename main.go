package main

import (
	"fmt"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/atsp"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/utils"
	"log"
)

func main() {
	G, err := utils.ReadGraphFromFile("data/tsp_10.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(G.ToString())

	var tsp atsp.ATSP
	tsp = atsp.NewBruteForceSolver(G)
	fmt.Println(tsp.Solve(0))

	tsp = atsp.NewBranchAndBoundSolver(G)
	fmt.Println(tsp.Solve(0))

	tsp = atsp.NewDynamicProgrammingSolver(G)
	fmt.Println(tsp.Solve(0))
}
