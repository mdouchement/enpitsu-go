package app

import(
  "gopkg.in/qml.v1"
)

type QmlAdapter struct {
  ctx *qml.Context
}

func NewQmlAdapter(ctx *qml.Context) *QmlAdapter {
  return &QmlAdapter{
    ctx: ctx,
  }
}
