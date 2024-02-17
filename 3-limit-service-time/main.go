//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID         int
	IsPremium  bool
	TimeUsedMs int64 // in ms
}

// FIXME: Can be adjusted for accuracy
var tickMs = 1 // in ms

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	var ticker = time.NewTicker(time.Duration(tickMs) * time.Millisecond)

	c := make(chan bool)
	go func() {
		process()
		c <- true
	}()
	for {
		select {
		case <-c:
			return true
		case <-time.After((10 * time.Second)):
			return false
		case <-ticker.C:
			u.TimeUsedMs += int64(tickMs)
			if u.TimeUsedMs > 10*1000 {
				return false
			}
		}
	}
}

func main() {
	RunMockServer()
}
