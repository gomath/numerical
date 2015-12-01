package adapt

import (
	"fmt"
)

// Surrogate is an interpolant for a function.
type Surrogate struct {
	Inputs    uint      // Number of inputs
	Outputs   uint      // Number of outputs
	Level     uint      // Interpolation level
	Nodes     uint      // Number of nodes
	Active    []uint    // Number of active nodes at each iteration
	Indices   []uint64  // Indices of the nodes
	Surpluses []float64 // Hierarchical surpluses of the nodes
}

func newSurrogate(ni, no uint) *Surrogate {
	return &Surrogate{
		Inputs:    ni,
		Outputs:   no,
		Active:    make([]uint, 0),
		Indices:   make([]uint64, 0, ni),
		Surpluses: make([]float64, 0, no),
	}
}

func (self *Surrogate) push(indices []uint64, surpluses []float64) {
	na := uint(len(indices)) / self.Inputs

	self.Nodes += na
	self.Active = append(self.Active, na)
	self.Indices = append(self.Indices, indices...)
	self.Surpluses = append(self.Surpluses, surpluses...)

	for _, level := range levelize(indices, self.Inputs) {
		if self.Level < level {
			self.Level = level
		}
	}
}

// String returns human-friendly information about the interpolant.
func (self *Surrogate) String() string {
	return fmt.Sprintf("Surrogate{inputs: %d, outputs: %d, level: %d, nodes: %d}",
		self.Inputs, self.Outputs, self.Level, self.Nodes)
}
