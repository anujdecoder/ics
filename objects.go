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
	RRule        []string
	ExRule       []string
	ExDate       []time.Time
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
	Rsvp   RSVP
}

type Role string

const (
	REQUIRED Role = "REQ-PARTICIPANT"
)

type AttendeeStatus string

const (
	AttendeeStatus_NeedAction AttendeeStatus = "NEEDS-ACTION"
	AttendeeStatus_TENTATIVE  AttendeeStatus = "TENTATIVE"
	AttendeeStatus_ACCEPTED   AttendeeStatus = "ACCEPTED"
	AttendeeStatus_DECLINED   AttendeeStatus = "DECLINED"
)

type CalendarUserType string

const (
	INDIVIDUAL CalendarUserType = "INDIVIDUAL"
)

// This is related to PartStat
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
type RSVP string

const (
	Rsvp_False RSVP = "FALSE"
	Rsvp_True  RSVP = "TRUE"
)
