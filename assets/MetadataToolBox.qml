import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.1

Rectangle {
  id: metadataToolBox
  Layout.minimumWidth: 304
  Layout.maximumWidth: 400
  color: 'lightblue'

  ColumnLayout {
    anchors.fill: parent

    Label {
      text: "Gallery's title:"
    }

    TextField {
      id: galleryTitle
      Layout.fillWidth: true
      text: gallery_title

      onTextChanged: {
        view.updateGalleryMeta(galleryTitle.text, galleryDescription.text);
      }
    }

    Label {
      text: "Gallery's description:"
    }

    TextArea {
      id: galleryDescription
      Layout.fillWidth: true
      text: gallery_description

      onTextChanged: {
        view.updateGalleryMeta(galleryTitle.text, galleryDescription.text);
      }
    }

    Label {
      text: "Image's title:"
    }

    TextField {
      id: imageTitle
      Layout.fillWidth: true
      text: image_title

      onTextChanged: {
        view.updateImageMeta(imageTitle.text, imageDescription.text);
      }
    }

    Label {
      text: "Gallery's description:"
    }

    TextArea {
      id: imageDescription
      Layout.fillWidth: true
      text: image_description

      onTextChanged: {
        view.updateImageMeta(imageTitle.text, imageDescription.text);
      }
    }

    RowLayout {
      Button {
        id: copy
        text: 'Copy'
      }

      Button {
        id: paste
        text: 'Paste'
      }

      Button {
        id: previous
        text: 'Previous'

        onClicked: {
          view.previousImage();
        }
      }

      Button {
        id: next
        text: 'Next'

        onClicked: {
          view.nextImage();
        }
      }
    }

    Button {
      id: generate
      text: 'Generate'

      onClicked: {
        view.generate();
      }
    }
  }
}
