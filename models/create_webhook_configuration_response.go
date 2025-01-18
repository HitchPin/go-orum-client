package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateWebhookConfigurationResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The webhook_configuration property
    webhook_configuration WebhookConfigurationSummaryable
}
// NewCreateWebhookConfigurationResponse instantiates a new CreateWebhookConfigurationResponse and sets the default values.
func NewCreateWebhookConfigurationResponse()(*CreateWebhookConfigurationResponse) {
    m := &CreateWebhookConfigurationResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreateWebhookConfigurationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateWebhookConfigurationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateWebhookConfigurationResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateWebhookConfigurationResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateWebhookConfigurationResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["webhook_configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookConfigurationSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebhookConfiguration(val.(WebhookConfigurationSummaryable))
        }
        return nil
    }
    return res
}
// GetWebhookConfiguration gets the webhook_configuration property value. The webhook_configuration property
// returns a WebhookConfigurationSummaryable when successful
func (m *CreateWebhookConfigurationResponse) GetWebhookConfiguration()(WebhookConfigurationSummaryable) {
    return m.webhook_configuration
}
// Serialize serializes information the current object
func (m *CreateWebhookConfigurationResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("webhook_configuration", m.GetWebhookConfiguration())
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
func (m *CreateWebhookConfigurationResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWebhookConfiguration sets the webhook_configuration property value. The webhook_configuration property
func (m *CreateWebhookConfigurationResponse) SetWebhookConfiguration(value WebhookConfigurationSummaryable)() {
    m.webhook_configuration = value
}
type CreateWebhookConfigurationResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetWebhookConfiguration()(WebhookConfigurationSummaryable)
    SetWebhookConfiguration(value WebhookConfigurationSummaryable)()
}
