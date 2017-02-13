// This file is used to test if folder model is working correctly.
// A folder is always linked to a folder
// He has basic folder to join
package datastores

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestFolderStore(t *testing.T) {
	ds := DbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db

	fsi := NewFolderStore()
	usi := NewUserStore()
	rsi := NewRoleStore()
	csi := NewChannelStore()
	msi := NewMessageStore()

	standartRole := Role{
		RoleName:      "teststandartazd",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	rsi.Save(&standartRole, ds)

	userTest := User{
		Username:  "TesTazd",
		Password:  "test",
		Email:     "test@popcube.fr",
		NickName:  "NickName",
		FirstName: "Test",
		LastName:  "L",
		IDRole:    standartRole.IDRole,
	}
	usi.Save(&userTest, ds)

	channelTest := Channel{
		ChannelName: "electrasomega",
		Type:        "audio",
		Private:     false,
		Description: "Testing channel description :O",
		Subject:     "Sujet",
		Avatar:      "jesuiscool.svg",
	}
	csi.Save(&channelTest, ds)

	message := Message{
		Content:   "Folder test",
		IDUser:    userTest.IDUser,
		IDChannel: channelTest.IDChannel,
	}
	msi.Save(&message, ds)

	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("folderStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("folderStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
		folder := Folder{
			IDMessage: message.IDMessage,
			Type:      "png",
			Link:      "link/zelda",
			Name:      "something",
		}

		Convey("Given a correct folder.", func() {
			appError := fsi.Save(&folder, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := fsi.Save(&folder, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&folder)
	})

	// Convey("Testing update function", t, func() {
	// 	dbError := u.NewLocAppError("folderStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
	// 	alreadyExistError := u.NewLocAppError("folderStoreImpl.Save", "save.transaction.create.already_exist", nil, "Folder Name: electras")
	// 	folder := Folder{
	// 		Foldername: "TesT2",
	// 		Password:    "test",
	// 		Email:       "test2@popcube.fr",
	// 		NickName:    "NickName",
	// 		FirstName:   "Test",
	// 		LastName:    "L",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	folderNew := Folder{
	// 		Foldername: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		Locale:      "vn_VN",
	// 		IDRole:      adminRole.IDRole,
	// 	}

	// 	appError := fsi.Save(&folder, ds)
	// 	So(appError, ShouldBeNil)
	// 	So(appError, ShouldNotResemble, dbError)
	// 	So(appError, ShouldNotResemble, alreadyExistError)

	// 	Convey("Provided correct Folder to modify should not return errors", func() {
	// 		appError := fsi.Update(&folder, &folderNew, ds)
	// 		folderShouldResemble := folderNew
	// 		folderShouldResemble.WebID = folder.WebID
	// 		folderShouldResemble.IDFolder = folder.IDFolder
	// 		folderShouldResemble.LastUpdate = folder.LastUpdate
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dbError)
	// 		So(appError, ShouldNotResemble, alreadyExistError)
	// 		So(folder, ShouldResemble, folderShouldResemble)
	// 	})

	// 	Convey("Provided wrong old Folder to modify should result in old_folder error", func() {
	// 		folder.WebID = "TesT"
	// 		Convey("Incorrect ID folder should return a folder invalid id", func() {
	// 			appError := fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.WebID.app_error", nil, ""))
	// 		})
	// 		folder.WebID = NewID()
	// 		Convey("Incorrect foldername folder should return error Invalid foldername", func() {
	// 			folder.Foldername = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folder.WebID))
	// 			folder.Foldername = ""
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folder.WebID))
	// 			folder.Foldername = "xD/"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folder.WebID))
	// 			folder.Foldername = "xD\\"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folder.WebID))
	// 			folder.Foldername = "xD*"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folder.WebID))
	// 			folder.Foldername = "xD{"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folder.WebID))
	// 		})

	// 		Convey("Password can]t be empty", func() {
	// 			folder.Password = ""
	// 			appError := fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderOld.PreSave", "model.folder.is_valid.auth_data_pwd.app_error", nil, "folder_webID="+folder.WebID))
	// 		})
	// 	})

	// 	Convey("Provided wrong new Folder to modify should result in old_folder error", func() {
	// 		Convey("Incorrect foldername folder should return error Invalid foldername", func() {
	// 			folderNew.Foldername = "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone"
	// 			appError := fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderNew.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folderNew.WebID))
	// 			folderNew.Foldername = ""
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderNew.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folderNew.WebID))
	// 			folderNew.Foldername = "xD/"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderNew.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folderNew.WebID))
	// 			folderNew.Foldername = "xD\\"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderNew.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folderNew.WebID))
	// 			folderNew.Foldername = "xD*"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderNew.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folderNew.WebID))
	// 			folderNew.Foldername = "xD{"
	// 			appError = fsi.Update(&folder, &folderNew, ds)
	// 			So(appError, ShouldNotBeNil)
	// 			So(appError, ShouldNotResemble, dbError)
	// 			So(appError, ShouldNotResemble, alreadyExistError)
	// 			So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Update.folderNew.PreSave", "model.folder.is_valid.Foldername.app_error", nil, "folder_webID="+folderNew.WebID))
	// 		})
	// 	})

	// 	db.Delete(&folder)
	// 	db.Delete(&folderNew)
	// })

	// Convey("Testing Getters", t, func() {
	// 	folder0 := Folder{
	// 		Foldername: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		Locale:      "en_EN",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	folder1 := Folder{
	// 		Foldername: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	folder2 := Folder{
	// 		Foldername: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		Locale:      "fr_FR",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	folder1New := Folder{
	// 		Foldername: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe, Jack, William, Avrell",
	// 		LastName:    "Dalton",
	// 		Locale:      "en_EN",
	// 		IDRole:      guestRole.IDRole,
	// 	}
	// 	folder3 := Folder{
	// 		Foldername: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	folder4 := Folder{
	// 		Foldername: "billythekid",
	// 		Password:    "chocolat",
	// 		Email:       "billythekid@popcube.fr",
	// 		NickName:    "Kid",
	// 		FirstName:   "Billy",
	// 		LastName:    "The Kid",
	// 		Locale:      "en_EN",
	// 		IDRole:      guestRole.IDRole,
	// 	}

	// 	fsi.Save(&folder0, ds)
	// 	fsi.Save(&folder1, ds)
	// 	fsi.Update(&folder1, &folder1New, ds)
	// 	fsi.Save(&folder2, ds)
	// 	fsi.Save(&folder3, ds)
	// 	fsi.Save(&folder4, ds)

	// 	// Have to be after save so ID are up to date :O
	// 	folderList := []Folder{
	// 		folder0,
	// 		folder1,
	// 		folder2,
	// 		folder3,
	// 		folder4,
	// 	}

	// 	admins := []Folder{folder0}
	// 	guests := []Folder{folder1, folder4}
	// 	emptyList := []Folder{}

	// 	Convey("We have to be able to find all folders in the db", func() {
	// 		folders := fsi.GetAll(ds)
	// 		So(folders, ShouldNotResemble, &emptyList)
	// 		So(folders, ShouldResemble, &folderList)
	// 	})

	// 	Convey("We have to be able to find a folder from is name", func() {
	// 		folder := fsi.GetByFolderName(folder0.Foldername, ds)
	// 		So(folder, ShouldNotResemble, &Folder{})
	// 		So(folder, ShouldResemble, &folder0)
	// 		folder = fsi.GetByFolderName(folder2.Foldername, ds)
	// 		So(folder, ShouldNotResemble, &Folder{})
	// 		So(folder, ShouldResemble, &folder2)
	// 		folder = fsi.GetByFolderName(folder3.Foldername, ds)
	// 		So(folder, ShouldNotResemble, &Folder{})
	// 		So(folder, ShouldResemble, &folder3)
	// 		folder = fsi.GetByFolderName(folder4.Foldername, ds)
	// 		So(folder, ShouldNotResemble, &Folder{})
	// 		So(folder, ShouldResemble, &folder4)
	// 		Convey("Should also work from updated value", func() {
	// 			folder = fsi.GetByFolderName(folder1New.Foldername, ds)
	// 			So(folder, ShouldNotResemble, &Folder{})
	// 			So(folder, ShouldResemble, &folder1)
	// 		})
	// 	})

	// 	Convey("We have to be able to find a folder from his email", func() {
	// 		folder := fsi.GetByEmail(folder0.Email, ds)
	// 		So(folder, ShouldNotResemble, &Folder{})
	// 		So(folder, ShouldResemble, &folder0)
	// 		folder = fsi.GetByEmail(folder2.Email, ds)
	// 		So(folder, ShouldNotResemble, &Folder{})
	// 		So(folder, ShouldResemble, &folder2)
	// 		folder = fsi.GetByEmail(folder3.Email, ds)
	// 		So(folder, ShouldResemble, &folder3)
	// 		folder = fsi.GetByEmail(folder4.Email, ds)
	// 		So(folder, ShouldNotResemble, &Folder{})
	// 		So(folder, ShouldResemble, &folder4)
	// 	})

	// 	Convey("We have to be able to find an folder from his Role", func() {
	// 		folders := fsi.GetByRole(&adminRole, ds)
	// 		So(folders, ShouldNotResemble, &Folder{})
	// 		So(folders, ShouldResemble, &admins)
	// 		folders = fsi.GetByRole(&guestRole, ds)
	// 		So(folders, ShouldNotResemble, &Folder{})
	// 		So(folders, ShouldResemble, &guests)

	// 	})

	// 	Convey("Searching for non existent folder should return empty", func() {
	// 		folder := fsi.GetByFolderName("fantome", ds)
	// 		So(folder, ShouldResemble, &Folder{})
	// 	})

	// 	db.Delete(&folder0)
	// 	db.Delete(&folder1)
	// 	db.Delete(&folder1New)
	// 	db.Delete(&folder2)
	// 	db.Delete(&folder3)

	// 	Convey("Searching all in empty table should return empty", func() {
	// 		folders := fsi.GetAll(ds)
	// 		So(folders, ShouldResemble, &[]Folder{})
	// 	})
	// })

	// Convey("Testing delete folder", t, func() {
	// 	dberror := u.NewLocAppError("folderStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
	// 	folder0 := Folder{
	// 		Foldername: "lucky",
	// 		Password:    "lucke",
	// 		Email:       "luckylucke@popcube.fr",
	// 		NickName:    "LL",
	// 		FirstName:   "Luky",
	// 		LastName:    "Luke",
	// 		Locale:      "en_EN",
	// 		IDRole:      adminRole.IDRole,
	// 	}
	// 	folder1 := Folder{
	// 		Foldername: "daltons",
	// 		Password:    "dalton",
	// 		Email:       "daltonsbrothers@popcube.fr",
	// 		NickName:    "thebrothers",
	// 		FirstName:   "Joe",
	// 		LastName:    "Dalton",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}
	// 	folder2 := Folder{
	// 		Foldername: "moris",
	// 		Password:    "gossiny",
	// 		Email:       "moris&gossiny@popcube.fr",
	// 		NickName:    "Moris",
	// 		FirstName:   "Moris",
	// 		LastName:    "Gossiny",
	// 		Locale:      "fr_FR",
	// 		IDRole:      ownerRole.IDRole,
	// 	}
	// 	folder3 := Folder{
	// 		Foldername: "jolly",
	// 		Password:    "jumper",
	// 		Email:       "jollyjumper@popcube.fr",
	// 		NickName:    "JJ",
	// 		FirstName:   "Jolly",
	// 		LastName:    "Jumper",
	// 		Locale:      "en_EN",
	// 		IDRole:      standartRole.IDRole,
	// 	}

	// 	fsi.Save(&folder0, ds)
	// 	fsi.Save(&folder1, ds)
	// 	fsi.Save(&folder2, ds)
	// 	fsi.Save(&folder3, ds)

	// 	// Have to be after save so ID are up to date :O
	// 	// folder3Old := folder3
	// 	// folderList1 := []Folder{
	// 	// 	folder0,
	// 	// 	folder1,
	// 	// 	folder2,
	// 	// 	folder3Old,
	// 	// }

	// 	Convey("Deleting a known folder should work", func() {
	// 		appError := fsi.Delete(&folder2, ds)
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dberror)
	// 		So(fsi.GetByFolderName("moris", ds), ShouldResemble, &Folder{})
	// 	})

	// 	// Convey("Trying to delete from non conform folder should return specific folder error and should not delete folders.", func() {
	// 	// 	folder3.FolderName = "Const"
	// 	// 	Convey("Too long or empty Name should return name error", func() {
	// 	// 		appError := fsi.Delete(&folder3, ds)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Delete.folder.PreSave", "model.folder.foldername.app_error", nil, ""))
	// 	// 		So(fsi.GetAll(ds), ShouldResemble, &folderList1)
	// 	// 		folder3.FolderName = "+alpha"
	// 	// 		appError = fsi.Delete(&folder3, ds)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Delete.folder.PreSave", "model.folder.foldername.app_error", nil, ""))
	// 	// 		So(fsi.GetAll(ds), ShouldResemble, &folderList1)
	// 	// 		folder3.FolderName = "alpha-numerique"
	// 	// 		appError = fsi.Delete(&folder3, ds)standartRole
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("folderStoreImpl.Delete.folder.PreSave", "model.folder.foldername.app_error", nil, ""))
	// 	// 		So(fsi.GetAll(ds), ShouldResemble, &folderList1)
	// 	// 	})
	// 	// })

	// 	db.Delete(&folder0)
	// 	db.Delete(&folder1)
	// 	db.Delete(&folder2)
	// 	db.Delete(&folder3)
	// })
	usi.Delete(&userTest, ds)
	csi.Delete(&channelTest, ds)
	rsi.Delete(&standartRole, ds)
}
