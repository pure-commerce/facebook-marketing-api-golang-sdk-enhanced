package types

import (
	"crypto/sha256"
	"encoding/hex"
)

// CutomerInfromation entity https://developers.facebook.com/docs/marketing-api/conversions-api/parameters/customer-information-parameters
type CustomerInformation struct {
	Email           []string `json:"em,omitempty"`
	PhoneNumber     []string `json:"ph,omitempty"`
	FirstName       string   `json:"fn,omitempty"`
	LastName        string   `json:"ln,omitempty"`
	City            string   `json:"ct,omitempty"`
	State           string   `json:"st,omitempty"`
	Zip             string   `json:"zp,omitempty"`
	Country         string   `json:"country,omitempty"`
	ClientIPAddress string   `json:"client_ip_address,omitempty"`
	ClientUserAgent string   `json:"client_user_agent,omitempty"`
	Fbc             string   `json:"fbc,omitempty"`
	Fbp             string   `json:"fbp,omitempty"`
}

func NewCustomerInformation() CustomerInformation {
	return CustomerInformation{}
}

// Adds already hashed email to the struct
func (ci CustomerInformation) WithHashedEmail(hashedEmail string) CustomerInformation {
	ci.Email = []string{hashedEmail}
	return ci
}

// Adds non-hashed email to the struct for it to be hashed.
func (ci CustomerInformation) WithEmail(email string) CustomerInformation {
	hashedEmailBytes := sha256.Sum256([]byte(email))
	hashedEmail := hex.EncodeToString(hashedEmailBytes[:])
	ci.Email = []string{hashedEmail}
	return ci
}

// Adds non-hashed phoneNumber to the struct for it to be hashed.
func (ci CustomerInformation) WithPhoneNumber(email string) CustomerInformation {
	hashedPhoneBytes := sha256.Sum256([]byte(email))
	hashedPhone := hex.EncodeToString(hashedPhoneBytes[:])
	ci.PhoneNumber = []string{hashedPhone}
	return ci
}

func (ci CustomerInformation) WithFirstName(first string) CustomerInformation {
	hashedFirstBytes := sha256.Sum256([]byte(first))
	hashedFirst := hex.EncodeToString(hashedFirstBytes[:])
	ci.FirstName = hashedFirst
	return ci
}

func (ci CustomerInformation) WithLastName(last string) CustomerInformation {
	hashedLastBytes := sha256.Sum256([]byte(last))
	hashedLast := hex.EncodeToString(hashedLastBytes[:])
	ci.LastName = hashedLast
	return ci
}

func (ci CustomerInformation) WithCity(city string) CustomerInformation {
	hashedCityBytes := sha256.Sum256([]byte(city))
	hashedCity := hex.EncodeToString(hashedCityBytes[:])
	ci.City = hashedCity
	return ci
}

func (ci CustomerInformation) WithState(state string) CustomerInformation {
	hashedStateBytes := sha256.Sum256([]byte(state))
	hashedState := hex.EncodeToString(hashedStateBytes[:])
	ci.State = hashedState
	return ci
}

func (ci CustomerInformation) WithZip(zip string) CustomerInformation {
	hashedZipBytes := sha256.Sum256([]byte(zip))
	hashedZip := hex.EncodeToString(hashedZipBytes[:])
	ci.Zip = hashedZip
	return ci
}

func (ci CustomerInformation) WithCountry(country string) CustomerInformation {
	hashedCountryBytes := sha256.Sum256([]byte(country))
	hashedCountry := hex.EncodeToString(hashedCountryBytes[:])
	ci.Country = hashedCountry
	return ci
}

// Adds already hashed email to the struct
func (ci CustomerInformation) WithHashedPhoneNumber(hashedPhoneNumber string) CustomerInformation {
	ci.PhoneNumber = []string{hashedPhoneNumber}
	return ci
}

func (ci CustomerInformation) WithFbc(fbc string) CustomerInformation {
	ci.Fbc = fbc
	return ci
}

func (ci CustomerInformation) WithFbp(fbp string) CustomerInformation {
	ci.Fbp = fbp
	return ci
}

func (ci CustomerInformation) WithClientIPAddress(clientIPAddress string) CustomerInformation {
	ci.ClientIPAddress = clientIPAddress
	return ci
}

func (ci CustomerInformation) WithClientUserAgent(clientUserAgent string) CustomerInformation {
	ci.ClientUserAgent = clientUserAgent
	return ci
}
