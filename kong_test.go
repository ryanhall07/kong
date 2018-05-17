package kong

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func mustNew(t *testing.T, cli interface{}) *Kong {
	t.Helper()
	parser, err := New("", "", cli)
	require.NoError(t, err)
	return parser
}

func TestArgument(t *testing.T) {
	/*
		app user create <id> <first> <last>
		app	user <id> delete
		app	user <id> rename <to>

	*/
	var cli struct {
		Create struct {
			Id    string `arg:"true"`
			First string `arg:"true"`
			Last  string `arg:"true"`
		}

		// Branching argument.
		Id struct {
			Id     int `arg:"true"`
			Flag   int
			Delete struct{}
			Rename struct {
				To string
			}
		} `arg:"true"`
	}
	p := mustNew(t, &cli)
	cmd, err := p.Parse([]string{"10", "delete"})
	require.NoError(t, err)
	require.Equal(t, 10, cli.Id.Id)
	require.Equal(t, "<id> delete", cmd)
}

func TestResetWithDefaults(t *testing.T) {
	var cli struct {
		Flag            string
		FlagWithDefault string `default:"default"`
	}
	cli.Flag = "BLAH"
	cli.FlagWithDefault = "BLAH"
	parser := mustNew(t, &cli)
	_, err := parser.Parse([]string{})
	require.NoError(t, err)
	require.Equal(t, "", cli.Flag)
	require.Equal(t, "default", cli.FlagWithDefault)
}

func TestSlice(t *testing.T) {
	var cli struct {
		Slice []int
	}
	parser := mustNew(t, &cli)
	_, err := parser.Parse([]string{"--slice=1,2,3"})
	require.NoError(t, err)
	require.Equal(t, []int{1, 2, 3}, cli.Slice)
}