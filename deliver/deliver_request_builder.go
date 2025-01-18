package deliver

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeliverRequestBuilder builds and executes requests for operations under \deliver
type DeliverRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Balance the balance property
// returns a *BalanceRequestBuilder when successful
func (m *DeliverRequestBuilder) Balance()(*BalanceRequestBuilder) {
    return NewBalanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Businesses the businesses property
// returns a *BusinessesRequestBuilder when successful
func (m *DeliverRequestBuilder) Businesses()(*BusinessesRequestBuilder) {
    return NewBusinessesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cards the cards property
// returns a *CardsRequestBuilder when successful
func (m *DeliverRequestBuilder) Cards()(*CardsRequestBuilder) {
    return NewCardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeliverRequestBuilderInternal instantiates a new DeliverRequestBuilder and sets the default values.
func NewDeliverRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeliverRequestBuilder) {
    m := &DeliverRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deliver", pathParameters),
    }
    return m
}
// NewDeliverRequestBuilder instantiates a new DeliverRequestBuilder and sets the default values.
func NewDeliverRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeliverRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeliverRequestBuilderInternal(urlParams, requestAdapter)
}
// Eligibility the eligibility property
// returns a *EligibilityRequestBuilder when successful
func (m *DeliverRequestBuilder) Eligibility()(*EligibilityRequestBuilder) {
    return NewEligibilityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// External the external property
// returns a *ExternalRequestBuilder when successful
func (m *DeliverRequestBuilder) External()(*ExternalRequestBuilder) {
    return NewExternalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Persons the persons property
// returns a *PersonsRequestBuilder when successful
func (m *DeliverRequestBuilder) Persons()(*PersonsRequestBuilder) {
    return NewPersonsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reports the reports property
// returns a *ReportsRequestBuilder when successful
func (m *DeliverRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TransferGroups the transferGroups property
// returns a *TransferGroupsRequestBuilder when successful
func (m *DeliverRequestBuilder) TransferGroups()(*TransferGroupsRequestBuilder) {
    return NewTransferGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transfers the transfers property
// returns a *TransfersRequestBuilder when successful
func (m *DeliverRequestBuilder) Transfers()(*TransfersRequestBuilder) {
    return NewTransfersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
