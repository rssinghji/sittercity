// Package paybyday to take year, day and list of public holidays as input to schedule payroll.
package paybyday

import (
	"fmt"
	"payrollhelper"
	"time"
	"utilities"
)

// SchedulePayrollByDay Funtion To schedule payroll based on year and day
// func : SchedulePayrollByDay
// inputs: year, day, publicHolidaysFilePath
// types: int, int, string
// output: for now error as nil
func SchedulePayrollByDay(year int, day int, publicHolidaysFilePath string) error {
	fmt.Println("\n-------------------------------------------------------------------")
	fmt.Println("Calculating Payroll Schedule Based on Year and Day")
	if publicHolidaysFilePath == "" {
		printPayrollSchedule(year, day)
	} else {
		printPayrollScheduleConsideringPublicHolidays(year, day, publicHolidaysFilePath)
	}
	fmt.Println("-------------------------------------------------------------------")
	return nil
}

// printPayrollSchedule Funtion To schedule payroll based on year and day without holidays
func printPayrollSchedule(year, day int) {
	fmt.Println("\nPrinting Payroll Schedule")
	// Set Date to first payroll date for January of the year required
	newDate := time.Date(year, time.January, day, 00, 00, 00, 00, time.UTC)
	previousMonth := newDate.Month()
	_ = payrollhelper.PrintPayrollSchedule(newDate, year, day, previousMonth)
}

// printPayrollSchedule Funtion To schedule payroll based on year and day with holidays
func printPayrollScheduleConsideringPublicHolidays(year, day int, publicHolidaysFilePath string) {
	publicHolidays := utilities.GetPublicHolidays(publicHolidaysFilePath)
	fmt.Println("\n\nPrinting Payroll Schedule")
	// Set Date to first payroll date for January of the year required
	newDate := time.Date(year, time.January, day, 00, 00, 00, 00, time.UTC)
	previousMonth := newDate.Month()
	_ = payrollhelper.PrintPayrollScheduleWithPublicHolidays(newDate, year, day, previousMonth,
		publicHolidays)
}
