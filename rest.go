package main

import (
	"fmt"
	"net/http"
)

import (
	"github.com/cgarvis/citizens/cases"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func RunServer() {
	handler := createHttpHandler()
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	http.ListenAndServe(addr, handler)
}

func createHttpHandler() http.Handler {
	m := martini.Classic()
	m.Use(render.Renderer())

	// Make sure server is still responsive
	m.Get("/ping", func() (int, string) {
		return http.StatusOK, "pong"
	})

	m.Get("/v1/citizens", func(params martini.Params, r render.Render) {
		citizens := cases.FindCitizens()
		response := make([]Citizen, len(citizens), len(citizens))
		for i, citizen := range citizens {
		    response[i].from(citizen)
		}
		r.JSON(http.StatusOK, response)
	})
	m.Post("/v1/citizens", binding.Json(Citizen{}), func(citizen Citizen, r render.Render) {
		if citizen, err := cases.CreateCitizen(citizen.UID, citizen.Secret); err != nil {
            r.JSON(http.StatusBadRequest, err)
		} else {
			response := Citizen{}
			response.from(citizen)
            r.JSON(http.StatusOK, response)
		}
	})

	m.Patch("/v1/citizens/:uid", binding.Json(Citizen{}), func(params martini.Params, attributes Citizen, r render.Render) {
	    if citizen, err := cases.UpdateCitizen(params["uid"], attributes.to()); err != nil {
	        r.JSON(http.StatusBadRequest, err)
	    } else {
	        response := Citizen{}
	        response.from(citizen)
	        r.JSON(http.StatusOK, response)
	    }
	})

	m.Get("/v1/citizens/:uid", func(params martini.Params, r render.Render) {
	    if citizen, err := cases.FindCitizen(params["uid"]); err != nil {
	        r.JSON(http.StatusBadRequest, err)
	    } else {
	        response := Citizen{}
	        response.from(citizen)
	        r.JSON(http.StatusOK, response)
	    }
	})

	m.Put("/v1/identify", func(params martini.Params) (int, string) {
		if token, err := cases.LoginCitizen(params["uid"], params["secret"]); err != nil {
			return http.StatusInternalServerError, ""
		} else {
			return http.StatusOK, token
		}
	})

	return m
}

type Citizen struct {
	UID string `json:"uid"`
	Secret string `json:",omitempty"`
}

func (c *Citizen) from(citizen cases.Citizen) {
	c.UID = citizen.UID
}

func (c *Citizen) to() cases.Citizen {
    return cases.Citizen{UID: c.UID}
}
