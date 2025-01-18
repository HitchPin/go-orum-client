package oauth

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// TokenRequestBuilder builds and executes requests for operations under \oauth\token
type TokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTokenRequestBuilderInternal instantiates a new TokenRequestBuilder and sets the default values.
func NewTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenRequestBuilder) {
    m := &TokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/oauth/token", pathParameters),
    }
    return m
}
// NewTokenRequestBuilder instantiates a new TokenRequestBuilder and sets the default values.
func NewTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post <a href="https://orum.io/contact/" target="_blank">Contact Orum</a> for access to an Orum Portal account, where you can set up your login credentials.<br /><br />Use this <code>/oauth/token</code> endpoint to exchange API credentials created in the Portal (client_id and client_secret) for an <code>access_token</code>.<br /><br /> To authenticate to all protected Orum endpoints, pass the <code>access_token</code> in the Authorization header <code>Bearer ${token}</code>.<br />The auth token is valid for one hour.
// returns a OauthTokenResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 401 status code
// returns a ErrorResponse error when the service returns a 403 status code
// returns a ErrorResponse error when the service returns a 429 status code
// returns a ErrorResponse error when the service returns a 503 status code
func (m *TokenRequestBuilder) Post(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.OauthTokenBodyable, requestConfiguration *TokenRequestBuilderPostRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.OauthTokenResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "401": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "403": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "429": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "503": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateOauthTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.OauthTokenResponseable), nil
}
// ToPostRequestInformation <a href="https://orum.io/contact/" target="_blank">Contact Orum</a> for access to an Orum Portal account, where you can set up your login credentials.<br /><br />Use this <code>/oauth/token</code> endpoint to exchange API credentials created in the Portal (client_id and client_secret) for an <code>access_token</code>.<br /><br /> To authenticate to all protected Orum endpoints, pass the <code>access_token</code> in the Authorization header <code>Bearer ${token}</code>.<br />The auth token is valid for one hour.
// returns a *RequestInformation when successful
func (m *TokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.OauthTokenBodyable, requestConfiguration *TokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TokenRequestBuilder when successful
func (m *TokenRequestBuilder) WithUrl(rawUrl string)(*TokenRequestBuilder) {
    return NewTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
