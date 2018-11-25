// Testing file for utility.go
package utilities

import (
	"log"
	"testing"
)

func TestGetPublicHolidays(test *testing.T) {
	test.Parallel()
	holidaysList := GetPublicHolidays("../payroll_scheduler/public_holidays_2018.json")
	if 0 == len(holidaysList) {
		log.Panic("\nTest Failed")
	}
}

func TestGetPublicHolidaysOpeningFailure(test *testing.T) {
	test.Parallel()
	holidaysList := GetPublicHolidays("")
	if 0 != len(holidaysList) {
		log.Panic("\nTest Failed")
	}
}

func TestGetPublicHolidaysWrongJSONFailure(test *testing.T) {
	test.Parallel()
	holidaysList := GetPublicHolidays("./testWrong.json")
	if 0 != len(holidaysList) {
		log.Panic("\nTest Failed")
	}
}
