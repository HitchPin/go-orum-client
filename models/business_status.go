package models
// Status of Business.
type BusinessStatus int

const (
    CREATED_BUSINESSSTATUS BusinessStatus = iota
    VERIFIED_BUSINESSSTATUS
    REJECTED_BUSINESSSTATUS
    RESTRICTED_BUSINESSSTATUS
    CLOSED_BUSINESSSTATUS
)

func (i BusinessStatus) String() string {
    return []string{"created", "verified", "rejected", "restricted", "closed"}[i]
}
func ParseBusinessStatus(v string) (any, error) {
    result := CREATED_BUSINESSSTATUS
    switch v {
        case "created":
            result = CREATED_BUSINESSSTATUS
        case "verified":
            result = VERIFIED_BUSINESSSTATUS
        case "rejected":
            result = REJECTED_BUSINESSSTATUS
        case "restricted":
            result = RESTRICTED_BUSINESSSTATUS
        case "closed":
            result = CLOSED_BUSINESSSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBusinessStatus(values []BusinessStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BusinessStatus) isMultiValue() bool {
    return false
}
