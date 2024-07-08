package concurrency

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func websiteCheckerMock(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	actual := CheckWebsites(websiteCheckerMock, websites)

	assert.Equal(t, expected, actual)
}

func slowWebsiteCheckerStub(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := lo.RepeatBy(100, func(_ int) string {
		return "a url"
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteCheckerStub, urls)
	}
}
