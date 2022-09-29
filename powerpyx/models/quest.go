package models

// Cyberpunk2077Quest is the model used for unmarshalling Cyberpunk 2077 quest response from powerpyx.
type Cyberpunk2077Quest struct {
	NameLinkPair map[string]string
	SubCategory  string
}
