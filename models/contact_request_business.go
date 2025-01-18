package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContactRequestBusiness contact information.
type ContactRequestBusiness struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *ContactTypeBusiness
    // Email, phone number or website.
    value *string
}
// NewContactRequestBusiness instantiates a new ContactRequestBusiness and sets the default values.
func NewContactRequestBusiness()(*ContactRequestBusiness) {
    m := &ContactRequestBusiness{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContactRequestBusinessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactRequestBusinessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContactRequestBusiness(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactRequestBusiness) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactRequestBusiness) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContactTypeBusiness)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*ContactTypeBusiness))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
// returns a *ContactTypeBusiness when successful
func (m *ContactRequestBusiness) GetTypeEscaped()(*ContactTypeBusiness) {
    return m.typeEscaped
}
// GetValue gets the value property value. Email, phone number or website.
// returns a *string when successful
func (m *ContactRequestBusiness) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ContactRequestBusiness) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ContactRequestBusiness) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ContactRequestBusiness) SetTypeEscaped(value *ContactTypeBusiness)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. Email, phone number or website.
func (m *ContactRequestBusiness) SetValue(value *string)() {
    m.value = value
}
type ContactRequestBusinessable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*ContactTypeBusiness)
    GetValue()(*string)
    SetTypeEscaped(value *ContactTypeBusiness)()
    SetValue(value *string)()
}
