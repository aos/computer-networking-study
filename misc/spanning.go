// Spanning tree construction for multiple nodes
// This is a WIP, still sort of broken
package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Name     string
	Address  int
	Root     *Node
	Distance int
	Receives <-chan Message
	Sends    chan<- Message
}

type Message struct {
	Root     *Node
	Address  int
	Distance int
}

var NodeList = make(map[string]struct {
	n         Node
	Neighbors []string
})

// Send sends messages to other nodes
func (n *Node) Send() {
	for {
		neighbors := NodeList[n.Name].Neighbors
		message := Message{
			Address:  n.Address,
			Root:     n.Root,
			Distance: n.Distance + 1,
		}

		for _, neighborName := range neighbors {
			neighborNode := NodeList[neighborName].n
			neighborNode.Msgs <- message
		}

	}
}

// Receive processes messages received from other nodes to its channel
func (n *Node) Receive() {
	for {
		msg := <-n.Msgs
		if msg.Address < n.Address && msg.Root != n.Root {
			n.Root = msg.Root
			n.Distance = msg.Distance
			fmt.Printf(
				"%s -- updated root to: %s\tupdated distance to: %d\n",
				n.Name, n.Root.Name, n.Distance)
		}
	}
}

func init() {
	adjacencyList := map[string][]string{
		"A": []string{"C", "E"},
		"B": []string{"C", "F"},
		"C": []string{"A", "B", "D"},
		"D": []string{"C", "E", "F"},
		"E": []string{"A", "D"},
		"F": []string{"B", "D"},
	}

	for i, v := range reflect.ValueOf(adjacencyList).MapKeys() {
		name := v.String()
		msgChan := make(chan Message)
		node := Node{Name: name, Address: i, Distance: 0, Msgs: msgChan}
		node.Root = &node

		nodeInMap := NodeList[name]
		nodeInMap.n = node
		nodeInMap.Neighbors = adjacencyList[name]
		NodeList[name] = nodeInMap
	}
}

func main() {
	// Each node sends periodic updates to neighbors with:
	// Message: its address, address of root, distance to root
	for _, node := range NodeList {
		fmt.Printf("Node: %v, Address: %d, Distance: %d\n",
			node.n.Name, node.n.Address, node.n.Distance)
	}

	for {
		for _, node := range NodeList {
			go func(n *Node) {
				n.Send()
			}(&(node.n))

			go func(n *Node) {
				n.Receive()
			}(&(node.n))
		}
	}
}
