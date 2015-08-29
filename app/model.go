package app

type Model struct {
  path string
  Observable
}

func NewModel() *Model {
  return new(Model)
}

func (m *Model) SetPath(path string) {
  m.path = path
  m.Publish("path", path)
}

func (m *Model) GetPath() string {
  return m.path
}
