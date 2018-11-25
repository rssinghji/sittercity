// Package payrollhelper contains helper functions for payroll scheduling by frequency or by day.
package payrollhelper

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"utilities"
)

// ParseHolidayMatch is to help in getting whether the current date is a public holiday or not
// func : ParseHolidayMatch
// inputs: holidaysList, date
// types: []string, time.Time
// output: true or false
func ParseHolidayMatch(holidaysList []string, date time.Time) bool {
	layout := utilities.DefaultDateLayout
	dayValue := strconv.Itoa(date.Day())
	if date.Day() < 10 {
		dayValue = "0" + dayValue
	}
	monthValue := strconv.Itoa(int(date.Month()))
	if int(date.Month()) < 10 {
		monthValue = "0" + monthValue
	}
	dateString := monthValue + "/" + dayValue + "/" + strconv.Itoa(date.Year())
	for index := 0; index < len(holidaysList); index++ {
		holiday, _ := time.Parse(layout, holidaysList[index])
		dateValue, _ := time.Parse(layout, dateString)
		if holiday == dateValue {
			return true
		}
	}
	return false
}

// PrintPayrollSchedule to print the payrollschedule without considering public holidays
// func : PrintPayrollSchedule
// inputs: newDate, year, day previousMonth
// types: time.Time, int, int, time.Month
// output: []string - list of date strings
func PrintPayrollSchedule(newDate time.Time, year, day int, previousMonth time.Month) []string {
	var result []string
	for {
		newDate = getNewDateBasedOnWeekendForPayByDay(newDate)
		if newDate.Year() != year {
			break
		}
		if newDate.Month() == previousMonth+1 {
			fmt.Println("\n\t", newDate.Weekday().String()+",", newDate.Month(), strconv.Itoa(newDate.Day())+",", newDate.Year())
			result = append(result, newDate.Weekday().String()+", "+newDate.Month().String()+" "+strconv.Itoa(newDate.Day())+", "+strconv.Itoa(newDate.Year()))
			newDate = time.Date(year, newDate.Month(), day, 00, 00, 00, 00, time.UTC)
			previousMonth++
			continue
		}
		fmt.Println("\n\t", newDate.Weekday().String()+",", newDate.Month(), strconv.Itoa(newDate.Day())+",", newDate.Year())
		result = append(result, newDate.Weekday().String()+", "+newDate.Month().String()+" "+strconv.Itoa(newDate.Day())+", "+strconv.Itoa(newDate.Year()))
		newDate = time.Date(year, newDate.Month()+1, day, 00, 00, 00, 00, time.UTC)
		previousMonth++
	}
	return result
}

// PrintPayrollScheduleWithPublicHolidays to print the payrollschedule considering public holidays
// func : PrintPayrollScheduleWithPublicHolidays
// inputs: newDate, year, day previousMonth, publicHolidays
// types: time.Time, int, int, time.Month, []string
// output: []string - list of date strings
func PrintPayrollScheduleWithPublicHolidays(newDate time.Time, year, day int,
	previousMonth time.Month, publicHolidays []string) []string {
	var result []string
	for {
		if newDate.Year() != year {
			break
		}
		isEnteredDateHoliday := ParseHolidayMatch(publicHolidays, newDate)
		if isEnteredDateHoliday {
			newDate = getNextPaymentDateOutOfHolidays(newDate, publicHolidays)
		} else {
			newDate = getNewDateBasedOnWeekendForPayByDayWithHolidays(newDate, publicHolidays)
		}
		if newDate.Month() == previousMonth+1 {
			fmt.Println("\n\t", newDate.Weekday().String()+",", newDate.Month(), strconv.Itoa(newDate.Day())+",", newDate.Year())
			result = append(result, newDate.Weekday().String()+", "+newDate.Month().String()+" "+strconv.Itoa(newDate.Day())+", "+strconv.Itoa(newDate.Year()))
			newDate = time.Date(year, newDate.Month(), day, 00, 00, 00, 00, time.UTC)
			previousMonth++
			continue
		}
		fmt.Println("\n\t", newDate.Weekday().String()+",", newDate.Month(), strconv.Itoa(newDate.Day())+",", newDate.Year())
		result = append(result, newDate.Weekday().String()+", "+newDate.Month().String()+" "+strconv.Itoa(newDate.Day())+", "+strconv.Itoa(newDate.Year()))
		newDate = time.Date(year, newDate.Month()+1, day, 00, 00, 00, 00, time.UTC)
		previousMonth++
	}
	return result
}

