package smtpemail

import (
	
	"log"
	"net/smtp"
)

func sendEmail(target string,subject string,body string) (string, error){
	auth:=smtp.PlainAuth(
		"",
		"techkaran5599@gmail.com",
		"qxav lynf wsmf ykgc",
		"smtp.gmail.com",

	)
	
	message:= subject + body
	err:=smtp.SendMail(
		"smtp.gmail.com:587",
        auth,
		"techkaran5599@gmail.com",
		[]string{target},
		[]byte(message),
		
	)
	if err != nil {
		log.Fatal(err)
		return "",err
	}
	return "meassge sent sucessfully",nil
}
