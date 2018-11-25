//Test cases file for package paybyfrequency
package paybyfrequency

import (
	"log"
	"testing"
)

func TestSchedulePayByFrequencyNoHolidays(test *testing.T) {
	test.Parallel()
	result := SchedulePayByFrequency("2 week", "01/15/2018", "")
	if result != nil {
		log.Panic("Some Error occured. Failed TestSchedulePayByFrequencyNoHolidays")
	}
}

func TestSchedulePayByFrequencyNoHolidaysWrongInput1(test *testing.T) {
	test.Parallel()
	result := SchedulePayByFrequency("x week", "01/15/2018", "")
	if result != nil {
		log.Panic("Some Error occured. Failed TestSchedulePayByFrequencyNoHolidaysWrongInput1")
	}
}

func TestSchedulePayByFrequencyWithHolidays(test *testing.T) {
	test.Parallel()
	result := SchedulePayByFrequency("2 week", "01/15/2018", "../payroll_scheduler/public_holidays_2018.json")
	if result != nil {
		log.Panic("Some Error occured. Failed TestSchedulePayByFrequencyWithHolidays")
	}
}
