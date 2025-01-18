package verify

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// AccountsAccountsItemRequestBuilder builds and executes requests for operations under \verify\accounts\{id}
type AccountsAccountsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccountsAccountsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccountsAccountsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccountsAccountsItemRequestBuilderInternal instantiates a new AccountsAccountsItemRequestBuilder and sets the default values.
func NewAccountsAccountsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountsAccountsItemRequestBuilder) {
    m := &AccountsAccountsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/verify/accounts/{id}", pathParameters),
    }
    return m
}
// NewAccountsAccountsItemRequestBuilder instantiates a new AccountsAccountsItemRequestBuilder and sets the default values.
func NewAccountsAccountsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccountsAccountsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccountsAccountsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Control the control property
// returns a *AccountsItemControlRequestBuilder when successful
func (m *AccountsAccountsItemRequestBuilder) Control()(*AccountsItemControlRequestBuilder) {
    return NewAccountsItemControlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Debits the debits property
// returns a *AccountsItemDebitsRequestBuilder when successful
func (m *AccountsAccountsItemRequestBuilder) Debits()(*AccountsItemDebitsRequestBuilder) {
    return NewAccountsItemDebitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a single verify account.
// returns a VerifyAccountResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 404 status code
func (m *AccountsAccountsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccountsAccountsItemRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.VerifyAccountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "404": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateVerifyAccountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.VerifyAccountResponseable), nil
}
// ToGetRequestInformation get a single verify account.
// returns a *RequestInformation when successful
func (m *AccountsAccountsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccountsAccountsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccountsAccountsItemRequestBuilder when successful
func (m *AccountsAccountsItemRequestBuilder) WithUrl(rawUrl string)(*AccountsAccountsItemRequestBuilder) {
    return NewAccountsAccountsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
