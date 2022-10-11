// Code generated by tutone: DO NOT EDIT
package nrtime

import (
	"github.com/RiotMingle/newrelic-client-go/internal/serialization"
)

// TimeWindowInput - Represents a time window input.
type TimeWindowInput struct {
	// The end time of the time window the number of milliseconds since the Unix epoch.
	EndTime EpochMilliseconds `json:"endTime"`
	// The start time of the time window the number of milliseconds since the Unix epoch.
	StartTime EpochMilliseconds `json:"startTime"`
}

// DateTime - The `DateTime` scalar represents a date and time. The `DateTime` appears as an ISO8601 formatted string.
type DateTime string

// EpochMilliseconds - The `EpochMilliseconds` scalar represents the number of milliseconds since the Unix epoch
type EpochMilliseconds serialization.EpochTime

// EpochSeconds - The `EpochSeconds` scalar represents the number of seconds since the Unix epoch
type EpochSeconds serialization.EpochTime

// Milliseconds - The `Milliseconds` scalar represents a duration in milliseconds
type Milliseconds int

// Minutes - The `Minutes` scalar represents a duration in minutes
type Minutes int

// NaiveDateTime - The `NaiveDateTime` scalar represents a date and time without a Time Zone. The `NaiveDateTime` appears as an ISO8601 formatted string.
type NaiveDateTime string

// Seconds - The `Seconds` scalar represents a duration in seconds
type Seconds string
