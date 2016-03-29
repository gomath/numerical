package hybrid

import (
	"testing"

	"github.com/ready-steady/adapt/algorithm/internal"
	"github.com/ready-steady/assert"
)

func TestBranin(t *testing.T) {
	fixture := &fixtureBranin
	interpolator := prepare(fixture)

	surrogate := interpolator.Compute(fixture.target)
	assert.Equal(internal.IsUnique(surrogate.Indices, surrogate.Inputs), true, t)
	assert.Equal(internal.IsAdmissible(surrogate.Indices, surrogate.Inputs), true, t)

	values := interpolator.Evaluate(surrogate, fixture.points)
	assert.EqualWithin(values, fixture.values, 100.0, t)
}
