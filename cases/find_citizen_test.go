package cases

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestFindCitizen(t *testing.T) {
    Convey("Find Citizen", t, func() {
        Store = NewMemoryStore()

        Convey("when does not exist", func() {
            _, err := FindCitizen("not@exist.com")

            Convey("returns error", func() {
                So(err, ShouldNotBeNil)
            })
        })

        Convey("when exists", func() {
            uid := "find@citizen.com"

            if err := Store.SaveCitizen(Citizen{UID:uid, encryptedSecret:[]byte("Secret")}); err != nil {
                panic("Creating a citizen for find citizens failed")
            }

            citizen, err := FindCitizen(uid)

            Convey("returns correct citizen", func() {
                So(err, ShouldBeNil)
                So(citizen.UID, ShouldEqual, uid)
            })
        })
    })
}

