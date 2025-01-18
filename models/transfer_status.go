package models
// Describes the current status of the transfer.
type TransferStatus int

const (
    COMPLETED_TRANSFERSTATUS TransferStatus = iota
    CREATED_TRANSFERSTATUS
    FAILED_TRANSFERSTATUS
    PENDING_TRANSFERSTATUS
    SETTLED_TRANSFERSTATUS
)

func (i TransferStatus) String() string {
    return []string{"completed", "created", "failed", "pending", "settled"}[i]
}
func ParseTransferStatus(v string) (any, error) {
    result := COMPLETED_TRANSFERSTATUS
    switch v {
        case "completed":
            result = COMPLETED_TRANSFERSTATUS
        case "created":
            result = CREATED_TRANSFERSTATUS
        case "failed":
            result = FAILED_TRANSFERSTATUS
        case "pending":
            result = PENDING_TRANSFERSTATUS
        case "settled":
            result = SETTLED_TRANSFERSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransferStatus(values []TransferStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransferStatus) isMultiValue() bool {
    return false
}
