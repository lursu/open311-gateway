package create

import (
	"CitySourcedAPI/logs"
	"bytes"
	"fmt"
	"text/template"

	"Gateway311/adapters/email/common"
	"Gateway311/adapters/email/data"
	"Gateway311/adapters/email/mail"
	"Gateway311/adapters/email/structs"
)

var (
	log = logs.Log
)

// ================================================================================================
//                                      CREATE
// ================================================================================================

// Request represents the XML payload for a report request to CitySourced.
// type Request struct {
// 	Sender            data.EmailSender
// 	RequestType       string  `json:"RequestType" xml:"RequestType"`
// 	RequestTypeID     int     `json:"RequestTypeId" xml:"RequestTypeId"`
// 	ImageURL          string  `json:"ImageUrl" xml:"ImageUrl"`
// 	Latitude          float64 `json:"Latitude" xml:"Latitude"`
// 	Longitude         float64 `json:"Longitude" xml:"Longitude"`
// 	Description       string  `json:"Description" xml:"Description"`
// 	AuthorNameFirst   string  `json:"AuthorNameFirst" xml:"AuthorNameFirst"`
// 	AuthorNameLast    string  `json:"AuthorNameLast" xml:"AuthorNameLast"`
// 	AuthorEmail       string  `json:"AuthorEmail" xml:"AuthorEmail"`
// 	AuthorTelephone   string  `json:"AuthorTelephone" xml:"AuthorTelephone"`
// 	AuthorIsAnonymous bool    `json:"AuthorIsAnonymous" xml:"AuthorIsAnonymous"`
// }

// Request represents the XML payload for a report request to CitySourced.
type Request struct {
	Sender data.EmailSender
	Body   *structs.Payload
}

// Process executes the request to create a new report.
func (r *Request) Process() (*Response, error) {
	fail := func(err error) (*Response, error) {
		response := Response{
			Message: fmt.Sprintf("unable to send email - %s", err),
		}
		return &response, err
	}

	if err := mail.Send(r.Sender, r.Body); err != nil {
		fail(err)
	}

	// to, from, subject := r.Sender.Address()
	// body, err := r.createEmail(r.Sender.Template())
	// if err != nil {
	// 	fail(err)
	// }
	//
	// address := &structs.Address{
	// 	To:   to,
	// 	From: from,
	// }
	// payload := structs.NewPayloadString(subject, &body)
	//
	// if err := mail.Send(address, payload); err != nil {
	// 	fail(err)
	// }

	return &Response{"Success"}, nil
}

// ------------------------------------------------------------------------------------------------

// Response is the response to creating or updating a report.
type Response struct {
	Message string `json:"Message" xml:"Message"`
}

// ================================================================================================
//                                      TEMPLATES
// ================================================================================================

// createEmail creates an email message from the request using the specified template
func (r *Request) createEmail(tmpl *template.Template) (string, error) {
	var doc bytes.Buffer
	// Apply the values we have initialized in our struct context to the template.
	if err := tmpl.Execute(&doc, r); err != nil {
		log.Error("error trying to execute email template ", err)
		return "", err
	}
	log.Debug("Doc:\n%s", doc.String())
	return doc.String(), nil
}

// ================================================================================================
//                                      STRINGS
// ================================================================================================

// String displays a Request
func (r Request) String() string {
	ls := new(common.LogString)
	ls.AddS("create.Request\n")
	t, f, s := r.Sender.Address()
	ls.AddF("Sender - to: %#v  from: %#v  subject: %q\n", t, f, s)
	ls.AddF("Message:\n%s\n", r.Body)
	return ls.Box(80)
}

// String displays a Response
func (r Response) String() string {
	ls := new(common.LogString)
	ls.AddS("create.Response\n")
	ls.AddF("Message: %v\n", r.Message)
	return ls.Box(80)
}