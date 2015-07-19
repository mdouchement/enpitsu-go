package main

type Model struct {
  path string
}

func NewModel() *Model {
  return new(Model)
}

func (m *Model) SetPath(path string) {
  m.path = path
}

func (m *Model) GetPath() string {
  return m.path
}
