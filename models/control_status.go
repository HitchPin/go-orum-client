package models
// Status of account control.
type ControlStatus int

const (
    PENDING_CONTROLSTATUS ControlStatus = iota
    VALID_CONTROLSTATUS
    CANCELED_CONTROLSTATUS
)

func (i ControlStatus) String() string {
    return []string{"pending", "valid", "canceled"}[i]
}
func ParseControlStatus(v string) (any, error) {
    result := PENDING_CONTROLSTATUS
    switch v {
        case "pending":
            result = PENDING_CONTROLSTATUS
        case "valid":
            result = VALID_CONTROLSTATUS
        case "canceled":
            result = CANCELED_CONTROLSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeControlStatus(values []ControlStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ControlStatus) isMultiValue() bool {
    return false
}
