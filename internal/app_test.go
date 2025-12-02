package app

import (
	"fmt"
	"testing"
	"time"
)

func makeTime(year, month, day, hour, minute, tz int) time.Time {
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.FixedZone("", tz*3600))
}

func TestParseShortHandTime(t *testing.T) {

	var tests = []struct {
		source   time.Time
		input    string
		wantErr  bool
		wantTime time.Time
	}{
		{source: makeTime(2000, 1, 1, 0, 0, 0), input: "220101:1122", wantTime: makeTime(2022, 1, 1, 11, 22, 0)},
		{source: makeTime(2000, 1, 1, 0, 0, 0), input: "221301:1122", wantTime: time.Time{}, wantErr: true},
		{source: makeTime(2000, 1, 1, 0, 0, -7), input: "200101:", wantTime: makeTime(2020, 1, 1, 0, 0, -7)},
		{source: makeTime(2000, 1, 1, 0, 0, 7), input: ":1015", wantTime: makeTime(2000, 1, 1, 10, 15, 7)},
	}

	for _, testCase := range tests {
		name := fmt.Sprintf("'%s' at %s", testCase.input, TimeToStr(testCase.source))
		t.Run(name, func(t *testing.T) {
			result, err := ParseShorthandTime(testCase.input, testCase.source)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("input %v produced no error", testCase.input)
				}
			} else {
				if err != nil {
					t.Errorf("%v", err)
				} else if !result.Equal(testCase.wantTime) {
					t.Errorf("result %v not matching expected: %v", result, testCase.wantTime)
				}
			}
		})
	}
}
