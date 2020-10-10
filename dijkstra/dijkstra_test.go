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
