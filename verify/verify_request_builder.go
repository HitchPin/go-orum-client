package verify

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// VerifyRequestBuilder builds and executes requests for operations under \verify
type VerifyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Accounts the accounts property
// returns a *AccountsRequestBuilder when successful
func (m *VerifyRequestBuilder) Accounts()(*AccountsRequestBuilder) {
    return NewAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVerifyRequestBuilderInternal instantiates a new VerifyRequestBuilder and sets the default values.
func NewVerifyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VerifyRequestBuilder) {
    m := &VerifyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/verify", pathParameters),
    }
    return m
}
// NewVerifyRequestBuilder instantiates a new VerifyRequestBuilder and sets the default values.
func NewVerifyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VerifyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVerifyRequestBuilderInternal(urlParams, requestAdapter)
}
