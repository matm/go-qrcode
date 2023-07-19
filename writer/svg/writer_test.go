package svg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/svg"
)

func Test_SVG_Write(t *testing.T) {
	qr, err := qrcode.New("https://github.com/yeqown/go-qrcode")
	require.NoError(t, err)
	require.NotNil(t, qr)

	// write into svg file
	w, err := svg.New("testdata/test.svg")
	require.NoError(t, err)
	require.NotNil(t, w)
	defer w.Close()

	err = qr.Save(w)
	assert.NoError(t, err)
}
