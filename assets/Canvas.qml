import QtQuick 2.2
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.1

Rectangle {
  id: canvas
  color: 'lightgray'

  ColumnLayout {
    Label {
      id: imageFile
      text: image_filename
    }

    Image {
      source: view.imagePath(imageFile.text)
      smooth: true
      cache: false
      fillMode: Image.PreserveAspectFit
      horizontalAlignment: Image.AlignHCenter
    }
  }
}
