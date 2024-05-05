package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
    tests := map[string]struct {
        input string
        want string
    }{
        "simple test": {input: "ApiKey 1234", want: "12345"},
    }

    for name, tc := range tests {
        header := make(http.Header)
        header.Add("Authorization", tc.input)
        got, err := GetAPIKey(header)
        if err != nil {
            t.Errorf("Failed %s: %v", name, err)
            continue
        }
        if got != tc.want {
            t.Errorf("%s| expected: %v, got: %v", name, tc.want, got)
            continue
        }
    }
}