// GetPayFrequencyNumberAndDateValues returns the pay frequency week value and parsed date
// func : GetPayFrequencyNumberAndDateValues
// inputs: payFrequency, startingDate
// types: int, string
// output: int, time.Month, int, int - frequency, month, year, day
func GetPayFrequencyNumberAndDateValues(payFrequency, startingDate string) (int, time.Month, int, int) {
	payFrequencyStrings := strings.Split(payFrequency, " ")
	payFrequencyValue, convError := strconv.Atoi(payFrequencyStrings[0])
	if convError != nil {
		log.Println("Error in Pay Frequency Value. Available options are: \n 1 week\n 2 week\n 4 week and\n 13 week")
		// Return default values
		return 2, time.January, 15, 2018
	}
	layout := utilities.DefaultDateLayout
	dateValue, _ := time.Parse(layout, startingDate)
	return payFrequencyValue, dateValue.Month(), dateValue.Day(), dateValue.Year()
}

// GetPayrollScheduleOnFrequency prints payroll schedule based on payfreuqency without considering public holidays
// func : GetPayrollScheduleOnFrequency
// inputs: nextPaymentDate, year, daysOffset
// types: time.Time, int, int
// output: []string - list of date strings
func GetPayrollScheduleOnFrequency(nextPaymentDate time.Time, year, daysOffset int) []string {
	var result []string
	for {
		if nextPaymentDate.Year() != year {
			break
		}
		nextPaymentDate = getNewDateBasedOnWeekendForPayByFrequency(nextPaymentDate)
		fmt.Println("\n\t", nextPaymentDate.Weekday().String()+",", nextPaymentDate.Month(), strconv.Itoa(nextPaymentDate.Day())+",", nextPaymentDate.Year())
		result = append(result, nextPaymentDate.Weekday().String()+", "+nextPaymentDate.Month().String()+" "+strconv.Itoa(nextPaymentDate.Day())+", "+strconv.Itoa(nextPaymentDate.Year()))
		nextPaymentDate = nextPaymentDate.AddDate(0, 0, daysOffset)
	}
	return result
}

// GetPayrollScheduleOnFrequencyWithPublicHolidays prints payroll schedule based on payfreuqency with public holidays
// func : GetPayrollScheduleOnFrequency
// inputs: nextPaymentDate, year, daysOffset, publicHolidays
// types: time.Time, int, int, []string
// output: []string - list of date strings
func GetPayrollScheduleOnFrequencyWithPublicHolidays(nextPaymentDate time.Time, year, daysOffset int,
	publicHolidays []string) []string {
	var result []string
	for {
		if nextPaymentDate.Year() != year {
			break
		}
		isEnteredDateHoliday := ParseHolidayMatch(publicHolidays, nextPaymentDate)
		if isEnteredDateHoliday {
			nextPaymentDate = getNextPaymentDateOutOfHolidays(nextPaymentDate, publicHolidays)
		} else {
			nextPaymentDate = getNewDateBasedOnWeekendForPayByFrequencyWithHolidays(nextPaymentDate, publicHolidays)
		}
		fmt.Println("\n\t", nextPaymentDate.Weekday().String()+",", nextPaymentDate.Month(), strconv.Itoa(nextPaymentDate.Day())+",", nextPaymentDate.Year())
		result = append(result, nextPaymentDate.Weekday().String()+", "+nextPaymentDate.Month().String()+" "+strconv.Itoa(nextPaymentDate.Day())+", "+strconv.Itoa(nextPaymentDate.Year()))
		nextPaymentDate = nextPaymentDate.AddDate(0, 0, daysOffset)
	}
	return result
}

