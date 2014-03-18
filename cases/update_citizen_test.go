package cases

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

import "github.com/cgarvis/citizens/entities"

func TestUpdateCitizen(t *testing.T) {
    entities.HashStrength = 1

    Convey("Update Citizen", t, func() {
        Store = NewMemoryStore()
        Store.SaveCitizen(Citizen{UID:"sue@example.com", encryptedSecret:[]byte("my secret")})
        citizen, _ := Store.FetchCitizenByUID("sue@example.com")

        Convey("with new secret", func() {
            attributes := Citizen{
                UID: "sue@example.com",
                encryptedSecret: []byte("super secret"),
            }
            updatedCitizen, err := UpdateCitizen(attributes.UID, attributes)

            So(err, ShouldBeNil)
            So(updatedCitizen.UID, ShouldEqual, citizen.UID)
            So(updatedCitizen.encryptedSecret, ShouldNotEqual, citizen.encryptedSecret)
        })

        Convey("with new uid", func() {
            attributes := Citizen{
                UID: "joe@example.com",
                encryptedSecret: []byte("super secret"),
            }
            updatedCitizen, err := UpdateCitizen("sue@example.com", attributes)

            So(err, ShouldBeNil)
            So(updatedCitizen.UID, ShouldEqual, "joe@example.com")
            So(updatedCitizen.encryptedSecret, ShouldEqual, citizen.encryptedSecret)
        })
    })
}
