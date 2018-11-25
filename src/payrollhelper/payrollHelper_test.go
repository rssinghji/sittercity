// Test File for all major functionalities created in payrollHelper.go
package payrollhelper

import (
	"log"
	"strings"
	"testing"
	"time"
)

func Test_getNewDateBasedOnWeekendForPayByDaySaturday(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 24, 00, 00, 00, 00, time.UTC)
	result := getNewDateBasedOnWeekendForPayByDay(date)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByDaySaturday")
	}
}

func Test_getNewDateBasedOnWeekendForPayByDaySunday(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 25, 00, 00, 00, 00, time.UTC)
	result := getNewDateBasedOnWeekendForPayByDay(date)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByDaySunday")
	}
}

func Test_getNewDateBasedOnWeekendForPayByDaySaturdayWithHolidays(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 24, 00, 00, 00, 00, time.UTC)
	publicHolidays := []string{"11/22/2018", "11/23/2018"}
	result := getNewDateBasedOnWeekendForPayByDayWithHolidays(date, publicHolidays)
	if result.Day() != 21 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByDaySaturdayWithHolidays")
	}
}

func Test_getNewDateBasedOnWeekendForPayByDaySundayWithHolidays(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 25, 00, 00, 00, 00, time.UTC)
	publicHolidays := []string{"11/22/2018", "11/23/2018"}
	result := getNewDateBasedOnWeekendForPayByDayWithHolidays(date, publicHolidays)
	if result.Day() != 21 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByDaySundayWithHolidays")
	}
}

func Test_getNewDateBasedOnWeekendForPayByDaySundayWithHolidays1(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 25, 00, 00, 00, 00, time.UTC)
	publicHolidays := []string{"11/22/2018"}
	result := getNewDateBasedOnWeekendForPayByDayWithHolidays(date, publicHolidays)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByDaySundayWithHolidays")
	}
}

func Test_getWeekdayDateFromSunday(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 25, 00, 00, 00, 00, time.UTC)
	result := getWeekdayDate(date, -2, 1)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getWeekdayDateFromSunday")
	}
}

func Test_getWeekdayDateFromSaturday(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 24, 00, 00, 00, 00, time.UTC)
	result := getWeekdayDate(date, -1, 2)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getWeekdayDateFromSaturday")
	}
}

func Test_getWeekdayDateFromSundayMonthChange(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.December, 2, 00, 00, 00, 00, time.UTC)
	result := getWeekdayDate(date, -2, 1)
	if result.Day() != 3 && result.Month() != time.December && result.Year() != 2018 {
		log.Panic("Failed Test_getWeekdayDateFromSundayMonthChange")
	}
}

func Test_getWeekdayDateFromSaturdayMonthChange(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.December, 1, 00, 00, 00, 00, time.UTC)
	result := getWeekdayDate(date, -1, 2)
	if result.Day() != 3 && result.Month() != time.December && result.Year() != 2018 {
		log.Panic("Failed Test_getWeekdayDateFromSaturdayMonthChange")
	}
}

func Test_getNewDateBasedOnWeekendForPayByFrequencySaturday(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 24, 00, 00, 00, 00, time.UTC)
	result := getNewDateBasedOnWeekendForPayByFrequency(date)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByFrequencySaturday")
	}
}

func Test_getNewDateBasedOnWeekendForPayByFrequencySunday(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 25, 00, 00, 00, 00, time.UTC)
	result := getNewDateBasedOnWeekendForPayByFrequency(date)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByFrequencySunday")
	}
}

func Test_getNewDateBasedOnWeekendForPayByFrequencySaturdayWithHolidays(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 24, 00, 00, 00, 00, time.UTC)
	publicHolidays := []string{"11/22/2018", "11/23/2018"}
	result := getNewDateBasedOnWeekendForPayByFrequencyWithHolidays(date, publicHolidays)
	if result.Day() != 21 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByFrequencySaturdayWithHolidays")
	}
}

func Test_getNewDateBasedOnWeekendForPayByFrequencySundayWithHolidays(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 25, 00, 00, 00, 00, time.UTC)
	publicHolidays := []string{"11/22/2018", "11/23/2018"}
	result := getNewDateBasedOnWeekendForPayByFrequencyWithHolidays(date, publicHolidays)
	if result.Day() != 21 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByFrequencySundayWithHolidays")
	}
}

func Test_getNewDateBasedOnWeekendForPayByFrequencySundayWithHolidays1(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.November, 25, 00, 00, 00, 00, time.UTC)
	publicHolidays := []string{"11/22/2018"}
	result := getNewDateBasedOnWeekendForPayByFrequencyWithHolidays(date, publicHolidays)
	if result.Day() != 23 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNewDateBasedOnWeekendForPayByFrequencySundayWithHolidays1")
	}
}

