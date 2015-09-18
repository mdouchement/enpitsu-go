package app

type Publisher interface {
	Publish(value interface{})
}

type Observer interface {
	Notify(key string, value interface{})
}

type ObserverFunc func(key string, value interface{})

func (fn ObserverFunc) Notify(key string, value interface{}) {
	fn(key, value)
}

type Observable []Observer

func (observers *Observable) AttachObserver(a Observer) {
	*observers = append(*observers, a)
}

func (observers Observable) Publish(key string, value interface{}) {
	for _, obs := range observers {
		obs.Notify(key, value)
	}
}
