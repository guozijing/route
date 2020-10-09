package route

import (
	"errors"
)

type Node struct {
	Key int64
	MMap []int64
	MValue []float64
}

type UG struct {
	Nodes map[int64]*Node
}

func NewUG() *UG {
	return &UG{
		Nodes: make(map[int64]*Node),
	}
}

func (ug *UG) AddNode(n *Node) error {
	if _, ok := ug.Nodes[n.Key]; ok {
		return errors.New("Can not add a exited node")
	}
	ug.Nodes[n.Key] = n
	return nil
}

func (ug *UG) AddNodes(n int64) {
	for i := (int64)0; i < n; i++ {
		ug.Nodes[i] = &Node{Key: i}
	}
}

func (ug *UG) AddMap(from int64, tos []int64, toValues []float64) {
	for _, to := range tos {
		ug.Nodes[from].MMap = append(ug.Nodes[from].MMap, to)
	}
	for _, toValue := range toValues {
		ug.Nodes[from].MValue = append(ug.Nodes[from].MValue, toValue)
	}
}
