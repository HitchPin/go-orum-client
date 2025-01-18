package models
// Type of ledger entry.
type LedgerEntryType int

const (
    RTP_LEDGERENTRYTYPE LedgerEntryType = iota
    ACH_LEDGERENTRYTYPE
    SDACH_LEDGERENTRYTYPE
    ACH_RETURN_LEDGERENTRYTYPE
    WIRE_LEDGERENTRYTYPE
    WIRE_RETURN_LEDGERENTRYTYPE
    OUTBOUND_FUNDS_LEDGERENTRYTYPE
    INBOUND_FUNDS_LEDGERENTRYTYPE
    BALANCE_TOP_UP_LEDGERENTRYTYPE
)

func (i LedgerEntryType) String() string {
    return []string{"rtp", "ach", "sdach", "ach_return", "wire", "wire_return", "outbound_funds", "inbound_funds", "balance_top_up"}[i]
}
func ParseLedgerEntryType(v string) (any, error) {
    result := RTP_LEDGERENTRYTYPE
    switch v {
        case "rtp":
            result = RTP_LEDGERENTRYTYPE
        case "ach":
            result = ACH_LEDGERENTRYTYPE
        case "sdach":
            result = SDACH_LEDGERENTRYTYPE
        case "ach_return":
            result = ACH_RETURN_LEDGERENTRYTYPE
        case "wire":
            result = WIRE_LEDGERENTRYTYPE
        case "wire_return":
            result = WIRE_RETURN_LEDGERENTRYTYPE
        case "outbound_funds":
            result = OUTBOUND_FUNDS_LEDGERENTRYTYPE
        case "inbound_funds":
            result = INBOUND_FUNDS_LEDGERENTRYTYPE
        case "balance_top_up":
            result = BALANCE_TOP_UP_LEDGERENTRYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLedgerEntryType(values []LedgerEntryType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LedgerEntryType) isMultiValue() bool {
    return false
}
