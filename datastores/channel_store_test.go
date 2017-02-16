// This file is used to test if user model is working correctly.
// A user is always linked to an channel
// He has bcsic channel to join
package datastores

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestChannelStore(t *testing.T) {
	store := NewStore()
	db := store.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")
	csi := store.Channel()
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("channelStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("channelStoreImpl.Save", "save.transaction.create.already_exist", nil, "Channel Name: electras")
		channel := Channel{
			ChannelName: "electras",
			Type:        "audio",
			Private:     false,
			Description: "Testing channel description :O",
			Subject:     "Sujet",
			Avatar:      "jesuiscool.svg",
		}
		Convey("Given a correct channel.", func() {
			appError := csi.Save(&channel, db)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := csi.Save(&channel, db)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&channel)
	})

	Convey("Testing update function", t, func() {
		dbError := u.NewLocAppError("channelStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("channelStoreImpl.Save", "save.transaction.create.already_exist", nil, "Channel Name: electras")
		// empty := Channel{}
		channel := Channel{
			ChannelName: "electras",
			Type:        "audio",
			Private:     false,
			Description: "Testing channel description :O",
			Subject:     "Sujet",
			Avatar:      "jesuiscool.svg",
		}
		channelNew := Channel{
			ChannelName: "elektra",
			Type:        "video",
			Private:     true,
			Description: "Testing channel description :O private firs ;o",
			Subject:     "New Sujet",
			Avatar:      "jesuistjscool.svg",
		}
		appError := csi.Save(&channel, db)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		So(appError, ShouldNotResemble, alreadyExistError)

		Convey("Provided correct Channel to modify should not return errors", func() {
			appError := csi.Update(&channel, &channelNew, db)
			channelShouldResemble := channelNew
			channelShouldResemble.WebID = channel.WebID
			channelShouldResemble.IDChannel = channel.IDChannel
			channelShouldResemble.LastUpdate = channel.LastUpdate
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
			So(channel, ShouldResemble, channelShouldResemble)
		})

		Convey("Provided wrong old Channel to modify should result in old_channel error", func() {
			channel.ChannelName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long channel name should return Too Long channel name error", func() {
				appError := csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebID))
			})
			Convey("Incorect Alpha Num channel name should be refused (no CAPS)", func() {
				channel.ChannelName = "JeSuisCaps"
				appError := csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = "?/+*"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = "("
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = "{"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = "}"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = ")"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = "["
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = "]"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
				channel.ChannelName = " "
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID))
			})
			channel.ChannelName = "electra"
			channel.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebID))
			})
			channel.Description = "Stoppppppp"
			channel.Subject = "Encore beaucoup de caractere pour rien .... mais un peu moins cette fois. Il n'en faut que 250 ........... Fait dodo, cola mon p'tit frere. Fais dodo, j'ai pêté un cable. Swing du null, Swing du null, c'est le swing du null ..... :guitare: :singer: :music: Je suis un main troll :O"
			Convey("Given a too long subject, should return too long description error :p", func() {
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebID))
			})
			channel.Subject = "Safe"
			channel.Type = "Outside of Range"
			Convey("Providing a wrong type should not work", func() {
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelOld.PreSave", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebID))
			})
		})

		Convey("Provided wrong new Channel to modify should result in new_channel error", func() {
			channelNew.ChannelName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long channel name should return Too Long channel name error", func() {
				appError := csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.channel_name.app_error", nil, "id="+channelNew.WebID))
			})
			Convey("Incorect Alpha Num channel name should be refused", func() {
				channelNew.ChannelName = "?/+*"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
				channelNew.ChannelName = "("
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
				channelNew.ChannelName = "{"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
				channelNew.ChannelName = "}"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
				channelNew.ChannelName = ")"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
				channelNew.ChannelName = "["
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
				channelNew.ChannelName = "]"
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
				channelNew.ChannelName = " "
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channelNew.WebID))
			})
			channelNew.ChannelName = "electra"
			channelNew.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.description.app_error", nil, "id="+channelNew.WebID))
			})
			channelNew.Description = "Stoppppppp"
			channelNew.Subject = "Encore beaucoup de caractere pour rien .... mais un peu moins cette fois. Il n'en faut que 250 ........... Fait dodo, cola mon p'tit frere. Fais dodo, j'ai pêté un cable. Swing du null, Swing du null, c'est le swing du null ..... :guitare: :singer: :music: Je suis un main troll :O"
			Convey("Given a too long subject, should return too long description error :p", func() {
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.subject.app_error", nil, "id="+channelNew.WebID))
			})
			channelNew.Subject = "Safe"
			channelNew.Type = "Outside of Range"
			Convey("Providing a wrong type should not work", func() {
				appError = csi.Update(&channel, &channelNew, db)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
				So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Update.channelNew.PreSave", "model.channel.is_valid.type.app_error", nil, "id="+channelNew.WebID))
			})
		})

		db.Delete(&channel)
		db.Delete(&channelNew)
	})

	Convey("Testing Getters", t, func() {
		channel0 := Channel{
			ChannelName: "electra",
			Type:        "video",
			Private:     false,
			Description: "Testing channel description :O",
			Subject:     "Sujet",
			Avatar:      "jesuiscool.svg",
		}
		channel1 := Channel{
			ChannelName: "mal",
			Type:        "audio",
			Private:     false,
			Description: "Speaking on Malsdjisqnju BD song from R. Sechan",
			Subject:     "Sujet1",
			Avatar:      "cover_mal.svg",
		}
		channel2 := Channel{
			ChannelName: "lagaffesfantasio",
			Type:        "direct",
			Private:     false,
			Avatar:      "gaston.svg",
		}
		channel1New := Channel{
			ChannelName: "malheur",
			Private:     true,
			Description: "Let's speak about the BD Mal",
			Subject:     "Mal",
			Avatar:      "cover_mal_efique.svg",
		}
		channel3 := Channel{
			ChannelName: "corsicarms",
			Type:        "audio",
			Private:     false,
			Description: "Speaking on Corsic Arms song from R. Sechan",
			Subject:     "Sujet",
			Avatar:      "cover_csa.svg",
		}

		csi.Save(&channel0, db)
		csi.Save(&channel1, db)
		// csi.Update(&channel1, &channel1New, db)
		csi.Save(&channel2, db)
		csi.Save(&channel3, db)

		// Have to be after save so ID are up to date :O
		channelList := []Channel{
			channel0,
			channel1,
			channel2,
			channel3,
		}

		audioList := []Channel{channel1, channel3}
		directList := []Channel{channel2}
		privateList := []Channel{channel2}
		publicList := []Channel{channel0, channel1, channel3}
		emptyList := []Channel{}

		Convey("We have to be able to find all channels in the db", func() {
			channels := csi.GetAll(db)
			So(channels, ShouldNotResemble, &emptyList)
			So(channels, ShouldResemble, &channelList)
		})

		Convey("We have to be able to find a channel from is name", func() {
			channel := csi.GetByName(channel0.ChannelName, db)
			So(channel, ShouldNotResemble, &Channel{})
			So(channel, ShouldResemble, &channel0)
			channel = csi.GetByName(channel2.ChannelName, db)
			So(channel, ShouldNotResemble, &Channel{})
			So(channel, ShouldResemble, &channel2)
			channel = csi.GetByName(channel3.ChannelName, db)
			So(channel, ShouldNotResemble, &Channel{})
			So(channel, ShouldResemble, &channel3)
			Convey("Should also work from updated value", func() {
				channel = csi.GetByName(channel1.ChannelName, db)
				So(channel, ShouldNotResemble, &Channel{})
				So(channel, ShouldResemble, &channel1)
			})
		})

		Convey("We have to be able to find channels from type", func() {
			channels := csi.GetByType("audio", db)
			So(channels, ShouldNotResemble, &Channel{})
			So(channels, ShouldResemble, &audioList)
			channels = csi.GetByType("direct", db)
			So(channels, ShouldNotResemble, &Channel{})
			So(channels, ShouldResemble, &directList)
		})

		Convey("We have to be able to find private or public channels list", func() {
			channels := csi.GetPrivate(db)
			So(channels, ShouldNotResemble, &Channel{})
			So(channels, ShouldResemble, &privateList)
			channels = csi.GetPublic(db)
			So(channels, ShouldNotResemble, &Channel{})
			So(channels, ShouldResemble, &publicList)
		})

		Convey("Searching for non existent channel should return empty", func() {
			channel := csi.GetByName("fantome", db)
			So(channel, ShouldResemble, &Channel{})
		})

		db.Delete(&channel0)
		db.Delete(&channel1)
		db.Delete(&channel1New)
		db.Delete(&channel2)
		db.Delete(&channel3)

		Convey("Searching all in empty table should return empty", func() {
			channels := csi.GetAll(db)
			So(channels, ShouldResemble, &[]Channel{})
		})
	})

	Convey("Testing delete channel", t, func() {
		dberror := u.NewLocAppError("channelStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
		channel0 := Channel{
			ChannelName: "electra",
			Type:        "video",
			Private:     false,
			Description: "Testing channel description :O",
			Subject:     "Sujet",
			Avatar:      "jesuiscool.svg",
		}
		channel1 := Channel{
			ChannelName: "mal",
			Type:        "audio",
			Private:     false,
			Description: "Speaking on Malsdjisqnju BD song from R. Sechan",
			Subject:     "Sujet1",
			Avatar:      "cover_mal.svg",
		}
		channel2 := Channel{
			ChannelName: "lagaffesfantasio",
			Type:        "direct",
			Private:     false,
			Avatar:      "gaston.svg",
		}
		channel3 := Channel{
			ChannelName: "corsicarms",
			Type:        "audio",
			Private:     false,
			Description: "Speaking on Corsic Arms song from R. Sechan",
			Subject:     "Sujet",
			Avatar:      "cover_csa.svg",
		}

		csi.Save(&channel0, db)
		csi.Save(&channel1, db)
		csi.Save(&channel2, db)
		csi.Save(&channel3, db)

		// Have to be after save so ID are up to date :O
		// channel3Old := channel3
		// channelList1 := []Channel{
		// 	channel0,
		// 	channel1,
		// 	channel2,
		// 	channel3Old,
		// }

		Convey("Deleting a known channel should work", func() {
			appError := csi.Delete(&channel2, db)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dberror)
			So(csi.GetByName("God", db), ShouldResemble, &Channel{})
		})

		// Convey("Trying to delete from non conform channel should return specific channel error and should not delete channels.", func() {
		// 	channel3.ChannelName = "Const"
		// 	Convey("Too long or empty Name should return name error", func() {
		// 		appError := csi.Delete(&channel3, db)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dberror)
		// 		So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Delete.channel.PreSave", "model.channel.channelname.app_error", nil, ""))
		// 		So(csi.GetAll(db), ShouldResemble, &channelList1)
		// 		channel3.ChannelName = "+alpha"
		// 		appError = csi.Delete(&channel3, db)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dberror)
		// 		So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Delete.channel.PreSave", "model.channel.channelname.app_error", nil, ""))
		// 		So(csi.GetAll(db), ShouldResemble, &channelList1)
		// 		channel3.ChannelName = "alpha-numerique"
		// 		appError = csi.Delete(&channel3, db)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dberror)
		// 		So(appError, ShouldResemble, u.NewLocAppError("channelStoreImpl.Delete.channel.PreSave", "model.channel.channelname.app_error", nil, ""))
		// 		So(csi.GetAll(db), ShouldResemble, &channelList1)
		// 	})
		// })

		db.Delete(&channel0)
		db.Delete(&channel1)
		db.Delete(&channel2)
		db.Delete(&channel3)
	})
}
