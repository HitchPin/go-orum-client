package models
// Currency code in ISO 4217 format. Only USD is supported.
type TransferCurrency int

const (
    USD_TRANSFERCURRENCY TransferCurrency = iota
)

func (i TransferCurrency) String() string {
    return []string{"USD"}[i]
}
func ParseTransferCurrency(v string) (any, error) {
    result := USD_TRANSFERCURRENCY
    switch v {
        case "USD":
            result = USD_TRANSFERCURRENCY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransferCurrency(values []TransferCurrency) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransferCurrency) isMultiValue() bool {
    return false
}
