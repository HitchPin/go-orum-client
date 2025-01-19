package webhooks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// SecretRequestBuilder builds and executes requests for operations under \webhooks\secret
type SecretRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecretRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecretRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecretRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecretRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecretRequestBuilderInternal instantiates a new SecretRequestBuilder and sets the default values.
func NewSecretRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecretRequestBuilder) {
    m := &SecretRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/webhooks/secret", pathParameters),
    }
    return m
}
// NewSecretRequestBuilder instantiates a new SecretRequestBuilder and sets the default values.
func NewSecretRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecretRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecretRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all webhook configurations
// returns a GetKeypairResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
func (m *SecretRequestBuilder) Get(ctx context.Context, requestConfiguration *SecretRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.GetKeypairResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateGetKeypairResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.GetKeypairResponseable), nil
}
// Post initialize secure webhook keypair
// returns a InitializeKeypairResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 403 status code
func (m *SecretRequestBuilder) Post(ctx context.Context, requestConfiguration *SecretRequestBuilderPostRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.InitializeKeypairResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "403": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateInitializeKeypairResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.InitializeKeypairResponseable), nil
}
// ToGetRequestInformation list all webhook configurations
// returns a *RequestInformation when successful
func (m *SecretRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecretRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation initialize secure webhook keypair
// returns a *RequestInformation when successful
func (m *SecretRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SecretRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SecretRequestBuilder when successful
func (m *SecretRequestBuilder) WithUrl(rawUrl string)(*SecretRequestBuilder) {
    return NewSecretRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
