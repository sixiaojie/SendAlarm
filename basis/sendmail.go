package basis

import(
	"gopkg.in/gomail.v2"
	"github.com/astaxie/beego/logs"
)

type Email_Server struct {
	Host string 	`json:"host"`
	User string 	`json:"user"`
	Password string `json:"password"`
	Port int		`json:"port"`
	//SSl bool		`json:"ssl"`
}


func (e *Email_Server) init(l *logs.BeeLogger){
	iniconf := Appconf()
	e.Host = iniconf.String("email_host")
	e.User = iniconf.String("email_user")
	e.Password = iniconf.String("email_password")
	e.Port,_ = iniconf.Int("email_port")
	//e.SSl,_ = iniconf.Bool("email_ssl")
}


//这里将发送的人，改成一个一个的发送。

func (e *Email_Server) SendMail(alias,subject,body string,to,cc []string,l *logs.BeeLogger){
	for i:=0 ;i<len(cc);i++{
		to = append(to,cc[i])
	}
	for i:= 0;i<len(to);i++{
		go Email_Accept(e,l,alias,subject,to[i],body)
	}
}

func Email_Accept(e *Email_Server,l *logs.BeeLogger,alias,subject,user,msg string){
	e.init(l)
	m := gomail.NewMessage()
	m.SetAddressHeader("From",e.User,alias)
	m.SetHeader("To",user)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)
	d := gomail.NewDialer(e.Host, e.Port, e.User, e.Password)
	if err := d.DialAndSend(m); err != nil {
		//Writefile("Sendmail "+to,errors.New("success"))
		msg := "Sendmail "+user+err.Error()
		l.Error(msg)
	}
}
