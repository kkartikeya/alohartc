package media

type VideoSource interface {
	Source

	Codec() string

	Width() int
	Height() int

	//GetFramerate() float32
	//GetBitrate() int

	//AdjustBitrate(bps int)
}
