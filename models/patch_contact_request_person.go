package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PatchContactRequestPerson contact information.
type PatchContactRequestPerson struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *ContactTypePerson
    // Email or phone number.
    value *string
}
// NewPatchContactRequestPerson instantiates a new PatchContactRequestPerson and sets the default values.
func NewPatchContactRequestPerson()(*PatchContactRequestPerson) {
    m := &PatchContactRequestPerson{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatchContactRequestPersonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePatchContactRequestPersonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatchContactRequestPerson(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PatchContactRequestPerson) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PatchContactRequestPerson) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContactTypePerson)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*ContactTypePerson))
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
// returns a *ContactTypePerson when successful
func (m *PatchContactRequestPerson) GetTypeEscaped()(*ContactTypePerson) {
    return m.typeEscaped
}
// GetValue gets the value property value. Email or phone number.
// returns a *string when successful
func (m *PatchContactRequestPerson) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *PatchContactRequestPerson) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PatchContactRequestPerson) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *PatchContactRequestPerson) SetTypeEscaped(value *ContactTypePerson)() {
    m.typeEscaped = value
}
// SetValue sets the value property value. Email or phone number.
func (m *PatchContactRequestPerson) SetValue(value *string)() {
    m.value = value
}
type PatchContactRequestPersonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTypeEscaped()(*ContactTypePerson)
    GetValue()(*string)
    SetTypeEscaped(value *ContactTypePerson)()
    SetValue(value *string)()
}
