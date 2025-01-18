package deliver

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73 "github.com/HitchPin/go-orum-client/models"
)

// PersonsRequestBuilder builds and executes requests for operations under \deliver\persons
type PersonsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PersonsRequestBuilderGetQueryParameters get all persons objects.
type PersonsRequestBuilderGetQueryParameters struct {
    // Filter by account number
    Account_number *string `uriparametername:"account_number"`
    // Filter by customer reference id
    Customer_reference_id *string `uriparametername:"customer_reference_id"`
    // Filter results created before this time
    End_time *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"end_time"`
    // The first name to filter persons by (case-insensitive, exact match).
    First_name *string `uriparametername:"first_name"`
    // Filter results by the person id
    Id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID `uriparametername:"id"`
    // Index for paginated results
    Index *int32 `uriparametername:"index"`
    // The last name to filter persons by (case-insensitive, exact match).
    Last_name *string `uriparametername:"last_name"`
    // Max number of results to return
    Size *int32 `uriparametername:"size"`
    // Filter results created at or after this time
    Start_time *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"start_time"`
    // Filter by customer statuses (multiple allowed)
    // Deprecated: This property is deprecated, use StatusAsPersonStatus instead
    Status []string `uriparametername:"status"`
    // Filter by customer statuses (multiple allowed)
    StatusAsPersonStatus []if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PersonStatus `uriparametername:"status"`
}
// PersonsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PersonsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PersonsRequestBuilderGetQueryParameters
}
// PersonsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PersonsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PersonsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PersonsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the github.com/HitchPin/go-orum-client.deliver.persons.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *PersonsPersonsItemRequestBuilder when successful
func (m *PersonsRequestBuilder) ById(id string)(*PersonsPersonsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewPersonsPersonsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdGuid gets an item from the github.com/HitchPin/go-orum-client.deliver.persons.item collection
// returns a *PersonsPersonsItemRequestBuilder when successful
func (m *PersonsRequestBuilder) ByIdGuid(id i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*PersonsPersonsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = id.String()
    return NewPersonsPersonsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPersonsRequestBuilderInternal instantiates a new PersonsRequestBuilder and sets the default values.
func NewPersonsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PersonsRequestBuilder) {
    m := &PersonsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver/persons{?account_number*,customer_reference_id*,end_time*,first_name*,id*,index*,last_name*,size*,start_time*,status*}", pathParameters),
    }
    return m
}
// NewPersonsRequestBuilder instantiates a new PersonsRequestBuilder and sets the default values.
func NewPersonsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PersonsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPersonsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all persons objects.
// returns a PersonsResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
func (m *PersonsRequestBuilder) Get(ctx context.Context, requestConfiguration *PersonsRequestBuilderGetRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PersonsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreatePersonsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PersonsResponseable), nil
}
// Post create a person object to represent a human customer.
// returns a PersonResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
func (m *PersonsRequestBuilder) Post(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PersonRequestable, requestConfiguration *PersonsRequestBuilderPostRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PersonResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreatePersonResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PersonResponseable), nil
}
// Put update a person object using its reference id.
// returns a PutPersonResponseable when successful
// returns a ErrorResponse error when the service returns a 400 status code
func (m *PersonsRequestBuilder) Put(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PutPersonRequestable, requestConfiguration *PersonsRequestBuilderPutRequestConfiguration)(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PutPersonResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreateErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.CreatePutPersonResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PutPersonResponseable), nil
}
// ToGetRequestInformation get all persons objects.
// returns a *RequestInformation when successful
func (m *PersonsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PersonsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a person object to represent a human customer.
// returns a *RequestInformation when successful
func (m *PersonsRequestBuilder) ToPostRequestInformation(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PersonRequestable, requestConfiguration *PersonsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update a person object using its reference id.
// returns a *RequestInformation when successful
func (m *PersonsRequestBuilder) ToPutRequestInformation(ctx context.Context, body if53ec6ffd026403c86e62c61aa37ac9a8820505adfed180e5be26eec36ab1c73.PutPersonRequestable, requestConfiguration *PersonsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *PersonsRequestBuilder when successful
func (m *PersonsRequestBuilder) WithUrl(rawUrl string)(*PersonsRequestBuilder) {
    return NewPersonsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
