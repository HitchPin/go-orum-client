package deliver

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// ReportsBalancesRequestBuilder builds and executes requests for operations under \deliver\reports\balances
type ReportsBalancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsBalancesRequestBuilderGetQueryParameters get a balance statement summary report for the requested start and end date.
type ReportsBalancesRequestBuilderGetQueryParameters struct {
    End_time *string `uriparametername:"end_time"`
    Start_time *string `uriparametername:"start_time"`
}
// ReportsBalancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsBalancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsBalancesRequestBuilderGetQueryParameters
}
// NewReportsBalancesRequestBuilderInternal instantiates a new ReportsBalancesRequestBuilder and sets the default values.
func NewReportsBalancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsBalancesRequestBuilder) {
    m := &ReportsBalancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/reports/balances?end_time={end_time}&start_time={start_time}", pathParameters),
    }
    return m
}
// NewReportsBalancesRequestBuilder instantiates a new ReportsBalancesRequestBuilder and sets the default values.
func NewReportsBalancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsBalancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsBalancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a balance statement summary report for the requested start and end date.
// returns a BalanceReportResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
// returns a ErrorResponse error when the service returns a 401 status code
// returns a ErrorResponse error when the service returns a 403 status code
func (m *ReportsBalancesRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsBalancesRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.BalanceReportResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "401": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
        "403": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateBalanceReportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.BalanceReportResponseable), nil
}
// ToGetRequestInformation get a balance statement summary report for the requested start and end date.
// returns a *RequestInformation when successful
func (m *ReportsBalancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsBalancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsBalancesRequestBuilder when successful
func (m *ReportsBalancesRequestBuilder) WithUrl(rawUrl string)(*ReportsBalancesRequestBuilder) {
    return NewReportsBalancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
