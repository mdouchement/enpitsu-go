package app

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var extensions []string = []string{"jpeg", "JPEG", "jpg", "JPG", "png", "PNG", "tif", "TIF", "tiff", "TIFF"}

type ImageBuffer struct {
	Filename    string `json:"filename"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MetadataBuffer struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Images      []ImageBuffer `json:"images"`
}

type item struct {
	ToFilePath string
	Buffer     MetadataBuffer
}

type Metadata struct {
	path       string
	buffer     *MetadataBuffer
	queue      chan item
	QuitBackup chan struct{}
}

func NewMetadata(path string) *Metadata {
	mtdt := &Metadata{
		path:  path,
		queue: make(chan item),
	}
	mtdt.initMetadata()

	go generate(mtdt.queue)
	mtdt.backup()

	return mtdt
}

func (mtdt *Metadata) HasRunningBackup() bool {
	return mtdt.QuitBackup != nil
}

func (mtdt *Metadata) GalleryHeaders() map[string]string {
	gh := make(map[string]string)
	gh["title"] = mtdt.buffer.Title
	gh["description"] = mtdt.buffer.Description
	return gh
}

func (mtdt *Metadata) Image(index int) map[string]string {
	image := make(map[string]string)
	image["title"] = mtdt.buffer.Images[index].Title
	image["description"] = mtdt.buffer.Images[index].Description
	image["filename"] = mtdt.buffer.Images[index].Filename
	return image
}

func (mtdt *Metadata) ImagePath(index int) string {
	return filepath.Join(mtdt.path, mtdt.buffer.Images[index].Filename)
}

func (mtdt *Metadata) UpdateAttributes(value map[string]string, index int) {
	if index == -666 {
		mtdt.buffer.Title = value["title"]
		mtdt.buffer.Description = value["description"]
	} else {
		mtdt.buffer.Images[index].Title = value["title"]
		mtdt.buffer.Images[index].Description = value["description"]
	}
}

func (mtdt *Metadata) NbOfImages() int {
	return len(mtdt.buffer.Images)
}

func (mtdt *Metadata) Generate(fname ...string) {
	filename := "metadata.json"
	if len(fname) > 0 {
		filename = fname[0]
	}

	mtdt.queue <- item{
		ToFilePath: filepath.Join(mtdt.path, filename),
		Buffer:     *mtdt.buffer,
	}
}

func (mtdt *Metadata) initMetadata() {
	if mtdt.hasFile() {
		mtdt.loadFromFile()
	} else {
		mtdt.newOne()
	}
}

func (mtdt *Metadata) hasFile() bool {
	_, err := os.Stat(filepath.Join(mtdt.path, "metadata.json"))
	return err == nil
}

func (mtdt *Metadata) loadFromFile() {
	rawBuffer, err := ioutil.ReadFile(filepath.Join(mtdt.path, "metadata.json"))
	check(err)

	mtdt.buffer = new(MetadataBuffer)
	err = json.Unmarshal(rawBuffer, mtdt.buffer)
	check(err)

	mtdt.consistencyCheck()
}

func (mtdt *Metadata) consistencyCheck() {
	filenames, exists := mtdt.loadImages()
	w := 0 // write index
	bufferedImages := map[string]bool{}

	// Remove from buffer images that no longer exist
	for _, image := range mtdt.buffer.Images {
		if !exists[image.Filename] {
			continue
		}
		bufferedImages[image.Filename] = true
		mtdt.buffer.Images[w] = image
		w++
	}
	mtdt.buffer.Images = mtdt.buffer.Images[:w]

	// Append new images into the buffer
	for _, filename := range filenames {
		if bufferedImages[filename] {
			continue
		}
		mtdt.buffer.Images = append(mtdt.buffer.Images, ImageBuffer{Filename: filename})
	}
}

func (mtdt *Metadata) newOne() {
	mtdt.buffer = new(MetadataBuffer)
	mtdt.buffer.Images = mtdt.buildImagesMetadata()
}

func (mtdt *Metadata) buildImagesMetadata() (ibs []ImageBuffer) {
	filenames, _ := mtdt.loadImages()
	for _, filename := range filenames {
		ibs = append(ibs, ImageBuffer{Filename: filename})
	}
	return
}

func (mtdt *Metadata) loadImages() (filenames []string, exists map[string]bool) {
	exists = map[string]bool{}
	files := []string{}
	for _, extension := range extensions {
		var err error = nil
		list, err := filepath.Glob(filepath.Join(mtdt.path, "*."+extension))
		check(err)
		files = append(files, list...)
	}

	for _, file := range files {
		_, filename := filepath.Split(file)
		filenames = append(filenames, filename)
		exists[filename] = true
	}
	return
}

func (mtdt *Metadata) backup() {
	ticker := time.NewTicker(5 * time.Minute)
	mtdt.QuitBackup = make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				filename := "metadata_backup_" + time.Now().Format("20060102150405") + ".json"
				mtdt.Generate(filename)
			case <-mtdt.QuitBackup:
				ticker.Stop()
				return
			}
		}
	}()
}

func generate(queue <-chan item) {
	for item := range queue {
		jsonData, err := json.MarshalIndent(item.Buffer, "", "  ")
		check(err)

		fo, err := os.Create(item.ToFilePath)
		check(err)
		defer fo.Close()

		w := bufio.NewWriter(fo)
		_, err = w.WriteString(string(jsonData))
		check(err)
		w.Flush()
	}
}
