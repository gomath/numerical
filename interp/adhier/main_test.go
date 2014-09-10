package adhier

import (
	"fmt"
	"testing"

	"github.com/go-math/numan/basis/linhat"
	"github.com/go-math/numan/grid/newcot"
	"github.com/go-math/support/assert"
)

func TestConstructStep(t *testing.T) {
	algorithm := makeSelf(1, 1, fixtureStep.surrogate.level)

	surrogate := algorithm.Compute(step)

	assert.Equal(surrogate, fixtureStep.surrogate, t)
}

func TestEvaluateStep(t *testing.T) {
	algorithm := makeSelf(1, 1, 0)

	values := algorithm.Evaluate(fixtureStep.surrogate, fixtureStep.points)

	assert.Equal(values, fixtureStep.values, t)
}

func TestConstructCube(t *testing.T) {
	algorithm := makeSelf(2, 1, fixtureCube.surrogate.level)

	surrogate := algorithm.Compute(cube)

	assert.Equal(surrogate, fixtureCube.surrogate, t)
}

func TestConstructBox(t *testing.T) {
	algorithm := makeSelf(2, 3, fixtureBox.surrogate.level)

	surrogate := algorithm.Compute(box)

	assert.Equal(surrogate, fixtureBox.surrogate, t)
}

func TestEvaluateBox(t *testing.T) {
	algorithm := makeSelf(2, 3, 0)

	values := algorithm.Evaluate(fixtureBox.surrogate, fixtureBox.points)

	assert.AlmostEqual(values, fixtureBox.values, t)
}

func BenchmarkHat(b *testing.B) {
	algorithm := makeSelf(1, 1, 0)

	for i := 0; i < b.N; i++ {
		_ = algorithm.Compute(hat)
	}
}

func BenchmarkCube(b *testing.B) {
	algorithm := makeSelf(2, 1, 0)

	for i := 0; i < b.N; i++ {
		_ = algorithm.Compute(cube)
	}
}

func BenchmarkBox(b *testing.B) {
	algorithm := makeSelf(2, 3, 0)

	for i := 0; i < b.N; i++ {
		_ = algorithm.Compute(box)
	}
}

func BenchmarkMany(b *testing.B) {
	algorithm := makeSelf(2, 1000, 0)
	function := many(2, 1000)

	for i := 0; i < b.N; i++ {
		_ = algorithm.Compute(function)
	}
}

// A one-input-one-output scenario with a non-smooth function.
func ExampleSelf_step() {
	const (
		inputs  = 1
		outputs = 1
	)

	grid := newcot.New(inputs)
	basis := linhat.New(inputs)

	config := DefaultConfig
	config.MaxLevel = 19
	algorithm := New(grid, basis, config, outputs)

	surrogate := algorithm.Compute(step)
	fmt.Println(surrogate)

	// Output:
	// Surrogate{ inputs: 1, outputs: 1, levels: 19, nodes: 38 }
}

// A one-input-one-output scenario with a smooth function.
func ExampleSelf_hat() {
	const (
		inputs  = 1
		outputs = 1
	)

	grid := newcot.New(inputs)
	basis := linhat.New(inputs)

	config := DefaultConfig
	config.MaxLevel = 9
	algorithm := New(grid, basis, config, outputs)

	surrogate := algorithm.Compute(hat)
	fmt.Println(surrogate)

	// Output:
	// Surrogate{ inputs: 1, outputs: 1, levels: 9, nodes: 305 }
}

// A multiple-input-one-output scenario with a non-smooth function.
func ExampleSelf_cube() {
	const (
		inputs  = 2
		outputs = 1
	)

	grid := newcot.New(inputs)
	basis := linhat.New(inputs)

	config := DefaultConfig
	config.MaxLevel = 9
	algorithm := New(grid, basis, config, outputs)

	surrogate := algorithm.Compute(cube)
	fmt.Println(surrogate)

	// Output:
	// Surrogate{ inputs: 2, outputs: 1, levels: 9, nodes: 377 }
}

// A multiple-input-many-output scenario with a non-smooth function.
func ExampleSelf_many() {
	const (
		inputs  = 2
		outputs = 1000
	)

	grid := newcot.New(inputs)
	basis := linhat.New(inputs)

	algorithm := New(grid, basis, DefaultConfig, outputs)

	surrogate := algorithm.Compute(many(inputs, outputs))
	fmt.Println(surrogate)

	// Output:
	// Surrogate{ inputs: 2, outputs: 1000, levels: 9, nodes: 362 }
}

func makeSelf(ic, oc uint16, ml uint8) *Self {
	config := DefaultConfig

	if ml > 0 {
		config.MaxLevel = ml
	}

	return New(newcot.New(ic), linhat.New(ic), config, oc)
}
