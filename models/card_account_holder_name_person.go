package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CardAccountHolderNamePerson name of account holder.
type CardAccountHolderNamePerson struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The first_name property
    first_name *string
    // The last_name property
    last_name *string
    // The middle_name property
    middle_name *string
    // Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
    suffix *string
}
// NewCardAccountHolderNamePerson instantiates a new CardAccountHolderNamePerson and sets the default values.
func NewCardAccountHolderNamePerson()(*CardAccountHolderNamePerson) {
    m := &CardAccountHolderNamePerson{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCardAccountHolderNamePersonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCardAccountHolderNamePersonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCardAccountHolderNamePerson(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CardAccountHolderNamePerson) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CardAccountHolderNamePerson) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["middle_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["suffix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuffix(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *CardAccountHolderNamePerson) GetFirstName()(*string) {
    return m.first_name
}
// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *CardAccountHolderNamePerson) GetLastName()(*string) {
    return m.last_name
}
// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *CardAccountHolderNamePerson) GetMiddleName()(*string) {
    return m.middle_name
}
// GetSuffix gets the suffix property value. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
// returns a *string when successful
func (m *CardAccountHolderNamePerson) GetSuffix()(*string) {
    return m.suffix
}
// Serialize serializes information the current object
func (m *CardAccountHolderNamePerson) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("first_name", m.GetFirstName())
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
    {
        err := writer.WriteStringValue("middle_name", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("suffix", m.GetSuffix())
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
func (m *CardAccountHolderNamePerson) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFirstName sets the first_name property value. The first_name property
func (m *CardAccountHolderNamePerson) SetFirstName(value *string)() {
    m.first_name = value
}
// SetLastName sets the last_name property value. The last_name property
func (m *CardAccountHolderNamePerson) SetLastName(value *string)() {
    m.last_name = value
}
// SetMiddleName sets the middle_name property value. The middle_name property
func (m *CardAccountHolderNamePerson) SetMiddleName(value *string)() {
    m.middle_name = value
}
// SetSuffix sets the suffix property value. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
func (m *CardAccountHolderNamePerson) SetSuffix(value *string)() {
    m.suffix = value
}
type CardAccountHolderNamePersonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirstName()(*string)
    GetLastName()(*string)
    GetMiddleName()(*string)
    GetSuffix()(*string)
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetMiddleName(value *string)()
    SetSuffix(value *string)()
}
