package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddr(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	assert.Equal(t, actual, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
