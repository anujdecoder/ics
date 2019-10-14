package ics

import (
	"time"
)

type Event struct {
	Class        CLASS
	Summary      string
	Description  string
	Status       EventStatus
	Geo          *GeoLocation
	Location     string
	DtEnd        time.Time
	DtStart      time.Time
	Transparency Transparency
	Attendees    []Attendee
	Organizer    Attendee
	UID          string

	dtStamp string
}

type EventStatus string

const (
	EventStatus_CONFIRMED EventStatus = "CONFIRMED"
	EventStatus_CANCELLED EventStatus = "CANCELLED"
	EventStatus_TENTATIVE EventStatus = "TENTATIVE"
)

type CLASS string

const (
	Classification_PUBLIC       CLASS = "PUBLIC"
	Classification_PRIVATE      CLASS = "PRIVATE"
	Classification_CONFIDENTIAL CLASS = "CONFIDENTIAL"
)

type GeoLocation struct {
	Latitude  float32
	Longitude float32
}

type Transparency string

const (
	TRANSAPARENT Transparency = "TRANSPARENT"
	OPAQUE       Transparency = "OPAQUE"
)

type Attendee struct {
	CommonName   string
	EmailAddress string
	Role         Role
	PartStatus   AttendeeStatus
	//RSVP is by default NO
	CuType CalendarUserType
}

type Role string

const (
	REQUIRED Role = "REQ-PARTICIPANT"
)

type AttendeeStatus string

const (
	AttendeeStatus_TENTATIVE AttendeeStatus = "TENTATIVE"
	AttendeeStatus_ACCEPTED  AttendeeStatus = "ACCEPTED"
)

type CalendarUserType string

const (
	INDIVIDUAL CalendarUserType = "INDIVIDUAL"
)
