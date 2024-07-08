package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	t.Run("greet with name", func(t *testing.T) {
		actual := Greet("Jane", "")
		expected := "Hello, Jane"

		assert.Equal(t, expected, actual)
	})

	t.Run("greet without name", func(t *testing.T) {
		actual := Greet("", "")
		expected := "Hello, world!"

		assert.Equal(t, expected, actual)
	})

	t.Run("greet in Spanish with name", func(t *testing.T) {
		actual := Greet("Elodie", "Spanish")
		expected := "Hola, Elodie!"

		assert.Equal(t, expected, actual)
	})

	t.Run("greet in French without name", func(t *testing.T) {
		actual := Greet("", "French")
		expected := "Bonjour, world!"

		assert.Equal(t, expected, actual)
	})
}
