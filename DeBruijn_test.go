package DeBruijin

import (
	"testing"
)

func TestCreation(t *testing.T) {
	t.Run("Graph Creation", func(t *testing.T) {
		graph := NewBruijnGraph()
		edges := graph.edges
		if len(edges) != 0 {
			t.Fail()
		}
	})
	t.Run("Node Creation", func(t *testing.T) {
		testValue := "ACTG"
		node := NewNode(testValue)
		if testValue != node.value {
			t.Fail()
		}
	})
}
func TestAddition(t *testing.T) {
	t.Run("Graph Creation and Node Addition", func(t *testing.T) {
		graph := NewBruijnGraph()
		edges := graph.edges
		if len(edges) != 0 {
			t.Fail()
		}
		testValue := "ACTG"
		node := NewNode(testValue)
		if testValue != node.value {
			t.Fail()
		}
		graph.AddNode(&node)
		if graph.nodes[0] != &node {
			t.Fail()
		}
	})
	t.Run("Graph Creation and Edge Addition", func(t *testing.T) {
		graph := NewBruijnGraph()
		edges := graph.edges
		if len(edges) != 0 {
			t.Fail()
		}
		testValueA := "A"
		testValueB := "B"
		node := NewNode(testValueA)
		if testValueA != node.value {
			t.Fail()
		}
		graph.AddNode(&node)
		if graph.nodes[0] != &node {
			t.Errorf("Graph node %s not contained", node.value)
		}
		node1 := NewNode(testValueB)
		if testValueB != node1.value {
			t.Fail()
		}
		graph.AddNode(&node1)
		if graph.nodes[1] != &node1 {
			t.Errorf("Graph node %s not contained", node1.value)
		}
		graph.AddEdge(&node, &node1)
		testString := "A -> B \nB -> A \n"
		if graph.String() != testString {
			t.Logf("String is %s not %s, runes = %d, %d", graph.String(), testString, []rune(graph.String()), []rune(testString))
			t.Fail()
		}
	})
}
