package bar

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatsBytesComplete(t *testing.T) {

	tests := []struct {
		current  int64
		total    int64
		expected string
	}{
		{
			current:  0,
			total:    0,
			expected: "0B/0B",
		},
		{
			current:  1,
			total:    2,
			expected: "1B/2B",
		},
		{
			current:  1,
			total:    2000,
			expected: "1B/2.0kB",
		},
		{
			current:  1500,
			total:    2000000,
			expected: "1.5kB/2.0MB",
		},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			actual := strings.TrimSpace(StatsBytesComplete(test.current, test.total, 0, nil))
			assert.Equal(t, test.expected, actual)
		})
	}
}
