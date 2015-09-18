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
  ctx.SetVar("gallery_title", "")
  ctx.SetVar("gallery_description", "")
  ctx.SetVar("image_title", "")
  ctx.SetVar("image_description", "")
  ctx.SetVar("image_filename", "")

  return view
}

func (view *View) Update(key string, value interface{}) {
  view.ctx.SetVar(key, value)
}

func (view *View) LoadPath(path string) {
  view.controller.LoadPath(path)
}

func (view *View) ImagePath() string {
  return view.controller.ImagePath()
}

func (view *View) NextImage() {
  view.controller.NextImage()
}

func (view *View) PreviousImage() {
  view.controller.PreviousImage()
}

func (view *View) Generate() {
  view.controller.Generate()
}

func (view *View) UpdateGalleryMeta(title, description string) {
  gh := map[string]string{
    "title": title,
    "description": description,
  }
  view.controller.UpdateMetadata("gallery_headers", gh)
}

func (view *View) UpdateImageMeta(title, description string) {
  image := map[string]string{
    "title": title,
    "description": description,
  }
  view.controller.UpdateMetadata("image", image)
}
