package flow1

import "github.com/longguikeji/arkfbp-go/node"

// TestLoopNode ...
type TestLoopNode struct {
	node.LoopNode
}

// ID ...
func (n *TestLoopNode) ID() string {
	return "TestLoopNode"
}
