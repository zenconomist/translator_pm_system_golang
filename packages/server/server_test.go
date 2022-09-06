package server_test

import (
	"bytes"
	dbc "dbconn"
	"dto"
	"encoding/json"
	"entities"
	"environment"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

var env environment.Env
var dbh dbc.TestPostgreDb // test/prod is here to decide

func TestCreateFirm(t *testing.T) {
	env.EnvFactory(&dbh)
	id := uuid.New().String()
	firm := dto.FirmRequestDTO{
		Name: "test" + id,
		FirmAddress: dto.AddressRequestDTO{
			CountryCode: "HU",
			PostCode:    "34534",
			City:        "valahol",
			Address:     "jani utca 3/b.",
		},
		MainEmail: "test@test" + id + ".com",
	}
	msg := doHttpRequest(http.MethodPost, "http://localhost/firm", firm)
	if msg[:28] != "{\"message\":\"New firm created" {
		t.Errorf("couldn't do request")
	}
	db := env.DbHandler.PassConnection()
	var firmEntity entities.Firm
	db.Last(&firmEntity)
	if firmEntity.Name != "test"+id {
		t.Errorf("names don't match of last inserted firm, and current id")
	}
	if firmEntity.FirmAddress.City != "valahol" {
		t.Errorf("city don't match of last inserted firm, and currently inserted firm dto")
	}
	if firmEntity.FirmAddress.Address != "jani utca 3/b." {
		t.Errorf("address don't match of last inserted firm, and currently inserted firm dto")
	}

	// should not be able to create a firm with the same parameters
	msgRepeat := doHttpRequest(http.MethodPost, "http://localhost/firm", firm)
	if msgRepeat[:28] == "{\"message\":\"New firm created" {
		t.Errorf("should not be able to create a firm with the same parameters")
	}

}

func TestCreateUser(t *testing.T) {
	id := uuid.New().String()
	user := dto.UserRequestDTO{
		Name:                id,
		UserName:            id,
		Email:               id + "@test.hu",
		Active:              true,
		BillingCurrencyName: "HUF",
		Specialities:        "medical, mechanical, engineering",
		Languages:           "HU->EN, EN->HU, HU->DE, DE->HU",
	}
	msg := doHttpRequest(http.MethodPost, "http://localhost/user", user)
	if msg[:28] != "{\"message\":\"New user created" {
		t.Errorf("couldn't do request")
	}
	// should not be able to create a firm with the same parameters
	msgRepeat := doHttpRequest(http.MethodPost, "http://localhost/user", user)
	if msgRepeat[:28] == "{\"message\":\"New user created" {
		t.Errorf("should not be able to create a firm with the same parameters")
	}
}

func doHttpRequest(method string, url string, v interface{}) string {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&v)
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return "error"
	}
	resp, err1 := client.Do(req)
	if err1 != nil {
		return "error"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}
