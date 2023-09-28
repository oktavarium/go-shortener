package base52

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encode(t *testing.T) {
	input := []byte("Privet!")
	output := encode(input)
	assert.Equal(t, "test", string(output))
}
