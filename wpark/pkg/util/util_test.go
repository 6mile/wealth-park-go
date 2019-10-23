package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUtilFunctions(t *testing.T) {
	// Create ID.
	id := CreateID()
	require.NotEmpty(t, id, "should create an ID")

	// Make timestamp.
	ts := MakeTimestamp()
	require.True(t, ts > 10000000, "should create a millisecond timestamp")

	timeNow := time.Now()
	ts = ToTimestamp(timeNow)
	require.True(t, ts > 10000000, "should create a millisecond timestamp")

	// Dump and return JSON.
	test := map[string]interface{}{
		"foo": "bar",
		"baz": 1337,
	}
	require.Contains(t, GetJSON(test), "baz", "should serialize JSON correctly")
	require.Contains(t, GetPrettyJSON(test), "baz", "should serialize JSON correctly")

}
