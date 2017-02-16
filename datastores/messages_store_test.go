// This file is used to test if message model is working correctly.
// A message is always linked to a message
// He has basic message to join
package datastores

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestMessageStore(t *testing.T) {
	store := NewStore()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")

	msi := store.Message()
	usi := store.User()
	rsi := store.Role()
	csi := store.Channel()

	standartRole := Role{
		RoleName:      randStringBytes(10),
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	rsi.Save(&standartRole, db)

	userTest := User{
		Username:  randStringBytes(10),
		Password:  "test",
		Email:     "test@popcube.fr",
		NickName:  "NickName",
		FirstName: "Test",
		LastName:  "L",
		IDRole:    standartRole.IDRole,
	}
	usi.Save(&userTest, db)

	channelTest := Channel{
		ChannelName: randStringBytes(10),
		Type:        "audio",
		Private:     false,
		Description: "Testing channel description :O",
		Subject:     "Sujet",
		Avatar:      "jesuiscool.svg",
	}
	csi.Save(&channelTest, db)

	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("messageStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("messageStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
		message := Message{
			Content:   "Message test",
			IDUser:    userTest.IDUser,
			IDChannel: channelTest.IDChannel,
		}

		Convey("Given a correct message.", func() {
			appError := msi.Save(&message, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := msi.Save(&message, db)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&message)
	})

	// Convey("Testing update function", t, func() {
	// 	dbError := u.NewLocAppError("messageStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
	// 	alreadyExistError := u.NewLocAppError("messageStoreImpl.Save", "save.transaction.create.already_exist", nil, "Message Name: electras")
	// 	message := Message{
	// 		Messagename: "TesT2",
	// 		Password:    "test",
	// 		Email:       "test2@popcube.fr",
	// 		NickName:    "NickName",
	// 		FirstName:   "Test",
	// 		LastName:    "L",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	messageNew := Message{
	// 		Messagename: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		Locale:      "vn_VN",
	// 		IDRole:      adminRole.IDRole,
	// 	}

	// 	appError := msi.Save(&message, db)
	// 	So(appError, ShouldBeNil)
	// 	So(appError, ShouldNotResemble, dbError)
	// 	So(appError, ShouldNotResemble, alreadyExistError)

	// 	Convey("Provided correct Message to modify should not return errors", func() {
	// 		appError := msi.Update(&message, &messageNew, db)
	// 		messageShouldResemble := messageNew
	// 		messageShouldResemble.WebID = message.WebID
	// 		messageShouldResemble.IDMessage = message.IDMessage
	// 		messageShouldResemble.LastUpdate = message.LastUpdate
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dbError)
	// 		So(appError, ShouldNotResemble, alreadyExistError)
	// 		So(message, ShouldResemble, messageShouldResemble)
	// 	})

	// 	Convey("Provided wrong old Message to modify should result in old_message error", func() {
	// 		message.WebID = "TesT"
	// 		Convey("Incorrect ID message should return a message invalid id", func() {
	// 			appError := msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.WebID.app_error", nil, ""))
	// 		})
	// 		message.WebID = NewID()
	// 		Convey("Incorrect messagename message should return error Invalid messagename", func() {
	// 			message.Messagename = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+message.WebID))
	// 			message.Messagename = ""
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+message.WebID))
	// 			message.Messagename = "xD/"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+message.WebID))
	// 			message.Messagename = "xD\\"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+message.WebID))
	// 			message.Messagename = "xD*"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+message.WebID))
	// 			message.Messagename = "xD{"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+message.WebID))
	// 		})

	// 		Convey("Password can]t be empty", func() {
	// 			message.Password = ""
	// 			appError := msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageOld.PreSave", "model.message.is_valid.auth_data_pwd.app_error", nil, "message_webID="+message.WebID))
	// 		})
	// 	})

	// 	Convey("Provided wrong new Message to modify should result in old_message error", func() {
	// 		Convey("Incorrect messagename message should return error Invalid messagename", func() {
	// 			messageNew.Messagename = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageNew.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+messageNew.WebID))
	// 			messageNew.Messagename = ""
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageNew.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+messageNew.WebID))
	// 			messageNew.Messagename = "xD/"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageNew.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+messageNew.WebID))
	// 			messageNew.Messagename = "xD\\"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageNew.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+messageNew.WebID))
	// 			messageNew.Messagename = "xD*"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageNew.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+messageNew.WebID))
	// 			messageNew.Messagename = "xD{"
	// 			appError = msi.Update(&message, &messageNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Update.messageNew.PreSave", "model.message.is_valid.Messagename.app_error", nil, "message_webID="+messageNew.WebID))
	// 		})
	// 	})

	// 	db.Delete(&message)
	// 	db.Delete(&messageNew)
	// })

	// Convey("Testing Getters", t, func() {
	// 	message0 := Message{
	// 		Messagename: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		Locale:      "en_EN",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	message1 := Message{
	// 		Messagename: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	message2 := Message{
	// 		Messagename: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		Locale:      "fr_FR",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	message1New := Message{
	// 		Messagename: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe, Jack, William, Avrell",
	// 		LastName:    "Dalton",
	// 		Locale:      "en_EN",
	// 		IDRole:      guestRole.IDRole,
	// 	}
	// 	message3 := Message{
	// 		Messagename: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	message4 := Message{
	// 		Messagename: "billythekid",
	// 		Password:    "chocolat",
	// 		Email:       "billythekid@popcube.fr",
	// 		NickName:    "Kid",
	// 		FirstName:   "Billy",
	// 		LastName:    "The Kid",
	// 		Locale:      "en_EN",
	// 		IDRole:      guestRole.IDRole,
	// 	}

	// 	msi.Save(&message0, db)
	// 	msi.Save(&message1, db)
	// 	msi.Update(&message1, &message1New, db)
	// 	msi.Save(&message2, db)
	// 	msi.Save(&message3, db)
	// 	msi.Save(&message4, db)

	// 	// Have to be after save so ID are up to date :O
	// 	messageList := []Message{
	// 		message0,
	// 		message1,
	// 		message2,
	// 		message3,
	// 		message4,
	// 	}

	// 	admins := []Message{message0}
	// 	guests := []Message{message1, message4}
	// 	emptyList := []Message{}

	// 	Convey("We have to be able to find all messages in the db", func() {
	// 		messages := msi.GetAll(db)
	// 		So(messages, ShouldNotResemble, &emptyList)
	// 		So(messages, ShouldResemble, &messageList)
	// 	})

	// 	Convey("We have to be able to find a message from is name", func() {
	// 		message := msi.GetByMessageName(message0.Messagename, db)
	// 		So(message, ShouldNotResemble, &Message{})
	// 		So(message, ShouldResemble, &message0)
	// 		message = msi.GetByMessageName(message2.Messagename, db)
	// 		So(message, ShouldNotResemble, &Message{})
	// 		So(message, ShouldResemble, &message2)
	// 		message = msi.GetByMessageName(message3.Messagename, db)
	// 		So(message, ShouldNotResemble, &Message{})
	// 		So(message, ShouldResemble, &message3)
	// 		message = msi.GetByMessageName(message4.Messagename, db)
	// 		So(message, ShouldNotResemble, &Message{})
	// 		So(message, ShouldResemble, &message4)
	// 		Convey("Should also work from updated value", func() {
	// 			message = msi.GetByMessageName(message1New.Messagename, db)
	// 			So(message, ShouldNotResemble, &Message{})
	// 			So(message, ShouldResemble, &message1)
	// 		})
	// 	})

	// 	Convey("We have to be able to find a message from his email", func() {
	// 		message := msi.GetByEmail(message0.Email, db)
	// 		So(message, ShouldNotResemble, &Message{})
	// 		So(message, ShouldResemble, &message0)
	// 		message = msi.GetByEmail(message2.Email, db)
	// 		So(message, ShouldNotResemble, &Message{})
	// 		So(message, ShouldResemble, &message2)
	// 		message = msi.GetByEmail(message3.Email, db)
	// 		So(message, ShouldResemble, &message3)
	// 		message = msi.GetByEmail(message4.Email, db)
	// 		So(message, ShouldNotResemble, &Message{})
	// 		So(message, ShouldResemble, &message4)
	// 	})

	// 	Convey("We have to be able to find an message from his Role", func() {
	// 		messages := msi.GetByRole(&adminRole, db)
	// 		So(messages, ShouldNotResemble, &Message{})
	// 		So(messages, ShouldResemble, &admins)
	// 		messages = msi.GetByRole(&guestRole, db)
	// 		So(messages, ShouldNotResemble, &Message{})
	// 		So(messages, ShouldResemble, &guests)

	// 	})

	// 	Convey("Searching for non existent message should return empty", func() {
	// 		message := msi.GetByMessageName("fantome", db)
	// 		So(message, ShouldResemble, &Message{})
	// 	})

	// 	db.Delete(&message0)
	// 	db.Delete(&message1)
	// 	db.Delete(&message1New)
	// 	db.Delete(&message2)
	// 	db.Delete(&message3)

	// 	Convey("Searching all in empty table should return empty", func() {
	// 		messages := msi.GetAll(db)
	// 		So(messages, ShouldResemble, &[]Message{})
	// 	})
	// })

	// Convey("Testing delete message", t, func() {
	// 	dberror := u.NewLocAppError("messageStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
	// 	message0 := Message{
	// 		Messagename: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		Locale:      "en_EN",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	message1 := Message{
	// 		Messagename: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	message2 := Message{
	// 		Messagename: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		Locale:      "fr_FR",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	message3 := Message{
	// 		Messagename: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}

	// 	msi.Save(&message0, db)
	// 	msi.Save(&message1, db)
	// 	msi.Save(&message2, db)
	// 	msi.Save(&message3, db)

	// 	// Have to be after save so ID are up to date :O
	// 	// message3Old := message3
	// 	// messageList1 := []Message{
	// 	// 	message0,
	// 	// 	message1,
	// 	// 	message2,
	// 	// 	message3Old,
	// 	// }

	// 	Convey("Deleting a known message should work", func() {
	// 		appError := msi.Delete(&message2, db)
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dberror)
	// 		So(msi.GetByMessageName("moris", db), ShouldResemble, &Message{})
	// 	})

	// 	// Convey("Trying to delete from non conform message should return specific message error and should not delete messages.", func() {
	// 	// 	message3.MessageName = "Const"
	// 	// 	Convey("Too long or empty Name should return name error", func() {
	// 	// 		appError := msi.Delete(&message3, db)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Delete.message.PreSave", "model.message.messagename.app_error", nil, ""))
	// 	// 		So(msi.GetAll(db), ShouldResemble, &messageList1)
	// 	// 		message3.MessageName = "+alpha"
	// 	// 		appError = msi.Delete(&message3, db)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Delete.message.PreSave", "model.message.messagename.app_error", nil, ""))
	// 	// 		So(msi.GetAll(db), ShouldResemble, &messageList1)
	// 	// 		message3.MessageName = "alpha-numerique"
	// 	// 		appError = msi.Delete(&message3, db)standartRole
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("messageStoreImpl.Delete.message.PreSave", "model.message.messagename.app_error", nil, ""))
	// 	// 		So(msi.GetAll(db), ShouldResemble, &messageList1)
	// 	// 	})
	// 	// })

	// 	db.Delete(&message0)
	// 	db.Delete(&message1)
	// 	db.Delete(&message2)
	// 	db.Delete(&message3)
	// })
	db.Delete(&userTest)
	db.Delete(&channelTest)
	db.Delete(&standartRole)
}
