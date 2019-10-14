package ics

import (
	"testing"
	"time"
)

func Test_formatDateTime(t *testing.T) {
	t.Run("double digits", func(t *testing.T) {
		date := time.Date(2019, time.December, 12, 12, 12, 12, 0, time.UTC)

		dateString := formatDateTime(date)
		expectedDateString := "20191212T121212Z"

		if dateString != expectedDateString {
			t.Fatalf("expected %s, got %s", expectedDateString, dateString)
		}
	})

	t.Run("single digits", func(t *testing.T) {
		date := time.Date(2019, time.January, 1, 9, 0, 0, 0, time.UTC)

		dateString := formatDateTime(date)
		expectedDateString := "20190101T090000Z"

		if dateString != expectedDateString {
			t.Fatalf("expected %s, got %s", expectedDateString, dateString)
		}
	})
}

func TestGenerate(t *testing.T) {
	event := &Event{
		Class:        Classification_PUBLIC,
		Summary:      "Event Name",
		Description:  "Event Description",
		Status:       EventStatus_CONFIRMED,
		Location:     "location",
		DtStart:      time.Date(2019, time.January, 1, 9, 0, 0, 0, time.UTC),
		DtEnd:        time.Date(2019, time.January, 1, 9, 30, 0, 0, time.UTC),
		Transparency: OPAQUE,
		Attendees: []Attendee{{
			CommonName:   "John Wick",
			EmailAddress: "john.wick@gmail.com",
			Role:         REQUIRED,
			PartStatus:   AttendeeStatus_ACCEPTED,
			CuType:       INDIVIDUAL,
		}},
		Organizer: Attendee{
			CommonName:   "My Calendar",
			EmailAddress: "my@calendar.com",
		},
		UID: "123-123-123",
	}

	gotIcs, err := Generate("com.calendar.my", event)
	if err != nil {
		t.Errorf("Generate() error = %v", err)
		return
	}

	expectedIcs := `BEGIN:VCALENDAR
PRODID:com.calendar.my
METHOD:REQUEST
VERSION:2.0
BEGIN:VEVENT
ORGANIZER;CN="My Calendar":mailto:my@calendar.com
ATTENDEE;CN="John Wick";RSVP=FALSE:mailto:john.wick@gmail.com;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED
LOCATION:location
DTSTAMP:` + event.dtStamp + `
DTSTART:20190101T090000Z
DTEND:20190101T093000Z
SUMMARY:Event Name
DESCRIPTION:Event Description
CLASS:PUBLIC
UID:123-123-123
STATUS:CONFIRMED
END:VEVENT
END:VCALENDAR`

	if gotIcs != expectedIcs {
		t.Errorf("Generate() = %v, want %v", gotIcs, expectedIcs)
	}

}

func TestEvent_Generate(t *testing.T) {
	event := &Event{
		Class:        Classification_PUBLIC,
		Summary:      "Event Name",
		Description:  "Event Description",
		Status:       EventStatus_CONFIRMED,
		Location:     "location",
		DtStart:      time.Date(2019, time.January, 1, 9, 0, 0, 0, time.UTC),
		DtEnd:        time.Date(2019, time.January, 1, 9, 30, 0, 0, time.UTC),
		Transparency: OPAQUE,
		Attendees: []Attendee{{
			CommonName:   "John Wick",
			EmailAddress: "john.wick@gmail.com",
			Role:         REQUIRED,
			PartStatus:   AttendeeStatus_ACCEPTED,
			CuType:       INDIVIDUAL,
		}},
		Organizer: Attendee{
			CommonName:   "My Calendar",
			EmailAddress: "my@calendar.com",
		},
		UID: "123-123-123",
	}

	gotIcs, err := event.Generate("com.calendar.my")
	if err != nil {
		t.Errorf("Generate() error = %v", err)
		return
	}

	expectedIcs := `BEGIN:VCALENDAR
PRODID:com.calendar.my
METHOD:REQUEST
VERSION:2.0
BEGIN:VEVENT
ORGANIZER;CN="My Calendar":mailto:my@calendar.com
ATTENDEE;CN="John Wick";RSVP=FALSE:mailto:john.wick@gmail.com;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED
LOCATION:location
DTSTAMP:` + event.dtStamp + `
DTSTART:20190101T090000Z
DTEND:20190101T093000Z
SUMMARY:Event Name
DESCRIPTION:Event Description
CLASS:PUBLIC
UID:123-123-123
STATUS:CONFIRMED
END:VEVENT
END:VCALENDAR`

	if gotIcs != expectedIcs {
		t.Errorf("Generate() = %v, want %v", gotIcs, expectedIcs)
	}
}
