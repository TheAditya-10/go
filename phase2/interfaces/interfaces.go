package interfaces
import "fmt"

type Notifier interface {
	Send(message string)
}

type  Email struct {}

func (e Email) Send(message string) {
	fmt.Println("Sending email with message:", message)
}

type SMS struct {}

func (s SMS) Send(message string) {
	fmt.Println("Sending SMS with message:", message)
}

func Notify(n Notifier) {
	n.Send("This is a notification.")
}