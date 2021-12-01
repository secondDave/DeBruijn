package DeBruij

import (
	"fmt"
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

func NewGraph() DeBruijn {
	return DeBruijn{
		nodes: make([]*Node, 0),
		edges: make(map[Node][]*Node),
		lock:  sync.RWMutex{},
	}
}

func (d *DeBruijn) AddNode(node *Node) {
	d.lock.Lock()
	d.nodes = append(d.nodes, node)
	d.lock.Unlock()
}

func (d *DeBruijn) AddEdge(node1, node2 *Node) {
	d.lock.Lock()
	d.edges[*node1] = append(d.edges[*node1], node2)
	d.edges[*node2] = append(d.edges[*node2], node1)
	d.lock.Unlock()
}

func (d *DeBruijn) String() {
	d.lock.RLock()
	s := ""
	for i := 0; i < len(d.nodes); i++ {
		s += d.nodes[i].String() + " -> "
		near := d.edges[*d.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	d.lock.RUnlock()
}
