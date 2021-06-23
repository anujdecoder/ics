# ics
ics file generator for golang

## Example
```go
event := &Event{
    Class:       Classification_PUBLIC,
    Summary:     "Meeting with John",
    Description: "About ICS",
    Status:      EventStatus_CONFIRMED,
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
        CommonName:   "Anuj",
        EmailAddress: "anuj@gmail.com",
    },
}

gotIcs, err := event.Generate("com.calendar.my")
if err != nil {
	/* handle error */
}
```
