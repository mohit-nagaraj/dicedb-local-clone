package resp

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDel(t *testing.T) {
	conn := getLocalConnection()
	defer conn.Close()

	testCases := []struct {
		name     string
		commands []string
		expected []interface{}
	}{
		{
			name:     "DEL with set key",
			commands: []string{"SET k1 v1", "DEL k1", "GET k1"},
			expected: []interface{}{"OK", int64(1), "(nil)"},
		},
		{
			name:     "DEL with multiple keys",
			commands: []string{"SET k1 v1", "SET k2 v2", "DEL k1 k2", "GET k1", "GET k2"},
			expected: []interface{}{"OK", "OK", int64(2), "(nil)", "(nil)"},
		},
		{
			name:     "DEL with key not set",
			commands: []string{"GET k3", "DEL k3"},
			expected: []interface{}{"(nil)", int64(0)},
		},
		{
			name:     "DEL with no keys or arguments",
			commands: []string{"DEL"},
			expected: []interface{}{"ERR wrong number of arguments for 'del' command"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, cmd := range tc.commands {
				result := FireCommand(conn, cmd)
				t.Logf("Command: %s, Result: %v, Expected: %v", cmd, result, tc.expected[i])
				assert.Equal(t, tc.expected[i], result)
			}
		})
	}
}
