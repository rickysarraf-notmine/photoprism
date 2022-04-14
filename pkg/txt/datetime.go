package txt

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Go regular expression tester: https://regoio.herokuapp.com/

var DateRegexp = regexp.MustCompile("\\D\\d{4}[\\-_]\\d{2}[\\-_]\\d{2,}")
var DatePathRegexp = regexp.MustCompile("\\D\\d{4}/\\d{1,2}/?\\d*")
var DateFilenameRegexp = regexp.MustCompile("\\d{4}\\d{2}\\d{2}")
var DateTimeRegexp = regexp.MustCompile("\\D\\d{4}[\\-_]\\d{2}[\\-_]\\d{2}.{1,4}\\d{2}\\D\\d{2}\\D\\d{2,}")
var DateIntRegexp = regexp.MustCompile("\\d{1,4}")
var YearRegexp = regexp.MustCompile("\\d{4,5}")
var IsDateRegexp = regexp.MustCompile("\\d{4}[\\-_]?\\d{2}[\\-_]?\\d{2}")
var IsDateTimeRegexp = regexp.MustCompile("\\d{4}[\\-_]?\\d{2}[\\-_]?\\d{2}.{1,4}\\d{2}\\D?\\d{2}\\D?\\d{2}")
var ExifDateTimeRegexp = regexp.MustCompile("((?P<year>\\d{4})|\\D{4})\\D((?P<month>\\d{2})|\\D{2})\\D((?P<day>\\d{2})|\\D{2})\\D((?P<h>\\d{2})|\\D{2})\\D((?P<m>\\d{2})|\\D{2})\\D((?P<s>\\d{2})|\\D{2})(\\.(?P<subsec>\\d+))?(?P<z>\\D)?(?P<zh>\\d{2})?\\D?(?P<zm>\\d{2})?")
var ExifDateTimeMatch = make(map[string]int)

func init() {
	names := ExifDateTimeRegexp.SubexpNames()
	for i := 0; i < len(names); i++ {
		if name := names[i]; name != "" {
			ExifDateTimeMatch[name] = i
		}
	}
}

var (
	YearMin = 1990
	YearMax = time.Now().Year() + 3
)

var (
	DefaultTime = time.Time{}
)

const (
	MonthMin = 1
	MonthMax = 12
	DayMin   = 1
	DayMax   = 31
	HourMin  = 0
	HourMax  = 24
	MinMin   = 0
	MinMax   = 59
	SecMin   = 0
	SecMax   = 59
)

// DateTime parses a string and returns a valid Exif timestamp if possible.
func DateTime(s, timeZone string) (t time.Time) {
	defer func() {
		if r := recover(); r != nil {
			// Panic? Return unknown time.
			t = time.Time{}
		}
	}()

	// Empty timestamp? Return unknown time.
	if s == "" {
		return time.Time{}
	}

	s = strings.TrimLeft(s, " ")

	// Timestamp too short?
	if len(s) < 4 {
		return time.Time{}
	} else if len(s) > 50 {
		// Clip to max length.
		s = s[:50]
	}

	// Pad short timestamp with whitespace at the end.
	s = fmt.Sprintf("%-19s", s)

	v := ExifDateTimeMatch
	m := ExifDateTimeRegexp.FindStringSubmatch(s)

	// Pattern doesn't match? Return unknown time.
	if len(m) == 0 {
		return time.Time{}
	}

	// Default to UTC.
	tz := time.UTC

	// Local time zone currently not supported (undefined).
	if timeZone == time.Local.String() {
		timeZone = ""
	}

	// Set time zone.
	loc, err := time.LoadLocation(timeZone)

	// Location found?
	if err == nil && timeZone != "" && tz != time.Local {
		tz = loc
		timeZone = tz.String()
	} else {
		timeZone = ""
	}

	// Does the timestamp contain a time zone offset?
	z := m[v["z"]]                     // Supported values, if not empty: Z, +, -
	zh := IntVal(m[v["zh"]], 0, 23, 0) // Hours.
	zm := IntVal(m[v["zm"]], 0, 59, 0) // Minutes.

	// Valid time zone offset found?
	if offset := (zh*60 + zm) * 60; offset > 0 && offset <= 86400 {
		// Offset timezone name example: UTC+03:30
		if z == "+" {
			// Positive offset relative to UTC.
			tz = time.FixedZone(fmt.Sprintf("UTC+%02d:%02d", zh, zm), offset)
		} else if z == "-" {
			// Negative offset relative to UTC.
			tz = time.FixedZone(fmt.Sprintf("UTC-%02d:%02d", zh, zm), -1*offset)
		}
	}

	var nsec int

	if subsec := m[v["subsec"]]; subsec != "" {
		nsec = Int(subsec + strings.Repeat("0", 9-len(subsec)))
	} else {
		nsec = 0
	}

	// Create rounded timestamp from parsed input values.
	t = time.Date(
		IntVal(m[v["year"]], 1, YearMax, time.Now().Year()),
		time.Month(IntVal(m[v["month"]], 1, 12, 1)),
		IntVal(m[v["day"]], 1, 31, 1),
		IntVal(m[v["h"]], 0, 23, 0),
		IntVal(m[v["m"]], 0, 59, 0),
		IntVal(m[v["s"]], 0, 59, 0),
		nsec,
		tz)

	if timeZone != "" && loc != nil && loc != tz {
		return t.In(loc)
	}

	return t
}

