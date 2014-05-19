package main
import (
	"fmt"
	"time"
)

func main() {
	for {
		t := time.Now()
		str := t.Format("2006/01/02 15:04:05")
		fmt.Print(str + "\r")
		time.Sleep(1 * time.Second)

	}

}
