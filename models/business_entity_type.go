package models
// Describes the type of business entity.
type BusinessEntityType int

const (
    SOLE_PROPRIETORSHIP_BUSINESSENTITYTYPE BusinessEntityType = iota
    PARTNERSHIP_BUSINESSENTITYTYPE
    LIMITED_LIABILITY_PARTNERSHIP_BUSINESSENTITYTYPE
    LIMITED_LIABILITY_COMPANY_BUSINESSENTITYTYPE
    C_CORPORATION_BUSINESSENTITYTYPE
    S_CORPORATION_BUSINESSENTITYTYPE
    B_CORPORATION_BUSINESSENTITYTYPE
    NONPROFIT_CORPORATION_BUSINESSENTITYTYPE
)

func (i BusinessEntityType) String() string {
    return []string{"sole_proprietorship", "partnership", "limited_liability_partnership", "limited_liability_company", "c_corporation", "s_corporation", "b_corporation", "nonprofit_corporation"}[i]
}
func ParseBusinessEntityType(v string) (any, error) {
    result := SOLE_PROPRIETORSHIP_BUSINESSENTITYTYPE
    switch v {
        case "sole_proprietorship":
            result = SOLE_PROPRIETORSHIP_BUSINESSENTITYTYPE
        case "partnership":
            result = PARTNERSHIP_BUSINESSENTITYTYPE
        case "limited_liability_partnership":
            result = LIMITED_LIABILITY_PARTNERSHIP_BUSINESSENTITYTYPE
        case "limited_liability_company":
            result = LIMITED_LIABILITY_COMPANY_BUSINESSENTITYTYPE
        case "c_corporation":
            result = C_CORPORATION_BUSINESSENTITYTYPE
        case "s_corporation":
            result = S_CORPORATION_BUSINESSENTITYTYPE
        case "b_corporation":
            result = B_CORPORATION_BUSINESSENTITYTYPE
        case "nonprofit_corporation":
            result = NONPROFIT_CORPORATION_BUSINESSENTITYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBusinessEntityType(values []BusinessEntityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BusinessEntityType) isMultiValue() bool {
    return false
}
