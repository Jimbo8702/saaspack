package util

import "time"

func DaysBetween(startDate, endDate time.Time) int {
	diff := endDate.Sub(startDate)
	days := int(diff.Hours() / 24)
	return days
}

func TotalPrice(length, price int) int {
	return length * price
}