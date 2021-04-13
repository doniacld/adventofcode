package passport

// Passport definition
type Passport map[string]string

// NewPassport returns a new passport
func NewPassport() Passport {
	return make(Passport, 0)
}

// IsValid returns id the passport is valid depending on the mandatory fields
func (p Passport) IsValid() bool {
	for key, validator := range mandatoryFields {
		value, ok := p[key]
		if !ok {
			return false
		}
		if !validator(value) {
			return false
		}
	}
	return true
}