// DateFromFilePath returns a string as time or the zero time instant in case it can not be converted.
func DateFromFilePath(s string) (result time.Time) {
	defer func() {
		if r := recover(); r != nil {
			result = DefaultTime
		}
	}()

	if len(s) < 6 {
		return DefaultTime
	}

	if !strings.HasPrefix(s, "/") {
		s = "/" + s
	}

	b := []byte(s)
	handlers := []func([]byte) time.Time{
		DateTimeHandler,
		DateHandler,
		DatePathHandler,
		DateFilenameHandler,
	}

	for _, handler := range handlers {
		result = handler(b)
		if !result.IsZero() {
			break
		}
	}

	return result.UTC()
}

// DateTimeHandler checks whether the data contains a datetime like "2020-01-30_09-57-18"
func DateTimeHandler(b []byte) time.Time {
	found := DateTimeRegexp.Find(b)

	if len(found) == 0 {
		return DefaultTime
	}

	n := DateIntRegexp.FindAll(found, -1)

	if len(n) != 6 {
		return DefaultTime
	}

	year := Int(string(n[0]))
	month := Int(string(n[1]))
	day := Int(string(n[2]))
	hour := Int(string(n[3]))
	min := Int(string(n[4]))
	sec := Int(string(n[5]))

	if !IsValidDate(year, month, day) || !IsValidTime(hour, min, sec) {
		return DefaultTime
	}

	return time.Date(
		year,
		time.Month(month),
		day,
		hour,
		min,
		sec,
		0,
		time.UTC)
}

// DateHandler checks whether the data contains a date like "2020-01-30"
func DateHandler(b []byte) time.Time {
	found := DateRegexp.Find(b)

	if len(found) == 0 {
		return DefaultTime
	}

	n := DateIntRegexp.FindAll(found, -1)

	if len(n) != 3 {
		return DefaultTime
	}

	year := Int(string(n[0]))
	month := Int(string(n[1]))
	day := Int(string(n[2]))

	if !IsValidDate(year, month, day) {
		return DefaultTime
	}

	return time.Date(
		year,
		time.Month(month),
		day,
		0,
		0,
		0,
		0,
		time.UTC)
}

// DatePathHandler checks whether the data contains a date like "2020/01/03"
func DatePathHandler(b []byte) time.Time {
	found := DatePathRegexp.Find(b)

	if len(found) == 0 {
		return DefaultTime
	}

	n := DateIntRegexp.FindAll(found, -1)

	if len(n) < 2 || len(n) > 3 {
		return DefaultTime
	}

	year := Int(string(n[0]))
	month := Int(string(n[1]))

	if !IsValidYearAndMonth(year, month) {
		return DefaultTime
	}

	if len(n) == 2 {
		return time.Date(
			year,
			time.Month(month),
			1,
			0,
			0,
			0,
			0,
			time.UTC)
	} else if day := Int(string(n[2])); IsValidDay(day) {
		return time.Date(
			year,
			time.Month(month),
			day,
			0,
			0,
			0,
			0,
			time.UTC)
	}

	return DefaultTime
}

// DateFilenameHandler checks whether the data contains a date like "20200130"
func DateFilenameHandler(b []byte) time.Time {
	found := DateFilenameRegexp.Find(b)

	if len(found) == 0 {
		return DefaultTime
	}

	match := string(found)

	year := Int(match[0:4])
	month := Int(match[4:6])
	day := Int(match[6:8])

	if !IsValidDate(year, month, day) {
		return DefaultTime
	}

	return time.Date(
		year,
		time.Month(month),
		day,
		0,
		0,
		0,
		0,
		time.UTC)
}

// IsValidTime tests if the given time is plausible
func IsValidTime(hour int, min int, sec int) bool {
	return IsValidHour(hour) && IsValidMinute(min) && IsValidSecond(sec)
}

// IsValidHour tests if the given hour value is plausible
func IsValidHour(hour int) bool {
	return hour >= HourMin && hour <= HourMax
}

// IsValidMinute tests if the given minutes value is plausible
func IsValidMinute(minute int) bool {
	return minute >= MinMin && minute <= MinMax
}

// IsValidSecond tests if the given second value is plausible
func IsValidSecond(second int) bool {
	return second >= SecMin && second <= SecMax
}

// IsValidDate tests if the given date is plausible
func IsValidDate(year int, month int, day int) bool {
	return IsValidYearAndMonth(year, month) && IsValidDay(day)
}

// IsValidYearAndMonth tests if the given year and month are plausible
func IsValidYearAndMonth(year int, month int) bool {
	return IsValidYear(year) && IsValidMonth(month)
}

// IsValidYear tests if the given year is plausible
func IsValidYear(year int) bool {
	return year >= YearMin && year <= YearMax
}

// IsValidMonth tests if the given month is plausible
func IsValidMonth(month int) bool {
	return month >= MonthMin && month <= MonthMax
}

// IsValidDay tests if the given day is plausible
func IsValidDay(day int) bool {
	return day >= DayMin && day <= DayMax
}

// IsTime tests if the string looks like a date and/or time.
func IsTime(s string) bool {
	if s == "" {
		return false
	} else if m := IsDateRegexp.FindString(s); m == s {
		return true
	} else if m := IsDateTimeRegexp.FindString(s); m == s {
		return true
	}

	return Time(s) != DefaultTime
}

// Year tries to find a matching year for a given string e.g. from a file oder directory name.
func Year(s string) int {
	b := []byte(s)

	found := YearRegexp.FindAll(b, -1)

	for _, match := range found {
		year := Int(string(match))

		if year > YearMin && year < YearMax {
			return year
		}
	}

	return 0
}
