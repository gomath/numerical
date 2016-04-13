// Package local provides an algorithm for hierarchical interpolation with local
// adaptation.
package local

import (
	"github.com/ready-steady/adapt/algorithm/external"
	"github.com/ready-steady/adapt/algorithm/internal"
)

// Basis is an interpolation basis.
type Basis interface {
	internal.BasisComputer
	internal.BasisIntegrator
}

// Grid is an interpolation grid.
type Grid interface {
	internal.GridComputer
	internal.GridRefiner
}

// Interpolator is an instance of the algorithm.
type Interpolator struct {
	ni uint
	no uint

	grid  Grid
	basis Basis
}

// New creates an interpolator.
func New(inputs, outputs uint, grid Grid, basis Basis) *Interpolator {
	return &Interpolator{
		ni: inputs,
		no: outputs,

		grid:  grid,
		basis: basis,
	}
}

// Compute constructs an interpolant for a function.
func (self *Interpolator) Compute(target external.Target,
	strategy external.Strategy) *external.Surrogate {

	ni, no := self.ni, self.no
	surrogate := external.NewSurrogate(ni, no)
	state := strategy.First()
	for strategy.Check(state, surrogate) {
		state.Volumes = internal.Measure(self.basis, state.Indices, ni)
		state.Nodes = self.grid.Compute(state.Indices)
		state.Observations = external.Invoke(target, state.Nodes, ni, no)
		state.Predictions = internal.Approximate(self.basis, surrogate.Indices,
			surrogate.Surpluses, state.Nodes, ni, no)
		state.Surpluses = internal.Subtract(state.Observations, state.Predictions)
		state.Scores = score(strategy, state, ni, no)
		surrogate.Push(state.Indices, state.Surpluses, state.Volumes)
		state = strategy.Next(state, surrogate)
	}
	return surrogate
}

// Evaluate computes the values of an interpolant at a set of points.
func (self *Interpolator) Evaluate(surrogate *external.Surrogate, points []float64) []float64 {
	return internal.Approximate(self.basis, surrogate.Indices, surrogate.Surpluses, points,
		surrogate.Inputs, surrogate.Outputs)
}

func score(strategy external.Strategy, state *external.State, ni, no uint) []float64 {
	nn := uint(len(state.Indices)) / ni
	scores := make([]float64, nn)
	for i := uint(0); i < nn; i++ {
		scores[i] = strategy.Score(&external.Element{
			Index:   state.Indices[i*ni : (i+1)*ni],
			Volume:  state.Volumes[i],
			Value:   state.Observations[i*no : (i+1)*no],
			Surplus: state.Surpluses[i*no : (i+1)*no],
		})
	}
	return scores
}
