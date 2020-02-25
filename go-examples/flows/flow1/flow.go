package flow1

import (
	"github.com/longguikeji/arkfbp-go/flow"
)

// Flow ...
type Flow struct {
	flow.Flow
}

// New ...
func New() *Flow {
	f := Flow{
		Flow: flow.Flow{
			CreateGraph: createGraph,
		},
	}

	return &f
}
