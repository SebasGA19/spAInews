package web

import (
	"github.com/SebasGA19/spAInews/api/internal/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	c := newTestController()
	defer c.Close()
	tests.ClearDB(c.SQL)
	engine := NewEngine(c)
	// Request register
	req := postRequest(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Register request failed")
	}
	if c.Email.ConfirmationCode == "" {
		t.Fatal("No registration code set")
	}
	// Confirm registration
	req = postRequest(ConfirmAccountURI, http.Header{ConfirmAccountCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Account confirmation failed")
	}
	// Login
	req = postRequest(LoginURI, nil, nil)
	req.SetBasicAuth("sulcud", "password")
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Login failed")
	}
}

func TestFailedRegister(t *testing.T) {
	c := newTestController()
	defer c.Close()
	tests.ClearDB(c.SQL)
	engine := NewEngine(c)
	// Request register
	req := postRequest(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Register request failed")
	}
	if c.Email.ConfirmationCode == "" {
		t.Fatal("No registration code set")
	}

	// Second code same information
	req = postRequest(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Register request failed")
	}
	if c.Email.ConfirmationCode == "" {
		t.Fatal("No registration code set")
	}
	secondCode := c.Email.ConfirmationCode
	// Confirm registration
	req = postRequest(ConfirmAccountURI, http.Header{ConfirmAccountCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Account confirmation failed")
	}
	// Confirm second Code
	// Confirm registration
	req = postRequest(ConfirmAccountURI, http.Header{ConfirmAccountCodeHeader: []string{secondCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Registration should fail")
	}
	// Existing username
	req = postRequest(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud2@email.com",
		Password: "password",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Register should fail")
	}
	// Existing email
	req = postRequest(RegisterURI, nil, Register{
		Username: "sulcud2",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Register should fail")
	}
}
