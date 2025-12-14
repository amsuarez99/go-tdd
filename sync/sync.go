package sync

type Counter struct {
	Count int
}

func (c *Counter) Inc() {
	c.Count++
}
