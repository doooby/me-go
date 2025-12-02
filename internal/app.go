package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func StrToInt64(input string) (int64, error) {
	num, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("failed to parse string '%s' to int64: %w", input, err)
	}

	return num, nil
}

const TimeLayout = "2006-01-02 15:04:05 -07:00"

func TimeToStr(t time.Time) string {
	return t.Format(TimeLayout)
}

func StrToTime(timeStr string) (time.Time, error) {
	return time.Parse(TimeLayout, timeStr)
}

const shorthandParseError = "failed to parse shorthand time '%s'"

func ParseShorthandTime(input string, source time.Time) (time.Time, error) {
	source.Truncate(time.Minute)
	var err error

	year := source.Year()
	month := int(source.Month())
	day := source.Day()
	hour := source.Hour()
	minute := source.Minute()
	_, tzOffset := source.Zone()

	splitIndex := strings.Index(input, ":")
	if splitIndex == -1 {
		return time.Time{}, fmt.Errorf(shorthandParseError, input)
	}
	dateStr := input[:splitIndex]
	timeStr := input[splitIndex+1:]

	if dateStr != "" {
		if len(dateStr) != 6 {
			return time.Time{}, fmt.Errorf(shorthandParseError, input)
		}
		year, err = strconv.Atoi(dateStr[0:2])
		if err != nil {
			return time.Time{}, fmt.Errorf(shorthandParseError, input)
		}
		year += 2000
		month, err = strconv.Atoi(dateStr[2:4])
		if err != nil {
			return time.Time{}, fmt.Errorf(shorthandParseError, input)
		}
		day, err = strconv.Atoi(dateStr[4:6])
		if err != nil {
			return time.Time{}, fmt.Errorf(shorthandParseError, input)
		}
	}

	if timeStr != "" {
		if len(timeStr) != 4 {
			return time.Time{}, fmt.Errorf(shorthandParseError, input)
		}
		hour, err = strconv.Atoi(timeStr[0:2])
		if err != nil {
			return time.Time{}, fmt.Errorf(shorthandParseError, input)
		}
		minute, err = strconv.Atoi(timeStr[2:4])
		if err != nil {
			return time.Time{}, fmt.Errorf(shorthandParseError, input)
		}
	}

	iso := fmt.Sprintf(
		"%04d-%02d-%02d %02d:%02d:%02d %s",
		year,
		month, // time.Month is implicitly convertible to an integer for %02d
		day,
		hour,
		minute,
		0,
		formatOffset(tzOffset),
	)

	resut, err := StrToTime(iso)
	if err != nil {
		return time.Time{}, fmt.Errorf(shorthandParseError, input)
	}
	// TODO this looks unnecessarry as the parser in the end doesn' parse invalid dates after all, like month=13 etc.
	// checkIso := TimeToStr(resut)
	// fmt.Fprintf(os.Stdout, "iso A : %s \n", iso)
	// fmt.Fprintf(os.Stdout, "iso B : %s \n", checkIso)

	return resut, nil
}

func formatOffset(sec int) string {
	if sec == 0 {
		return "+00:00"
	}

	absSec := sec
	sign := "+"
	if sec < 0 {
		absSec = -sec
		sign = "-"
	}

	hours := absSec / 3600
	minutes := (absSec % 3600) / 60
	return fmt.Sprintf("%s%02d:%02d", sign, hours, minutes)
}
