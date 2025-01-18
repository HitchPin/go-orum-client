package models
// Status of card in Orum system.
type CardStatus int

const (
    CREATED_CARDSTATUS CardStatus = iota
    VERIFIED_CARDSTATUS
    CLOSED_CARDSTATUS
)

func (i CardStatus) String() string {
    return []string{"created", "verified", "closed"}[i]
}
func ParseCardStatus(v string) (any, error) {
    result := CREATED_CARDSTATUS
    switch v {
        case "created":
            result = CREATED_CARDSTATUS
        case "verified":
            result = VERIFIED_CARDSTATUS
        case "closed":
            result = CLOSED_CARDSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCardStatus(values []CardStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CardStatus) isMultiValue() bool {
    return false
}
