package observer

type NewsPublisher struct{
	newsListeners []newsListener
}

func NewNewsPublisher() *NewsPublisher{
	return &NewsPublisher{make([]newsListener, 0)}
}

func (e *NewsPublisher)Subscribe(subscriber newsListener){
	e.newsListeners = append(e.newsListeners,subscriber)
}

func (e *NewsPublisher)NotifyAll(news string){
	for _,subscriber:=range e.newsListeners{
		subscriber.Update(news)
	}
}