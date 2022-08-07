package web

import (
	"encoding/json"
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
	req := putRequest(RegisterURI, nil, Register{
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmAccountCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Account confirmation failed")
	}
	// Login
	req = getRequest(SessionURI, nil, nil)
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
	req := putRequest(RegisterURI, nil, Register{
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
	req = putRequest(RegisterURI, nil, Register{
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmAccountCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Account confirmation failed")
	}
	// Confirm second Code
	// Confirm registration
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmAccountCodeHeader: []string{secondCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Registration should fail")
	}
	// Existing username
	req = putRequest(RegisterURI, nil, Register{
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
	req = putRequest(RegisterURI, nil, Register{
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

func TestDeleteSession(t *testing.T) {
	c := newTestController()
	defer c.Close()
	tests.ClearDB(c.SQL)
	engine := NewEngine(c)
	// Request register
	req := putRequest(RegisterURI, nil, Register{
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmAccountCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Account confirmation failed")
	}
	// Login
	req = getRequest(SessionURI, nil, nil)
	req.SetBasicAuth("sulcud", "password")
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Login failed")
	}
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	tests.FailOnError(decodeError)
	// Delete Session
	req = deleteRequest(SessionURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Delete session failed")
	}
	// Delete session twice
	req = deleteRequest(SessionURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Delete should fail")
	}
}
