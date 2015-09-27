import QtQuick 2.2
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.1

Rectangle {
  id: canvas
  color: 'gray'

  ColumnLayout {
    anchors.fill: parent

    Label {
      id: imageFile
      anchors {
        top: parent.top
        margins: 4
        horizontalCenter: parent.horizontalCenter
      }
      text: image_filename
    }

    Image {
      anchors {
        top: imageFile.bottom
        bottom: parent.bottom
        left: parent.left
        right: parent.right
      }
      source: view.imagePath(imageFile.text)
      smooth: true
      cache: false
      fillMode: Image.PreserveAspectFit
      horizontalAlignment: Image.AlignHCenter
    }
  }
}
