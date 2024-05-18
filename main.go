package main

import (
	"github.com/darshDM/notify-me-lld/observable"
	"github.com/darshDM/notify-me-lld/observer"
)

func main() {

	user1 := &observer.EmailNotificationObserver{EmailId: "test@gmail.com", Id: 1}
	user2 := &observer.SMSNotificationObserver{Phone: "0987654321", Id: 2}
	user3 := &observer.EmailNotificationObserver{EmailId: "user@gyahoo.com", Id: 3}

	var item observable.Observable
	item.Add(user1)
	item.Add(user2)
	item.Add(user3)
	item.Remove(user1)

	item.SetQuantity(2)

}
