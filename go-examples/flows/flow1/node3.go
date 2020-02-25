package flow1

import "github.com/longguikeji/arkfbp-go/node"

// Node3 ...
type Node3 struct {
	node.APINode
}

// ID ...
func (n *Node3) ID() string {
	return "Node3"
}

// Next ...
func (n *Node3) Next() string {
	return "Node4"
}
