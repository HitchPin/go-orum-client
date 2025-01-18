package models
// Status of account verification.
type VerificationStatus int

const (
    PENDING_VERIFICATIONSTATUS VerificationStatus = iota
    VALID_VERIFICATIONSTATUS
    CLOSED_VERIFICATIONSTATUS
    INVALID_VERIFICATIONSTATUS
    FAILED_VERIFICATIONSTATUS
)

func (i VerificationStatus) String() string {
    return []string{"pending", "valid", "closed", "invalid", "failed"}[i]
}
func ParseVerificationStatus(v string) (any, error) {
    result := PENDING_VERIFICATIONSTATUS
    switch v {
        case "pending":
            result = PENDING_VERIFICATIONSTATUS
        case "valid":
            result = VALID_VERIFICATIONSTATUS
        case "closed":
            result = CLOSED_VERIFICATIONSTATUS
        case "invalid":
            result = INVALID_VERIFICATIONSTATUS
        case "failed":
            result = FAILED_VERIFICATIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeVerificationStatus(values []VerificationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VerificationStatus) isMultiValue() bool {
    return false
}
