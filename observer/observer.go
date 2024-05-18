package observer

import "fmt"

type NotifyObserver interface {
	Update(qty int)
	GetID() int
}

// Email related struct and functions
type EmailNotificationObserver struct {
	Id      int
	EmailId string
}

func (eno *EmailNotificationObserver) Update(qty int) {
	fmt.Printf("> Sending mail to: %v, stock available:%v\n", eno.EmailId, qty)
}
func (eno *EmailNotificationObserver) GetID() int {
	return eno.Id
}

// Phone related struct and functions
type SMSNotificationObserver struct {
	Id    int
	Phone string
}

func (pno *SMSNotificationObserver) Update(qty int) {
	fmt.Printf("> Sending SMS to: %v, stock available:%v\n", pno.Phone, qty)
}

func (pno *SMSNotificationObserver) GetID() int {
	return pno.Id
}
