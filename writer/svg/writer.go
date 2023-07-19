package svg

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yeqown/go-qrcode/v2"
)

var _ qrcode.Writer = (*Writer)(nil)

// Writer is a writer to output QRCode into SVG format.
type Writer struct {
	*option

	file *os.File
}

func New(filename string, opts ...Option) (*Writer, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	do := defaultOption()
	for _, applyOpt := range opts {
		applyOpt(do)
	}

	return &Writer{
		option: do,
		file:   file,
	}, nil
}

func (w *Writer) Write(mat qrcode.Matrix) error {
	buf := bytes.NewBuffer(nil)

	width, height := mat.Width()*w.blockPixel, mat.Height()*w.blockPixel
	line := fmt.Sprintf("<svg width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\">\n", width, height)

	_, err := buf.WriteString(line)
	if err != nil {
		return err
	}
	mat.Iterate(qrcode.IterDirection_COLUMN, func(x, y int, v qrcode.QRValue) {
		if v.IsSet() {
			line := fmt.Sprintf("<rect x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\"/>\n", x*w.blockPixel, y*w.blockPixel, w.blockPixel, w.blockPixel)
			_, _ = buf.WriteString(line)
		}
	})
	_, err = buf.WriteString("</svg>")
	if err != nil {
		return err
	}

	if _, err = w.file.WriteString(buf.String()); err != nil {
		return err
	}

	return nil
}

func (w *Writer) Close() error {
	return w.file.Close()
}
