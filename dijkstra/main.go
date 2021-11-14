package main

import (
	"fmt"
	"math"
)

func findLowestCostNode(costMap map[string]int, processedSet map[string]bool) string {
	lowestCost := math.MaxInt
	lowestCostNode := ""
	var cost int
	for k, v := range costMap {
		cost = v
		if cost < lowestCost && !processedSet[k] {
			lowestCost = cost
			lowestCostNode = k
		}
	}
	return lowestCostNode

}

func dijkstra() int {
	//map to store a graph - nodes/neigbours and its costs
	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["A"] = make(map[string]int)
	graph["B"] = make(map[string]int)
	graph["end"] = make(map[string]int)
	graph["start"]["A"] = 6
	graph["start"]["B"] = 2
	graph["A"]["end"] = 1
	graph["B"]["A"] = 3
	graph["B"]["end"] = 5
	graph["end"] = nil

	//map to store costs - it will be modified during algorithm run (costs of all currently visible nodes)
	costs := make(map[string]int)
	costs["A"] = 6
	costs["B"] = 2
	costs["end"] = math.MaxInt

	//parents of every visible node
	parents := make(map[string]string)
	parents["A"] = "start"
	parents["B"] = "start"
	parents["end"] = ""

	// current neighbour map
	neigbours := make(map[string]int)

	var node string
	var cost int
	var newCost int
	var nextHop string
	//set to store already processed nodes
	processed := make(map[string]bool)

	node = findLowestCostNode(costs, processed) // find node with lowest costs among unprocessed
	for node != "" {                            //do this until all the nodes will be processed
		cost = costs[node]
		neigbours = graph[node]
		for k, v := range neigbours { //iterate over every neighbour of current node
			newCost = cost + v
			if costs[k] > newCost { //if neighbour can be reached "faster" via current node
				costs[k] = newCost //refresh its cost
				parents[k] = node  //and current node become parent of the neghbour
			}
		}
		processed[node] = true                      // mark node as processed
		node = findLowestCostNode(costs, processed) // find next node to process
	}
	nextHop = "end"
	fmt.Println("Shortest route:")
	for nextHop != "" {
		if nextHop == "start" {
			fmt.Printf("( %v )\n\n", nextHop)
		} else {
			fmt.Printf("( %v )<---", nextHop)
		}
		nextHop = parents[nextHop]
	}
	return costs["end"]
}

func main() {
	fmt.Println("Shortest path length is", dijkstra())
}
