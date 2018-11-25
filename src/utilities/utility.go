// Package utilities contains utility functions, structs and constants.
package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// PublicHolidays A structure to parse the public holiday's info from JSON file
type PublicHolidays struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

// HolidaysList A structure to parse the list of public holidays from JSON file
type HolidaysList struct {
	PublicHolidaysList []PublicHolidays `json:"public_holidays"`
}

// Constants to help in Flags Parsing and other payroll scheduling activities
const (
	YearFlag            = `year`
	DefaultYearValue    = 2018
	DayFlag             = `day`
	DefaultDayNumber    = 30
	PayFrequencyFlag    = `pay_frequency`
	DefaultPayFrequency = `2 week`
	StartingDateFlag    = `starting`
	DefaultStartingDay  = `01/15/2018`
	PublicHolidaysFlag  = `public_holidays`
	PublicHolidaysFile  = `a file path like - ./public_holidays_2018.json`
	DefaultFilePath     = ``
	IntFlagDesc         = `an int`
	StringFlagDesc      = `a string`
	ArgumentsOffset     = 1
	DefaultDateLayout   = "01/02/2006"
	DaysConstant        = 7
)

// GetPublicHolidays return the array of PublicHolidays structure, which contains
// all info about public holidays
func GetPublicHolidays(publicHolidaysFilePath string) []string {
	var publicHolidaysList HolidaysList
	var holidaysDates []string
	publicHolidaysFilePath = strings.TrimLeft(publicHolidaysFilePath, "=")
	jsonFile, openError := os.Open(publicHolidaysFilePath)
	if openError != nil {
		log.Println("Error in opening JSON file: ", openError)
		return holidaysDates
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonParserError := json.Unmarshal([]byte(byteValue), &publicHolidaysList)
	if jsonParserError != nil {
		log.Println("Error in parsing JSON file: ", jsonParserError)
		return holidaysDates
	}

	for index := 0; index < len(publicHolidaysList.PublicHolidaysList); index++ {
		holidaysDates = append(holidaysDates, publicHolidaysList.PublicHolidaysList[index].Date)
	}
	return holidaysDates
}
