package main

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := rectangle.Perimeter()
	expected := 40.0

	assert.Equal(t, expected, actual)
}

func TestArea(t *testing.T) {
	type Area struct {
		name     string
		shape    Shape
		expected float64
	}

	areas := []Area{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, expected: 72.0},
		{name: "Circle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	lo.ForEach(areas, func(area Area, _ int) {
		t.Run(area.name, func(t *testing.T) {
			actual := area.shape.Area()
			expected := area.expected

			assert.Equal(t, expected, actual)
		})
	})
}
