package cases

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestFindCitizens(t *testing.T) {
    Store = NewMemoryStore()

    Convey("Find Citizens", t, func() {
        Convey("when exists", func() {
            if err := Store.SaveCitizen(Citizen{UID:"1234567890", encryptedSecret:[]byte("Secret")}); err != nil {
                panic("Creating a citizen for find citizens failed")
            }

            citizens := FindCitizens()

            Convey("returns correct citizen", func() {
                So(len(citizens), ShouldNotEqual, 0)

                citizen := citizens[0]

                So(citizen.UID, ShouldEqual, "1234567890")
            })
        })
    })
}
