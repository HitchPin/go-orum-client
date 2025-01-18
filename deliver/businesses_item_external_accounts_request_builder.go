package deliver

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// BusinessesItemExternalAccountsRequestBuilder builds and executes requests for operations under \deliver\businesses\{id}\external\accounts
type BusinessesItemExternalAccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessesItemExternalAccountsRequestBuilderGetQueryParameters get all external account objects, associated with a specific business
type BusinessesItemExternalAccountsRequestBuilderGetQueryParameters struct {
    // Index for paginated results
    Index *int32 `uriparametername:"index"`
    // Max number of results to return
    Size *int32 `uriparametername:"size"`
}
// BusinessesItemExternalAccountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessesItemExternalAccountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessesItemExternalAccountsRequestBuilderGetQueryParameters
}
// NewBusinessesItemExternalAccountsRequestBuilderInternal instantiates a new BusinessesItemExternalAccountsRequestBuilder and sets the default values.
func NewBusinessesItemExternalAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessesItemExternalAccountsRequestBuilder) {
    m := &BusinessesItemExternalAccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/businesses/{id}/external/accounts{?index*,size*}", pathParameters),
    }
    return m
}
// NewBusinessesItemExternalAccountsRequestBuilder instantiates a new BusinessesItemExternalAccountsRequestBuilder and sets the default values.
func NewBusinessesItemExternalAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessesItemExternalAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessesItemExternalAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all external account objects, associated with a specific business
// returns a ExternalAccountsResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 404 status code
func (m *BusinessesItemExternalAccountsRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessesItemExternalAccountsRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.ExternalAccountsResponseable, error) {
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
// ToGetRequestInformation get all external account objects, associated with a specific business
// returns a *RequestInformation when successful
func (m *BusinessesItemExternalAccountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessesItemExternalAccountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BusinessesItemExternalAccountsRequestBuilder when successful
func (m *BusinessesItemExternalAccountsRequestBuilder) WithUrl(rawUrl string)(*BusinessesItemExternalAccountsRequestBuilder) {
    return NewBusinessesItemExternalAccountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
