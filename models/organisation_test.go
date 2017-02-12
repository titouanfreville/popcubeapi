package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"strings"
	"testing"
	u "github.com/titouanfreville/popcubeapi/utils"
)

func TestOrganisationModel(t *testing.T) {
	Convey("Testing json vs organisation conversions", t, func() {
		Convey("Given a organisation", func() {
			organisation := Organisation{OrganisationName: NewID()}
			Convey("Converting organisation to json then json to organisation should provide same organisation information", func() {
				json := organisation.ToJSON()
				testOrganisation := OganisationFromJSON(strings.NewReader(json))
				So(organisation.OrganisationName, ShouldEqual, testOrganisation.OrganisationName)
			})
		})
	})

	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct organisation. Organisation should be validate", func() {
			organisation := Organisation{
				IDOrganisation:   0,
				DockerStack:      1,
				OrganisationName: "electra",
				Description:      "Testing organisation description :O",
				Avatar:           "jesuiscool.svg",
				Domain:           "electra.popcube",
			}
			So(organisation.IsValid(), ShouldBeNil)
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
		})
		Convey("Given an incorrect organisation. Organisation should be refused", func() {
			empty := Organisation{}
			organisation := Organisation{
				IDOrganisation: 0,
				DockerStack:    1,
				Description:    "Testing organisation description :O",
				Avatar:         "jesuiscool.svg",
				Domain:         "electra.popcube",
			}
			Convey("Empty organisation or no Organisation Name organisation should return No name error", func() {
				So(empty.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.OrganisationName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long organisation name should return Too Long organisation name error", func() {
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			Convey("Incorect Alpha Num organisation name should be refused (no CAPS)", func() {
				organisation.OrganisationName = "JeSuisCaps"
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "?/+*"
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "("
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "{"
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "}"
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = ")"
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "["
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = "]"
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
				organisation.OrganisationName = " "
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.OrganisationName = "electra"

			organisation.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {
				So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			})
			organisation.Description = "Stoppppppp"
		})
	})

	Convey("Testing PreSave function", t, func() {
		organisation := Organisation{}
		Convey("If organisation is empty, should fill some fields - webID, OrganisationName, LastUpdate, Avatar and type, and organisation should not be valid", func() {
			organisation.PreSave()
			So(organisation.IsValid(), ShouldResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.Avatar, ShouldEqual, "default_organisation_avatar.svg")
		})
		Convey("If provided OrganisationName contain caps, they should be lowered", func() {
			organisation.OrganisationName = "JeSuisCaps"
			organisation.PreSave()
			So(organisation.IsValid(), ShouldBeNil)
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.OrganisationName, ShouldEqual, "jesuiscaps")
			organisation.OrganisationName = "nocapsshouldnotbemodified"
			organisation.PreSave()
			So(organisation.IsValid(), ShouldBeNil)
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.IsValid(), ShouldNotResemble, u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10)))
			So(organisation.OrganisationName, ShouldEqual, "nocapsshouldnotbemodified")
		})
	})
}
