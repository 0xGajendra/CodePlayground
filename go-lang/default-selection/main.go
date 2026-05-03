//the default case in a select is run if no other case is ready

//use a default case to try a send or receive without blocking

// select {
// 	case i := <-c:
// 		//use i
// 	default:
// 		//receiving from c would block
// }


package main

import (
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Microsecond)
	}
	for {
		select {
		case <- tick:
			fmt.Printf("[%6s] tick.\n", elapsed())

		case <- boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())

		default:
			fmt.Printf("[%6s]      .\n", elapsed())
			time.Sleep(50*time.Millisecond)
		}
	}
}