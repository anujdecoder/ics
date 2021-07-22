package ics

const ics = `BEGIN:VCALENDAR
PRODID:{{.ProdId}}
METHOD:REQUEST
VERSION:2.0
{{range $ve := .Events}}{{$ve}}
{{end}}END:VCALENDAR`

const vevent = `BEGIN:VEVENT
ORGANIZER;CN="{{.Organizer.CommonName}}":mailto:{{.Organizer.EmailAddress}}
{{range $at := .Attendees}}ATTENDEE;CUTYPE={{$at.CuType}};ROLE={{$at.Role}};PARTSTAT={{$at.PartStatus}};CN="{{$at.CommonName}}";RSVP={{$at.Rsvp}}:mailto:{{$at.EmailAddress}}
{{end}}LOCATION:{{.Location}}
DTSTAMP:{{.DtStamp}}
DTSTART:{{.DtStart}}
DTEND:{{.DtEnd}}{{range $ru := .RRule}}
RRULE:{{$ru}}{{end}}{{range $ru := .ExRule}}
EXRULE:{{$ru}}{{end}}{{range $ru := .ExDate}}
EXDATE:{{$ru}}{{end}}
SUMMARY:{{.Summary}}
DESCRIPTION:{{.Description}}
CLASS:{{.Class}}
UID:{{.UID}}
STATUS:{{.Status}}
END:VEVENT`

type vEvent struct {
	*Event
	DtStamp     string
	DtEnd       string
	DtStart     string
	ExDate      []string
	Description string
}
