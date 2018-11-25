# Payroll Scheduler For Sittercity

## The intent of this code base is to calculate payroll schedule according to the following ***problem statement***

### Write a payroll scheduler that calculates the pay date for a salary that is paid from January through to December of any given year. The application must be able to calculate the pay date for each month based on the users input. The user can either supply the payment date as the day of the month, i.e. always on the 30th day; or a payment frequency, i.e. every 2 weeks

#### For Example

To get the payroll schedule for 2018, with the payment day on the 30th of each month, the executable signature must conform to this interface

    payroll_scheduler --year 2018 --day 30

To get the payroll schedule for 2018 on a bi-weekly schedule starting on the 15th January 2018, the executable signature must conform to this interface

    payroll_scheduler --pay_frequency "2 week" --starting "01/15/2018"

## Acceptance Criteria

---------------
The payroll scheduler must conform to the following criteria:

### PayBy Day

    ● The payment date and year must be supplied to the scheduler from the command line. 
    ● All payment dates must fall within the year being processed. 
    ● All payment dates must be made during a standard working week; e.g. Monday through Friday. 
    ● Any payment date that falls on a weekend should be made on the preceding Friday. 
    ● Any payment date that falls on a public holiday should be made on the preceding Friday, unless the public holiday is a Friday.
    ● If the public holiday is a Friday then the payment should be made on the first preceding day that is not a public holiday or weekend.
    ● Each payment date must be printed to the command line in the following format; Monday, January 1 2018 and include a trailing new line.

### Pay By Frequency

    ● There must be one payment only per calendar month. No calendar month should have less or more than one payment.
    ● If the payment date falls on a weekend, and the preceding available day is outside of the current month, then the payment should be made on the first available day after the current date.
    ● Payment frequency options are:
        1. "1 week"
        2. "2 week"
        3. "4 week"
        4. "13 week"

### Bonus! Pay by frequency could handle public holidays

A new requirement came in for the payroll by frequency to handle public holidays:
Additional AC:
Pay by frequency calculator can take a json file including public holidays with the following signature:

    payroll_scheduler --pay_frequency "2 week" --starting "01/15/2018" --public_holidays="./public_holidays_2018.json"

Any payment date that falls on a public holiday should be made on the preceding Friday, unless the public holiday is a Friday. If the public holiday is a Friday then the payment should be made on the first preceding day that is not a public holiday or weekend.

#### List of Public Holidays

``` JSON
{
    "public_holidays": [
    {
        "date": "11/22/2018",
        "name": "Thanksgiving"
    },
    {
        "date": "11/23/2018",
        "name": "Day after Thanksgiving"
    },
    {
        "date": "12/25/2018",
        "name": "Christmas Day"
    },
    {
        "date": "07/06/2018",
        "name": "Test Holiday"
    },
    {
        "date": "11/30/2018",
        "name": "Test Holiday"
    },
    {
        "date": "11/22/2029",
        "name": "Test Holiday"
    },
    {
        "date": "12/21/2018",
        "name": "Test Holiday"
    },
    {
        "date": "11/26/2018",
        "name": "Test Holiday"
    }
    ]
}
```

## Code Organization

---------------

The code is oranized as Sittercity/src/*. All the packages are mentioned with comments and test cases are included within respective packages. The major hierarchy of code is:

    --sittercity
        |--src
            |-- paybyday
                    |-- paybyday.go
                    |-- paybyday_test.go
            |-- paybyfrequency
                    |-- paybyfrequency.go
                    |-- paybyfrequency_test.go
            |-- payroll_scheduler
                    |-- main.go
                    |-- public_holidays_2018.json
            |-- payrollhelper
                    |-- payrollHelper.go
                    |-- payrollHelper_test.go
            |-- utilities
                    |-- utility.go
                    |-- utility_test.go
            |-- coverage.html
            |-- coverage.out
            |-- Readme.md

***payroll_scheduler*** is the package which contains the main function and package main. The execution starts here. Separate calls are made to **paybyday** and **paybyfrequency** exported functions based on flags passed. These packages in turn uses **payrollHelper** package to finally print the payroll schedule. All above mentioned packages uses the ***utilities*** package for constants and some utility functions.

### Working with this code

---------------

#### Prerequisites

    1. GO installed. Preferrably latest version.
    2. VS Code / Terminal / vim or any code explorer.

First get the code from github.com using the following command:

    go get github.com/rssinghji/sittercity/src/payroll_scheduler

***Make sure to add the code folder to your GOPATH***

Now change directory to **sittercity/src/payroll_scheduler**

Building the code:

    go build

This will create an exe/binary under *payroll_scheduler* package

Alternatively:

    go install

Will create the binary under bin folder if GOBIN has been properly set.

### Running the command

Once the binary is generated the code can be run as mentioned before from the terminal as:

    payroll_scheduler --year 2018 --day 25

Or

    payroll_scheduler --year 2018 --day 22 --public_holidays="./public_holidays_2018.json"

If you're running binary from bin folder, make sure you give the correct file path for JSON.

Another commands which can be used are:

    payroll_scheduler --pay_frequency "2 week" --starting "01/22/2018"

&

    payroll_scheduler --pay_frequency "2 week" --starting "01/22/2018" --public_holidays "./public_holidays_2018.json"

## Code Metrics

---------------

Change directory to sittercity\src and run command:

    go test ./... -coverprofile=coverage.out

It'll test all the packages with test files and generate a coverage.out file under src\

Now run command:

    go tool cover -html=coverage.out

This will generate an HTML report similar to coverage.html under src\ folder.

### Current Code Metrics are at 100% Code Coverage

Alternatively, if you just want overall picture of code coverage inside terminal, run:

    go tool cover -func=coverage.out

This will give you coverage of all files by their individual functions.

### Note: No third party packages/libraries have been used in this code. Hence, you should be able to run it with usual GO steps

Finally, if you need to run individual test cases, change directory to particular package or folder and run:

    go test

***Author: Ravneet Singh***
*********************************************************************************************************