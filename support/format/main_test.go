package format

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFormatFloat(t *testing.T) {

	tests := map[float64]string{
		0: "0.0000000",
		0.0000001: "0.0000001",
		1.0000001: "1.0000001",
		123: "123.0000000",
	}

	for f, s := range tests{
		assert.Equal(t, s, FormatFloat(f))
	}
}
