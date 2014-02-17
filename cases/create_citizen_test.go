package cases

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

import "github.com/cgarvis/citizens/entities"

func TestCreateCitizen(t *testing.T) {
    entities.HashStrength = 1

    Convey("Create Citizen", t, func() {
        Store = NewMemoryStore()

        Convey("with valid uid and secret", func() {
            uid := "valid@citizen.com"
            secret := "my secret"

            citizen, err := CreateCitizen(uid, secret)
            So(err, ShouldBeNil)

            Convey("returns corrent citizen", func() {
                So(citizen.UID, ShouldEqual, uid)
                So(string(citizen.secret), ShouldNotEqual, "")
                So(string(citizen.secret), ShouldNotEqual, secret)
            })
        })

        Convey("without uid", func() {
            _, err := CreateCitizen("", "some secret")
            Convey("returns an error", func() {
                So(err, ShouldNotBeNil)
            })
        })

        Convey("without secret", func() {
            _, err := CreateCitizen("withoutsecret@citizen.com", "")
            Convey("returns an error", func() {
                So(err, ShouldNotBeNil)
            })
        })

        Convey("when uid is already taken", func() {
            CreateCitizen("smith@example.com", "secret")
            _, err := CreateCitizen("smith@example.com", "secret")
            Convey("returns an error", func() {
                So(err, ShouldNotBeNil)
            })
        })
    })
}
