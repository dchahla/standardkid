package standardkid

import (
	"testing"
)

func TestExtractKidFromJWT(t *testing.T) {
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImtleV9pZCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    expectedKid := "key_id"

    kid, err := ExtractKidFromJWT(token)
    if err != nil {
        t.Errorf("ExtractKidFromJWT returned an error: %v", err)
    }

    if kid != expectedKid {
        t.Errorf("Expected kid %s, got %s", expectedKid, kid)
    }
}
