package json

import (
	"strconv"
	"time"
)

// AppendTime formats the input time with the given format
// and appends the encoded string to the input byte slice.
func AppendTime(dst []byte, t time.Time, format string) []byte {
	if format == "" {
		return AppendInt64(dst, t.Unix())
	}
	return append(t.AppendFormat(append(dst, '"'), format), '"')
}

// AppendTimes converts the input times with the given format
// and appends the encoded string list to the input byte slice.
func AppendTimes(dst []byte, vals []time.Time, format string) []byte {
	if format == "" {
		return appendUnixTimes(dst, vals)
	}
	if len(vals) == 0 {
		return append(dst, '[', ']')
	}
	dst = append(dst, '[')
	dst = append(vals[0].AppendFormat(append(dst, '"'), format), '"')
	if len(vals) > 1 {
		for _, t := range vals[1:] {
			dst = append(t.AppendFormat(append(dst, ',', '"'), format), '"')
		}
	}
	dst = append(dst, ']')
	return dst
}

func appendUnixTimes(dst []byte, vals []time.Time) []byte {
	if len(vals) == 0 {
		return append(dst, '[', ']')
	}
	dst = append(dst, '[')
	dst = strconv.AppendInt(dst, vals[0].Unix(), 10)
	if len(vals) > 1 {
		for _, t := range vals[1:] {
			dst = strconv.AppendInt(dst, t.Unix(), 10)
		}
	}
	dst = append(dst, ']')
	return dst
}

// AppendDuration formats the input duration with the given unit & format
// and appends the encoded string to the input byte slice.
func AppendDuration(dst []byte, d time.Duration, unit time.Duration, useInt bool) []byte {
	if useInt {
		return strconv.AppendInt(dst, int64(d/unit), 10)
	}
	return AppendFloat64(dst, float64(d)/float64(unit))
}

// AppendDurations formats the input durations with the given unit & format
// and appends the encoded string list to the input byte slice.
func AppendDurations(dst []byte, vals []time.Duration, unit time.Duration, useInt bool) []byte {
	if len(vals) == 0 {
		return append(dst, '[', ']')
	}
	dst = append(dst, '[')
	dst = AppendDuration(dst, vals[0], unit, useInt)
	if len(vals) > 1 {
		for _, d := range vals[1:] {
			dst = AppendDuration(append(dst, ','), d, unit, useInt)
		}
	}
	dst = append(dst, ']')
	return dst
}
