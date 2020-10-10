package route

import (
	"errors"
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

func (ug *UG) GetRoutes(from int, to int) (map[int][]int, error) {
	_, ok1 := ug.Nodes[from]
	_, ok2 := ug.Nodes[to]
	if !ok1 || !ok2 {
		return nil, errors.New("The node should be exited")
	}

        var isVisited []bool
        for i := 0; i < len(ug.Nodes); i++ {
                isVisited = append(isVisited, false)
        }

	var resM = make(map[int][]int)
	num := 0

	var dfs func(path []int, t int, visited []bool)
	dfs = func(path []int, t int, visited []bool) {
		if t == to {
			resM[num] = append(path, t)
			num++
			return
		}
		if visited[t] == true {
			return
		}
		for _, v := range ug.Nodes[t].MMap {
			var temp = make([]bool, len(visited))
			temp = visited
			temp[t] = true
			dfs(append(path, t), v, temp)
		}
	}
	dfs([]int{}, from, isVisited)
	return resM, nil
}
