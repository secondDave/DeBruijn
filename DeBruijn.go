package DeBruijin

import (
	"fmt"
	"strings"
	"sync"
)

type Node struct {
	value string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func NewNode(value string) Node {
	return Node{
		value: value,
	}
}

type DeBruijn struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

func NewBruijnGraph() DeBruijn {
	return DeBruijn{
		nodes: make([]*Node, 0),
		edges: make(map[Node][]*Node),
		lock:  sync.RWMutex{},
	}
}

// TODO: Check if the new nodes suffix can also be anothers prefix to add the edge in the other sense as well.
// TODO: Add the nodes that don't have any suffix or prefix matches to a hashmap and check if they have a match in following checks.
// TODO: Create more than 1 node edges and relations to structure this more like a puzzle with parallel checks,
func (d *DeBruijn) canAddNode(node *Node, completion func(bool)) {
	// Add starting node
	count := len(d.nodes)
	if count == 0 {
		completion(true)
	}
	// Check if a new node starts with a previous nodes suffix
	for index, currentNode := range d.nodes {
		suffix := string(currentNode.String()[len(currentNode.String())-1])
		if strings.HasPrefix(node.String(), suffix) {
			completion(true)
		}
		if count == index {
			completion(false)
		}
	}
}

// AddNode first checks if we can add a node and then adds it.
func (d *DeBruijn) AddNode(node *Node) {
	d.lock.Lock()
	canAdd := func(can bool) {
		if can {
			d.nodes = append(d.nodes, node)
		}
	}
	d.canAddNode(node, canAdd)
	d.lock.Unlock()
}

func (d *DeBruijn) AddEdge(node1, node2 *Node) {
	d.lock.Lock()
	d.edges[*node1] = append(d.edges[*node1], node2)
	d.edges[*node2] = append(d.edges[*node2], node1)
	d.lock.Unlock()
}

func (d *DeBruijn) String() string {
	d.lock.RLock()
	s := ""

	for _, node := range d.nodes {
		s += node.String() + " -> "
		near := d.edges[*node]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	d.lock.RUnlock()
	return s
}
