package models
type CreateWebhookConfiguration_enabled int

const (
    TRUE_CREATEWEBHOOKCONFIGURATION_ENABLED CreateWebhookConfiguration_enabled = iota
    FALSE_CREATEWEBHOOKCONFIGURATION_ENABLED
)

func (i CreateWebhookConfiguration_enabled) String() string {
    return []string{"true", "false"}[i]
}
func ParseCreateWebhookConfiguration_enabled(v string) (any, error) {
    result := TRUE_CREATEWEBHOOKCONFIGURATION_ENABLED
    switch v {
        case "true":
            result = TRUE_CREATEWEBHOOKCONFIGURATION_ENABLED
        case "false":
            result = FALSE_CREATEWEBHOOKCONFIGURATION_ENABLED
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCreateWebhookConfiguration_enabled(values []CreateWebhookConfiguration_enabled) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CreateWebhookConfiguration_enabled) isMultiValue() bool {
    return false
}
