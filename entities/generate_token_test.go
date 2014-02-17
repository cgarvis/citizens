package entities

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestGenerateToken(t *testing.T) {
    Convey("Generate Token", t, func() {
        token := GenerateToken([]byte("API KEY"), "1234567890")

        So(token, ShouldEqual, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiIxMjM0NTY3ODkwIn0.jbxNX9WA4kdVHsvU3ob26biL_8vICytZHfTZ4DlFpec")
    })
}
