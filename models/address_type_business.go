package models
// Type of address. Always 'legal' for business object.
type AddressTypeBusiness int

const (
    LEGAL_ADDRESSTYPEBUSINESS AddressTypeBusiness = iota
)

func (i AddressTypeBusiness) String() string {
    return []string{"legal"}[i]
}
func ParseAddressTypeBusiness(v string) (any, error) {
    result := LEGAL_ADDRESSTYPEBUSINESS
    switch v {
        case "legal":
            result = LEGAL_ADDRESSTYPEBUSINESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAddressTypeBusiness(values []AddressTypeBusiness) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AddressTypeBusiness) isMultiValue() bool {
    return false
}
