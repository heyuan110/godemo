import QtQuick 2.5
import QtQuick.Controls 2.0
import QtQuick.Dialogs 1.2

ApplicationWindow {
    id: loginWindow
    visible: true
    width: 640
    height: 480
    color: "#62b9d9"
    title: qsTr("OA Management System")
    flags: Qt.Window | Qt.FramelessWindowHint

    Image {
        id: image
        anchors.fill: parent
        fillMode: Image.Stretch
        source: "qrc:/assets/images/bg.jpg"
    }

    MouseArea{
            anchors.fill: parent
            property point clickPos:"0,0"
            onPressed: {
                clickPos = Qt.point(mouse.x,mouse.y)
            }
            onPositionChanged: {
                var delta = Qt.point(mouse.x-clickPos.x,mouse.y-clickPos.y)
                loginWindow.setX(loginWindow.x+delta.x)
                loginWindow.setY(loginWindow.y+delta.y)
            }
    }

    Rectangle {
        id: rectangle
        x: 112
        y: 40
        width: 434
        height: 261
        color: "#81ffffff"
        radius: 6
        anchors.verticalCenterOffset: 30
        anchors.verticalCenter: parent.verticalCenter
        anchors.horizontalCenterOffset: 0
        border.color: "#23df8282"
        anchors.horizontalCenter: parent.horizontalCenter

        Column {
            id: column
            x: 90
            y: 30
            width: 276
            height: 200
            spacing: 20

            Row {
                id: row
                width: 200
                height: 40
                spacing: 10
                anchors.top: parent.top
                anchors.topMargin: 0

                Text {
                    id: username_lbl
                    text: qsTr("UserName")
                    anchors.verticalCenter: parent.verticalCenter
                    font.pixelSize: 14
                }

                Rectangle {
                    id: rect1
                    width: 200
                    height: 40
                    color: "#ecf6e3"
                    border.width: 1

                    TextInput {
                        id: username_txt
                        x: 10
                        y: 0
                        width: 180
                        height: 40
                        text: qsTr("")
                        clip: true
                        focus: true
                        wrapMode: Text.NoWrap
                        anchors.verticalCenterOffset: 0
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignLeft
                        renderType: Text.QtRendering
                        anchors.verticalCenter: parent.verticalCenter
                        font.pixelSize: 14
                    }
                }
            }

            Row {
                id: row1
                y: 60
                width: 200
                height: 40
                spacing: 10

                Text {
                    id: pwd_lbl
                    text: qsTr("Password")
                    anchors.verticalCenter: parent.verticalCenter
                    horizontalAlignment: Text.AlignLeft
                    font.pixelSize: 14
                }

                Rectangle {
                    id: rect2
                    width: 200
                    height: 40
                    color: "#ecf6e3"
                    anchors.left: parent.left
                    anchors.leftMargin: 77
                    border.width: 1
                    TextInput {
                        id: pwd_txt
                        x: 10
                        y: 0
                        width: 180
                        height: 40
                        text: qsTr("")
                        clip: true
                        wrapMode: Text.NoWrap
                        echoMode: TextInput.Password
                        anchors.verticalCenter: parent.verticalCenter
                        font.pixelSize: 14
                        renderType: Text.QtRendering
                        anchors.verticalCenterOffset: 0
                        horizontalAlignment: Text.AlignLeft
                        verticalAlignment: Text.AlignVCenter
                    }
                }
            }



        }

        Button {
            id: button
            x: 244
            y: 192
            width: 100
            height: 40
            text: qsTr("Login")
            onClicked: {
                    alert_dia.open()
            }
        }

        Button {
            id: button1
            x: 91
            y: 192
            text: qsTr("Cancel")
            onClicked: {
                close()
            }
        }

        CheckBox {
            id: checkBox
            x: 261
            y: 146
            width: 115
            height: 33
            text: qsTr("Remember")
        }
    }

    Text {
        id: element
        x: 308
        y: 62
        text: qsTr("OA Management System")
        anchors.bottom: rectangle.top
        anchors.bottomMargin: 40
        anchors.horizontalCenter: parent.horizontalCenter
        font.italic: false
        font.bold: true
        font.pixelSize: 36
    }

    Dialog {
        id: alert_dia
        visible: false
        title: "Tips"
        Text {
            id: hello
            font.pixelSize: 14
            text: qsTr("User name："+username_txt.text+" \ncontinue to login?")
            anchors.centerIn: parent
        }
        height: 150
        standardButtons: StandardButton.Yes | StandardButton.No
        modality: Qt.ApplicationModal
        onYes:{
            alert_dia.close()
            busyIndicator.visible = true
            rectangle.enabled = false
            bridge.sendToGo("["+username_txt.text+","+pwd_txt.text+"]")
            //Qt.quit();
         }
    }

    Dialog {
        id: alert_loginerror
        visible: false
        title: "Tips"
        Text {
            id: alert_loginerror_txt
            font.pixelSize: 14
            text: qsTr("User name or password wrong!")
            anchors.centerIn: parent
        }
        height: 150
        standardButtons: StandardButton.Ok
        modality: Qt.ApplicationModal
    }

    BusyIndicator {
        id: busyIndicator
        x: 24
        y: 405
        visible: false
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.verticalCenter: parent.verticalCenter
    }

    Connections {
        target: bridge
        onSendToQml:{
            if(data == "login_error"){
                busyIndicator.visible = false
                rectangle.enabled = true
                alert_loginerror.visible=true
            }
        }
    }
}
