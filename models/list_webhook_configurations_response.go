package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ListWebhookConfigurationsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The webhook_configurations property
    webhook_configurations []WebhookConfigurationSummaryable
}
// NewListWebhookConfigurationsResponse instantiates a new ListWebhookConfigurationsResponse and sets the default values.
func NewListWebhookConfigurationsResponse()(*ListWebhookConfigurationsResponse) {
    m := &ListWebhookConfigurationsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateListWebhookConfigurationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateListWebhookConfigurationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListWebhookConfigurationsResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ListWebhookConfigurationsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ListWebhookConfigurationsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["webhook_configurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebhookConfigurationSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebhookConfigurationSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WebhookConfigurationSummaryable)
                }
            }
            m.SetWebhookConfigurations(res)
        }
        return nil
    }
    return res
}
// GetWebhookConfigurations gets the webhook_configurations property value. The webhook_configurations property
// returns a []WebhookConfigurationSummaryable when successful
func (m *ListWebhookConfigurationsResponse) GetWebhookConfigurations()([]WebhookConfigurationSummaryable) {
    return m.webhook_configurations
}
// Serialize serializes information the current object
func (m *ListWebhookConfigurationsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetWebhookConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebhookConfigurations()))
        for i, v := range m.GetWebhookConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("webhook_configurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ListWebhookConfigurationsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWebhookConfigurations sets the webhook_configurations property value. The webhook_configurations property
func (m *ListWebhookConfigurationsResponse) SetWebhookConfigurations(value []WebhookConfigurationSummaryable)() {
    m.webhook_configurations = value
}
type ListWebhookConfigurationsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetWebhookConfigurations()([]WebhookConfigurationSummaryable)
    SetWebhookConfigurations(value []WebhookConfigurationSummaryable)()
}
