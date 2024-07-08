package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("test a stack of integers", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		assert.True(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)

		assert.False(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)

		value, _ := myStackOfInts.Pop()

		assert.Equal(t, value, 456)

		value, _ = myStackOfInts.Pop()

		assert.Equal(t, value, 123)
		assert.True(t, myStackOfInts.IsEmpty())
	})

	t.Run("test a stack of strings", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		assert.True(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("123")

		assert.False(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("456")

		value, _ := myStackOfStrings.Pop()

		assert.Equal(t, value, "456")

		value, _ = myStackOfStrings.Pop()

		assert.Equal(t, value, "123")
		assert.True(t, myStackOfStrings.IsEmpty())
	})

	t.Run("ensure good stack DX", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		assert.Equal(t, firstNum+secondNum, 3)
	})
}
