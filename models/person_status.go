package models
// Status of person in Orum system.
type PersonStatus int

const (
    CREATED_PERSONSTATUS PersonStatus = iota
    VERIFIED_PERSONSTATUS
    REJECTED_PERSONSTATUS
    RESTRICTED_PERSONSTATUS
    CLOSED_PERSONSTATUS
)

func (i PersonStatus) String() string {
    return []string{"created", "verified", "rejected", "restricted", "closed"}[i]
}
func ParsePersonStatus(v string) (any, error) {
    result := CREATED_PERSONSTATUS
    switch v {
        case "created":
            result = CREATED_PERSONSTATUS
        case "verified":
            result = VERIFIED_PERSONSTATUS
        case "rejected":
            result = REJECTED_PERSONSTATUS
        case "restricted":
            result = RESTRICTED_PERSONSTATUS
        case "closed":
            result = CLOSED_PERSONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePersonStatus(values []PersonStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PersonStatus) isMultiValue() bool {
    return false
}
