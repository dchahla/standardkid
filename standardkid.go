package standardkid

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Extracts the kid (key ID) from a JWT (JSON Web Token)
func ExtractKidFromJWT(token string) (string, error) {
    // Split the token into its parts [](header, payload, signature)
    parts := strings.Split(token, ".")
    if len(parts) != 3 {
        return "", fmt.Errorf("invalid JWT format: must have three parts")
    }

    // Decode the header JSON
    headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
    if err != nil {
        return "", fmt.Errorf("error decoding header: %v", err)
    }

    // Parse the header JSON
    var header map[string]interface{}
    if err := json.Unmarshal(headerJSON, &header); err != nil {
        return "", fmt.Errorf("error parsing header JSON: %v", err)
    }

    // Extract the kid (key ID) from the header
    kid, ok := header["kid"].(string)
    if !ok {
        return "", fmt.Errorf("kid not found in JWT header")
    }

    return kid, nil
}
