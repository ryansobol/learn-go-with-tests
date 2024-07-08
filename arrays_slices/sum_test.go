package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("sum a collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		assert.Equal(t, actual, expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum a collection of collections", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		actual := SumAll(numbers1, numbers2)
		expected := []int{3, 9}

		assert.Equal(t, actual, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum the tails of a collection of collections", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		actual := SumAllTails(numbers1, numbers2)
		expected := []int{2, 9}

		assert.Equal(t, actual, expected)
	})

	t.Run("safely sum the tails of a collection of empty slices", func(t *testing.T) {
		numbers1 := []int{}
		numbers2 := []int{3, 4, 5}

		actual := SumAllTails(numbers1, numbers2)
		expected := []int{0, 9}

		assert.Equal(t, actual, expected)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiply integers", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		actual := Reduce([]int{1, 2, 3}, multiply, 1)
		expected := 6

		assert.Equal(t, expected, actual)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}

		actual := "abc"
		expected := Reduce([]string{"a", "b", "c"}, concat, "")

		assert.Equal(t, expected, actual)
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	assert.Equal(t, 200.0, newBalanceFor(riya))
	assert.Equal(t, 0.0, newBalanceFor(chris))
	assert.Equal(t, 175.0, newBalanceFor(adil))
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		assert.True(t, found)
		assert.Equal(t, 2, firstEvenNumber)
	})

	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		person, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		assert.True(t, found)
		assert.Equal(t, Person{Name: "Chris James"}, person)
	})
}
