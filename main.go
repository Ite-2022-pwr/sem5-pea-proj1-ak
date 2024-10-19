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

	tsp := atsp.NewBruteForceSolver(G)
	fmt.Println(tsp.Solve(0))

	pq := utils.NewPriorityQueue(func(a, b int) bool {
		return a < b
	})

	pq.Push(2)
	pq.Push(1)
	pq.Push(3)
	pq.Push(7)
	pq.Push(5)
	pq.Push(4)
	pq.Push(6)

	for !pq.IsEmpty() {
		//fmt.Println(pq.GetElements())
		fmt.Println(pq.Pop())
	}
}
