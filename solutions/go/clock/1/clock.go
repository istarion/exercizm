package clock

import "fmt"

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	daytime := (h*60 + m) % (60 * 24)
	if daytime < 0 {
		daytime += (60 * 24)
	}
	return Clock{minutes: daytime}
}

func (c Clock) Add(m int) Clock {
	daytime := (c.minutes + m) % (60 * 24)
	if daytime < 0 {
		daytime += (60 * 24)
	}
	return Clock{minutes: daytime}
}

func (c Clock) Subtract(m int) Clock {
	daytime := (c.minutes - m) % (60 * 24)
	if daytime < 0 {
		daytime += (60 * 24)
	}
	return Clock{minutes: daytime}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
