package models
// Status of account debit.
type DebitStatus int

const (
    PENDING_DEBITSTATUS DebitStatus = iota
    VALID_DEBITSTATUS
    BLOCKED_DEBITSTATUS
    FAILED_DEBITSTATUS
)

func (i DebitStatus) String() string {
    return []string{"pending", "valid", "blocked", "failed"}[i]
}
func ParseDebitStatus(v string) (any, error) {
    result := PENDING_DEBITSTATUS
    switch v {
        case "pending":
            result = PENDING_DEBITSTATUS
        case "valid":
            result = VALID_DEBITSTATUS
        case "blocked":
            result = BLOCKED_DEBITSTATUS
        case "failed":
            result = FAILED_DEBITSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDebitStatus(values []DebitStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DebitStatus) isMultiValue() bool {
    return false
}
