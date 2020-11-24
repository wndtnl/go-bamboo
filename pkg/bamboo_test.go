package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewDefaultClient(t *testing.T) {
	bamboo, err := NewDefaultClient()
	assert.Nil(t, err, err)
	assert.NotNil(t, bamboo, "Could not initialize client")
}
