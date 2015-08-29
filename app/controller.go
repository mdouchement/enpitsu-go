package app

import(
  "gopkg.in/qml.v1"
  "strings"
)

type Controller struct {
  Model *Model
  View *View
}

func NewController(ctx *qml.Context) *Controller {
  ctrl := &Controller{}

  view := NewView(ctrl, ctx)

  model := NewModel()
  model.AttachObserver(ObserverFunc(view.Update))

  ctrl.Model = model
  ctrl.View = view
  return ctrl
}

func (ctrl *Controller) SetPath(path string) {
  ctrl.Model.SetPath(strings.Replace(path, "file://", "", 1))
}

//  https://gist.github.com/icambridge/9708081
// https://talks.golang.org/2014/organizeio.slide#1
