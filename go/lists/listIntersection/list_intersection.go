package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// the basic node of SLL
type Node struct {
	val  string
	next *Node
}

func DoLinkedListsIntersect(nodes []*Node, test string) bool {
	tempMap := make(map[string]int)
	splits := strings.Split(test, ",")

	for _, val := range splits {
		//look for the node matching the value
		dp := SearchNodes(nodes, val)
		if dp == nil {
			continue //element not in list
		}

		//from this node go to the end of list and add all element to the map and set the value as 1
		for dp.next != nil {
			if tempMap[dp.val] == 1 {
				return true
			} else {
				tempMap[dp.val] = 1
			}
			dp = dp.next
		}
		//last node
		if tempMap[dp.val] == 1 {
			return true
		} else {
			tempMap[dp.val] = 1
		}

	}
	return false
}

func SearchNodes(nodes []*Node, val string) *Node {
	for _, node := range nodes {
		if node.val == val {
			return node
		}
	}
	return nil
}

func buildList(s, d string, allNodes []*Node) []*Node {
	dest := Node{
		val: d,
	}
	src := Node{
		val:  s,
		next: &dest,
	}
	sp := SearchNodes(allNodes, src.val)
	dp := SearchNodes(allNodes, dest.val)
	if sp == nil && dp == nil { //source and dest are new so a new list
		allNodes = append(allNodes, &src)
		allNodes = append(allNodes, &dest)
		return allNodes
	} else if sp != nil && dp == nil { //source is found in all node collection and dest is missing, append dest to the end of existing source
		sp.next = &dest
		allNodes = append(allNodes, &dest)
	} else if sp == nil && dp != nil { //source is not in allNodes collection but dest is present, make a link between source and existing node
		allNodes = append(allNodes, &src)
		src.next = dp
	} else { //both sourse and dest are found, just make a link between them
		sp.next = dp
		//if both source and dest are found once the link is made, check if from the dest we find the end of list of destination itself
		val := dp.val
		node := dp.next
		for node.next != nil {
			if node.val == val { //if we reach a node that has same value as dest its is cyclic
				fmt.Println("Error Thrown!")
				os.Exit(0)
			}
			node = node.next
		}
	}
	return allNodes
}

func main() {

	//collection to have all the node of the list
	var allNodes []*Node

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		text := scanner.Text()
		if strings.Contains(text, "->") { //nodes of list
			splits := strings.Split(text, "->")
			allNodes = buildList(splits[0], splits[1], allNodes)

		} else { //the string to test
			if DoLinkedListsIntersect(allNodes, text) {
				fmt.Println("True")
			} else {
				fmt.Println("False")
			}

		}
	}
}
