// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has basic channel to join
package datastores

import (
	"testing"

	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
)

// Test correction test for user ;)

// Test Password functionalities from user Model
func TestDbStores(t *testing.T) {
	Convey("Testing DB initialisation and closure", t, func() {
		Convey("Given a datastore", func() {
			ds := Store()
			Convey("Initialising should provide a db", func() {
				resDb := ds.InitConnection("root", "popcube_test", "popcube_dev", "database", "3306")
				db, _ := gorm.Open("mysql", "root:popcube_dev@/?charset=utf8&parseTime=True&loc=Local")
				// Should nor resemble cause of Channel IDs. 2 Objects have 2 IDs :'('
				So(resDb, ShouldNotResemble, db)
				So(resDb.Value, ShouldEqual, db.Value)
				So(resDb.Error, ShouldEqual, db.Error)
				So(resDb.RowsAffected, ShouldEqual, db.RowsAffected)

				ds.InitConnection("test_user", "popcube_test", "test", "database", "3306")
				db, _ = gorm.Open("mysql", "test_user:test@/?popcube_test?charset=utf8&parseTime=True&loc=Local")
				// Should nor resemble cause of Channel IDs. 2 Objects have 2 IDs :'('
				So(resDb, ShouldNotResemble, db)
				So(resDb.Value, ShouldEqual, db.Value)
				So(resDb.Error, ShouldEqual, db.Error)
				So(resDb.RowsAffected, ShouldEqual, db.RowsAffected)
			})
			// Convey("Stoping the connection should destroy the Db stored.", func() {
			// 	resDb := ds.InitConnection("test_user", "popcube_test", "test", "database", "3306")
			// 	ds.CloseConnection(resDb)
			// 	So(resDb, ShouldResemble, &gorm.DB{})
			// })
		})
	})
}
