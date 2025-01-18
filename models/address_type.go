package models
// Type of address.
type AddressType int

const (
    HOME_ADDRESSTYPE AddressType = iota
    LEGAL_ADDRESSTYPE
)

func (i AddressType) String() string {
    return []string{"home", "legal"}[i]
}
func ParseAddressType(v string) (any, error) {
    result := HOME_ADDRESSTYPE
    switch v {
        case "home":
            result = HOME_ADDRESSTYPE
        case "legal":
            result = LEGAL_ADDRESSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAddressType(values []AddressType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AddressType) isMultiValue() bool {
    return false
}
