package main

import (
	"fmt"
	"sort"
	"strings"
)

//Node is a node in the graph
type Node struct {
	value    string
	lockedBy map[string]bool
	count    int
}

var nodes map[string]*Node
var letterToInt map[string]int

func init() {
	letterToInt = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}
}

const baseNum = 60
const numWorkers = 5

func readInput() {
	rowSlice := strings.Split(inputStr, "\n")
	nodes = make(map[string]*Node)
	for _, row := range rowSlice {
		strvec := strings.Split(row, " ")
		fromVal := strvec[1]
		toVal := strvec[7]

		//make to node if it doesn't already exist, otherwise add lock
		if _, ok := nodes[toVal]; !ok {
			lockMap := make(map[string]bool)
			lockMap[fromVal] = true
			nodes[toVal] = &Node{value: toVal,
				lockedBy: lockMap,
				count:    baseNum + letterToInt[toVal],
			}
		} else {
			nodes[toVal].lockedBy[fromVal] = true
		}
		//Make from node if it doesn't already exist
		if _, ok := nodes[fromVal]; !ok {
			lockMap := make(map[string]bool)
			nodes[fromVal] = &Node{value: fromVal,
				lockedBy: lockMap,
				count:    baseNum + letterToInt[fromVal],
			}
		}

	}
}

func findUnlocked() []*Node {
	unlocked := []*Node{}
	for _, nodePtr := range nodes {
		//fmt.Printf("Length of nodePtr %s is %d\n", nodePtr.value, len(nodePtr.lockedBy))
		if len(nodePtr.lockedBy) == 0 {
			unlocked = append(unlocked, nodePtr)
		}
	}
	return unlocked
}

func resolveAndPrint() {
	unlocked := findUnlocked()

	timeVal := 0
	for len(unlocked) > 0 {
		sort.SliceStable(unlocked, func(i, j int) bool {
			return unlocked[i].value < unlocked[j].value
		})

		//Announce and free work that is finished
		for _, node := range unlocked {
			if node.count == 0 {
				fmt.Printf("Finished node %s\n", node.value)
				removeNode(node)
			}
		}
		unlocked = findUnlocked()
		fmt.Printf("At time %d we have workers working on ", timeVal)
		for i := 0; i < numWorkers; i++ {
			if len(unlocked) <= i {
				continue
			}
			nextNode := unlocked[i]
			nextNode.count--
			fmt.Printf(" %s", nextNode.value)
		}
		fmt.Printf("\n")
		timeVal++
	}
}

func removeNode(delNode *Node) {
	for _, n := range nodes {
		delete(n.lockedBy, delNode.value)
	}
	delete(nodes, delNode.value)
}

func main() {
	readInput()
	resolveAndPrint()
	fmt.Printf("\n\nDONE!\n")
}

/*
const inputStr = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`
*/
const inputStr = `Step A must be finished before step R can begin.
Step J must be finished before step B can begin.
Step D must be finished before step B can begin.
Step X must be finished before step Z can begin.
Step H must be finished before step M can begin.
Step B must be finished before step F can begin.
Step Q must be finished before step I can begin.
Step U must be finished before step O can begin.
Step T must be finished before step W can begin.
Step V must be finished before step S can begin.
Step N must be finished before step P can begin.
Step P must be finished before step O can begin.
Step E must be finished before step C can begin.
Step F must be finished before step O can begin.
Step G must be finished before step I can begin.
Step Y must be finished before step Z can begin.
Step M must be finished before step K can begin.
Step C must be finished before step W can begin.
Step L must be finished before step W can begin.
Step W must be finished before step S can begin.
Step Z must be finished before step O can begin.
Step K must be finished before step S can begin.
Step S must be finished before step R can begin.
Step R must be finished before step I can begin.
Step O must be finished before step I can begin.
Step A must be finished before step Q can begin.
Step Z must be finished before step R can begin.
Step T must be finished before step R can begin.
Step M must be finished before step O can begin.
Step Q must be finished before step Z can begin.
Step V must be finished before step C can begin.
Step Y must be finished before step W can begin.
Step N must be finished before step F can begin.
Step J must be finished before step D can begin.
Step D must be finished before step N can begin.
Step B must be finished before step M can begin.
Step P must be finished before step I can begin.
Step W must be finished before step Z can begin.
Step Q must be finished before step V can begin.
Step V must be finished before step K can begin.
Step B must be finished before step Z can begin.
Step M must be finished before step I can begin.
Step G must be finished before step C can begin.
Step K must be finished before step O can begin.
Step E must be finished before step O can begin.
Step C must be finished before step I can begin.
Step X must be finished before step G can begin.
Step B must be finished before step T can begin.
Step B must be finished before step I can begin.
Step E must be finished before step F can begin.
Step N must be finished before step K can begin.
Step D must be finished before step W can begin.
Step R must be finished before step O can begin.
Step V must be finished before step I can begin.
Step T must be finished before step O can begin.
Step B must be finished before step Q can begin.
Step T must be finished before step L can begin.
Step M must be finished before step C can begin.
Step A must be finished before step M can begin.
Step F must be finished before step L can begin.
Step X must be finished before step T can begin.
Step G must be finished before step K can begin.
Step C must be finished before step L can begin.
Step D must be finished before step Z can begin.
Step H must be finished before step L can begin.
Step P must be finished before step Z can begin.
Step A must be finished before step V can begin.
Step G must be finished before step R can begin.
Step E must be finished before step G can begin.
Step D must be finished before step P can begin.
Step X must be finished before step L can begin.
Step U must be finished before step C can begin.
Step Z must be finished before step K can begin.
Step E must be finished before step W can begin.
Step B must be finished before step Y can begin.
Step J must be finished before step I can begin.
Step U must be finished before step P can begin.
Step Y must be finished before step L can begin.
Step N must be finished before step L can begin.
Step L must be finished before step S can begin.
Step H must be finished before step P can begin.
Step P must be finished before step S can begin.
Step J must be finished before step S can begin.
Step J must be finished before step U can begin.
Step H must be finished before step T can begin.
Step L must be finished before step I can begin.
Step N must be finished before step Z can begin.
Step A must be finished before step G can begin.
Step H must be finished before step S can begin.
Step S must be finished before step I can begin.
Step H must be finished before step E can begin.
Step W must be finished before step R can begin.
Step B must be finished before step G can begin.
Step U must be finished before step Y can begin.
Step J must be finished before step G can begin.
Step M must be finished before step L can begin.
Step G must be finished before step Z can begin.
Step N must be finished before step W can begin.
Step D must be finished before step E can begin.
Step A must be finished before step W can begin.
Step G must be finished before step Y can begin.`
