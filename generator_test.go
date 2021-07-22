package ics

import (
	"testing"
	"time"
)

func Test_formatDateTime(t *testing.T) {
	t.Run("double digits", func(t *testing.T) {
		date := time.Date(2019, time.December, 12, 12, 12, 12, 0, time.UTC)

		dateString := FormatDateTime(date)
		expectedDateString := "20191212T121212Z"

		if dateString != expectedDateString {
			t.Fatalf("expected %s, got %s", expectedDateString, dateString)
		}
	})

	t.Run("single digits", func(t *testing.T) {
		date := time.Date(2019, time.January, 1, 9, 0, 0, 0, time.UTC)

		dateString := FormatDateTime(date)
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
		UID: "123123123",
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
ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN="John Wick";RSVP=FALSE:mailto:john.wick@gmail.com
LOCATION:location
DTSTAMP:` + event.dtStamp + `
DTSTART:20190101T090000Z
DTEND:20190101T093000Z
SUMMARY:Event Name
DESCRIPTION:Event Description
CLASS:PUBLIC
UID:313233313233313233
STATUS:CONFIRMED
END:VEVENT
END:VCALENDAR`

	if gotIcs != expectedIcs {
		t.Errorf("Generate() = %v, want %v", gotIcs, expectedIcs)
	}

}

func TestGenerateWithRecurrence(t *testing.T) {
	event := &Event{
		Class:       Classification_PUBLIC,
		Summary:     "Event Name",
		Description: "Event Description",
		Status:      EventStatus_CONFIRMED,
		Location:    "location",
		DtStart:     time.Date(2019, time.January, 1, 9, 0, 0, 0, time.UTC),
		DtEnd:       time.Date(2019, time.January, 1, 9, 30, 0, 0, time.UTC),
		RRule: []string{
			"FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,WE,FR",
			"FREQ=WEEKLY;INTERVAL=2;BYDAY=SU",
		},
		ExRule:       []string{},
		ExDate:       []time.Time{},
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
		UID: "123123123",
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
ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN="John Wick";RSVP=FALSE:mailto:john.wick@gmail.com
LOCATION:location
DTSTAMP:` + event.dtStamp + `
DTSTART:20190101T090000Z
DTEND:20190101T093000Z
RRULE:FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,WE,FR
RRULE:FREQ=WEEKLY;INTERVAL=2;BYDAY=SU
SUMMARY:Event Name
DESCRIPTION:Event Description
CLASS:PUBLIC
UID:313233313233313233
STATUS:CONFIRMED
END:VEVENT
END:VCALENDAR`

	if gotIcs != expectedIcs {
		t.Errorf("Generate() = %v, want %v", gotIcs, expectedIcs)
	}

}

func TestGenerateWithExceptions(t *testing.T) {
	event := &Event{
		Class:       Classification_PUBLIC,
		Summary:     "Event Name",
		Description: "Event Description",
		Status:      EventStatus_CONFIRMED,
		Location:    "location",
		DtStart:     time.Date(2019, time.January, 1, 9, 0, 0, 0, time.UTC),
		DtEnd:       time.Date(2019, time.January, 1, 9, 30, 0, 0, time.UTC),
		RRule: []string{
			"FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,WE,FR",
		},
		ExRule: []string{
			"FREQ=WEEKLY;INTERVAL=2;BYDAY=FR",
		},
		ExDate: []time.Time{
			time.Date(2019, time.January, 31, 9, 0, 0, 0, time.UTC),
			time.Date(2019, time.February, 28, 9, 0, 0, 0, time.UTC),
		},
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
		UID: "123123123",
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
ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN="John Wick";RSVP=FALSE:mailto:john.wick@gmail.com
LOCATION:location
DTSTAMP:` + event.dtStamp + `
DTSTART:20190101T090000Z
DTEND:20190101T093000Z
RRULE:FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,WE,FR
EXRULE:FREQ=WEEKLY;INTERVAL=2;BYDAY=FR
EXDATE:20190131T090000Z
EXDATE:20190228T090000Z
SUMMARY:Event Name
DESCRIPTION:Event Description
CLASS:PUBLIC
UID:313233313233313233
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
		UID: "123123123",
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
ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN="John Wick";RSVP=FALSE:mailto:john.wick@gmail.com
LOCATION:location
DTSTAMP:` + event.dtStamp + `
DTSTART:20190101T090000Z
DTEND:20190101T093000Z
SUMMARY:Event Name
DESCRIPTION:Event Description
CLASS:PUBLIC
UID:313233313233313233
STATUS:CONFIRMED
END:VEVENT
END:VCALENDAR`

	if gotIcs != expectedIcs {
		t.Errorf("Generate() = %v, want %v", gotIcs, expectedIcs)
	}
}

func TestGenerateWithRSVP(t *testing.T) {
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
			Rsvp:         Rsvp_True,
		}},
		Organizer: Attendee{
			CommonName:   "My Calendar",
			EmailAddress: "my@calendar.com",
		},
		UID: "123123123",
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
ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN="John Wick";RSVP=TRUE:mailto:john.wick@gmail.com
LOCATION:location
DTSTAMP:` + event.dtStamp + `
DTSTART:20190101T090000Z
DTEND:20190101T093000Z
SUMMARY:Event Name
DESCRIPTION:Event Description
CLASS:PUBLIC
UID:313233313233313233
STATUS:CONFIRMED
END:VEVENT
END:VCALENDAR`

	if gotIcs != expectedIcs {
		t.Errorf("Generate() = %v, want %v", gotIcs, expectedIcs)
	}

}
