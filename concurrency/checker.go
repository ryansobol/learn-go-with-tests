package concurrency

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// func urlsToChannel(wc WebsiteChecker, urls []string) <-chan result {
// 	results := make(chan result, 10)

// 	for _, url := range urls {
// 		go func(_url string) {
// 			results <- result{_url, wc(_url)}
// 		}(url)
// 	}

// 	return results
// }

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	channel := urlsToChannel(wc, urls)

// 	returning := make(map[string]bool)

// 	for i := 0; i < len(urls); i++ {
// 		result := <-channel
// 		returning[result.string] = result.bool
// 	}

// 	return returning
// }

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := lop.Map(urls, func(url string, _ int) result {
		// Parallelize invocations of wc WebsiteChecker as it may be slow
		return result{url, wc(url)}
	})

	return lo.Associate(results, func(result result) (string, bool) {
		return result.string, result.bool
	})
}
