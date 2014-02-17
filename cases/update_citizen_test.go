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
        Store.SaveCitizen(Citizen{"sue@example.com", []byte("my secret")})
        citizen, _ := Store.FetchCitizenByUID("sue@example.com")

        Convey("with new secret", func() {
            attributes := Citizen{
                UID: "sue@example.com",
                secret: []byte("super secret"),
            }
            updatedCitizen, err := UpdateCitizen(attributes.UID, attributes)

            So(err, ShouldBeNil)
            So(citizen.UID, ShouldEqual, updatedCitizen.UID)
            So(citizen.secret, ShouldNotEqual, updatedCitizen.secret)
        })

        Convey("with new uid", func() {
        })
    })
}
