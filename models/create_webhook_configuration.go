package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateWebhookConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enabled property
    enabled *CreateWebhookConfiguration_enabled
    // The event_types property
    event_types []WebhookEventType
    // The url property
    url *string
}
// NewCreateWebhookConfiguration instantiates a new CreateWebhookConfiguration and sets the default values.
func NewCreateWebhookConfiguration()(*CreateWebhookConfiguration) {
    m := &CreateWebhookConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreateWebhookConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateWebhookConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateWebhookConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateWebhookConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *CreateWebhookConfiguration_enabled when successful
func (m *CreateWebhookConfiguration) GetEnabled()(*CreateWebhookConfiguration_enabled) {
    return m.enabled
}
// GetEventTypes gets the event_types property value. The event_types property
// returns a []WebhookEventType when successful
func (m *CreateWebhookConfiguration) GetEventTypes()([]WebhookEventType) {
    return m.event_types
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateWebhookConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCreateWebhookConfiguration_enabled)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val.(*CreateWebhookConfiguration_enabled))
        }
        return nil
    }
    res["event_types"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseWebhookEventType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebhookEventType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*WebhookEventType))
                }
            }
            m.SetEventTypes(res)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *CreateWebhookConfiguration) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CreateWebhookConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnabled() != nil {
        cast := (*m.GetEnabled()).String()
        err := writer.WriteStringValue("enabled", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEventTypes() != nil {
        err := writer.WriteCollectionOfStringValues("event_types", SerializeWebhookEventType(m.GetEventTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *CreateWebhookConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *CreateWebhookConfiguration) SetEnabled(value *CreateWebhookConfiguration_enabled)() {
    m.enabled = value
}
// SetEventTypes sets the event_types property value. The event_types property
func (m *CreateWebhookConfiguration) SetEventTypes(value []WebhookEventType)() {
    m.event_types = value
}
// SetUrl sets the url property value. The url property
func (m *CreateWebhookConfiguration) SetUrl(value *string)() {
    m.url = value
}
type CreateWebhookConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*CreateWebhookConfiguration_enabled)
    GetEventTypes()([]WebhookEventType)
    GetUrl()(*string)
    SetEnabled(value *CreateWebhookConfiguration_enabled)()
    SetEventTypes(value []WebhookEventType)()
    SetUrl(value *string)()
}
