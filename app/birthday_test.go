package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCSV(t *testing.T) {
	for _, tc := range []struct {
		input string
		err   string
		out   []Birthday
	}{
		{
			input: "john, 27, 11, john's birthday\njane,19,04,jane's birthday\nnodesc,1,1",
			out: []Birthday{
				Birthday{
					Name:        "john",
					Day:         27,
					Month:       11,
					Description: "john's birthday",
				},
				Birthday{
					Name:        "jane",
					Day:         19,
					Month:       4,
					Description: "jane's birthday",
				},
				Birthday{
					Name:  "nodesc",
					Day:   1,
					Month: 1,
				},
			},
		},
		{
			input: "",
			out:   []Birthday{},
		},
		{
			input: "john,a,b,c",
			err:   "record at line 1: invalid value for day: 'a'",
		},
		{
			input: "john,10,b,c",
			err:   "record at line 1: invalid value for month: 'b'",
		},
		{
			input: "john,10,13,c",
			err:   "record at line 1: invalid value for month: '13'",
		},
		{
			input: "john,10,0,c",
			err:   "record at line 1: invalid value for month: '0'",
		},
		{
			input: "john,0,10,c",
			err:   "record at line 1: invalid value for day: '0'",
		},
		{
			input: "john,32,10,c",
			err:   "record at line 1: invalid value for day: '32'",
		},
	} {
		reader := strings.NewReader(tc.input)
		out, err := fromCSV(reader)
		if len(tc.err) != 0 {
			assert.True(t, strings.Contains(fmt.Sprint(err), tc.err), fmt.Sprintf("%s should contain %s", err, tc.err))
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, tc.out, out)
	}
}
