package responders

type ResponseConfiguration struct {
	mTtlSecondsOverride *uint32
}

func (c ResponseConfiguration) Adjust(r Responder) {
	if c.mTtlSecondsOverride != nil {
		r.SetTtl(*c.mTtlSecondsOverride)
	}
}

func (c *ResponseConfiguration) SetTtlSecondsOverride(i uint32) {
	c.mTtlSecondsOverride = &i
}
