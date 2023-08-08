package helper

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func LoginPetros(username, password string) (bool, error) {
	response, err := http.PostForm(os.Getenv("PETROS_URL")+"/wp-login.php",
		url.Values{"log": {username}, "pwd": {password}})
	if err != nil {
		return false, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	return !strings.Contains(string(responseData), "<strong>Error</strong>"), nil
}
