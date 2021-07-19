package bizerror

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSilence(t *testing.T) {
	for _, err := range []error{
		Silence(nil),
		Silence(errors.New("error")),
		Wrap(100, "100", Silence(errors.New("error"))),
	} {
		assert.True(t, IsSilence(err))
	}
}
