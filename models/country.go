package models
// 2-character ISO country code.
type Country int

const (
    US_COUNTRY Country = iota
)

func (i Country) String() string {
    return []string{"US"}[i]
}
func ParseCountry(v string) (any, error) {
    result := US_COUNTRY
    switch v {
        case "US":
            result = US_COUNTRY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCountry(values []Country) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Country) isMultiValue() bool {
    return false
}
