package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PatchContactRequestBusiness contact information.
type PatchContactRequestBusiness struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Type of contact information associated with a business - 'email', 'phone' or 'website'.
    typeEscaped *ContactTypeBusiness
    // Email, phone number or website.
    value *string
}
// NewPatchContactRequestBusiness instantiates a new PatchContactRequestBusiness and sets the default values.
func NewPatchContactRequestBusiness()(*PatchContactRequestBusiness) {
    m := &PatchContactRequestBusiness{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatchContactRequestBusinessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePatchContactRequestBusinessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatchContactRequestBusiness(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PatchContactRequestBusiness) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PatchContactRequestBusiness) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTypeEscaped gets the type property value. Type of contact information associated with a business - 'email', 'phone' or 'website'.
// returns a *ContactTypeBusiness when successful
func (m *PatchContactRequestBusiness) GetTypeEscaped()(*ContactTypeBusiness) {
    return m.typeEscaped
}
// GetValue gets the value property value. Email, phone number or website.
// returns a *string when successful
func (m *PatchContactRequestBusiness) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *PatchContactRequestBusiness) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PatchContactRequestBusiness) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. Type of contact information associated with a business - 'email', 'phone' or 'website'.
func (m *PatchContactRequestBusiness) SetTypeEscaped(value *ContactTypeBusiness)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. Email, phone number or website.
func (m *PatchContactRequestBusiness) SetValue(value *string)() {
    m.value = value
}
type PatchContactRequestBusinessable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*ContactTypeBusiness)
    GetValue()(*string)
    SetTypeEscaped(value *ContactTypeBusiness)()
    SetValue(value *string)()
}
