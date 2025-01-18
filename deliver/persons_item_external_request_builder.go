package deliver

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PersonsItemExternalRequestBuilder builds and executes requests for operations under \deliver\persons\{id}\external
type PersonsItemExternalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Accounts the accounts property
// returns a *PersonsItemExternalAccountsRequestBuilder when successful
func (m *PersonsItemExternalRequestBuilder) Accounts()(*PersonsItemExternalAccountsRequestBuilder) {
    return NewPersonsItemExternalAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPersonsItemExternalRequestBuilderInternal instantiates a new PersonsItemExternalRequestBuilder and sets the default values.
func NewPersonsItemExternalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PersonsItemExternalRequestBuilder) {
    m := &PersonsItemExternalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/persons/{id}/external", pathParameters),
    }
    return m
}
// NewPersonsItemExternalRequestBuilder instantiates a new PersonsItemExternalRequestBuilder and sets the default values.
func NewPersonsItemExternalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PersonsItemExternalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPersonsItemExternalRequestBuilderInternal(urlParams, requestAdapter)
}
