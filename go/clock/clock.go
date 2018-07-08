package clock

import (
	"fmt"
	"time"
)

//Clock - it's a clock yo
type Clock struct {
	hours   int
	minutes int
}

//New - creates a new clock
func New(hours int, minutes int) Clock {
	t := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	t = t.Add(time.Hour * time.Duration(hours))
	t = t.Add(time.Minute * time.Duration(minutes))

	return Clock{t.Hour(), t.Minute()}
}

//String - stringer for the clock
func (clock Clock) String() string {
	strHours := zeroify(clock.hours)
	strMinutes := zeroify(clock.minutes)
	return fmt.Sprintf("%s:%s", strHours, strMinutes)
}

//Add - Adds 'minutes' to the clock and returns a new clock
func (clock Clock) Add(minutes int) Clock {
	t := time.Date(0, 0, 0, clock.hours, clock.minutes, 0, 0, time.UTC)
	t = t.Add(time.Minute * time.Duration(minutes))
	return New(t.Hour(), t.Minute())
}

//Subtract - Removes 'minutes' from the clock and returns a new clock
func (clock Clock) Subtract(minutes int) Clock {
	t := time.Date(0, 0, 0, clock.hours, clock.minutes, 0, 0, time.UTC)
	t = t.Add(time.Minute * time.Duration(minutes*-1))
	return New(t.Hour(), t.Minute())
}

func zeroify(value int) string {

	var stringed string
	if value < 10 || value == 0 {
		stringed = fmt.Sprintf("0%d", value)
	} else {
		stringed = fmt.Sprintf("%d", value)
	}
	return stringed
}
