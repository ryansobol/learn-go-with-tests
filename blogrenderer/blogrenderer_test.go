package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
	"ryansobol.com/learn-go-with-tests/blogrenderer"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("convert a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		renderer, err := blogrenderer.NewPostRenderer()

		assert.NoError(t, err)

		err = renderer.Render(&buf, aPost)

		assert.NoError(t, err)

		actual := buf.String()

		approvals.VerifyString(t, actual)
	})

	t.Run("render an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		renderer, err := blogrenderer.NewPostRenderer()

		assert.NoError(t, err)

		err = renderer.RenderIndex(&buf, posts)

		assert.NoError(t, err)

		actual := buf.String()

		approvals.VerifyString(t, actual)
	})
}

func BenchmarkRender(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	renderer, err := blogrenderer.NewPostRenderer()

	assert.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		renderer.Render(io.Discard, aPost)
	}
}
