package main

type Controller struct {
  model *Model
  qmlAdapter *QmlAdapter
}

func NewController() *Controller {
  model := NewModel()
  return &Controller{
    model: model,
    qmlAdapter: NewQmlAdapter(),
  }
}
//  https://gist.github.com/icambridge/9708081
