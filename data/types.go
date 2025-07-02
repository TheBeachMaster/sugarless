package data

import "encoding/json"

type ExtensionsNode struct {
	Eid              string   `graphql:"eid" redis:"eid" faker:"uuid_digit"`
	ExtensionName    string   `graphql:"extensionName" redis:"extensionName" faker:"word"`
	ExtensionTenancy []string `graphql:"extensionTenancy" redis:"extensionTenancy" faker:"oneof: CLOUD, DATA CENTER, SERVER"`
	IsListed         bool     `graphql:"isListed" redis:"isListed"`
	Logo             string   `graphql:"extensionLogo" redis:"extensionLogo" faker:"url"`
	ShortDescription string   `graphql:"extensionDescription" redis:"extensionDescription" faker:"sentence"`
	LongDescription  string   `graphql:"extensionLongDescription" redis:"extensionLongDescription" faker:"paragraph"`
	IsPrivileged     bool     `graphql:"extensionIsPrevileged" redis:"extensionIsPrevileged"`
	IsEssential      bool     `graphql:"extensionIsEssential" redis:"extensionIsEssential"`
	Project          struct {
		VendorId string `graphql:"vendorAccountID" redis:"vendorAccountID" faker:"uuid_hyphenated"`
	}
}

type Edges []struct {
	Node ExtensionsNode `graphql:"node" redis:"node"`
}

type ExtensionList struct {
	Extensions []ExtensionsNode `redis:"extensions"`
}

func (m *ExtensionList) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *ExtensionList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}
