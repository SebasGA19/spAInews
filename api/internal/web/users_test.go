package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
}

func TestFailedRegister(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	firstConfirmCode := engine.Controller.SMTP.ConfirmationCode
	// Second code same information
	response = engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	secondConfirmCode := engine.Controller.SMTP.ConfirmationCode
	// Confirm first code
	response = engine.Post(ConfirmRegistrationURI, map[string]string{
		ConfirmCodeHeader: firstConfirmCode,
	}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Confirm second Code
	response = engine.Post(ConfirmRegistrationURI, map[string]string{
		ConfirmCodeHeader: secondConfirmCode,
	}, nil)
	assert.NotEqual(t, http.StatusOK, response.StatusCode, "Account confirmation should fail")
	// Existing username
	response = engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud2@email.com",
		Password: "password",
	})
	assert.NotEqual(t, http.StatusOK, response.StatusCode, "Register should fail")
	// Existing email
	response = engine.Put(RegisterURI, nil, Register{
		Username: "sulcud2",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	assert.NotEqual(t, http.StatusOK, response.StatusCode, "Register should fail")
}

func TestDeleteSession(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})
	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
	// Delete Session
	response = engine.Delete(SessionURI, map[string]string{SessionHeader: session.Session}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Delete session failed")
	// Delete session twice
	response = engine.Delete(SessionURI, map[string]string{SessionHeader: session.Session}, nil)
	assert.NotEqual(t, http.StatusOK, response.StatusCode, "Delete session should fail")
}

func TestChangePassword(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
	// Change Password
	response = engine.Post(PasswordURI, map[string]string{
		SessionHeader:  session.Session,
		"Content-Type": "application/json",
	}, ChangePassword{
		OldPassword: "password",
		NewPassword: "secret-password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Update password failed")

	// Delete old session
	response = engine.Delete(SessionURI, map[string]string{SessionHeader: session.Session}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Delete session failed")
	// Login old credentials
	header = map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)
	assert.NotEqual(t, http.StatusOK, response.StatusCode, "Login should fail")
	// Login new credentials
	header = map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:secret-password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	decodeError = json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
}

func TestChangePasswordInvalidSession(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
	// Delete old session
	response = engine.Delete(SessionURI, map[string]string{SessionHeader: session.Session}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Delete session failed")
	// Change Password
	response = engine.Post(PasswordURI, map[string]string{
		SessionHeader:  session.Session,
		"Content-Type": "application/json",
	}, ChangePassword{
		OldPassword: "password",
		NewPassword: "secret-password",
	})
	assert.NotEqual(t, http.StatusOK, response.StatusCode, "Update should fail")
	// Login
	header = map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
}

func TestAccountInformation(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
	// Request account information
	response = engine.Get(AccountURI, map[string]string{SessionHeader: session.Session}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Cannot get account information")
	var accountInformation AccountInformationResponse
	decodeError = json.NewDecoder(response.Body).Decode(&accountInformation)
	assert.ErrorIs(t, decodeError, nil)
	assert.Equal(t, "sulcud", accountInformation.Username, "Invalid username")
	assert.Equal(t, "sulcud@email.com", accountInformation.Email, "Invalid email")
}

func TestAccountInformationInvalidSession(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
	// Delete session
	response = engine.Delete(SessionURI, map[string]string{SessionHeader: session.Session}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Delete session failed")
	// Request account information
	response = engine.Get(AccountURI, map[string]string{SessionHeader: session.Session}, nil)
	assert.NotEqual(t, http.StatusOK, response.StatusCode, "Reuse of deleted session should fail")
}

func TestUpdateWords(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
	// Update words
	response = engine.Post(WordsURI, map[string]string{
		SessionHeader:  session.Session,
		"Content-Type": "application/json",
	}, UpdateWords{
		Automatic: true,
		Words:     []string{"Venezuela", "Colombia"},
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Update words failed")
	// Get words
	response = engine.Get(WordsURI, map[string]string{
		SessionHeader: session.Session,
	}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Get words failed")
	var getWords GetWordsResponse
	decodeError = json.NewDecoder(response.Body).Decode(&getWords)
	assert.ErrorIs(t, decodeError, nil)
	if getWords.Automatic != true {
		t.Fatal("Invalid automatic")
	}
	if len(getWords.Words) == 2 && getWords.Words[0] != "Venezuela" && getWords.Words[1] != "Colombia" {
		t.Fatal("Invalid words")
	}
}

func TestResetPassword(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Request Password reset
	response = engine.Post(ResetPasswordURI, map[string]string{"Content-Type": "application/json"}, RequestPasswordReset{
		Email: "sulcud@email.com",
	})

	// Confirm password reset
	response = engine.Post(ConfirmResetPasswordURI, map[string]string{
		ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode,
		"Content-Type":    "application/json",
	}, ResetPassword{
		NewPassword: "my-new-password",
	})

	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:my-new-password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
}

func TestUpdateEmail(t *testing.T) {
	engine := NewTestEngine()
	defer engine.Close()
	// Request register
	response := engine.Put(RegisterURI, nil, Register{
		Username: "sulcud",
		Email:    "sulcud@email.com",
		Password: "password",
	})

	assert.Equal(t, http.StatusOK, response.StatusCode, "Register request failed")
	assert.Greater(t, len(engine.Controller.SMTP.ConfirmationCode), 0, "Not registration code set")
	// Confirm registration
	response = engine.Post(ConfirmRegistrationURI, map[string]string{ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Account confirmation failed")
	// Login
	header := map[string]string{"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("sulcud:password")))}
	response = engine.Get(SessionURI, header, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Login failed")
	var session SessionResponse
	decodeError := json.NewDecoder(response.Body).Decode(&session)
	assert.ErrorIs(t, decodeError, nil)
	// Update email
	response = engine.Post(EmailURI, map[string]string{
		SessionHeader:  session.Session,
		"Content-Type": "application/json",
	}, UpdateEmail{
		Password: "password",
		NewEmail: "sulcud@my.new.email.com",
	})

	// Confirm new email
	response = engine.Post(ConfirmUpdateEmailURI,
		map[string]string{
			ConfirmCodeHeader: engine.Controller.SMTP.ConfirmationCode,
		}, nil)

	// Account information
	response = engine.Get(AccountURI, map[string]string{SessionHeader: session.Session}, nil)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Cannot get account information")
	var accountInformation AccountInformationResponse
	decodeError = json.NewDecoder(response.Body).Decode(&accountInformation)
	assert.ErrorIs(t, decodeError, nil)
	assert.Equal(t, "sulcud", accountInformation.Username, "Invalid username")
	assert.Equal(t, "sulcud@my.new.email.com", accountInformation.Email, "Invalid email")
}
