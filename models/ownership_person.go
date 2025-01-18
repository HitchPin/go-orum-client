package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OwnershipPerson ownership details for a person
type OwnershipPerson struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The first_name property
    first_name *string
    // Result of Ownership match
    first_name_match *OwnershipResult
    // The last_name property
    last_name *string
    // Result of Ownership match
    last_name_match *OwnershipResult
}
// NewOwnershipPerson instantiates a new OwnershipPerson and sets the default values.
func NewOwnershipPerson()(*OwnershipPerson) {
    m := &OwnershipPerson{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOwnershipPersonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOwnershipPersonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOwnershipPerson(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OwnershipPerson) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OwnershipPerson) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["first_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["first_name_match"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnershipResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstNameMatch(val.(*OwnershipResult))
        }
        return nil
    }
    res["last_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["last_name_match"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnershipResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastNameMatch(val.(*OwnershipResult))
        }
        return nil
    }
    return res
}
// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *OwnershipPerson) GetFirstName()(*string) {
    return m.first_name
}
// GetFirstNameMatch gets the first_name_match property value. Result of Ownership match
// returns a *OwnershipResult when successful
func (m *OwnershipPerson) GetFirstNameMatch()(*OwnershipResult) {
    return m.first_name_match
}
// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *OwnershipPerson) GetLastName()(*string) {
    return m.last_name
}
// GetLastNameMatch gets the last_name_match property value. Result of Ownership match
// returns a *OwnershipResult when successful
func (m *OwnershipPerson) GetLastNameMatch()(*OwnershipResult) {
    return m.last_name_match
}
// Serialize serializes information the current object
func (m *OwnershipPerson) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("first_name", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    if m.GetFirstNameMatch() != nil {
        cast := (*m.GetFirstNameMatch()).String()
        err := writer.WriteStringValue("first_name_match", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_name", m.GetLastName())
        if err != nil {
            return err
        }
    }
    if m.GetLastNameMatch() != nil {
        cast := (*m.GetLastNameMatch()).String()
        err := writer.WriteStringValue("last_name_match", &cast)
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
func (m *OwnershipPerson) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFirstName sets the first_name property value. The first_name property
func (m *OwnershipPerson) SetFirstName(value *string)() {
    m.first_name = value
}
// SetFirstNameMatch sets the first_name_match property value. Result of Ownership match
func (m *OwnershipPerson) SetFirstNameMatch(value *OwnershipResult)() {
    m.first_name_match = value
}
// SetLastName sets the last_name property value. The last_name property
func (m *OwnershipPerson) SetLastName(value *string)() {
    m.last_name = value
}
// SetLastNameMatch sets the last_name_match property value. Result of Ownership match
func (m *OwnershipPerson) SetLastNameMatch(value *OwnershipResult)() {
    m.last_name_match = value
}
type OwnershipPersonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirstName()(*string)
    GetFirstNameMatch()(*OwnershipResult)
    GetLastName()(*string)
    GetLastNameMatch()(*OwnershipResult)
    SetFirstName(value *string)()
    SetFirstNameMatch(value *OwnershipResult)()
    SetLastName(value *string)()
    SetLastNameMatch(value *OwnershipResult)()
}
