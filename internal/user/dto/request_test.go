package dto

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
)

func TestCreateUserRequestBindsMultipartForm(t *testing.T) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	fields := map[string]string{
		"name":     "Sarkar",
		"email":    "saurav@example.com",
		"password": "12345678",
	}
	for name, value := range fields {
		if err := writer.WriteField(name, value); err != nil {
			t.Fatalf("WriteField(%q) error = %v", name, err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest("POST", "/api/v1/auth/register", &body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var got CreateUserRequest
	if err := c.Bind(&got); err != nil {
		t.Fatalf("Bind() error = %v", err)
	}

	if got.Name != fields["name"] {
		t.Errorf("Name = %q, want %q", got.Name, fields["name"])
	}
	if got.Email != fields["email"] {
		t.Errorf("Email = %q, want %q", got.Email, fields["email"])
	}
	if got.Password != fields["password"] {
		t.Errorf("Password = %q, want %q", got.Password, fields["password"])
	}
}
