package models
// Type of bank account - checking or savings.
type AccountType int

const (
    CHECKING_ACCOUNTTYPE AccountType = iota
    SAVINGS_ACCOUNTTYPE
)

func (i AccountType) String() string {
    return []string{"checking", "savings"}[i]
}
func ParseAccountType(v string) (any, error) {
    result := CHECKING_ACCOUNTTYPE
    switch v {
        case "checking":
            result = CHECKING_ACCOUNTTYPE
        case "savings":
            result = SAVINGS_ACCOUNTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAccountType(values []AccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccountType) isMultiValue() bool {
    return false
}
