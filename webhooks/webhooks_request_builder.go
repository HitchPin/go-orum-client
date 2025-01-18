package webhooks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WebhooksRequestBuilder builds and executes requests for operations under \webhooks
type WebhooksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Configurations the configurations property
// returns a *ConfigurationsRequestBuilder when successful
func (m *WebhooksRequestBuilder) Configurations()(*ConfigurationsRequestBuilder) {
    return NewConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWebhooksRequestBuilderInternal instantiates a new WebhooksRequestBuilder and sets the default values.
func NewWebhooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WebhooksRequestBuilder) {
    m := &WebhooksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/webhooks", pathParameters),
    }
    return m
}
// NewWebhooksRequestBuilder instantiates a new WebhooksRequestBuilder and sets the default values.
func NewWebhooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WebhooksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWebhooksRequestBuilderInternal(urlParams, requestAdapter)
}
// Invoke the invoke property
// returns a *InvokeRequestBuilder when successful
func (m *WebhooksRequestBuilder) Invoke()(*InvokeRequestBuilder) {
    return NewInvokeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Secret the secret property
// returns a *SecretRequestBuilder when successful
func (m *WebhooksRequestBuilder) Secret()(*SecretRequestBuilder) {
    return NewSecretRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
