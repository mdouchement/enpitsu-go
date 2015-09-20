package app

import (
	"gopkg.in/qml.v1"
	"strings"
)

type Controller struct {
	Model *Model
	View  *View
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

func (ctrl *Controller) LoadPath(path string) {
	ctrl.Model.LoadPath(strings.Replace(path, "file://", "", 1))
}

func (ctrl *Controller) ImagePath() string {
	var path string
	if ctrl.Model.HaveMetadata() {
		path = "file://" + ctrl.Model.ImagePath()
	} else {
		path = "qrc:///assets/oawsooo.png"
	}
	return path
}

func (ctrl *Controller) Copy(buffer map[string]string) {
	ctrl.Model.Copy(buffer)
}

func (ctrl *Controller) Paste() {
	ctrl.Model.Paste()
}

func (ctrl *Controller) NextImage() {
	if ctrl.Model.HaveMetadata() && ctrl.Model.IsIndexValid("next") {
		ctrl.Model.NextImage()
	}
}

func (ctrl *Controller) PreviousImage() {
	if ctrl.Model.HaveMetadata() && ctrl.Model.IsIndexValid("previous") {
		ctrl.Model.PreviousImage()
	}
}

func (ctrl *Controller) Generate() {
	if ctrl.Model.HaveMetadata() {
		ctrl.Model.Generate()
	}
}

func (ctrl *Controller) UpdateMetadata(kind string, value map[string]string) {
	if ctrl.Model.HaveMetadata() {
		ctrl.Model.UpdateAttributes(kind, value)
	}
}
