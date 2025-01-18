package models
// Status of account ownership.
type OwnershipStatus int

const (
    PENDING_OWNERSHIPSTATUS OwnershipStatus = iota
    FULL_OWNERSHIPSTATUS
    PARTIAL_OWNERSHIPSTATUS
    NOT_A_MATCH_OWNERSHIPSTATUS
    NOT_FOUND_OWNERSHIPSTATUS
    CANCELED_OWNERSHIPSTATUS
    FAILED_OWNERSHIPSTATUS
)

func (i OwnershipStatus) String() string {
    return []string{"pending", "full", "partial", "not_a_match", "not_found", "canceled", "failed"}[i]
}
func ParseOwnershipStatus(v string) (any, error) {
    result := PENDING_OWNERSHIPSTATUS
    switch v {
        case "pending":
            result = PENDING_OWNERSHIPSTATUS
        case "full":
            result = FULL_OWNERSHIPSTATUS
        case "partial":
            result = PARTIAL_OWNERSHIPSTATUS
        case "not_a_match":
            result = NOT_A_MATCH_OWNERSHIPSTATUS
        case "not_found":
            result = NOT_FOUND_OWNERSHIPSTATUS
        case "canceled":
            result = CANCELED_OWNERSHIPSTATUS
        case "failed":
            result = FAILED_OWNERSHIPSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOwnershipStatus(values []OwnershipStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OwnershipStatus) isMultiValue() bool {
    return false
}
