package route

import (
	"errors"
	"github.com/guozijing/stack"
)

type Node struct {
	Key int
	MMap []int
	MValue []float64
}

type UG struct {
	Nodes map[int]*Node
}

func NewUG() *UG {
	return &UG{
		Nodes: make(map[int]*Node),
	}
}

func (ug *UG) AddNode(n *Node) error {
	if _, ok := ug.Nodes[n.Key]; ok {
		return errors.New("Can not add a exited node")
	}
	ug.Nodes[n.Key] = n
	return nil
}

func (ug *UG) AddNodes(n int) {
	for i := int(0); i < n; i++ {
		ug.Nodes[i] = &Node{Key: i}
	}
}

func (ug *UG) AddMap(from int, tos []int, toValues []float64) error {
	if _, ok := ug.Nodes[from]; !ok {
		return errors.New("The node should be exited")
	}
	for _, to := range tos {
		if _, ok := ug.Nodes[to]; !ok {
			return errors.New("The node should be exited")
		}
		ug.Nodes[from].MMap = append(ug.Nodes[from].MMap, to)
	}
	for _, toValue := range toValues {
		ug.Nodes[from].MValue = append(ug.Nodes[from].MValue, toValue)
	}
	return nil
}

func (ug *UG) GetRoutes(from int, to int) (map[int][]int, []float64, error) {
	ok1 := ug.Nodes[from]
	ok2 := ug.Nodes[to]
	if !ok1 || !ok2 {
		return nil, nil, errors.New("The node should be exited")
	}

	var isVisited = make(map[int]bool)
	var temp []int
	var resM = make(map[int][]int)
	var resV []float64

	st := stack.New()
	st.Push(from)
	isVisited[from] = true

	for {
		if st.IsEmpty() == true {
			break
		}
	}
}
