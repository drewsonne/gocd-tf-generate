package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshallLinkFieldFail(t *testing.T) {
	d := linkField{}
	err := unmarshallLinkField(d, "missing-field", nil)
	assert.EqualError(t, err, "'missing-field' was not present in `map[]`")
}
