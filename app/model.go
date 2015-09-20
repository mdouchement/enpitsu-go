package app

type Model struct {
	metadata  *Metadata
	index     int
	clipboard map[string]string
	Observable
}

func NewModel() *Model {
	return new(Model)
}

func (m *Model) LoadPath(path string) {
	m.metadata = NewMetadata(path)
	m.index = 0
	m.Publish("path", path)
	m.updateGallery()
}

func (m *Model) ImagePath() string {
	return m.metadata.ImagePath(m.index)
}

func (m *Model) Copy(buffer map[string]string) {
	m.clipboard = buffer
}

func (m *Model) Paste() {
	for key, value := range m.clipboard {
		m.Publish("image_"+key, value)
	}
}

func (m *Model) NextImage() {
	m.index += 1
	m.updateGallery()
}

func (m *Model) PreviousImage() {
	m.index -= 1
	m.updateGallery()
}

func (m *Model) HaveMetadata() bool {
	return m.metadata != nil
}

func (m *Model) IsIndexValid(action string) bool {
	index := m.index
	switch action {
	case "previous":
		index -= 1
	case "next":
		index += 1
	}
	return index > -1 && index < m.metadata.NbOfImages()
}

func (m *Model) UpdateAttributes(kind string, value map[string]string) {
	switch kind {
	case "gallery_headers":
		m.metadata.UpdateAttributes(value, -666)
	case "image":
		m.metadata.UpdateAttributes(value, m.index)
	}
}

func (m *Model) Generate() {
	m.metadata.Generate()
}

func (m *Model) updateGallery() {
	for key, value := range m.metadata.GalleryHeaders() {
		m.Publish("gallery_"+key, value)
	}
	for key, value := range m.metadata.Image(m.index) {
		m.Publish("image_"+key, value)
	}
}
