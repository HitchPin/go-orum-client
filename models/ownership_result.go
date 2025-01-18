package models
// Result of Ownership match
type OwnershipResult int

const (
    MATCH_OWNERSHIPRESULT OwnershipResult = iota
    NOT_A_MATCH_OWNERSHIPRESULT
    NOT_FOUND_OWNERSHIPRESULT
)

func (i OwnershipResult) String() string {
    return []string{"match", "not_a_match", "not_found"}[i]
}
func ParseOwnershipResult(v string) (any, error) {
    result := MATCH_OWNERSHIPRESULT
    switch v {
        case "match":
            result = MATCH_OWNERSHIPRESULT
        case "not_a_match":
            result = NOT_A_MATCH_OWNERSHIPRESULT
        case "not_found":
            result = NOT_FOUND_OWNERSHIPRESULT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOwnershipResult(values []OwnershipResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OwnershipResult) isMultiValue() bool {
    return false
}