// getNextPaymentDate to return the date which is not a holiday
func getNextPaymentDate(nextPaymentDate time.Time, publicHolidays []string) time.Time {
	for {
		nextPaymentDate = nextPaymentDate.AddDate(0, 0, -1)
		if !ParseHolidayMatch(publicHolidays, nextPaymentDate) &&
			nextPaymentDate.Weekday() != time.Saturday && nextPaymentDate.Weekday() != time.Sunday {
			return nextPaymentDate
		}
	}
}

// getNextPaymentDateOutOfHolidays to return the date which is not a holiday considering Friday
func getNextPaymentDateOutOfHolidays(nextPaymentDate time.Time, publicHolidays []string) time.Time {
	if nextPaymentDate.Weekday() != time.Friday {
		for {
			nextPaymentDate = nextPaymentDate.AddDate(0, 0, -1)
			if nextPaymentDate.Weekday() == time.Friday {
				if ParseHolidayMatch(publicHolidays, nextPaymentDate) {
					nextPaymentDate = getNextPaymentDate(nextPaymentDate, publicHolidays)
					break
				}
				break
			}
		}
	} else {
		nextPaymentDate = getNextPaymentDate(nextPaymentDate, publicHolidays)
	}
	return nextPaymentDate
}

// getNewDateBasedOnWeekendForPayByDay to find date which is not a weekend
func getNewDateBasedOnWeekendForPayByDay(newDate time.Time) time.Time {
	if newDate.Weekday() == time.Saturday {
		newDate = newDate.AddDate(0, 0, -1)
	} else if newDate.Weekday() == time.Sunday {
		newDate = newDate.AddDate(0, 0, -2)
	}
	return newDate
}

// getNewDateBasedOnWeekendForPayByDayWithHolidays to find date which is not a weekend considering holidays
func getNewDateBasedOnWeekendForPayByDayWithHolidays(newDate time.Time, publicHolidays []string) time.Time {
	newDate = getNewDateBasedOnWeekendForPayByDay(newDate)
	isEnteredDateHoliday := ParseHolidayMatch(publicHolidays, newDate)
	if isEnteredDateHoliday {
		newDate = getNextPaymentDateOutOfHolidays(newDate, publicHolidays)
	}
	return newDate
}

// getWeekdayDate finds date which is not a weekend in the same month
func getWeekdayDate(nextPaymentDate time.Time, subOffset, addOffset int) time.Time {
	tempDate := nextPaymentDate.AddDate(0, 0, subOffset)
	if tempDate.Month() != nextPaymentDate.Month() {
		nextPaymentDate = nextPaymentDate.AddDate(0, 0, addOffset)
	} else {
		nextPaymentDate = tempDate
	}
	return nextPaymentDate
}

// getNewDateBasedOnWeekendForPayByFrequency finds date which is not a weekend in the same month
func getNewDateBasedOnWeekendForPayByFrequency(nextPaymentDate time.Time) time.Time {
	if nextPaymentDate.Weekday() == time.Saturday {
		nextPaymentDate = getWeekdayDate(nextPaymentDate, -1, 2)
	} else if nextPaymentDate.Weekday() == time.Sunday {
		nextPaymentDate = getWeekdayDate(nextPaymentDate, -2, 1)
	}
	return nextPaymentDate
}

// getNewDateBasedOnWeekendForPayByFrequencyWithHolidays finds next date which is in same month and not a holiday or weekend
func getNewDateBasedOnWeekendForPayByFrequencyWithHolidays(nextPaymentDate time.Time,
	publicHolidays []string) time.Time {
	nextPaymentDate = getNewDateBasedOnWeekendForPayByFrequency(nextPaymentDate)
	isEnteredDateHoliday := ParseHolidayMatch(publicHolidays, nextPaymentDate)
	if isEnteredDateHoliday {
		nextPaymentDate = getNextPaymentDateOutOfHolidays(nextPaymentDate, publicHolidays)
	}
	return nextPaymentDate
}
