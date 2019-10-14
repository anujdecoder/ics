package ics

const ics = `BEGIN:VCALENDAR
PRODID:{{.ProdId}}
METHOD:REQUEST
VERSION:2.0
{{range $ve := .Events}}{{$ve}}
{{end}}END:VCALENDAR`

const vevent = `BEGIN:VEVENT
ORGANIZER;CN="{{.Organizer.CommonName}}":mailto:{{.Organizer.EmailAddress}}
{{range $at := .Attendees}}ATTENDEE;CN="{{$at.CommonName}}";RSVP=FALSE:mailto:{{$at.EmailAddress}};CUTYPE={{$at.CuType}};ROLE={{$at.Role}};PARTSTAT={{$at.PartStatus}}
{{end}}LOCATION:{{.Location}}
DTSTAMP:{{.DtStamp}}
DTSTART:{{.DtStart}}
DTEND:{{.DtEnd}}
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
	Description string
}
