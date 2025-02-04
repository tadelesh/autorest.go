// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package enumdiscgroup

// Cobra model
type Cobra struct {
	// CONSTANT; discriminator property
	// Field has constant value SnakeKindCobra, any specified value is ignored.
	Kind *SnakeKind

	// REQUIRED; Length of the snake
	Length *int32
}

// GetSnake implements the SnakeClassification interface for type Cobra.
func (c *Cobra) GetSnake() *Snake {
	return &Snake{
		Kind:   c.Kind,
		Length: c.Length,
	}
}

// Test extensible enum type for discriminator
type Dog struct {
	// REQUIRED; discriminator property
	Kind *DogKind

	// REQUIRED; Weight of the dog
	Weight *int32
}

// GetDog implements the DogClassification interface for type Dog.
func (d *Dog) GetDog() *Dog { return d }

// Golden dog model
type Golden struct {
	// CONSTANT; discriminator property
	// Field has constant value DogKindGolden, any specified value is ignored.
	Kind *DogKind

	// REQUIRED; Weight of the dog
	Weight *int32
}

// GetDog implements the DogClassification interface for type Golden.
func (g *Golden) GetDog() *Dog {
	return &Dog{
		Kind:   g.Kind,
		Weight: g.Weight,
	}
}

// Test fixed enum type for discriminator
type Snake struct {
	// REQUIRED; discriminator property
	Kind *SnakeKind

	// REQUIRED; Length of the snake
	Length *int32
}

// GetSnake implements the SnakeClassification interface for type Snake.
func (s *Snake) GetSnake() *Snake { return s }
