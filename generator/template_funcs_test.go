package generator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTemplateFuncsSuccess(t *testing.T) {
	for _, test := range []struct {
		raw      []string
		expected string
	}{
		{raw: []string{}, expected: ""},
		{raw: []string{"hello", "world"}, expected: `"hello",
"world"`},
		{raw: []string{"$hello", "$world"}, expected: `"$$hello",
"$$world"`},
	} {
		actual, err := templateStringJoin(test.raw)
		assert.Equal(t, test.expected, actual)
		assert.NoError(t, err)
	}
}
