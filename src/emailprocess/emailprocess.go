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

func Receive() {
	log.Println("Connecting to server...")

	mailClient := login()
	defer mailClient.Logout()

	// mbox := getInbox(mailClient)

	// // Get the last message
	// if mbox.Messages == 0 {
	// 	log.Fatal("No message in mailbox")
	// }
	// to := mbox.Messages
	// from := to - 9
	// seqSet := new(imap.SeqSet)
	// seqSet.AddRange(from, to)
	// // seqSet.AddNum(mbox.Messages)

	// // Get the whole message body
	// section := &imap.BodySectionName{}
	// items := []imap.FetchItem{section.FetchItem()}

	// messages := make(chan *imap.Message, 10)
	// go func() {
	// 	err := mailClient.Fetch(seqSet, items, messages)
	// 	errorFatal(err)
	// }()

	section := &imap.BodySectionName{}
	for msg := range getMessages(10, mailClient) {
		if msg == nil {
			log.Fatal("Server didn't returned message")
		}

		r := msg.GetBody(section)
		if r == nil {
			log.Fatal("Server didn't returned message body")
		}

		// Create a new mail reader
		mr, err := mail.CreateReader(r)
		errorFatal(err)

		// Print some info about the message
		header := mr.Header
		if date, err := header.Date(); err == nil {
			log.Println("Date:", date)
		}
		// if from, err := header.AddressList("From"); err == nil {
		// 	log.Println("From:", from)
		// }
		// if to, err := header.AddressList("To"); err == nil {
		// 	log.Println("To:", to)
		// }
		if subject, err := header.Subject(); err == nil {
			log.Println("Subject:", subject)
		}

		// Process each message's part
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else {
				errorFatal(err)
			}

			switch h := p.Header.(type) {
			case mail.TextHeader:
				// This is the message's text (can be plain-text or HTML)
				// b, _ := ioutil.ReadAll(p.Body)
				// log.Println("Got text: %v", string(b))
			case mail.AttachmentHeader:
				// This is an attachment
				filename, _ := h.Filename()
				b, _ := ioutil.ReadAll(p.Body)
				fo, _ := os.Create(filename)
				fo.Write(b)
				log.Println("Got attachment: %v", filename)
			}
		}
	}

}

func getMessages(count uint32, mailClient *client.Client) chan *imap.Message {
	mbox := getInbox(mailClient)
	// Get the last message
	to := mbox.Messages
	from := to - count
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	} else if mbox.Messages < count {
		from = 1
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, to)

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 10)
	go func() {
		err := mailClient.Fetch(seqSet, items, messages)
		errorFatal(err)
	}()
	return messages
}

func getInbox(mailClient *client.Client) *imap.MailboxStatus {
	mbox, err := mailClient.Select("INBOX", false)
	errorFatal(err)
	return mbox
}

func login() *client.Client {
	c := connect()

	user := conf.Get().MailBox.User
	password := conf.Get().MailBox.Password

	err := c.Login(user, password)
	errorFatal(err)
	log.Println("Logged in")

	return c
}

func connect() *client.Client {
	mailBox := conf.Get().MailBox
	result, err := client.DialTLS(mailBox.ServerAddress, nil)
	errorFatal(err)
	log.Println("Connected")
	return result
}

func errorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
