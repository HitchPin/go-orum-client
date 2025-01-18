package deliver

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BusinessesItemExternalRequestBuilder builds and executes requests for operations under \deliver\businesses\{id}\external
type BusinessesItemExternalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Accounts the accounts property
// returns a *BusinessesItemExternalAccountsRequestBuilder when successful
func (m *BusinessesItemExternalRequestBuilder) Accounts()(*BusinessesItemExternalAccountsRequestBuilder) {
    return NewBusinessesItemExternalAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewBusinessesItemExternalRequestBuilderInternal instantiates a new BusinessesItemExternalRequestBuilder and sets the default values.
func NewBusinessesItemExternalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessesItemExternalRequestBuilder) {
    m := &BusinessesItemExternalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/businesses/{id}/external", pathParameters),
    }
    return m
}
// NewBusinessesItemExternalRequestBuilder instantiates a new BusinessesItemExternalRequestBuilder and sets the default values.
func NewBusinessesItemExternalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessesItemExternalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessesItemExternalRequestBuilderInternal(urlParams, requestAdapter)
}
