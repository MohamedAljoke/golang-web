package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"webapp/pkg/data"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})
	for _, e := range tests {
		handlerToTest := app.addIPToContext(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		if e.emptyAddr {
			req.RemoteAddr = ""
		}
		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}
		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContexT(t *testing.T) {

	ctx := context.WithValue(context.Background(), contextUserKey, "ip")

	result := app.ipFromContext(ctx)
	if !strings.EqualFold("ip", result) {
		t.Error("should be ip")
	}
}

func Test_app_auth(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
	var tests = []struct {
		name   string
		isAuth bool
	}{
		{"logged in", true},
		{"not logged in", false},
	}
	for _, e := range tests {
		handlerToTest := app.auth(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)
		if e.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)
		if e.isAuth && rr.Code != http.StatusOK {
			t.Errorf("%s: expected status code of 200 but got %d", e.name, rr.Code)
		}
		if !e.isAuth && rr.Code != http.StatusTemporaryRedirect {
			t.Errorf("%s: expected status 307 but got %d", e.name, rr.Code)
		}
	}
}
