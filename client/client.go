package client

import (
	"fmt"
	"github.com/FXHibon/domotic-hue-manager/config"
	"io"
	"log"
	"net/http"
	"strings"
)

type UnreachableApi error

type InvalidCredentials struct {
	Status  int
	Details string
}

func (i InvalidCredentials) Error() string {
	return fmt.Sprintf("API doesn't accept the provided credentials and responded with status=%v body=%v", i.Status, i.Details)
}

func ValidateCredentials(credentials config.ApiConfiguration) error {
	resp, err := http.Get(fmt.Sprintf("%v/%v", credentials.Uri, credentials.User))

	if err != nil {
		log.Panic(UnreachableApi(err))
	}

	bodyWriter := strings.Builder{}
	_, err = io.Copy(&bodyWriter, resp.Body)
	if err != nil {
		log.Panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		log.Panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Panic(InvalidCredentials{resp.StatusCode, bodyWriter.String()})
	}

	return nil
}
