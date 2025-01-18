package models
// Type of address. Always 'home' for person object.
type AddressTypePerson int

const (
    HOME_ADDRESSTYPEPERSON AddressTypePerson = iota
)

func (i AddressTypePerson) String() string {
    return []string{"home"}[i]
}
func ParseAddressTypePerson(v string) (any, error) {
    result := HOME_ADDRESSTYPEPERSON
    switch v {
        case "home":
            result = HOME_ADDRESSTYPEPERSON
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAddressTypePerson(values []AddressTypePerson) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AddressTypePerson) isMultiValue() bool {
    return false
}
