package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

// UnmarshalTimeString convert time string to local time
func UnmarshalTimeString(timeString string) (error, LocalTime) {
	localTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	return err, LocalTime(localTime)
}

// MarshalTimeString convert local time to string
func MarshalTimeString(localTime LocalTime) string {
	timeString := fmt.Sprintf("%v", time.Time(localTime).Format("2006-01-02 15:04:05"))
	return timeString
}

/*
	$ localTime sql handler
*/

// Value the handler of staging localTime struct to database as time
func (t LocalTime) Value() (driver.Value, error) {
	sqlTime := time.Time(t)
	return sqlTime.Format("2006-01-02 15:04:05"), nil
}

// Scan the handler of converting the time data from database to LocalTime
func (t *LocalTime) Scan(val interface{}) error {
	if time, ok := val.(time.Time); ok {
		*t = LocalTime(time)
		return nil
	}
	return fmt.Errorf("couldn't convert %v to timestamp", val)
}
