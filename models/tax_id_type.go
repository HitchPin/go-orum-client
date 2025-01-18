package models
// The type of tax ID number associated with the business. Note: This field is required for payouts, deposits, or account-to-account transfers.
type TaxIdType int

const (
    TIN_TAXIDTYPE TaxIdType = iota
    EIN_TAXIDTYPE
)

func (i TaxIdType) String() string {
    return []string{"tin", "ein"}[i]
}
func ParseTaxIdType(v string) (any, error) {
    result := TIN_TAXIDTYPE
    switch v {
        case "tin":
            result = TIN_TAXIDTYPE
        case "ein":
            result = EIN_TAXIDTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTaxIdType(values []TaxIdType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TaxIdType) isMultiValue() bool {
    return false
}
