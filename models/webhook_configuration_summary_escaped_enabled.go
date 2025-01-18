package models
type WebhookConfigurationSummary_enabled int

const (
    TRUE_WEBHOOKCONFIGURATIONSUMMARY_ENABLED WebhookConfigurationSummary_enabled = iota
    FALSE_WEBHOOKCONFIGURATIONSUMMARY_ENABLED
)

func (i WebhookConfigurationSummary_enabled) String() string {
    return []string{"true", "false"}[i]
}
func ParseWebhookConfigurationSummary_enabled(v string) (any, error) {
    result := TRUE_WEBHOOKCONFIGURATIONSUMMARY_ENABLED
    switch v {
        case "true":
            result = TRUE_WEBHOOKCONFIGURATIONSUMMARY_ENABLED
        case "false":
            result = FALSE_WEBHOOKCONFIGURATIONSUMMARY_ENABLED
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWebhookConfigurationSummary_enabled(values []WebhookConfigurationSummary_enabled) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhookConfigurationSummary_enabled) isMultiValue() bool {
    return false
}
