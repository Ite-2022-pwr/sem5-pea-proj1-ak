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

	tsp := atsp.NewATSP(G)
	fmt.Println(tsp.BruteForce(0))
}
