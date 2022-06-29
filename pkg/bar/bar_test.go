package bar

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"

func TestSimpleBar(t *testing.T) {
    assert.Equal(t, 1, 1)
    require.NoError(t, nil)
}
