//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

func accessGranted() {
	fmt.Println("Access granted")
}

func accessDenied() {
	fmt.Println("Access denied")
}

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func isAllowedAtAnyTime(role int) bool {
	if role == 10 || role == 20 {
		return true
	} else {
		return false
	}
}

func isWeekDay(day int) bool {
	if day <= 4 {
		return true
	} else {
		return false
	}
}

func isMondayWednesdayOrFriday(day int) bool {
	if day == 0 || day == 2 || day == 4 {
		return true
	} else {
		return false
	}
}

func main() {

	// The day and role. Change these to check if condition is working
	today, role := Monday, Contractor
	accessCondition := isAllowedAtAnyTime(role) || (role == Contractor && !isWeekDay(today)) || (role == Member && isWeekDay(today)) || (role == Guest && isMondayWednesdayOrFriday(today))

	if accessCondition {
		accessGranted()
	} else {
		accessDenied()
	}

}
