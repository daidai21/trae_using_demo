package identity

import (
	"fmt"
	"strings"
)

const (
	CountryCN = "CN"
	CountryID = "ID"
	CountryUS = "US"
)

const (
	ModeNormal   = "normal"
	ModePreSale  = "pre_sale"
	ModeAuction  = "auction"
)

type BusinessIdentity struct {
	Country string
	Mode    string
}

func NewBusinessIdentity(country, mode string) *BusinessIdentity {
	return &BusinessIdentity{
		Country: country,
		Mode:    mode,
	}
}

func (id *BusinessIdentity) String() string {
	return fmt.Sprintf("%s.%s", id.Country, id.Mode)
}

func (id *BusinessIdentity) IsValid() bool {
	validCountries := map[string]bool{
		CountryCN: true,
		CountryID: true,
		CountryUS: true,
	}
	validModes := map[string]bool{
		ModeNormal:  true,
		ModePreSale: true,
		ModeAuction: true,
	}
	return validCountries[id.Country] && validModes[id.Mode]
}

func Parse(identityStr string) (*BusinessIdentity, error) {
	parts := strings.SplitN(identityStr, ".", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid identity format: %s", identityStr)
	}

	id := &BusinessIdentity{
		Country: parts[0],
		Mode:    parts[1],
	}

	if !id.IsValid() {
		return nil, fmt.Errorf("invalid identity: %s", identityStr)
	}

	return id, nil
}
