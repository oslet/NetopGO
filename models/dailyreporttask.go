package models

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"time"

	"net/smtp"
	//"os"
	"strings"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/robfig/cron"
)

var (
	mailserver  string = beego.AppConfig.String("mailserver")
	mailport    string = beego.AppConfig.String("mailport")
	mailuser    string = beego.AppConfig.String("mailuser")
	mailpass    string = beego.AppConfig.String("mailpass")
	mailsubject string = beego.AppConfig.String("mailsubject")
	username    string = beego.AppConfig.String("db_user")
	userpwd     string = beego.AppConfig.String("db_passwd")
	dbhost      string = beego.AppConfig.String("db_host")
	dbport      string = beego.AppConfig.String("db_port")
	dbname      string = beego.AppConfig.String("db_schema")
)

type Dailyreporttask struct {
	Appname      string
	Appcontent   string
	Applicant    string
	Publisher    string
	Department   string
	Publishtime  string
	Followstatus string
	Followman    string
}

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
	Mtype   string
}

type SmtpServer struct {
	Host      string
	Port      string
	TlsConfig *tls.Config
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	var content_type string
	if mail.Mtype == "html" {
		content_type = "Content-Type: text/" + mail.Mtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += fmt.Sprintf("%s\r\n", content_type)

	header += "\r\n\r\n" + mail.Body
	//fmt.Println("header:", header)
	return header
}

func TasksForDailyReport() {
	curdate := time.Now().Format("2006-01-02")
	mail := Mail{}
	mail.Sender = mailuser
	mail.To = []string{mailuser}
	mail.Cc = []string{}
	mail.Bcc = []string{}
	mail.Subject = mailsubject + "(" + curdate + ")"
	mail.Mtype = "html"

	mail.Body = GetDailyReport()

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{Host: mailserver, Port: mailport}
	smtpServer.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	auth := smtp.PlainAuth("", mail.Sender, mailpass, smtpServer.Host)
	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.TlsConfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		log.Panic(err)
	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDailyReport() string {
	schemaUrl := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_passwd") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_schema") + "?charset=utf8"
	db, err := sql.Open("mysql", schemaUrl)
	if err != nil {
		log.Println(err.Error()) //仅仅是显示异常
	}

	defer db.Close()
	rows, err := db.Query("SELECT appname, appcontent, applicant, publisher, department, publishtime, followstatus, followman FROM dailyreport where date(publishtime) = curdate()")
	var appname, appcontent, applicant, publisher, department, publishtime, followstatus, followman string
	locals := make(map[string]interface{})
	dailyreports := []Dailyreporttask{}

	for rows.Next() {
		err = rows.Scan(&appname, &appcontent, &applicant, &publisher, &department, &publishtime, &followstatus, &followman)
		if err == nil {
			//	log.Println(appname, appcontent, applicant, publisher, department, publishtime, followstatus, followman)
			dailyreports = append(dailyreports, Dailyreporttask{appname, appcontent, applicant, publisher, department, publishtime, followstatus, followman})
		}
	}
	locals["reports"] = dailyreports
	t, _ := template.ParseFiles("views/report_dailydeploy_mail.html")
	var body bytes.Buffer
	t.Execute(&body, dailyreports)
	return fmt.Sprintf(body.String())
}

func getDB(username, userpwd, dbhost, dbport, dbname string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, userpwd, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(err.Error()) //仅仅是显示异常
		return nil, err
	}
	return db, nil
}
