package deliver

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// PersonsItemExternalAccountsRequestBuilder builds and executes requests for operations under \deliver\persons\{id}\external\accounts
type PersonsItemExternalAccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PersonsItemExternalAccountsRequestBuilderGetQueryParameters get all external account objects, associated with a specific person
type PersonsItemExternalAccountsRequestBuilderGetQueryParameters struct {
    // Index for paginated results
    Index *int32 `uriparametername:"index"`
    // Max number of results to return
    Size *int32 `uriparametername:"size"`
}
// PersonsItemExternalAccountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PersonsItemExternalAccountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PersonsItemExternalAccountsRequestBuilderGetQueryParameters
}
// NewPersonsItemExternalAccountsRequestBuilderInternal instantiates a new PersonsItemExternalAccountsRequestBuilder and sets the default values.
func NewPersonsItemExternalAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PersonsItemExternalAccountsRequestBuilder) {
    m := &PersonsItemExternalAccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/persons/{id}/external/accounts{?index*,size*}", pathParameters),
    }
    return m
}
// NewPersonsItemExternalAccountsRequestBuilder instantiates a new PersonsItemExternalAccountsRequestBuilder and sets the default values.
func NewPersonsItemExternalAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PersonsItemExternalAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPersonsItemExternalAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all external account objects, associated with a specific person
// returns a ExternalAccountsResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 404 status code
func (m *PersonsItemExternalAccountsRequestBuilder) Get(ctx context.Context, requestConfiguration *PersonsItemExternalAccountsRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.ExternalAccountsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "404": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateExternalAccountsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.ExternalAccountsResponseable), nil
}
// ToGetRequestInformation get all external account objects, associated with a specific person
// returns a *RequestInformation when successful
func (m *PersonsItemExternalAccountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PersonsItemExternalAccountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PersonsItemExternalAccountsRequestBuilder when successful
func (m *PersonsItemExternalAccountsRequestBuilder) WithUrl(rawUrl string)(*PersonsItemExternalAccountsRequestBuilder) {
    return NewPersonsItemExternalAccountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
