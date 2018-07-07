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
	e.init(l)
	m := gomail.NewMessage()
	//m.SetHeader("From", Email_user)
	m.SetAddressHeader("From",e.User,alias)
	for i:=0 ;i<len(cc);i++{
		to = append(to,cc[i])
	}
	for i:= 0;i<len(to);i++{
		go func(){
			m.SetHeader("To",to[i])
			m.SetHeader("Subject", subject)
			d := gomail.NewDialer(e.Host, e.Port, e.User, e.Password)
			if err := d.DialAndSend(m); err != nil {
				//Writefile("Sendmail "+to,errors.New("success"))
				msg := "Sendmail "+to[i]+err.Error()
				l.Error(msg)
			}
		}()
	}
}
