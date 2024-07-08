package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Chris")

	actual := buffer.String()
	expected := "Hello, Chris"

	assert.Equal(t, expected, actual)
}
