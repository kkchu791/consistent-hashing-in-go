package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

type Ring struct {
	positions []int
	nodes     map[int]string
	vnodes    int
}

func NewRing(vnodes int) *Ring {
	return &Ring{
		positions: []int{},
		nodes:     map[int]string{},
		vnodes:    vnodes,
	}
}

func hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}

func (r *Ring) AddNode(name string) {
	// // create 3 vnodes with has function
	// use fnv hash function
	// append them intot he slice
	// sort the slice
	for i := 0; i < r.vnodes; i++ {

		vnodename := fmt.Sprintf("%s-v%d", name, i)
		vnodehash := hash(vnodename)
		r.positions = append(r.positions, vnodehash)
		r.nodes[vnodehash] = name
	}

	sort.Ints(r.positions)
}

func (r *Ring) GetNode(key string) string {
	// hash the key to get a position
	hash := hash(key) //123456789

	// binary search r.positions for the first position >= hash
	idx := sort.SearchInts(r.positions, hash)

	// if you go past the end of slice, wrap around to index 0
	if idx == len(r.positions) {
		idx = 0
	}

	// look up that position in r.nodes and return the node name
	return r.nodes[r.positions[idx]]
}

func main() {
	r := NewRing(3)
	r.AddNode("node1")
	r.AddNode("node2")
	r.AddNode("node3")

	r.GetNode("lakers_score")
	r.GetNode("rockets_score")
	r.GetNode("celtics_score")
}
