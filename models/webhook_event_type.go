package models
type WebhookEventType int

const (
    SUBSCRIBE_ALL_WEBHOOKEVENTTYPE WebhookEventType = iota
    PERSON_CREATED_WEBHOOKEVENTTYPE
    PERSON_VERIFIED_WEBHOOKEVENTTYPE
    PERSON_REJECTED_WEBHOOKEVENTTYPE
    PERSON_RESTRICTED_WEBHOOKEVENTTYPE
    PERSON_UNVERIFIED_WEBHOOKEVENTTYPE
    PERSON_CLOSED_WEBHOOKEVENTTYPE
    PERSON_ALL_WEBHOOKEVENTTYPE
    BUSINESS_CREATED_WEBHOOKEVENTTYPE
    BUSINESS_VERIFIED_WEBHOOKEVENTTYPE
    BUSINESS_REJECTED_WEBHOOKEVENTTYPE
    BUSINESS_RESTRICTED_WEBHOOKEVENTTYPE
    BUSINESS_UNVERIFIED_WEBHOOKEVENTTYPE
    BUSINESS_CLOSED_WEBHOOKEVENTTYPE
    BUSINESS_ALL_WEBHOOKEVENTTYPE
    EXTERNAL_ACCOUNT_CREATED_WEBHOOKEVENTTYPE
    EXTERNAL_ACCOUNT_VERIFIED_WEBHOOKEVENTTYPE
    EXTERNAL_ACCOUNT_REJECTED_WEBHOOKEVENTTYPE
    EXTERNAL_ACCOUNT_RESTRICTED_WEBHOOKEVENTTYPE
    EXTERNAL_ACCOUNT_UNVERIFIED_WEBHOOKEVENTTYPE
    EXTERNAL_ACCOUNT_CLOSED_WEBHOOKEVENTTYPE
    EXTERNAL_ACCOUNT_ALL_WEBHOOKEVENTTYPE
    TRANSFER_UPDATED_WEBHOOKEVENTTYPE
    TRANSFER_ALL_WEBHOOKEVENTTYPE
    VERIFY_ACCOUNT_UPDATED_WEBHOOKEVENTTYPE
)

func (i WebhookEventType) String() string {
    return []string{"subscribe_all", "person_created", "person_verified", "person_rejected", "person_restricted", "person_unverified", "person_closed", "person_all", "business_created", "business_verified", "business_rejected", "business_restricted", "business_unverified", "business_closed", "business_all", "external_account_created", "external_account_verified", "external_account_rejected", "external_account_restricted", "external_account_unverified", "external_account_closed", "external_account_all", "transfer_updated", "transfer_all", "verify_account_updated"}[i]
}
func ParseWebhookEventType(v string) (any, error) {
    result := SUBSCRIBE_ALL_WEBHOOKEVENTTYPE
    switch v {
        case "subscribe_all":
            result = SUBSCRIBE_ALL_WEBHOOKEVENTTYPE
        case "person_created":
            result = PERSON_CREATED_WEBHOOKEVENTTYPE
        case "person_verified":
            result = PERSON_VERIFIED_WEBHOOKEVENTTYPE
        case "person_rejected":
            result = PERSON_REJECTED_WEBHOOKEVENTTYPE
        case "person_restricted":
            result = PERSON_RESTRICTED_WEBHOOKEVENTTYPE
        case "person_unverified":
            result = PERSON_UNVERIFIED_WEBHOOKEVENTTYPE
        case "person_closed":
            result = PERSON_CLOSED_WEBHOOKEVENTTYPE
        case "person_all":
            result = PERSON_ALL_WEBHOOKEVENTTYPE
        case "business_created":
            result = BUSINESS_CREATED_WEBHOOKEVENTTYPE
        case "business_verified":
            result = BUSINESS_VERIFIED_WEBHOOKEVENTTYPE
        case "business_rejected":
            result = BUSINESS_REJECTED_WEBHOOKEVENTTYPE
        case "business_restricted":
            result = BUSINESS_RESTRICTED_WEBHOOKEVENTTYPE
        case "business_unverified":
            result = BUSINESS_UNVERIFIED_WEBHOOKEVENTTYPE
        case "business_closed":
            result = BUSINESS_CLOSED_WEBHOOKEVENTTYPE
        case "business_all":
            result = BUSINESS_ALL_WEBHOOKEVENTTYPE
        case "external_account_created":
            result = EXTERNAL_ACCOUNT_CREATED_WEBHOOKEVENTTYPE
        case "external_account_verified":
            result = EXTERNAL_ACCOUNT_VERIFIED_WEBHOOKEVENTTYPE
        case "external_account_rejected":
            result = EXTERNAL_ACCOUNT_REJECTED_WEBHOOKEVENTTYPE
        case "external_account_restricted":
            result = EXTERNAL_ACCOUNT_RESTRICTED_WEBHOOKEVENTTYPE
        case "external_account_unverified":
            result = EXTERNAL_ACCOUNT_UNVERIFIED_WEBHOOKEVENTTYPE
        case "external_account_closed":
            result = EXTERNAL_ACCOUNT_CLOSED_WEBHOOKEVENTTYPE
        case "external_account_all":
            result = EXTERNAL_ACCOUNT_ALL_WEBHOOKEVENTTYPE
        case "transfer_updated":
            result = TRANSFER_UPDATED_WEBHOOKEVENTTYPE
        case "transfer_all":
            result = TRANSFER_ALL_WEBHOOKEVENTTYPE
        case "verify_account_updated":
            result = VERIFY_ACCOUNT_UPDATED_WEBHOOKEVENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhookEventType(values []WebhookEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhookEventType) isMultiValue() bool {
    return false
}
