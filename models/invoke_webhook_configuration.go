package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InvokeWebhookConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data property
    data InvokeWebhookConfiguration_dataable
    // The event_type property
    event_type *WebhookEventType
}
// NewInvokeWebhookConfiguration instantiates a new InvokeWebhookConfiguration and sets the default values.
func NewInvokeWebhookConfiguration()(*InvokeWebhookConfiguration) {
    m := &InvokeWebhookConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInvokeWebhookConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInvokeWebhookConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvokeWebhookConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InvokeWebhookConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The data property
// returns a InvokeWebhookConfiguration_dataable when successful
func (m *InvokeWebhookConfiguration) GetData()(InvokeWebhookConfiguration_dataable) {
    return m.data
}
// GetEventType gets the event_type property value. The event_type property
// returns a *WebhookEventType when successful
func (m *InvokeWebhookConfiguration) GetEventType()(*WebhookEventType) {
    return m.event_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InvokeWebhookConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInvokeWebhookConfiguration_dataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(InvokeWebhookConfiguration_dataable))
        }
        return nil
    }
    res["event_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWebhookEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventType(val.(*WebhookEventType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *InvokeWebhookConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    if m.GetEventType() != nil {
        cast := (*m.GetEventType()).String()
        err := writer.WriteStringValue("event_type", &cast)
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
func (m *InvokeWebhookConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *InvokeWebhookConfiguration) SetData(value InvokeWebhookConfiguration_dataable)() {
    m.data = value
}
// SetEventType sets the event_type property value. The event_type property
func (m *InvokeWebhookConfiguration) SetEventType(value *WebhookEventType)() {
    m.event_type = value
}
type InvokeWebhookConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()(InvokeWebhookConfiguration_dataable)
    GetEventType()(*WebhookEventType)
    SetData(value InvokeWebhookConfiguration_dataable)()
    SetEventType(value *WebhookEventType)()
}
