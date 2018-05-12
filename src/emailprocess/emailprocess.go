package emailprocess

import (
	// "io"
	// "io/ioutil"
	"io"
	"io/ioutil"
	"log"
	"os"

	"../conf"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

const checkMailSubject = "#CHECK"

func Receive2() {
	log.Println("Connecting to server...")

	// Connect to server
	c := connect()
	// Don't forget to logout
	defer c.Logout()

	// Login
	user := conf.Get().MailBox.User
	password := conf.Get().MailBox.Password

	if err := c.Login(user, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(mbox.Messages)

	// Get the whole message body
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Fatal("Server didn't returned message")
	}

	r := msg.GetBody(section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// Print some info about the message
	header := mr.Header
	if date, err := header.Date(); err == nil {
		log.Println("Date:", date)
	}
	if from, err := header.AddressList("From"); err == nil {
		log.Println("From:", from)
	}
	if to, err := header.AddressList("To"); err == nil {
		log.Println("To:", to)
	}
	if subject, err := header.Subject(); err == nil {
		log.Println("Subject:", subject)
	}

	// Process each message's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch h := p.Header.(type) {
		case mail.TextHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(p.Body)
			log.Println("Got text: %v", string(b))
		case mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			b, _ := ioutil.ReadAll(p.Body)
			fo, _ := os.Create("output.txt")
			fo.Write(b)
			log.Println("Got attachment: %v", filename)
		}
	}
}

func connect() *client.Client {
	mailBox := conf.Get().MailBox
	result, err := client.DialTLS(mailBox.ServerAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	return result
}
