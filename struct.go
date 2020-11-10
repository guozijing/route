package simpleRoute

import (
	"fmt"
	"errors"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Node struct {
	Key int
	MMap []int
}

type UG struct {
	Nodes map[int]*Node
}

type StructGraph struct {
	Graph []*StructMap `json:"graph"`
}

type StructMap struct {
	Index int `json:"index"`
	LinkTo []int `json:"linkTo"`
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

func (ug *UG) AddMap(from int, tos []int) error {
	if _, ok := ug.Nodes[from]; !ok {
		return errors.New("The node should be exited")
	}
	for _, to := range tos {
		if _, ok := ug.Nodes[to]; !ok {
			return errors.New("The node should be exited")
		}
		ug.Nodes[from].MMap = append(ug.Nodes[from].MMap, to)
	}
	return nil
}

func (ug *UG) AddMapFromFile(fileName string) error {
	var content []byte
	var err error
	content, err = file_get_contents(fileName)
	if err != nil {
		return err
	}

	var c StructGraph
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		return err
	}
	
	ug.AddNodes(len(c.Graph))

	for _, fromTos := range c.Graph {
		err := ug.AddMap(fromTos.Index, fromTos.LinkTo)
		if err != nil {
			return err
		}
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
		fmt.Println(path, t, to)
                if t == to {
                        resM[num] = append(path, t)
                        num++
                        return
                }
                if visited[t] == true {
                        return
                }
                for _, v := range ug.Nodes[t].MMap {
                        var temp []bool
                        for _, tf := range visited {
                                temp = append(temp, tf)
                        }
                        temp[t] = true
			fmt.Println(t, v)
                        dfs(append(path, t), v, temp)
                }
        }
	dfs([]int{}, from, isVisited)
	return resM, nil
}

func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
