package svg

type Option func(*option)

type option struct {
	// BlockPixel is the size of each block in pixel. default is 5.
	blockPixel int
}

func defaultOption() *option {
	return &option{
		blockPixel: 5,
	}
}

// WithBlockPixel set the size of each block in pixel.
func WithBlockPixel(pixel int) Option {
	return func(o *option) {
		o.blockPixel = pixel
	}
}
