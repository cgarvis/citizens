package entities

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestCompareSecret(t *testing.T) {
    Convey("Compare Secret", t, func() {
        plain := []byte("my secret")
        encrypted := []byte("$2a$10$uc55jB/0kzQuflZbACddp.5ZSHsxn/Jtr1BS6luJ/1kH5trquN33O")

        Convey("when plain matches encrypted", func() {
            result := CompareSecret(encrypted, plain)

            Convey("returns true", func() {
                So(result, ShouldBeTrue)
            })
        })

        Convey("when plain does not matches encrypted", func() {
            result := CompareSecret(encrypted, []byte("not my secret"))

            Convey("returns true", func() {
                So(result, ShouldBeFalse)
            })
        })
    })
}
