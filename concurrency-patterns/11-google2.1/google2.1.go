package _0_google2_1

import (
	"fmt"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(1 * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) []Result {
	c := make(chan Result)

	// each search performs in a goroutine
	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Image(query)
	}()
	go func() {
		c <- Video(query)
	}()

	var results []Result

	timeout := time.After(50 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timeout")
			return results
		}
	}

	return results
}
