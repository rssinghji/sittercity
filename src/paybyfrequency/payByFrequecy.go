// Package paybyfrequency which takes pay frequency, starting date and list of public holidays to schedule payroll.
package paybyfrequency

import (
	"fmt"
	"payrollhelper"
	"time"
	"utilities"
)

// SchedulePayByFrequency To Schedule the payroll by frequency with a Starting Date
// func : SchedulePayByFrequency
// inputs: payFrequency, startingDate, publicHolidaysFilePath
// types: int, string, string
// output: for now error as nil
func SchedulePayByFrequency(payFrequency string, startingDate string, publicHolidaysFilePath string) error {
	fmt.Println("\n-------------------------------------------------------------------")
	fmt.Println("Calculating Payroll Schedule Based on Pay Frquency and Starting Date")
	if publicHolidaysFilePath == "" {
		printPayrollSchedule(payFrequency, startingDate)
	} else {
		printPayrollScheduleConsideringPublicHolidays(payFrequency, startingDate, publicHolidaysFilePath)
	}
	fmt.Println("\n-------------------------------------------------------------------")
	return nil
}

// printPayrollSchedule To Schedule the payroll by frequency with a Starting Date without holidays
func printPayrollSchedule(payFrequency, startingDate string) {
	fmt.Println("\nPrinting Payroll Schedule")
	payFrequencyValue, month, day, year := payrollhelper.GetPayFrequencyNumberAndDateValues(payFrequency,
		startingDate)
	daysOffset := payFrequencyValue * utilities.DaysConstant
	date := time.Date(year, month, day, 00, 00, 00, 00, time.UTC)
	_ = payrollhelper.GetPayrollScheduleOnFrequency(date, year, daysOffset)
}

// printPayrollScheduleConsideringPublicHolidays To Schedule the payroll by frequency with a Starting Date with holidays
func printPayrollScheduleConsideringPublicHolidays(payFrequency, startingDate,
	publicHolidaysFilePath string) {
	publicHolidays := utilities.GetPublicHolidays(publicHolidaysFilePath)
	fmt.Println("\nPrinting Payroll Schedule")
	payFrequencyValue, month, day, year := payrollhelper.GetPayFrequencyNumberAndDateValues(payFrequency,
		startingDate)
	daysOffset := payFrequencyValue * utilities.DaysConstant
	date := time.Date(year, month, day, 00, 00, 00, 00, time.UTC)
	_ = payrollhelper.GetPayrollScheduleOnFrequencyWithPublicHolidays(date, year, daysOffset, publicHolidays)
}
