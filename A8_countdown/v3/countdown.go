package v3

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type DefaultSleeper struct{}
type Sleeper interface {
	Sleep()
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
