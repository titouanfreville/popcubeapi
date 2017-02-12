package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestFolderModel(t *testing.T) {
	userTest := User{
		WebID:              NewID(),
		LastUpdate:          10,
		Deleted:            true,
		Username:           "l",
		Password:           "test",
		Email:              "test@popcube.fr",
		EmailVerified:      true,
		NickName:           "NickName",
		FirstName:          "Test",
		LastName:           "L",
		Role:               Owner,
		LastPasswordUpdate: 20,
		FailedAttempts:     1,
		Locale:             "vi",
	}

	channelTest := Channel{
		WebID:       NewID(),
		ChannelName: "electra",
		LastUpdate:   GetMillis(),
		Type:        "audio",
		Private:     false,
		Description: "Testing channel description :O",
		Subject:     "Sujet",
		Avatar:      "jesuiscool.svg",
	}

	messageTest := Message{
		Date:    GetMillis(),
		Content: "Message test",
		Creator: userTest,
		Channel: channelTest,
	}

	Convey("Testing IsValid function", t, func() {
		nameError := u.NewLocAppError("Folder.IsValid", "model.folder.name.app_error", nil, "")
		linkError := u.NewLocAppError("Folder.IsValid", "model.folder.link.app_error", nil, "")
		typeError := u.NewLocAppError("Folder.IsValid", "model.folder.type.app_error", nil, "")
		messageError := u.NewLocAppError("Folder.IsValid", "model.folder.message.app_error", nil, "")
		Convey("Given a correct folders. Should be validated", func() {
			folder := Folder{
				Name:    "Je suis .... Cailloux",
				Link:    "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type:    "Video",
				Message: messageTest,
			}
			So(folder.IsValid(), ShouldBeNil)
			So(folder.IsValid(), ShouldNotResemble, nameError)
			So(folder.IsValid(), ShouldNotResemble, linkError)
			So(folder.IsValid(), ShouldNotResemble, typeError)
			So(folder.IsValid(), ShouldNotResemble, messageError)
		})

		Convey("Given incorrect folders. Should be refused", func() {
			folder := Folder{
				Name:    "Je suis .... Cailloux",
				Link:    "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type:    "Video",
				Message: messageTest,
			}
			empty := Folder{}
			folder.Name = ""

			Convey("empty Name or folder should return name error", func() {
				So(folder.IsValid(), ShouldNotBeNil)
				So(folder.IsValid(), ShouldResemble, nameError)
				So(folder.IsValid(), ShouldNotResemble, linkError)
				So(folder.IsValid(), ShouldNotResemble, typeError)
				So(folder.IsValid(), ShouldNotResemble, messageError)
				So(empty.IsValid(), ShouldNotBeNil)
				So(empty.IsValid(), ShouldResemble, nameError)
				So(empty.IsValid(), ShouldNotResemble, linkError)
				So(empty.IsValid(), ShouldNotResemble, typeError)
				So(empty.IsValid(), ShouldNotResemble, messageError)
			})

			folder.Name = "Correct Name"
			folder.Link = ""

			Convey("Empty link should result in link error", func() {
				So(folder.IsValid(), ShouldNotBeNil)
				So(folder.IsValid(), ShouldNotResemble, nameError)
				So(folder.IsValid(), ShouldResemble, linkError)
				So(folder.IsValid(), ShouldNotResemble, typeError)
				So(folder.IsValid(), ShouldNotResemble, messageError)
			})

			folder.Link = "folder/corretc.xml"
			folder.Type = ""

			Convey("Empty type should result in type error", func() {
				So(folder.IsValid(), ShouldNotBeNil)
				So(folder.IsValid(), ShouldNotResemble, nameError)
				So(folder.IsValid(), ShouldNotResemble, linkError)
				So(folder.IsValid(), ShouldResemble, typeError)
				So(folder.IsValid(), ShouldNotResemble, messageError)
			})

			folder.Type = "xml"
			folder.Message = Message{}

			Convey("Empty message should result in message", func() {
				So(folder.IsValid(), ShouldNotBeNil)
				So(folder.IsValid(), ShouldNotResemble, nameError)
				So(folder.IsValid(), ShouldNotResemble, linkError)
				So(folder.IsValid(), ShouldNotResemble, typeError)
				So(folder.IsValid(), ShouldResemble, messageError)
			})
		})
	})

	Convey("Testing json VS folder transformations", t, func() {
		Convey("Given an folder", func() {
			folder := Folder{
				Name: "Je suis .... Cailloux",
				Link: "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type: "Video",
			}
			Convey("Transforming it in JSON then back to FOLDER should provide similar objects", func() {
				json := folder.ToJSON()
				newFolder := FolderFromJSON(strings.NewReader(json))
				So(newFolder, ShouldResemble, &folder)
			})
		})

		Convey("Given an folder list", func() {
			folder1 := Folder{
				Name: "Je suis .... Cailloux",
				Link: "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type: "Video",
			}
			folder2 := Folder{
				Name: "Somethi,g",
				Link: "folders/something.sql",
				Type: "sql",
			}
			folder3 := Folder{
				Name: "Should automatize type recon",
				Link: "folders/facepalm.svg",
				Type: "facepalm?",
			}
			folderList := []*Folder{&folder1, &folder2, &folder3}

			Convey("Transfoming it in JSON then back to FOLDER LIST shoud give ressembling objects", func() {
				json := FolderListToJSON(folderList)
				newFolderList := FolderListFromJSON(strings.NewReader(json))
				So(newFolderList, ShouldResemble, folderList)
			})

		})
	})

}
