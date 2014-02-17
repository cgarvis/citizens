package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

import "github.com/cgarvis/citizens/cases"

func TestREST(t *testing.T) {
    cases.Store = cases.NewMemoryStore()
	server := httptest.NewServer(createHttpHandler())
	defer server.Close()

	get := func(path string) (*http.Response, error) {
		resp, err := http.DefaultClient.Get(server.URL + path)
		return resp, err
	}

	patch := func(path, json string) (*http.Response, error) {
		b := strings.NewReader(json)
		req, _ := http.NewRequest("PATCH", server.URL+path, b)
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		return resp, err
	}

	post := func(path, json string) (*http.Response, error) {
		b := strings.NewReader(json)
		resp, err := http.DefaultClient.Post(server.URL+path, "application/json", b)
		return resp, err
	}

	Convey("REST API", t, func() {
		Convey("when method doesn't exist", func() {
			resp, _ := get("/not_method")

			Convey("status code is 404", func() {
				So(resp.StatusCode, ShouldEqual, 404)
			})
		})

		Convey("when creating a citizen", func() {
            resp, _ := post("/v1/citizens", `{"uid": "joe@example.com", "secret": "kittens"}`)

			Convey("returns new citizen", func() {
				So(resp.Header.Get("Content-Type"), ShouldEqual, "application/json; charset=UTF-8")
				So(resp.StatusCode, ShouldEqual, 200)
				if bs, err := ioutil.ReadAll(resp.Body); err == nil {
                    So(string(bs), ShouldEqual, `{"uid":"joe@example.com"}`)
				}
			})
		})

		Convey("when listing citizens", func() {
            resp, _ := get("/v1/citizens")

		    Convey("returns a json list", func() {
                So(resp.Header.Get("Content-Type"), ShouldEqual, "application/json; charset=UTF-8")
                So(resp.StatusCode, ShouldEqual, 200)
				if bs, err := ioutil.ReadAll(resp.Body); err == nil {
                    So(string(bs), ShouldEqual, `[{"uid":"joe@example.com"}]`)
				}
		    })
		})

		Convey("when looking up a citizen", func() {
            resp, _ := get("/v1/citizens/joe@example.com")

		    Convey("returns citizen", func() {
				So(resp.Header.Get("Content-Type"), ShouldEqual, "application/json; charset=UTF-8")
				So(resp.StatusCode, ShouldEqual, 200)

				if bs, err := ioutil.ReadAll(resp.Body); err == nil {
                    So(string(bs), ShouldEqual, `{"uid":"joe@example.com"}`)
				}
		    })
		})

		Convey("when updating citizen", func() {
		    resp, _ := patch("/v1/citizens/joe@example.com", `{"secret": "puppies"}`)

		    Convey("returns updated citizen", func() {
				So(resp.Header.Get("Content-Type"), ShouldEqual, "application/json; charset=UTF-8")
				So(resp.StatusCode, ShouldEqual, 200)
				if bs, err := ioutil.ReadAll(resp.Body); err == nil {
                    So(string(bs), ShouldEqual, `{"uid":"joe@example.com"}`)
				}
		    })
		})
	})
}
