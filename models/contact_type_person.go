package models
// Type of contact information associated with a person - 'email' or 'phone'.
type ContactTypePerson int

const (
    EMAIL_CONTACTTYPEPERSON ContactTypePerson = iota
    PHONE_CONTACTTYPEPERSON
)

func (i ContactTypePerson) String() string {
    return []string{"email", "phone"}[i]
}
func ParseContactTypePerson(v string) (any, error) {
    result := EMAIL_CONTACTTYPEPERSON
    switch v {
        case "email":
            result = EMAIL_CONTACTTYPEPERSON
        case "phone":
            result = PHONE_CONTACTTYPEPERSON
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeContactTypePerson(values []ContactTypePerson) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContactTypePerson) isMultiValue() bool {
    return false
}
