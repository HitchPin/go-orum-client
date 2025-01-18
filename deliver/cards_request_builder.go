package deliver

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// CardsRequestBuilder builds and executes requests for operations under \deliver\cards
type CardsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CardsRequestBuilderGetQueryParameters get all cards by filter.
type CardsRequestBuilderGetQueryParameters struct {
    // Filter by customer reference id
    Customer_reference_id *string `uriparametername:"customer_reference_id"`
    // Index for paginated results
    Index *int32 `uriparametername:"index"`
    // Max number of results to return
    Size *int32 `uriparametername:"size"`
}
// CardsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CardsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CardsRequestBuilderGetQueryParameters
}
// CardsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CardsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HitchPin/go-orum-client.deliver.cards.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CardsCardsItemRequestBuilder when successful
func (m *CardsRequestBuilder) ById(id string)(*CardsCardsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewCardsCardsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HitchPin/go-orum-client.deliver.cards.item collection
// returns a *CardsCardsItemRequestBuilder when successful
func (m *CardsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CardsCardsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewCardsCardsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCardsRequestBuilderInternal instantiates a new CardsRequestBuilder and sets the default values.
func NewCardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CardsRequestBuilder) {
    m := &CardsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/cards{?customer_reference_id*,index*,size*}", pathParameters),
    }
    return m
}
// NewCardsRequestBuilder instantiates a new CardsRequestBuilder and sets the default values.
func NewCardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CardsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCardsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all cards by filter.
// returns a CardsResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
func (m *CardsRequestBuilder) Get(ctx context.Context, requestConfiguration *CardsRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CardsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateCardsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CardsResponseable), nil
}
// Post create a card object.
// returns a CardResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
func (m *CardsRequestBuilder) Post(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CardRequestable, requestConfiguration *CardsRequestBuilderPostRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CardResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateCardResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CardResponseable), nil
}
// ToGetRequestInformation get all cards by filter.
// returns a *RequestInformation when successful
func (m *CardsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CardsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a card object.
// returns a *RequestInformation when successful
func (m *CardsRequestBuilder) ToPostRequestInformation(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CardRequestable, requestConfiguration *CardsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CardsRequestBuilder when successful
func (m *CardsRequestBuilder) WithUrl(rawUrl string)(*CardsRequestBuilder) {
    return NewCardsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
