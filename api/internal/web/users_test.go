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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Account confirmation failed")
	}
	// Confirm second Code
	// Confirm registration
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{secondCode}}, nil)
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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

func TestChangePassword(t *testing.T) {
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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
	// Change Password
	req = postRequest(PasswordURI, http.Header{
		SessionHeader:  []string{session.Session},
		"Content-Type": []string{"application/json"},
	}, ChangePassword{
		OldPassword: "password",
		NewPassword: "secret-password",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Update password failed")
	}
	// Delete old session
	req = deleteRequest(SessionURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Delete session failed")
	}
	// Login old credentials
	req = getRequest(SessionURI, nil, nil)
	req.SetBasicAuth("sulcud", "password")
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Login should fail")
	}
	// Login new credentials
	req = getRequest(SessionURI, nil, nil)
	req.SetBasicAuth("sulcud", "secret-password")
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Login failed")
	}
	decodeError = json.NewDecoder(response.Body).Decode(&session)
	tests.FailOnError(decodeError)
}

func TestChangePasswordInvalidSession(t *testing.T) {
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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
	// Delete old session
	req = deleteRequest(SessionURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Delete session failed")
	}
	// Change Password
	req = postRequest(PasswordURI, http.Header{
		SessionHeader:  []string{session.Session},
		"Content-Type": []string{"application/json"},
	}, ChangePassword{
		OldPassword: "password",
		NewPassword: "secret-password",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Update should fail")
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

func TestAccountInformation(t *testing.T) {
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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
	// Request account information
	req = getRequest(AccountURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Cannot get account information")
	}
	var accountInformation AccountInformationResponse
	decodeError = json.NewDecoder(response.Body).Decode(&accountInformation)
	tests.FailOnError(decodeError)
	if accountInformation.Username != "sulcud" {
		t.Fatal("Invalid username")
	}
	if accountInformation.Email != "sulcud@email.com" {
		t.Fatal("Invalid email")
	}
}

func TestAccountInformationInvalidSession(t *testing.T) {
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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
	// Delete session
	req = deleteRequest(SessionURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Delete session failed")
	}
	// Request account information
	req = getRequest(AccountURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode == http.StatusOK {
		t.Fatal("Should fail")
	}
}

func TestUpdateWords(t *testing.T) {
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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
	// Update words
	req = postRequest(WordsURI, http.Header{
		SessionHeader:  []string{session.Session},
		"Content-Type": []string{"application/json"},
	}, UpdateWords{
		Automatic: true,
		Words:     []string{"Venezuela", "Colombia"},
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Update words failed")
	}
	// Get words
	req = getRequest(WordsURI, http.Header{
		SessionHeader: []string{session.Session},
	}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Get words failed")
	}
	var getWords GetWordsResponse
	decodeError = json.NewDecoder(response.Body).Decode(&getWords)
	tests.FailOnError(decodeError)
	if getWords.Automatic != true {
		t.Fatal("Invalid automatic")
	}
	if len(getWords.Words) == 2 && getWords.Words[0] != "Venezuela" && getWords.Words[1] != "Colombia" {
		t.Fatal("Invalid words")
	}
}

func TestResetPassword(t *testing.T) {
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Account confirmation failed")
	}
	// Request Password reset
	req = postRequest(ResetPasswordURI, http.Header{"Content-Type": []string{"application/json"}}, RequestPasswordReset{
		Email: "sulcud@email.com",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	// Confirm password reset
	req = postRequest(ConfirmResetPasswordURI, http.Header{
		ConfirmCodeHeader: []string{c.Email.ConfirmationCode},
		"Content-Type":    []string{"application/json"},
	}, ResetPassword{
		NewPassword: "my-new-password",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	// Login
	req = getRequest(SessionURI, nil, nil)
	req.SetBasicAuth("sulcud", "my-new-password")
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Login failed")
	}
}

func TestUpdateEmail(t *testing.T) {
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
	req = postRequest(ConfirmRegistrationURI, http.Header{ConfirmCodeHeader: []string{c.Email.ConfirmationCode}}, nil)
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
	// Update email
	req = postRequest(EmailURI, http.Header{
		SessionHeader:  []string{session.Session},
		"Content-Type": []string{"application/json"},
	}, UpdateEmail{
		Password: "password",
		NewEmail: "sulcud@my.new.email.com",
	})
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	// Confirm new email
	req = postRequest(ConfirmUpdateEmailURI,
		http.Header{
			ConfirmCodeHeader: []string{c.Email.ConfirmationCode},
		}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	// Account information
	req = getRequest(AccountURI, http.Header{SessionHeader: []string{session.Session}}, nil)
	w = httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	response = w.Result()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Cannot get account information")
	}
	var accountInformation AccountInformationResponse
	decodeError = json.NewDecoder(response.Body).Decode(&accountInformation)
	tests.FailOnError(decodeError)
	if accountInformation.Username != "sulcud" {
		t.Fatal("Invalid username")
	}
	if accountInformation.Email != "sulcud@my.new.email.com" {
		t.Fatal("Invalid email")
	}
}
