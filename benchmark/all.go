package benchmark

import (
	"fmt"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/atsp"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/generator"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/utils"
	"log"
	"path/filepath"
	"runtime/debug"
)

func All() {
	promptBB := utils.BlueColor("Branch and Bound")
	promptDP := utils.BlueColor("Dynamic programming")
	totalTimeBB, totalTimeDP := 0.0, 0.0
	var resultBB, resultDP [][]string

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Algorytmów"))
	for numOfCities := MinVertices; numOfCities <= MaxVertices; numOfCities++ {
		for i := 0; i < NumberOfGraphs; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, test: %d/%d", numOfCities, i+1, NumberOfGraphs)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
			var tsp atsp.ATSP
			tsp = atsp.NewBranchAndBoundSolver(G)
			totalTimeBB += MeasureSolveTime(tsp, promptBB)

			tsp = atsp.NewDynamicProgrammingSolver(G)
			totalTimeDP += MeasureSolveTime(tsp, promptDP)
			debug.FreeOSMemory()
		}
		avgTimeBB := totalTimeBB / float64(NumberOfGraphs)
		avgTimeDP := totalTimeDP / float64(NumberOfGraphs)
		resultBB = append(resultBB, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%d", int64(avgTimeBB))})
		resultDP = append(resultDP, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%d", int64(avgTimeDP))})
		utils.SaveCSV(filepath.Join(OutputDirectory, "branch_and_bound2.csv"), resultBB)
		utils.SaveCSV(filepath.Join(OutputDirectory, "dynamic_programming2.csv"), resultDP)
		totalTimeBB, totalTimeDP = 0.0, 0.0
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "branch_and_bound2.csv"), resultBB)
	utils.SaveCSV(filepath.Join(OutputDirectory, "dynamic_programming2.csv"), resultDP)
}
