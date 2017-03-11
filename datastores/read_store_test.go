// This file is used to test if read model is working correctly.
// A read is always linked to a read
// He has basic read to join
package datastores

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestReadStore(t *testing.T) {
	store := Store()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")

	resi := store.Read()
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

	messageTest := Message{
		Content:   "Message test",
		IDUser:    userTest.IDUser,
		IDChannel: channelTest.IDChannel,
	}
	msi.Save(&messageTest, db)

	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("readStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("readStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
		read := Read{
			IDUser:    userTest.IDUser,
			IDChannel: channelTest.IDChannel,
			IDMessage: messageTest.IDMessage,
		}

		Convey("Given a correct read.", func() {
			appError := resi.Save(&read, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := resi.Save(&read, db)
				So(appError2, ShouldNotBeNil)
				// So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&read)
	})

	// Convey("Testing update function", t, func() {
	// 	dbError := u.NewLocAppError("readStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
	// 	alreadyExistError := u.NewLocAppError("readStoreImpl.Save", "save.transaction.create.already_exist", nil, "Read Name: electras")
	// 	read := Read{
	// 		Readname: "TesT2",
	// 		Password:    "test",
	// 		Email:       "test2@popcube.fr",
	// 		NickName:    "NickName",
	// 		FirstName:   "Test",
	// 		LastName:    "L",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	readNew := Read{
	// 		Readname: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		IDRole:      adminRole.IDRole,
	// 	}

	// 	appError := resi.Save(&read, db)
	// 	So(appError, ShouldBeNil)
	// 	So(appError, ShouldNotResemble, dbError)
	// 	So(appError, ShouldNotResemble, alreadyExistError)

	// 	Convey("Provided correct Read to modify should not return errors", func() {
	// 		appError := resi.Update(&read, &readNew, db)
	// 		readShouldResemble := readNew
	// 		readShouldResemble.WebID = read.WebID
	// 		readShouldResemble.IDRead = read.IDRead
	// 		readShouldResemble.LastUpdate = read.LastUpdate
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dbError)
	// 		So(appError, ShouldNotResemble, alreadyExistError)
	// 		So(read, ShouldResemble, readShouldResemble)
	// 	})

	// 	Convey("Provided wrong old Read to modify should result in old_read error", func() {
	// 		read.WebID = "TesT"
	// 		Convey("Incorrect ID read should return a read invalid id", func() {
	// 			appError := resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.WebID.app_error", nil, ""))
	// 		})
	// 		read.WebID = NewID()
	// 		Convey("Incorrect readname read should return error Invalid readname", func() {
	// 			read.Readname = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+read.WebID))
	// 			read.Readname = ""
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+read.WebID))
	// 			read.Readname = "xD/"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+read.WebID))
	// 			read.Readname = "xD\\"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+read.WebID))
	// 			read.Readname = "xD*"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+read.WebID))
	// 			read.Readname = "xD{"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+read.WebID))
	// 		})

	// 		Convey("Password can]t be empty", func() {
	// 			read.Password = ""
	// 			appError := resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readOld.PreSave", "model.read.is_valid.auth_data_pwd.app_error", nil, "read_webID="+read.WebID))
	// 		})
	// 	})

	// 	Convey("Provided wrong new Read to modify should result in old_read error", func() {
	// 		Convey("Incorrect readname read should return error Invalid readname", func() {
	// 			readNew.Readname = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readNew.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+readNew.WebID))
	// 			readNew.Readname = ""
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readNew.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+readNew.WebID))
	// 			readNew.Readname = "xD/"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readNew.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+readNew.WebID))
	// 			readNew.Readname = "xD\\"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readNew.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+readNew.WebID))
	// 			readNew.Readname = "xD*"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readNew.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+readNew.WebID))
	// 			readNew.Readname = "xD{"
	// 			appError = resi.Update(&read, &readNew, db)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Update.readNew.PreSave", "model.read.is_valid.Readname.app_error", nil, "read_webID="+readNew.WebID))
	// 		})
	// 	})

	// 	db.Delete(&read)
	// 	db.Delete(&readNew)
	// })

	// Convey("Testing Getters", t, func() {
	// 	read0 := Read{
	// 		Readname: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	read1 := Read{
	// 		Readname: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	read2 := Read{
	// 		Readname: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	read1New := Read{
	// 		Readname: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe, Jack, William, Avrell",
	// 		LastName:    "Dalton",
	// 		IDRole:      guestRole.IDRole,
	// 	}
	// 	read3 := Read{
	// 		Readname: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	read4 := Read{
	// 		Readname: "billythekid",
	// 		Password:    "chocolat",
	// 		Email:       "billythekid@popcube.fr",
	// 		NickName:    "Kid",
	// 		FirstName:   "Billy",
	// 		LastName:    "The Kid",
	// 		IDRole:      guestRole.IDRole,
	// 	}

	// 	resi.Save(&read0, db)
	// 	resi.Save(&read1, db)
	// 	resi.Update(&read1, &read1New, db)
	// 	resi.Save(&read2, db)
	// 	resi.Save(&read3, db)
	// 	resi.Save(&read4, db)

	// 	// Have to be after save so ID are up to date :O
	// 	readList := []Read{
	// 		read0,
	// 		read1,
	// 		read2,
	// 		read3,
	// 		read4,
	// 	}

	// 	admins := []Read{read0}
	// 	guests := []Read{read1, read4}
	// 	emptyList := []Read{}

	// 	Convey("We have to be able to find all reads in the db", func() {
	// 		reads := resi.GetAll(db)
	// 		So(reads, ShouldNotResemble, &emptyList)
	// 		So(reads, ShouldResemble, &readList)
	// 	})

	// 	Convey("We have to be able to find a read from is name", func() {
	// 		read := resi.GetByReadName(read0.Readname, db)
	// 		So(read, ShouldNotResemble, &Read{})
	// 		So(read, ShouldResemble, &read0)
	// 		read = resi.GetByReadName(read2.Readname, db)
	// 		So(read, ShouldNotResemble, &Read{})
	// 		So(read, ShouldResemble, &read2)
	// 		read = resi.GetByReadName(read3.Readname, db)
	// 		So(read, ShouldNotResemble, &Read{})
	// 		So(read, ShouldResemble, &read3)
	// 		read = resi.GetByReadName(read4.Readname, db)
	// 		So(read, ShouldNotResemble, &Read{})
	// 		So(read, ShouldResemble, &read4)
	// 		Convey("Should also work from updated value", func() {
	// 			read = resi.GetByReadName(read1New.Readname, db)
	// 			So(read, ShouldNotResemble, &Read{})
	// 			So(read, ShouldResemble, &read1)
	// 		})
	// 	})

	// 	Convey("We have to be able to find a read from his email", func() {
	// 		read := resi.GetByEmail(read0.Email, db)
	// 		So(read, ShouldNotResemble, &Read{})
	// 		So(read, ShouldResemble, &read0)
	// 		read = resi.GetByEmail(read2.Email, db)
	// 		So(read, ShouldNotResemble, &Read{})
	// 		So(read, ShouldResemble, &read2)
	// 		read = resi.GetByEmail(read3.Email, db)
	// 		So(read, ShouldResemble, &read3)
	// 		read = resi.GetByEmail(read4.Email, db)
	// 		So(read, ShouldNotResemble, &Read{})
	// 		So(read, ShouldResemble, &read4)
	// 	})

	// 	Convey("We have to be able to find an read from his Role", func() {
	// 		reads := resi.GetByRole(&adminRole, db)
	// 		So(reads, ShouldNotResemble, &Read{})
	// 		So(reads, ShouldResemble, &admins)
	// 		reads = resi.GetByRole(&guestRole, db)
	// 		So(reads, ShouldNotResemble, &Read{})
	// 		So(reads, ShouldResemble, &guests)

	// 	})

	// 	Convey("Searching for non existent read should return empty", func() {
	// 		read := resi.GetByReadName("fantome", db)
	// 		So(read, ShouldResemble, &Read{})
	// 	})

	// 	db.Delete(&read0)
	// 	db.Delete(&read1)
	// 	db.Delete(&read1New)
	// 	db.Delete(&read2)
	// 	db.Delete(&read3)

	// 	Convey("Searching all in empty table should return empty", func() {
	// 		reads := resi.GetAll(db)
	// 		So(reads, ShouldResemble, &[]Read{})
	// 	})
	// })

	// Convey("Testing delete read", t, func() {
	// 	dberror := u.NewLocAppError("readStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
	// 	read0 := Read{
	// 		Readname: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	read1 := Read{
	// 		Readname: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	read2 := Read{
	// 		Readname: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	read3 := Read{
	// 		Readname: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		IDRole:      standartRole.IDRole,
	// 	}

	// 	resi.Save(&read0, db)
	// 	resi.Save(&read1, db)
	// 	resi.Save(&read2, db)
	// 	resi.Save(&read3, db)

	// 	// Have to be after save so ID are up to date :O
	// 	// read3Old := read3
	// 	// readList1 := []Read{
	// 	// 	read0,
	// 	// 	read1,
	// 	// 	read2,
	// 	// 	read3Old,
	// 	// }

	// 	Convey("Deleting a known read should work", func() {
	// 		appError := resi.Delete(&read2, db)
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dberror)
	// 		So(resi.GetByReadName("moris", db), ShouldResemble, &Read{})
	// 	})

	// 	// Convey("Trying to delete from non conform read should return specific read error and should not delete reads.", func() {
	// 	// 	read3.ReadName = "Const"
	// 	// 	Convey("Too long or empty Name should return name error", func() {
	// 	// 		appError := resi.Delete(&read3, db)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Delete.read.PreSave", "model.read.readname.app_error", nil, ""))
	// 	// 		So(resi.GetAll(db), ShouldResemble, &readList1)
	// 	// 		read3.ReadName = "+alpha"
	// 	// 		appError = resi.Delete(&read3, db)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Delete.read.PreSave", "model.read.readname.app_error", nil, ""))
	// 	// 		So(resi.GetAll(db), ShouldResemble, &readList1)
	// 	// 		read3.ReadName = "alpha-numerique"
	// 	// 		appError = resi.Delete(&read3, db)standartRole
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("readStoreImpl.Delete.read.PreSave", "model.read.readname.app_error", nil, ""))
	// 	// 		So(resi.GetAll(db), ShouldResemble, &readList1)
	// 	// 	})
	// 	// })

	// 	db.Delete(&read0)
	// 	db.Delete(&read1)
	// 	db.Delete(&read2)
	// 	db.Delete(&read3)
	// })
	db.Delete(&userTest)
	db.Delete(&channelTest)
	db.Delete(&standartRole)
	db.Delete(&messageTest)
}
