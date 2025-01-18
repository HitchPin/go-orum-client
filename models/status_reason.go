package models
// Status reason for failed, invalid, and closed accounts.
type StatusReason int

const (
    BLOCKED_ACCOUNT_STATUSREASON StatusReason = iota
    CLOSED_ACCOUNT_STATUSREASON
    DECEASED_PARTY_STATUSREASON
    DUPLICATED_TRANSFER_STATUSREASON
    INVALID_ACCOUNT_STATUSREASON
    INVALID_FIELD_STATUSREASON
    INVALID_ROUTING_STATUSREASON
    REGULATORY_ERROR_STATUSREASON
    UNAVAILABLE_FINANCIAL_INSTITUTION_STATUSREASON
    UNEXPECTED_ERROR_STATUSREASON
    UNSUPPORTED_TRANSFER_STATUSREASON
)

func (i StatusReason) String() string {
    return []string{"blocked_account", "closed_account", "deceased_party", "duplicated_transfer", "invalid_account", "invalid_field", "invalid_routing", "regulatory_error", "unavailable_financial_institution", "unexpected_error", "unsupported_transfer"}[i]
}
func ParseStatusReason(v string) (any, error) {
    result := BLOCKED_ACCOUNT_STATUSREASON
    switch v {
        case "blocked_account":
            result = BLOCKED_ACCOUNT_STATUSREASON
        case "closed_account":
            result = CLOSED_ACCOUNT_STATUSREASON
        case "deceased_party":
            result = DECEASED_PARTY_STATUSREASON
        case "duplicated_transfer":
            result = DUPLICATED_TRANSFER_STATUSREASON
        case "invalid_account":
            result = INVALID_ACCOUNT_STATUSREASON
        case "invalid_field":
            result = INVALID_FIELD_STATUSREASON
        case "invalid_routing":
            result = INVALID_ROUTING_STATUSREASON
        case "regulatory_error":
            result = REGULATORY_ERROR_STATUSREASON
        case "unavailable_financial_institution":
            result = UNAVAILABLE_FINANCIAL_INSTITUTION_STATUSREASON
        case "unexpected_error":
            result = UNEXPECTED_ERROR_STATUSREASON
        case "unsupported_transfer":
            result = UNSUPPORTED_TRANSFER_STATUSREASON
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeStatusReason(values []StatusReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StatusReason) isMultiValue() bool {
    return false
}
