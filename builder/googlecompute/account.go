package googlecompute

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

type ServiceAccount struct {
	jsonKey []byte
	jwt     *jwt.Config
}

func ProcessAccountFile(text string) (*ServiceAccount, error) {
	// Assume text is a JSON string
	// This func is used for validation now to avoid causing errors in NewClientGCE function
	var err error
	var data []byte
	conf, err := google.JWTConfigFromJSON([]byte(text), DriverScopes...)
	if err != nil {
		// If text was not JSON, assume it is a file path instead
		if _, err = os.Stat(text); os.IsNotExist(err) {
			return nil, fmt.Errorf(
				"account_file path does not exist: %s",
				text)
		}
		data, err = ioutil.ReadFile(text)
		if err != nil {
			return nil, fmt.Errorf(
				"Error reading account_file from path '%s': %s",
				text, err)
		}
		conf, err = google.JWTConfigFromJSON(data, DriverScopes...)
		if err != nil {
			return nil, fmt.Errorf("Error parsing account_file: %s", err)
		}
	}
	return &ServiceAccount{
		jsonKey: data,
		jwt:     conf,
	}, nil
}
