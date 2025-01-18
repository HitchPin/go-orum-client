package deliver

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// TransferGroupsTransferGroupsItemRequestBuilder builds and executes requests for operations under \deliver\transfer-groups\{id}
type TransferGroupsTransferGroupsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TransferGroupsTransferGroupsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransferGroupsTransferGroupsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTransferGroupsTransferGroupsItemRequestBuilderInternal instantiates a new TransferGroupsTransferGroupsItemRequestBuilder and sets the default values.
func NewTransferGroupsTransferGroupsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransferGroupsTransferGroupsItemRequestBuilder) {
    m := &TransferGroupsTransferGroupsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/transfer-groups/{id}", pathParameters),
    }
    return m
}
// NewTransferGroupsTransferGroupsItemRequestBuilder instantiates a new TransferGroupsTransferGroupsItemRequestBuilder and sets the default values.
func NewTransferGroupsTransferGroupsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransferGroupsTransferGroupsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransferGroupsTransferGroupsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a Transfer Group object by ID
// returns a TransferGroupResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 404 status code
func (m *TransferGroupsTransferGroupsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TransferGroupsTransferGroupsItemRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.TransferGroupResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "404": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateTransferGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.TransferGroupResponseable), nil
}
// ToGetRequestInformation get a Transfer Group object by ID
// returns a *RequestInformation when successful
func (m *TransferGroupsTransferGroupsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TransferGroupsTransferGroupsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TransferGroupsTransferGroupsItemRequestBuilder when successful
func (m *TransferGroupsTransferGroupsItemRequestBuilder) WithUrl(rawUrl string)(*TransferGroupsTransferGroupsItemRequestBuilder) {
    return NewTransferGroupsTransferGroupsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
