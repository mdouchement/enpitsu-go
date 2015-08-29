import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.1
import QtQuick.Dialogs 1.1

Rectangle {
  id: browser
  height: 80
  color: 'lightsteelblue'

  ColumnLayout {
    anchors.fill: parent

    Label {
      text: "Gallery's path:"
    }

    RowLayout {
      TextField {
        id: filePath
        Layout.fillWidth: true
        text: path
      }

      Button {
        id: browse
        text: 'Browse'
        onClicked: {
          fileDialog.open()
        }
      }
    }
  }

  FileDialog {
    id: fileDialog
    title: 'Please choose a file'
    selectFolder: true
    // folder: shortcuts.home
    onAccepted: {
      view.setPath(fileDialog.folder)
    }
    Component.onCompleted: visible = false
  }
}