func TestParseHolidayMatchTrueCase1(test *testing.T) {
	test.Parallel()
	holidayList := []string{"11/22/2018", "11/02/2018"}
	date := time.Date(2018, time.November, 2, 00, 00, 00, 00, time.UTC)
	result := ParseHolidayMatch(holidayList, date)
	if !result {
		log.Panic("Failed TestParseHolidayMatch")
	}
}

func TestParseHolidayMatchTrueCase2(test *testing.T) {
	test.Parallel()
	holidayList := []string{"11/22/2018", "09/07/2018"}
	date := time.Date(2018, time.September, 7, 00, 00, 00, 00, time.UTC)
	result := ParseHolidayMatch(holidayList, date)
	if !result {
		log.Panic("Failed TestParseHolidayMatch")
	}
}

func TestParseHolidayMatchFalseCase(test *testing.T) {
	test.Parallel()
	holidayList := []string{"11/22/2018", "11/23/2018", "12/25/2018", "07/04/2018"}
	date := time.Date(2018, time.November, 30, 00, 00, 00, 00, time.UTC)
	result := ParseHolidayMatch(holidayList, date)
	if result {
		log.Panic("Failed TestParseHolidayMatch")
	}
}

func Test_getNextPaymentDateOutOfHolidaysWeekdayIsNotFriday(test *testing.T) {
	test.Parallel()
	holidayList := []string{"11/22/2018", "11/23/2018", "12/25/2018", "07/04/2018"}
	date := time.Date(2018, time.November, 21, 00, 00, 00, 00, time.UTC)
	result := getNextPaymentDateOutOfHolidays(date, holidayList)
	if result.Day() != 16 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNextPaymentDateOutOfHolidaysWeekdayIsNotFriday")
	}
}

func Test_getNextPaymentDateOutOfHolidaysWeekdayIsFriday(test *testing.T) {
	test.Parallel()
	holidayList := []string{"12/25/2018", "07/04/2018"}
	date := time.Date(2018, time.November, 23, 00, 00, 00, 00, time.UTC)
	result := getNextPaymentDateOutOfHolidays(date, holidayList)
	if result.Day() != 22 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNextPaymentDateOutOfHolidaysWeekdayIsFriday")
	}
}

func Test_getNextPaymentDateOutOfHolidaysWeekdayIsFridayAfterPublicHoliday(test *testing.T) {
	test.Parallel()
	holidayList := []string{"11/22/2018", "12/25/2018", "07/04/2018"}
	date := time.Date(2018, time.November, 23, 00, 00, 00, 00, time.UTC)
	result := getNextPaymentDateOutOfHolidays(date, holidayList)
	if result.Day() != 21 && result.Month() != time.November && result.Year() != 2018 {
		log.Panic("Failed Test_getNextPaymentDateOutOfHolidaysWeekdayIsFridayAfterPublicHoliday")
	}
}

func TestGetPayFrequencyNumberAndDateValues(test *testing.T) {
	test.Parallel()
	frequency, month, day, year := GetPayFrequencyNumberAndDateValues("4 week", "02/17/2018")
	if frequency != 4 || month != time.February || day != 17 || year != 2018 {
		log.Panic("Failed TestGetPayFrequencyNumberAndDateValues")
	}
}

func TestGetPayFrequencyNumberAndDateValuesErrorCase(test *testing.T) {
	test.Parallel()
	frequency, month, day, year := GetPayFrequencyNumberAndDateValues("x week", "")
	if frequency != 2 || month != time.January || day != 15 || year != 2018 {
		log.Panic("Failed TestGetPayFrequencyNumberAndDateValues")
	}
}
func TestPrintPayrollScheduleWithoutPublicHolidays(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.January, 30, 00, 00, 00, 00, time.UTC)
	result := PrintPayrollSchedule(date, 2018, 30, time.January)
	tempResultArray := []string{
		"Tuesday, January 30, 2018",
		"Friday, March 2, 2018",
		"Friday, March 30, 2018",
		"Monday, April 30, 2018",
		"Wednesday, May 30, 2018",
		"Friday, June 29, 2018",
		"Monday, July 30, 2018",
		"Thursday, August 30, 2018",
		"Friday, September 28, 2018",
		"Tuesday, October 30, 2018",
		"Friday, November 30, 2018",
		"Friday, December 28, 2018",
	}

	for index := 0; index < len(result); index++ {
		if 0 != strings.Compare(result[index], tempResultArray[index]) {
			log.Panic("Failed TestPrintPayrollScheduleWithoutPublicHolidays for date: ", result[index])
		}
	}
}

