package util

import (
	"time"
	"github.com/jinzhu/now"
	"strings"
	"strconv"
)

func GetFromAndToDate(month,year int) (time.Time, time.Time) {
	i := time.January
	switch month {
	case 1 :
		i = time.January
	case 2 :
		i = time.February
	case 3 :
		i = time.March
	case 4 :
		i = time.April
	case 5 :
		i = time.May
	case 6 :
		i = time.June
	case 7 :
		i = time.July
	case 8 :
		i = time.August
	case 9 :
		i = time.September
	case 10 :
		i = time.October
	case 11 :
		i = time.November
	case 12 :
		i = time.December
	default:
		i = time.January
	}

	t := time.Date(year,i,1,0,0,0,0,time.UTC)

	return now.New(t).BeginningOfMonth(), now.New(t).EndOfMonth()
}

func StringToDate(date string) (time.Time,error) {
	dateArr := strings.Split(date,"-")
	day, err := strconv.Atoi(dateArr[0])
	if err != nil {

	}
	month, err := strconv.Atoi(dateArr[1])
	if err != nil {

	}
	year, err := strconv.Atoi(dateArr[2])
	if err != nil {

	}

	i := time.January
	switch month {
	case 1 :
		i = time.January
	case 2 :
		i = time.February
	case 3 :
		i = time.March
	case 4 :
		i = time.April
	case 5 :
		i = time.May
	case 6 :
		i = time.June
	case 7 :
		i = time.July
	case 8 :
		i = time.August
	case 9 :
		i = time.September
	case 10 :
		i = time.October
	case 11 :
		i = time.November
	case 12 :
		i = time.December
	default:
		i = time.January
	}

	return time.Date(year,i,day,0,0,0,0,time.UTC), nil

}