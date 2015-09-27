import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.1
import "."

ApplicationWindow {
  visible: true
  title: 'Enpitsu'

  width: 1280
  height: 720
  minimumHeight: 400
  minimumWidth: 600
  color: 'lightgray'

  SplitView {
    orientation: Qt.Vertical
    Layout.fillWidth: true
    anchors.fill: parent

    Browser {
      id: browser
      anchors.top: parent.top
    }

    SplitView {
      anchors.bottom: parent.bottom

      MetadataToolBox {
        id: metadataToolBox
        anchors.left: parent.left
      }

      Canvas {
        id: canvas
        anchors.right: parent.right
      }
    }
  }
}
