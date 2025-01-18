package models
// Type of contact information associated with a business - 'email', 'phone' or 'website'.
type ContactTypeBusiness int

const (
    EMAIL_CONTACTTYPEBUSINESS ContactTypeBusiness = iota
    PHONE_CONTACTTYPEBUSINESS
    WEBSITE_CONTACTTYPEBUSINESS
)

func (i ContactTypeBusiness) String() string {
    return []string{"email", "phone", "website"}[i]
}
func ParseContactTypeBusiness(v string) (any, error) {
    result := EMAIL_CONTACTTYPEBUSINESS
    switch v {
        case "email":
            result = EMAIL_CONTACTTYPEBUSINESS
        case "phone":
            result = PHONE_CONTACTTYPEBUSINESS
        case "website":
            result = WEBSITE_CONTACTTYPEBUSINESS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeContactTypeBusiness(values []ContactTypeBusiness) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContactTypeBusiness) isMultiValue() bool {
    return false
}
