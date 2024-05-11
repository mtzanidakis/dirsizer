package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mtzanidakis/dirsizer/internal/util"
	"gopkg.in/gomail.v2"
)

type DirSizer struct {
	Identifier string
	Directory  string
	MailFrom   string
	MailTo     string
	Threshold  int64
	*gomail.Dialer
}

// NewDirSizer returns a new DirSizer.
func NewDirSizer(
	identifier, directory, mailFrom, mailTo, smtpServer, threshold string,
) (*DirSizer, error) {
	smtpServerSplit := strings.Split(smtpServer, ":")
	smtpHost := smtpServerSplit[0]
	smtpPort, err := strconv.Atoi(smtpServerSplit[1])
	if err != nil {
		return &DirSizer{}, err
	}

	d := gomail.Dialer{Host: smtpHost, Port: smtpPort}

	thresholdInt, err := util.IECToBytes(threshold)
	if err != nil {
		return &DirSizer{}, err
	}

	return &DirSizer{
		Identifier: identifier,
		Directory:  directory,
		MailFrom:   mailFrom,
		MailTo:     mailTo,
		Threshold:  thresholdInt,
		Dialer:     &d,
	}, nil
}

// Count returns the size of the given directory in bytes.
func (d *DirSizer) Count() (int64, error) {
	var size int64
	err := filepath.Walk(d.Directory, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// SendMail sends an email.
func (d *DirSizer) SendMail(subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", d.MailFrom)
	m.SetHeader("To", d.MailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return d.Dialer.DialAndSend(m)
}

const bodyTmpl = `
<h1>Directory size exceeded</h1>
<p><b>Size</b> %s</p>
<p><b>Threshold</b> %s</p>`
const logTmpl = "Directory size exceeded: %s"
const subjectTmpl = "Directory size exceeded for %s"

// Run runs the DirSizer.
func (d *DirSizer) Run() error {
	size, err := d.Count()
	if err != nil {
		return err
	}

	if size > d.Threshold {
		log.Printf(logTmpl, util.ByteCountIEC(size))

		subject := fmt.Sprintf(subjectTmpl, d.Identifier)
		body := fmt.Sprintf(
			bodyTmpl,
			util.ByteCountIEC(size),
			util.ByteCountIEC(d.Threshold),
		)

		return d.SendMail(subject, body)
	}

	return nil
}

func main() {
	d, err := NewDirSizer(
		util.EnvOrDefault("IDENTIFIER", "dir on localhost"),
		util.EnvOrDefault("DIRECTORY", "."),
		util.EnvOrDefault("MAIL_FROM", "dirsizer@localhost"),
		util.EnvOrDefault("MAIL_TO", "root"),
		util.EnvOrDefault("SMTP_SERVER", "localhost:25"),
		util.EnvOrDefault("THRESHOLD", "500M"),
	)
	if err != nil {
		panic(err)
	}

	log.Printf("Starting DirSizer for %s", d.Identifier)

	ticker := time.NewTicker(4 * time.Hour)

	for ; true; <-ticker.C {
		if err = d.Run(); err != nil {
			panic(err)
		}
	}
}
