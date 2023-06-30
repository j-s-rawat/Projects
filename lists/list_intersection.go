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
		idx := SearchNodes(nodes, val)
		if idx == -1 {
			continue //element not in list
		}

		node := nodes[idx]
		//from this node go to the end of list and add all element to the map and set the value as 1
		for node.next != nil {
			if tempMap[node.val] == 1 {
				return true
			} else {
				tempMap[node.val] = 1
			}
			node = node.next
		}
		//last node
		if tempMap[node.val] == 1 {
			return true
		} else {
			tempMap[node.val] = 1
		}

	}
	return false
}

// Note: this function can be improved to return the node and not the index
func SearchNodes(nodes []*Node, val string) (index int) {
	for index, node := range nodes {
		if node.val == val {
			return index
		}
	}
	return -1
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
			dest := Node{
				val: splits[1],
			}
			src := Node{
				val:  splits[0],
				next: &dest,
			}
			iSrc := SearchNodes(allNodes, src.val)
			iDest := SearchNodes(allNodes, dest.val)
			if iSrc < 0 && iDest < 0 { //source and dest are new so a new list
				allNodes = append(allNodes, &src)
				allNodes = append(allNodes, &dest)
				continue
			} else if iSrc >= 0 && iDest < 0 { //source is found in all node collection and dest is missing, append dest to the end of existing source
				allNodes[iSrc].next = &dest
				allNodes = append(allNodes, &dest)
			} else if iSrc < 0 && iDest >= 0 { //source is not in allNodes collection but dest is present, make a link between source and existing node
				allNodes = append(allNodes, &src)
				src.next = allNodes[iDest]
			} else { //both sourse and dest are found, just make a link between them
				allNodes[iSrc].next = allNodes[iDest]
				//if both source and dest are found once the link is made, check if from the dest we find the end of list of destination itself
				val := allNodes[iDest].val
				node := allNodes[iDest].next
				for node.next != nil {
					if node.val == val { //if we reach a node that has same value as dest its is cyclic
						fmt.Println("Error Thrown!")
						os.Exit(0)
					}
					node = node.next
				}

			}
		} else { //the string to test
			if DoLinkedListsIntersect(allNodes, text) {
				fmt.Println("True")
			} else {
				fmt.Println("False")
			}

		}
	}
}
