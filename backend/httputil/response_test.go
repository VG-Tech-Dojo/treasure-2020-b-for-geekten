package httputil

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestReturnErrorWhen5xx(t *testing.T) {

	w := httptest.NewRecorder()
	RespondErrorJson(w, 500, errors.New("query failed. select id from user where password="))

	resp := w.Result()
	if a, e := resp.StatusCode, 500; a != e {
		t.Fatalf("Expect %v but got %v", e, a)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if a, e := string(body), `{"message":"query failed. select id from user where password="}`; a != e {
		t.Fatalf("Expect %v but got %v", e, a)
	}
}

func TestReturnErrorWhen4xx(t *testing.T) {

	w := httptest.NewRecorder()
	RespondErrorJson(w, 400, errors.New("client side error"))

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if a, e := string(body), `{"message":"client side error"}`; a != e {
		t.Fatalf("Expect %v but got %v", e, a)
	}
}

func TestRespondJSONReturn500WhenNil(t *testing.T) {

	w := httptest.NewRecorder()
	RespondJSON(w, 200, nil)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if a, e := resp.StatusCode, 200; a != e {
		t.Fatalf("Expect %v but got %v", e, a)
	}

	if a, e := string(body), "{}"; a != e {
		t.Fatalf("Expect %v but got %v", e, a)
	}
}
