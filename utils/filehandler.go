package utils

import (
	"bufio"
	"fmt"
	"github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadGraphFromFile(filePath string) (graph.Graph, error) {
	log.Println(YellowColor("[*] Czytanie danych z pliku: ", filePath))
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(RedColor(fmt.Sprintf("[!!] Nie można otworzyć pliku: %v", err)))
		return nil, err
	}
	defer f.Close()

	rdr := bufio.NewReader(f)
	var line string
	line, err = rdr.ReadString('\n')
	if err != nil {
		log.Println(RedColor(fmt.Sprintf("[!!] Błąd czytania linii '%v': %v", line, err)))
		return nil, err
	}
	line = strings.TrimSpace(line)

	vertices, err := strconv.Atoi(line)
	if err != nil {
		log.Println(RedColor(fmt.Sprintf("[!!] Błąd konwersji liczby wierzchołków: %v", err)))
		return nil, err
	}

	if vertices < 1 {
		log.Println(RedColor(fmt.Sprintf("[!!] Liczba wierzchołków musi być większa od 0, podano: %d", vertices)))
		return nil, fmt.Errorf("wrong vertices number: %d", vertices)
	}

	G, err := graph.NewAdjacencyMatrix(vertices)

	for i := 0; i < vertices; i++ {
		line, err = rdr.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Println(RedColor(fmt.Sprintf("[!!] Błąd czytania linii '%v': %v", line, err)))
			return nil, err
		}

		edges := strings.Fields(strings.TrimSpace(line))
		for j, edge := range edges {
			weight, err := strconv.Atoi(edge)
			if err != nil {
				log.Println(RedColor(fmt.Sprintf("[!!] Błąd konwersji wagi krawędzi: %v", err)))
				return nil, err
			}
			if weight > 0 {
				err = G.PutEdge(i, j, weight)
				if err != nil {
					log.Println(RedColor(fmt.Sprintf("[!!] Błąd dodawania krawędzi: %v", err)))
					return nil, err
				}
			}
		}
	}

	return G, nil
}
