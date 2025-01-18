package deliver

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// TransfersItemCancelRequestBuilder builds and executes requests for operations under \deliver\transfers\{id}\cancel
type TransfersItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TransfersItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransfersItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTransfersItemCancelRequestBuilderInternal instantiates a new TransfersItemCancelRequestBuilder and sets the default values.
func NewTransfersItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransfersItemCancelRequestBuilder) {
    m := &TransfersItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/transfers/{id}/cancel", pathParameters),
    }
    return m
}
// NewTransfersItemCancelRequestBuilder instantiates a new TransfersItemCancelRequestBuilder and sets the default values.
func NewTransfersItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransfersItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransfersItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiates a transfer cancellation request if the transfer is eligible.
// returns a TransferCancelResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 404 status code
func (m *TransfersItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *TransfersItemCancelRequestBuilderPostRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.TransferCancelResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "404": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateTransferCancelResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.TransferCancelResponseable), nil
}
// ToPostRequestInformation initiates a transfer cancellation request if the transfer is eligible.
// returns a *RequestInformation when successful
func (m *TransfersItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TransfersItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TransfersItemCancelRequestBuilder when successful
func (m *TransfersItemCancelRequestBuilder) WithUrl(rawUrl string)(*TransfersItemCancelRequestBuilder) {
    return NewTransfersItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
