package app

import(
  "gopkg.in/qml.v1"
)

type Controller struct {
  model *Model
  view *QmlAdapter
}

func NewController(ctx *qml.Context) *Controller {
  view := NewQmlAdapter(ctx)
  model := NewModel()
  return &Controller{
    model: model,
    view: view,
  }
}
//  https://gist.github.com/icambridge/9708081
// https://talks.golang.org/2014/organizeio.slide#1