func TestPrintPayrollScheduleWithPublicHolidays(test *testing.T) {
	test.Parallel()
	holidayList := []string{"11/22/2018", "11/23/2018", "12/25/2018", "07/06/2018", "11/30/2018",
		"11/22/2029", "12/21/2018", "11/26/2018"}
	date := time.Date(2018, time.January, 30, 00, 00, 00, 00, time.UTC)
	result := PrintPayrollScheduleWithPublicHolidays(date, 2018, 30, time.January, holidayList)
	tempResultArray := []string{
		"Tuesday, January 30, 2018",
		"Friday, March 2, 2018",
		"Friday, March 30, 2018",
		"Monday, April 30, 2018",
		"Wednesday, May 30, 2018",
		"Friday, June 29, 2018",
		"Monday, July 30, 2018",
		"Thursday, August 30, 2018",
		"Friday, September 28, 2018",
		"Tuesday, October 30, 2018",
		"Thursday, November 29, 2018",
		"Friday, December 28, 2018",
	}

	for index := 0; index < len(result); index++ {
		if 0 != strings.Compare(result[index], tempResultArray[index]) {
			log.Panic("Failed TestPrintPayrollScheduleWithPublicHolidays for date: ", result[index])
		}
	}
}

func TestGetPayrollScheduleOnFrequency(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.January, 15, 00, 00, 00, 00, time.UTC)
	result := GetPayrollScheduleOnFrequency(date, 2018, 14)
	tempResultArray := []string{
		"Monday, January 15, 2018",
		"Monday, January 29, 2018",
		"Monday, February 12, 2018",
		"Monday, February 26, 2018",
		"Monday, March 12, 2018",
		"Monday, March 26, 2018",
		"Monday, April 9, 2018",
		"Monday, April 23, 2018",
		"Monday, May 7, 2018",
		"Monday, May 21, 2018",
		"Monday, June 4, 2018",
		"Monday, June 18, 2018",
		"Monday, July 2, 2018",
		"Monday, July 16, 2018",
		"Monday, July 30, 2018",
		"Monday, August 13, 2018",
		"Monday, August 27, 2018",
		"Monday, September 10, 2018",
		"Monday, September 24, 2018",
		"Monday, October 8, 2018",
		"Monday, October 22, 2018",
		"Monday, November 5, 2018",
		"Monday, November 19, 2018",
		"Monday, December 3, 2018",
		"Monday, December 17, 2018",
		"Monday, December 31, 2018",
	}

	for index := 0; index < len(result); index++ {
		if 0 != strings.Compare(result[index], tempResultArray[index]) {
			log.Panic("Failed TestGetPayrollScheduleOnFrequency for date: ", result[index])
		}
	}
}

func TestGetPayrollScheduleOnFrequencyWithPublicHolidays(test *testing.T) {
	test.Parallel()
	date := time.Date(2018, time.January, 22, 00, 00, 00, 00, time.UTC)
	holidayList := []string{"11/22/2018", "11/23/2018", "12/25/2018", "07/06/2018", "11/30/2018",
		"11/22/2029", "12/21/2018", "11/26/2018"}
	result := GetPayrollScheduleOnFrequencyWithPublicHolidays(date, 2018, 14, holidayList)
	tempResultArray := []string{
		"Monday, January 22, 2018",
		"Monday, February 5, 2018",
		"Monday, February 19, 2018",
		"Monday, March 5, 2018",
		"Monday, March 19, 2018",
		"Monday, April 2, 2018",
		"Monday, April 16, 2018",
		"Monday, April 30, 2018",
		"Monday, May 14, 2018",
		"Monday, May 28, 2018",
		"Monday, June 11, 2018",
		"Monday, June 25, 2018",
		"Monday, July 9, 2018",
		"Monday, July 23, 2018",
		"Monday, August 6, 2018",
		"Monday, August 20, 2018",
		"Monday, September 3, 2018",
		"Monday, September 17, 2018",
		"Monday, October 1, 2018",
		"Monday, October 15, 2018",
		"Monday, October 29, 2018",
		"Monday, November 12, 2018",
		"Wednesday, November 21, 2018",
		"Wednesday, December 5, 2018",
		"Wednesday, December 19, 2018",
	}

	for index := 0; index < len(result); index++ {
		if 0 != strings.Compare(result[index], tempResultArray[index]) {
			log.Panic("Failed TestGetPayrollScheduleOnFrequencyWithPublicHolidays for date: ", result[index])
		}
	}
}
