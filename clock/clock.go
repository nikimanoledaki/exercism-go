package clock

import "fmt"

// Clock struct has an hour and a minute field, where both are type integer.
type Clock struct {
	hour   int
	minute int
}

// New creates a new clock struct.
func New(hour, minute int) Clock {
	return Clock{hour, minute}.CorrectTime()
}

// String converts clock struct to type string.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds given minutes of type int to a Clock struct's minutes field.
func (c Clock) Add(minutes int) (clock Clock) {
	c.minute += minutes
	return New(c.hour, c.minute)
}

// Subtract subtracts given minutes of type int to a Clock struct's minutes field.
func (c Clock) Subtract(minutes int) (clock Clock) {
	c.minute -= minutes
	return New(c.hour, c.minute)
}

// CorrectTime adjusts the time according to the conventions for writing the time.
func (c Clock) CorrectTime() (clock Clock) {
	for c.minute < 0 {
		c.minute += 60
		c.hour--
	}

	for c.hour < 0 {
		c.hour += 24
	}

	c.hour += int(c.minute / 60)
	c.minute = c.minute % 60
	c.hour = c.hour % 24

	return Clock{c.hour, c.minute}
}
