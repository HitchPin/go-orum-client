package models
// Status of the external account.
type ExternalAccountStatus int

const (
    CREATED_EXTERNALACCOUNTSTATUS ExternalAccountStatus = iota
    VERIFIED_EXTERNALACCOUNTSTATUS
    REJECTED_EXTERNALACCOUNTSTATUS
    RESTRICTED_EXTERNALACCOUNTSTATUS
    CLOSED_EXTERNALACCOUNTSTATUS
)

func (i ExternalAccountStatus) String() string {
    return []string{"created", "verified", "rejected", "restricted", "closed"}[i]
}
func ParseExternalAccountStatus(v string) (any, error) {
    result := CREATED_EXTERNALACCOUNTSTATUS
    switch v {
        case "created":
            result = CREATED_EXTERNALACCOUNTSTATUS
        case "verified":
            result = VERIFIED_EXTERNALACCOUNTSTATUS
        case "rejected":
            result = REJECTED_EXTERNALACCOUNTSTATUS
        case "restricted":
            result = RESTRICTED_EXTERNALACCOUNTSTATUS
        case "closed":
            result = CLOSED_EXTERNALACCOUNTSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeExternalAccountStatus(values []ExternalAccountStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ExternalAccountStatus) isMultiValue() bool {
    return false
}
