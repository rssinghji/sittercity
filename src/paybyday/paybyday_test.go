//Test cases file for the package paybyday
package paybyday

import (
	"log"
	"testing"
)

func TestSchedulePayrollByDayNoHolidays(test *testing.T) {
	test.Parallel()
	result := SchedulePayrollByDay(2018, 30, "")
	if result != nil {
		log.Panic("Some Error occured. Failed TestSchedulePayrollByDayNoHolidays")
	}
}

func TestSchedulePayrollByDayWithHolidays(test *testing.T) {
	test.Parallel()
	result := SchedulePayrollByDay(2018, 30, "../payroll_scheduler/public_holidays_2018.json")
	if result != nil {
		log.Panic("Some Error occured. Failed TestSchedulePayrollByDayWithHolidays")
	}
}
