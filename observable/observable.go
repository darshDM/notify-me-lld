package observable

import (
	"github.com/darshDM/notify-me-lld/observer"
)

type QuantityObservable interface {
	Add(no observer.NotifyObserver)
	Remove(no observer.NotifyObserver)
	NotifyAll(qty int)
	SetQuantity(qty int)
}

type Observable struct {
	qty  int
	List []observer.NotifyObserver
}

func (obs *Observable) Add(no observer.NotifyObserver) {
	obs.List = append(obs.List, no)
}

func (obs *Observable) Remove(no observer.NotifyObserver) {
	for i, observer := range obs.List {
		if observer.GetID() == no.GetID() {
			obs.List = append(obs.List[:i], obs.List[i+1:]...)
		}
	}
}

func (obs *Observable) NotifyAll() {
	for _, no := range obs.List {
		no.Update(obs.qty)
	}
}

func (obs *Observable) SetQuantity(qty int) {
	obs.qty = qty
	obs.NotifyAll()
}
