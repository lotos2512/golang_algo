package dijkstra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	A = Node{
		Name: "A",
	}
	B = Node{
		Name: "B",
	}
	C = Node{
		Name: "C",
	}
	D = Node{
		Name: "D",
	}
	E = Node{
		Name: "E",
	}
	G = Node{
		Name: "G",
	}
	F = Node{
		Name: "F",
	}
	M = Node{
		Name: "M",
	}
)

func init() {
	A.SetRelation(Relation{
		Node:  &B,
		Price: 5,
	})
	A.SetRelation(Relation{
		Node:  &C,
		Price: 2,
	})

	B.SetRelation(Relation{
		Node:  &A,
		Price: 5,
	})
	B.SetRelation(Relation{
		Node:  &C,
		Price: 2,
	})
	B.SetRelation(Relation{
		Node:  &E,
		Price: 10,
	})
	E.SetRelation(Relation{
		Node:  &B,
		Price: 10,
	})

	C.SetRelation(Relation{
		Node:  &A,
		Price: 2,
	})
	C.SetRelation(Relation{
		Node:  &B,
		Price: 2,
	})
	C.SetRelation(Relation{
		Node:  &D,
		Price: 8,
	})

	D.SetRelation(Relation{
		Node:  &C,
		Price: 8,
	})
	D.SetRelation(Relation{
		Node:  &E,
		Price: 10,
	})
	E.SetRelation(Relation{
		Node:  &D,
		Price: 10,
	})

	D.SetRelation(Relation{
		Node:  &F,
		Price: 1,
	})
	F.SetRelation(Relation{
		Node:  &D,
		Price: 1,
	})

	F.SetRelation(Relation{
		Node:  &E,
		Price: 1,
	})
	E.SetRelation(Relation{
		Node:  &F,
		Price: 1,
	})

	M.SetRelation(Relation{
		Node:  &A,
		Price: 10,
	})
	M.SetRelation(Relation{
		Node:  &B,
		Price: 10,
	})
	M.SetRelation(Relation{
		Node:  &D,
		Price: 20,
	})
}

func TestSearchBestWayAToE(t *testing.T) {
	way := SearchBestWay(&A, &E)
	assert.Equal(t, way, Result{
		Nodes: []*Node{&E, &F, &D, &C, &A},
		Price: 12,
	})
}

func TestSearchBestWayEToA(t *testing.T) {
	way := SearchBestWay(&A, &E)
	assert.Equal(t, way, Result{
		Nodes: []*Node{&E, &F, &D, &C, &A},
		Price: 12,
	})
}

func TestSearchBestWayEToC(t *testing.T) {
	way := SearchBestWay(&E, &C)
	assert.Equal(t, way, Result{
		Nodes: []*Node{&C, &D, &F, &E},
		Price: 10,
	})
}

func TestSearchBestWayCToB(t *testing.T) {
	way := SearchBestWay(&C, &B)
	assert.Equal(t, way, Result{
		Nodes: []*Node{&B, &C},
		Price: 2,
	})
}

func TestSearchBestWayAToG(t *testing.T) {
	way := SearchBestWay(&A, &G)
	assert.Equal(t, way, Result{})
}

func TestSearchBestWayMToE(t *testing.T) {
	way := SearchBestWay(&M, &E)
	assert.Equal(t, way, Result{
		Nodes: []*Node{&E, &B, &M},
		Price: 20,
	})
}

func TestSearchBestWay(t *testing.T) {
	A := &Node{Name: "A"}
	B := &Node{Name: "B"}
	C := &Node{Name: "C"}
	D := &Node{Name: "D"}

	A.SetRelation(Relation{Node: B, Price: 5})
	A.SetRelation(Relation{Node: C, Price: 7})
	B.SetRelation(Relation{Node: C, Price: 2})
	C.SetRelation(Relation{Node: D, Price: 3})

	tests := []struct {
		name     string
		start    *Node
		end      *Node
		expected Result
	}{
		{
			name:  "Path from A to D",
			start: A,
			end:   D,
			expected: Result{
				Nodes: []*Node{D, C, B, A},
				Price: 10,
			},
		},
		{
			name:  "Path from A to C",
			start: A,
			end:   C,
			expected: Result{
				Nodes: []*Node{C, B, A},
				Price: 7,
			},
		},
		{
			name:  "Path from A to A",
			start: A,
			end:   A,
			expected: Result{
				Nodes: []*Node{A},
				Price: 0,
			},
		},
		{
			name:  "Path from B to D",
			start: B,
			end:   D,
			expected: Result{
				Nodes: []*Node{D, C, B},
				Price: 5,
			},
		},
		{
			name:  "Path from C to A",
			start: C,
			end:   A,
			expected: Result{
				Nodes: nil, // Нет пути от C до A
				Price: 0,
			},
		},
		{
			name:  "No Path from D to A",
			start: D,
			end:   A,
			expected: Result{
				Price: 0,
			},
		},
	}

	for _, test := range tests {
		result := SearchBestWay(test.start, test.end)

		// Проверка длины пути и стоимости
		if result.Price != test.expected.Price {
			t.Errorf("[%s] Некорректная стоимость пути. Ожидаемый результат: %d, полученный результат: %d",
				test.name, test.expected.Price, result.Price)
		}

		// Проверка последовательности узлов
		if len(result.Nodes) != len(test.expected.Nodes) {
			t.Errorf("[%s] Некорректная последовательность узлов. Ожидаемый результат: %v, полученный результат: %v",
				test.name, test.expected.Nodes, result.Nodes)
		} else {
			for i, node := range result.Nodes {
				if node.Name != test.expected.Nodes[i].Name {
					t.Errorf("[%s] Некорректная последовательность узлов. Ожидаемый результат: %v, полученный результат: %v",
						test.name, test.expected.Nodes, result.Nodes)
					break
				}
			}
		}
	}
}
