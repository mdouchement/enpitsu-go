import QtQuick 2.4
import QtQuick.Controls 1.3
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
    }

    Label {
      text: "Gallery's description:"
    }

    TextArea {
      id: galleryDescription
      Layout.fillWidth: true
    }

    Label {
      text: "Image's title:"
    }

    TextField {
      id: imageTitle
      Layout.fillWidth: true
    }

    Label {
      text: "Gallery's description:"
    }

    TextArea {
      id: imageDescription
      Layout.fillWidth: true
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
      }

      Button {
        id: next
        text: 'Next'
      }
    }

    Button {
      id: generate
      text: 'Generate'
    }
  }
}
