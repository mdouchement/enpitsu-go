package app

import(
  "gopkg.in/qml.v1"
)

type View struct {
  controller *Controller
  ctx *qml.Context
}

func NewView(controller *Controller, ctx *qml.Context) *View {
  view := &View{
    controller: controller,
    ctx: ctx,
  }

  // Initialize QML templates
  ctx.SetVar("view", view)
  ctx.SetVar("path", "")

  return view
}

func (view *View) Update(key string, value interface{}) {
  view.ctx.SetVar(key, value)
}

func (view *View) SetPath(path string) {
  view.controller.SetPath(path)
}
