package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	actual, err := dictionary.Search(word)

	assert.NoError(t, err)

	expected := definition

	assert.Equal(t, expected, actual)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search known word", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("serach unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assert.Equal(t, ErrUnknownWord, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("test", "this is another test")

		assertDefinition(t, dictionary, "test", "this is another test")
	})

	t.Run("add existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is another test"}

		err := dictionary.Add("test", "this is a different definition")

		assert.Equal(t, ErrWordExists, err)

		assertDefinition(t, dictionary, "test", "this is another test")
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("update existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assert.NoError(t, err)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assert.Equal(t, ErrWordDoesNotExist, err)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assert.Equal(t, ErrUnknownWord, err)
}
