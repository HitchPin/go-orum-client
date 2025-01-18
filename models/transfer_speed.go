package models
// Speed of transfer.
type TransferSpeed int

const (
    ASAP_TRANSFERSPEED TransferSpeed = iota
    STANDARD_TRANSFERSPEED
    SAME_DAY_TRANSFERSPEED
    WIRE_TRANSFERSPEED
)

func (i TransferSpeed) String() string {
    return []string{"asap", "standard", "same_day", "wire"}[i]
}
func ParseTransferSpeed(v string) (any, error) {
    result := ASAP_TRANSFERSPEED
    switch v {
        case "asap":
            result = ASAP_TRANSFERSPEED
        case "standard":
            result = STANDARD_TRANSFERSPEED
        case "same_day":
            result = SAME_DAY_TRANSFERSPEED
        case "wire":
            result = WIRE_TRANSFERSPEED
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransferSpeed(values []TransferSpeed) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransferSpeed) isMultiValue() bool {
    return false
}
