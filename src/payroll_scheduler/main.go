// Package main which starts the execution for payroll scheduling.
package main

import (
	"flag"
	"fmt"
	"os"
	"paybyday"
	"paybyfrequency"
	"strings"
	"utilities"
)

// main function to kick off scheduling
func main() {
	fmt.Println("\n\n***************** Welcome to Payroll Schedule Calculator **********************")
	fmt.Println("\nYou requested the Payroll Scheduling as per following details: ")

	// Declaring pointers to store the parsed flags
	payFrequencyPointer := flag.String(utilities.PayFrequencyFlag, utilities.DefaultPayFrequency, utilities.StringFlagDesc)
	startingDatePointer := flag.String(utilities.StartingDateFlag, utilities.DefaultStartingDay, utilities.StringFlagDesc)
	yearFlagPointer := flag.Int(utilities.YearFlag, utilities.DefaultYearValue, utilities.IntFlagDesc)
	dayFlagPointer := flag.Int(utilities.DayFlag, utilities.DefaultDayNumber, utilities.IntFlagDesc)
	publicHolidaysFilePath := flag.String(utilities.PublicHolidaysFlag, utilities.DefaultFilePath, utilities.PublicHolidaysFile)

	flag.Parse()
	if *publicHolidaysFilePath != "" {
		fmt.Println("Public Holidays File Path Selected: ", *publicHolidaysFilePath)
	}

	passedArgumentsFlags := os.Args[utilities.ArgumentsOffset:]
	argumentString := strings.Join(passedArgumentsFlags, "")

	if strings.Contains(argumentString, utilities.PayFrequencyFlag) {
		fmt.Println("Pay Frequecny Selected:", *payFrequencyPointer)
		fmt.Println("Starting Date Selected:", *startingDatePointer)
		paybyfrequency.SchedulePayByFrequency(*payFrequencyPointer, *startingDatePointer, *publicHolidaysFilePath)
	} else {
		fmt.Println("Year Selected:", *yearFlagPointer)
		fmt.Println("Day Selected:", *dayFlagPointer)
		paybyday.SchedulePayrollByDay(*yearFlagPointer, *dayFlagPointer, *publicHolidaysFilePath)
	}
}
