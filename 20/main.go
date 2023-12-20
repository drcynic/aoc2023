package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	part1()
	part2()
}

type OpType int

const (
	bc OpType = iota
	ff
	con
)

type Module struct {
	op          OpType
	targets     []string
	inputValues map[string]bool
}

type Op struct {
	source string
	target string
	signal bool
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	states := make(map[string]*Module)
	for _, line := range lines {
		lineParts := strings.Split(line, " -> ")
		sourceParts := lineParts[0]
		targetParts := lineParts[1]
		targets := strings.Split(targetParts, ", ")
		op := bc
		source := sourceParts
		if sourceParts[0] == '%' {
			op = ff
			source = sourceParts[1:]
		} else if sourceParts[0] == '&' {
			op = con
			source = sourceParts[1:]
		}
		if e, ok := states[source]; ok {
			e.targets = append(e.targets, targets...)
		} else {
			states[source] = &Module{op, targets, make(map[string]bool)}
		}
	}
	for ki, vi := range states {
		for _, t := range vi.targets {
			targetMod := states[t]
			if targetMod == nil {
				continue
			}
			if targetMod.op == con {
				targetMod.inputValues[ki] = false
			}
		}
	}

	for i := 0; i < 1000; i++ {
		queue := make([]Op, 0)
		queue = append(queue, createOps(&states, Op{"button", "broadcaster", false})...)
		for len(queue) > 0 {
			op := queue[0]
			queue = queue[1:]
			queue = append(queue, createOps(&states, op)...)
		}
	}

	//fmt.Printf("signals: %v\n", signals)
	numLow, numHigh := 0, 0
	for _, s := range signals {
		if s {
			numHigh++
		} else {
			numLow++
		}
	}
	println("part1:", numLow*numHigh)

}

var signals = make([]bool, 0)

func createOps(states *map[string]*Module, op Op) []Op {
	//println(op.source, "-", op.signal, "->", op.target)
	moduleName := op.target
	input := op.signal
	signals = append(signals, input)
	mod := (*states)[moduleName]
	if mod == nil {
		return make([]Op, 0)
	}

	ops := make([]Op, 0)
	sendOut, outputSignal := false, false
	switch mod.op {
	case bc:
		sendOut, outputSignal = true, input
	case ff:
		if !input {
			mod.inputValues[moduleName] = !mod.inputValues[moduleName]
		}
		sendOut, outputSignal = input == false, mod.inputValues[moduleName]
	case con:
		mod.inputValues[op.source] = input
		all := true
		for _, v := range mod.inputValues {
			all = all && v
		}

		// part 2 hardcoded cycle detection
		if moduleName == "ff" && !all {
			println("ff diff to last:", buttonPresses-ffLast)
			ffLast = buttonPresses
		}
		if moduleName == "fk" && !all {
			println("fk diff to last:", buttonPresses-fkLast)
			fkLast = buttonPresses
		}
		if moduleName == "lh" && !all {
			println("lh diff to last:", buttonPresses-lhLast)
			lhLast = buttonPresses
		}
		if moduleName == "mm" && !all {
			println("mm diff to last:", buttonPresses-mmLast)
			mmLast = buttonPresses
		}

		sendOut, outputSignal = true, !all
	}

	if sendOut {
		for _, target := range mod.targets {
			ops = append(ops, Op{moduleName, target, outputSignal})
		}
	}
	return ops
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	states := make(map[string]*Module)
	for _, line := range lines {
		lineParts := strings.Split(line, " -> ")
		sourceParts := lineParts[0]
		targetParts := lineParts[1]
		targets := strings.Split(targetParts, ", ")
		op := bc
		source := sourceParts
		if sourceParts[0] == '%' {
			op = ff
			source = sourceParts[1:]
		} else if sourceParts[0] == '&' {
			op = con
			source = sourceParts[1:]
		}
		if e, ok := states[source]; ok {
			e.targets = append(e.targets, targets...)
		} else {
			states[source] = &Module{op, targets, make(map[string]bool)}
		}
	}

	nrInputs := make([]string, 0)
	for ki, vi := range states {
		for _, t := range vi.targets {
			if t == "nr" {
				nrInputs = append(nrInputs, ki)
			}
			targetMod := states[t]
			if targetMod == nil {
				continue
			}
			if targetMod.op == con {
				targetMod.inputValues[ki] = false
			}
		}
	}
	fmt.Printf("nr inputs: %v\n", nrInputs)
	// [ff fk lh mm] -> nr -> rx

	for i := 1; i < 5000; i++ {
		buttonPresses = i
		queue := make([]Op, 0)
		queue = append(queue, createOps(&states, Op{"button", "broadcaster", false})...)
		for len(queue) > 0 {
			op := queue[0]
			queue = queue[1:]
			queue = append(queue, createOps(&states, op)...)
		}
	}
	lcm := LCM(3797, 4079, 3761, 3919)
	println("part2:", lcm)
}

var buttonPresses = 0
var ffLast = 0
var fkLast = 0
var lhLast = 0
var mmLast = 0

// [ff fk lh mm] -> nr -> rx

// LCM taken from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
