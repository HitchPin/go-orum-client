package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WebhookConfigurationSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *string
    // The created_by property
    created_by *string
    // The data_version property
    data_version *string
    // The enabled property
    enabled *bool
    // The enterprise_name property
    enterprise_name *string
    // The event_types property
    event_types []WebhookEventType
    // The id property
    id *string
    // The updated_at property
    updated_at *string
    // The updated_by property
    updated_by *string
    // The url property
    url *string
}
// NewWebhookConfigurationSummary instantiates a new WebhookConfigurationSummary and sets the default values.
func NewWebhookConfigurationSummary()(*WebhookConfigurationSummary) {
    m := &WebhookConfigurationSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebhookConfigurationSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebhookConfigurationSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebhookConfigurationSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WebhookConfigurationSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetCreatedAt()(*string) {
    return m.created_at
}
// GetCreatedBy gets the created_by property value. The created_by property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetCreatedBy()(*string) {
    return m.created_by
}
// GetDataVersion gets the data_version property value. The data_version property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetDataVersion()(*string) {
    return m.data_version
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *WebhookConfigurationSummary) GetEnabled()(*bool) {
    return m.enabled
}
// GetEnterpriseName gets the enterprise_name property value. The enterprise_name property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetEnterpriseName()(*string) {
    return m.enterprise_name
}
// GetEventTypes gets the event_types property value. The event_types property
// returns a []WebhookEventType when successful
func (m *WebhookConfigurationSummary) GetEventTypes()([]WebhookEventType) {
    return m.event_types
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebhookConfigurationSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["created_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["data_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataVersion(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["enterprise_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseName(val)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["updated_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedBy(val)
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
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetId()(*string) {
    return m.id
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetUpdatedAt()(*string) {
    return m.updated_at
}
// GetUpdatedBy gets the updated_by property value. The updated_by property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetUpdatedBy()(*string) {
    return m.updated_by
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *WebhookConfigurationSummary) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *WebhookConfigurationSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_by", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("data_version", m.GetDataVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enterprise_name", m.GetEnterpriseName())
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
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_by", m.GetUpdatedBy())
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
func (m *WebhookConfigurationSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *WebhookConfigurationSummary) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetCreatedBy sets the created_by property value. The created_by property
func (m *WebhookConfigurationSummary) SetCreatedBy(value *string)() {
    m.created_by = value
}
// SetDataVersion sets the data_version property value. The data_version property
func (m *WebhookConfigurationSummary) SetDataVersion(value *string)() {
    m.data_version = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *WebhookConfigurationSummary) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetEnterpriseName sets the enterprise_name property value. The enterprise_name property
func (m *WebhookConfigurationSummary) SetEnterpriseName(value *string)() {
    m.enterprise_name = value
}
// SetEventTypes sets the event_types property value. The event_types property
func (m *WebhookConfigurationSummary) SetEventTypes(value []WebhookEventType)() {
    m.event_types = value
}
// SetId sets the id property value. The id property
func (m *WebhookConfigurationSummary) SetId(value *string)() {
    m.id = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *WebhookConfigurationSummary) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
// SetUpdatedBy sets the updated_by property value. The updated_by property
func (m *WebhookConfigurationSummary) SetUpdatedBy(value *string)() {
    m.updated_by = value
}
// SetUrl sets the url property value. The url property
func (m *WebhookConfigurationSummary) SetUrl(value *string)() {
    m.url = value
}
type WebhookConfigurationSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetCreatedBy()(*string)
    GetDataVersion()(*string)
    GetEnabled()(*bool)
    GetEnterpriseName()(*string)
    GetEventTypes()([]WebhookEventType)
    GetId()(*string)
    GetUpdatedAt()(*string)
    GetUpdatedBy()(*string)
    GetUrl()(*string)
    SetCreatedAt(value *string)()
    SetCreatedBy(value *string)()
    SetDataVersion(value *string)()
    SetEnabled(value *bool)()
    SetEnterpriseName(value *string)()
    SetEventTypes(value []WebhookEventType)()
    SetId(value *string)()
    SetUpdatedAt(value *string)()
    SetUpdatedBy(value *string)()
    SetUrl(value *string)()
}
