package utils

import (
	gomail "gopkg.in/mail.v2"
)

func Sendmail(email,secretcode string){

	abc := gomail.NewMessage()

	abc.SetHeader("From", "buyoyaibrahim92@gmail.com")
	abc.SetHeader("To",email )
	abc.SetHeader("Subject", "Acount validation")
	abc.SetBody("text/html", "Code de verification :" + secretcode)

	a := gomail.NewDialer("smtp.gmail.com",587,"buyoyaibrahim92@gmail.com","Ibrahim_mouhammad&2021")

	if err := a.DialAndSend(abc); err != nil {
		// fmt.Printf(err.Error())
		panic(err)
	}
	
}