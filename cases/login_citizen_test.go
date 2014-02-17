package cases

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestLoginCitizen(t *testing.T) {
    Store = NewMemoryStore()

    Convey("Login Citizen", t, func() {
        uid := "1234567890"
        secret := "my secret"
        encrypted := []byte("$2a$10$mA7yxta9XDomJUBKhNpVkuNeImkeOrl0h9TfGZBbohJqOa7qqCFX6")

        Store.SaveCitizen(Citizen{uid, encrypted})

        Convey("with valid uid and secret", func() {
            jwt, err := LoginCitizen(uid, secret)
            So(err, ShouldBeNil)
            So(jwt, ShouldNotEqual, "")
        })

        Convey("with invalid secret", func() {
            jwt, err := LoginCitizen(uid, "not my secret")
            So(err, ShouldNotBeNil)
            So(jwt, ShouldEqual, "")
        })

        Convey("when citizen doesn't exist", func() {
            jwt, err := LoginCitizen("123", secret)
            So(err, ShouldNotBeNil)
            So(jwt, ShouldEqual, "")
        })
    })
}
