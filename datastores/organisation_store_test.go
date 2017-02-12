// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has bosic channel to join
package datastores

import (
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/titouanfreville/popcubeapi/models"
	"strconv"
	"strings"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestOrganisationStore(t *testing.T) {
	ds := dbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db
	osi := OrganisationStoreImpl{}
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("organisationStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyexistError := u.NewLocAppError("organisationStoreImpl.Save", "save.transaction.create.already_exist", nil, "Organisation Name: zeus")
		organisation := Organisation{
			IDOrganisation:   0,
			DockerStack:      1,
			OrganisationName: "zeus",
			Description:      "Testing organisation description :O",
			Avatar:           "zeus.svg",
			Domain:           "zeus.popcube",
		}
		Convey("Given a correct organisation.", func() {
			appError := osi.Save(&organisation, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := osi.Save(&organisation, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyexistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		Convey("Given an incorrect organisation.", func() {
			empty := Organisation{}
			organisation.OrganisationName = ""
			Convey("Empty organisation or no Organisation Name organisation should return No name error", func() {
				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				appError = osi.Save(&empty, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.OrganisationName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long organisation name should return Too Long organisation name error", func() {
				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			Convey("Incorect Alpha Num organisation name should be refused ", func() {
				organisation.OrganisationName = "?/+*"
				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "("
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "{"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "}"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = ")"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "["
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "]"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = " "
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.OrganisationName = "electra"

			organisation.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {

				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Save.organisation.PreSave", "model.organisation.is_valid.description.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.Description = "Stoppppppp"
		})
		db.Delete(&organisation)
	})
	Convey("Testing update function", t, func() {
		organisation := Organisation{
			IDOrganisation:   0,
			DockerStack:      1,
			OrganisationName: "zeus",
			Description:      "Testing organisation description :O",
			Avatar:           "zeus.svg",
			Domain:           "zeus.popcube",
		}
		newOrganisation := Organisation{
			DockerStack:      4,
			OrganisationName: "NewZeus",
		}
		appError := osi.Save(&organisation, ds)
		dbError := u.NewLocAppError("organisationStoreImpl.Update", "update.transaction.updates.encounterError", nil, "")
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		Convey("Providing a correct user to update", func() {
			appError = osi.Update(&organisation, &newOrganisation, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
		})
		Convey("Providing an incorrect user as new should result in errors", func() {
			empty := Organisation{}
			newOrganisation.OrganisationName = ""
			Convey("Empty organisation or no Organisation Name organisation should return No name error", func() {
				appError := osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				appError = osi.Update(&organisation, &empty, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(empty.IDOrganisation, 10)))
			})
			newOrganisation.OrganisationName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long organisation name should return Too Long organisation name error", func() {
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
			})
			Convey("Incorect Alpha Num organisation name should be refused ", func() {
				newOrganisation.OrganisationName = "?/+*"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				newOrganisation.OrganisationName = "("
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				newOrganisation.OrganisationName = "{"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				newOrganisation.OrganisationName = "}"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				newOrganisation.OrganisationName = ")"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				newOrganisation.OrganisationName = "["
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				newOrganisation.OrganisationName = "]"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
				newOrganisation.OrganisationName = " "
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
			})
			newOrganisation.OrganisationName = "electra"

			newOrganisation.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationNew.PreSave", "model.organisation.is_valid.description.app_error",
					nil, "id="+strconv.FormatUint(newOrganisation.IDOrganisation, 10)))
			})
			newOrganisation.Description = "Stoppppppp"

		})
		Convey("Providing an incorrect user as old should result in errors", func() {
			empty := Organisation{}
			organisation.OrganisationName = ""
			Convey("Empty organisation or no Organisation Name organisation should return No name error", func() {
				appError := osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				appError = osi.Update(&empty, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(empty.IDOrganisation, 10)))
			})
			organisation.OrganisationName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long organisation name should return Too Long organisation name error", func() {
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			Convey("Incorect Alpha Num organisation name should be refused ", func() {
				organisation.OrganisationName = "?/+*"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "("
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "{"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "}"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = ")"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "["
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "]"
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = " "
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.OrganisationName = "electra"

			organisation.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {
				appError = osi.Update(&organisation, &newOrganisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldResemble, u.NewLocAppError("organisationStoreImpl.Update.organisationOld.PreSave", "model.organisation.is_valid.description.app_error",
					nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.Description = "Stoppppppp"
		})
		db.Delete(&organisation)
		db.Delete(&newOrganisation)
	})
	Convey("Testing Get function", t, func() {
		organisation := Organisation{
			IDOrganisation:   0,
			DockerStack:      1,
			OrganisationName: "zeus",
			Description:      "Testing organisation description :O",
			Avatar:           "zeus.svg",
			Domain:           "zeus.popcube",
		}
		Convey("Trying to get organisation from empty DB should return empty", func() {
			So(&Organisation{}, ShouldResemble, osi.Get(ds))
		})
		appError := osi.Save(&organisation, ds)
		So(appError, ShouldBeNil)
		Convey("Trying to get organisation from non empty DB should return a correct organisation object", func() {
			got := osi.Get(ds)
			So(&organisation, ShouldResemble, got)
			So(got.IsValid(), ShouldBeNil)
		})
		db.Delete(&organisation)
	})
}
