package models
// Type of customer resource - business, person, or enterprise.
type EndCustomerResourceType int

const (
    BUSINESS_ENDCUSTOMERRESOURCETYPE EndCustomerResourceType = iota
    PERSON_ENDCUSTOMERRESOURCETYPE
    ENTERPRISE_ENDCUSTOMERRESOURCETYPE
)

func (i EndCustomerResourceType) String() string {
    return []string{"business", "person", "enterprise"}[i]
}
func ParseEndCustomerResourceType(v string) (any, error) {
    result := BUSINESS_ENDCUSTOMERRESOURCETYPE
    switch v {
        case "business":
            result = BUSINESS_ENDCUSTOMERRESOURCETYPE
        case "person":
            result = PERSON_ENDCUSTOMERRESOURCETYPE
        case "enterprise":
            result = ENTERPRISE_ENDCUSTOMERRESOURCETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEndCustomerResourceType(values []EndCustomerResourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EndCustomerResourceType) isMultiValue() bool {
    return false
}
