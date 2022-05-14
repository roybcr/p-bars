package bar

import (
	"fmt"
	"strings"
	"time"
)

const PROGRESS_WIDTH int = 40

func Run(duration int, c chan<- string, name string) {

	time_duration := time.Duration(duration / 100)
	p_bar := fmt.Sprintf("\r[%%-%vs] %s", PROGRESS_WIDTH, name)

	for i := 0; i < PROGRESS_WIDTH; i++ {

		c <- fmt.Sprintf(p_bar, strings.Repeat("=", i)+">")
		time.Sleep(time.Millisecond * time_duration)

	}

	c <- fmt.Sprintf(p_bar+" Loaded.", strings.Repeat("=", PROGRESS_WIDTH))
	close(c)
}
