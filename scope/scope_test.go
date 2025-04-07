package scope

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseScope(t *testing.T) {
	scope, err := ParseScope("read:user:query.id")
	require.NoError(t, err)

	assert.Equal(t, &Scope{
		Action:     "read",
		ObjectType: "user",
		ObjectID: ObjectID{
			In:  ObjectIDLocatorQuery,
			Key: "id",
		},
	}, scope)
}
