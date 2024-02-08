//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, chanTweet chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}

		chanTweet <- tweet
	}
	close(chanTweet)
}

func consumer(chanTweet chan *Tweet) {
	for t := range chanTweet {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	chanTweet := make(chan *Tweet)

	// Producer
	go producer(stream, chanTweet)

	// Consumer
	consumer(chanTweet)

	fmt.Printf("Process took %s\n", time.Since(start))
}
