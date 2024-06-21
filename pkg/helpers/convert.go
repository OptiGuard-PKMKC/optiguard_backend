package helpers

import (
	"strconv"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers/customtypes"
)

func StringToInt64(s string) (*int64, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &num, nil
}

func GetDaysOfWeek(start, end customtypes.Date) ([]int, error) {
	daysOfWeekMap := make(map[int]bool)
	var daysOfWeek []int
	for d := start.Time; !d.After(end.Time); d = d.AddDate(0, 0, 1) {
		day := int(d.Weekday())
		if day == 0 {
			day = 7
		}
		if !daysOfWeekMap[day] {
			daysOfWeekMap[day] = true
			daysOfWeek = append(daysOfWeek, day)
		}
	}

	return daysOfWeek, nil
}

func GetWorkYears(start customtypes.Date, end customtypes.Date) int {
	var year, month int
	year = end.Time.Year() - start.Time.Year()
	if year <= 1 {
		month = (int(end.Time.Month()) + 12) - int(start.Time.Month())
		if month <= 12 {
			year = 0
		}
	}

	return year
}
