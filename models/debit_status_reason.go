package models
// Status reason for failed and blocked accounts.
type DebitStatusReason int

const (
    BLOCKED_ACCOUNT_DEBITSTATUSREASON DebitStatusReason = iota
    CLOSED_ACCOUNT_DEBITSTATUSREASON
    DECEASED_PARTY_DEBITSTATUSREASON
    DUPLICATED_TRANSFER_DEBITSTATUSREASON
    INVALID_ACCOUNT_DEBITSTATUSREASON
    INVALID_FIELD_DEBITSTATUSREASON
    INVALID_ROUTING_DEBITSTATUSREASON
    REGULATORY_ERROR_DEBITSTATUSREASON
    UNAVAILABLE_FINANCIAL_INSTITUTION_DEBITSTATUSREASON
    UNEXPECTED_ERROR_DEBITSTATUSREASON
    UNSUPPORTED_TRANSFER_DEBITSTATUSREASON
    DEBIT_BLOCKED_DEBITSTATUSREASON
    UNAUTHORIZED_TRANSFER_DEBITSTATUSREASON
)

func (i DebitStatusReason) String() string {
    return []string{"blocked_account", "closed_account", "deceased_party", "duplicated_transfer", "invalid_account", "invalid_field", "invalid_routing", "regulatory_error", "unavailable_financial_institution", "unexpected_error", "unsupported_transfer", "debit_blocked", "unauthorized_transfer"}[i]
}
func ParseDebitStatusReason(v string) (any, error) {
    result := BLOCKED_ACCOUNT_DEBITSTATUSREASON
    switch v {
        case "blocked_account":
            result = BLOCKED_ACCOUNT_DEBITSTATUSREASON
        case "closed_account":
            result = CLOSED_ACCOUNT_DEBITSTATUSREASON
        case "deceased_party":
            result = DECEASED_PARTY_DEBITSTATUSREASON
        case "duplicated_transfer":
            result = DUPLICATED_TRANSFER_DEBITSTATUSREASON
        case "invalid_account":
            result = INVALID_ACCOUNT_DEBITSTATUSREASON
        case "invalid_field":
            result = INVALID_FIELD_DEBITSTATUSREASON
        case "invalid_routing":
            result = INVALID_ROUTING_DEBITSTATUSREASON
        case "regulatory_error":
            result = REGULATORY_ERROR_DEBITSTATUSREASON
        case "unavailable_financial_institution":
            result = UNAVAILABLE_FINANCIAL_INSTITUTION_DEBITSTATUSREASON
        case "unexpected_error":
            result = UNEXPECTED_ERROR_DEBITSTATUSREASON
        case "unsupported_transfer":
            result = UNSUPPORTED_TRANSFER_DEBITSTATUSREASON
        case "debit_blocked":
            result = DEBIT_BLOCKED_DEBITSTATUSREASON
        case "unauthorized_transfer":
            result = UNAUTHORIZED_TRANSFER_DEBITSTATUSREASON
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDebitStatusReason(values []DebitStatusReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DebitStatusReason) isMultiValue() bool {
    return false
}
