package main

import (
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"io/ioutil"
	"log"
	"strings"
)

var connections = make(map[string]map[string]bool)
var nodes = make([]string, 0)

func main() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		//println(line)
		split := strings.Fields(line)
		comp := split[0][0:3]
		for i := 1; i < len(split); i++ {
			n2 := split[i]
			if _, ok := connections[comp]; !ok {
				connections[comp] = make(map[string]bool)
			}
			connections[comp][n2] = true
			if _, ok := connections[n2]; !ok {
				connections[n2] = make(map[string]bool)
			}
			connections[n2][comp] = true
		}
	}

	// generate neato graph to visually inspect the 3 connections to remove
	//drawGraph(connections)

	for k, _ := range connections {
		nodes = append(nodes, k)
	}
	numAllNodes := len(nodes)

	// remove cons and count
	//consToRemove := [][2]string{{"hfx", "pzl"}, {"bvb", "cmg"}, {"nvd", "jqt"}} // test input
	consToRemove := [][2]string{{"bmx", "zlv"}, {"xsl", "tpb"}, {"qpg", "lrd"}}

	for _, ctr := range consToRemove {
		delete(connections[ctr[0]], ctr[1])
		delete(connections[ctr[1]], ctr[0])
	}

	found := make(map[string]bool)
	collect(&found, consToRemove[0][0])

	numSet1 := len(found)
	println("set1:", len(found))

	numSet2 := numAllNodes - numSet1
	println("set2:", numSet2)

	println("Part 1:", numSet1*numSet2)
}

func collect(found *map[string]bool, n string) bool {
	if _, ok := (*found)[n]; ok {
		return false
	}
	(*found)[n] = true

	for c, _ := range connections[n] {
		collect(found, c)
	}
	return false
}

func drawGraph(connections map[string]map[string]bool) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	nodes := make(map[string]*cgraph.Node)
	used := make(map[string]bool)
	for k, _ := range connections {
		nodes[k], _ = graph.CreateNode(k)
	}
	for k, v := range connections {
		for k2, _ := range v {
			_, ok1 := used[k+k2]
			_, ok2 := used[k2+k]
			if ok1 || ok2 {
				continue
			}
			used[k+k2] = true
			n1 := nodes[k]
			n2 := nodes[k2]
			_, _ = graph.CreateEdge("", n1, n2)
		}
	}
	g.SetLayout(graphviz.NEATO)
	if err = g.RenderFilename(graph, graphviz.SVG, "graph.svg"); err != nil {
		log.Fatal(err)
	}
}
