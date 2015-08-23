import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.1
import "content"

ApplicationWindow {
  visible: true
  title: 'Enpitsu'

  width: 1280
  height: 720
  minimumHeight: 400
  minimumWidth: 600

  SplitView {
    anchors.fill: parent
    orientation: Qt.Vertical
    Layout.fillWidth: true

    Browser { id: browser }

    SplitView {
      MetadataToolBox { id: metadataToolBox }
      Canvas { id: canvas }
    }
  }
}
