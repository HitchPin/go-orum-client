package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OwnershipPersonRequest ownership details for a person
type OwnershipPersonRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The first_name property
    first_name *string
    // The last_name property
    last_name *string
}
// NewOwnershipPersonRequest instantiates a new OwnershipPersonRequest and sets the default values.
func NewOwnershipPersonRequest()(*OwnershipPersonRequest) {
    m := &OwnershipPersonRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOwnershipPersonRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOwnershipPersonRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOwnershipPersonRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OwnershipPersonRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OwnershipPersonRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *OwnershipPersonRequest) GetFirstName()(*string) {
    return m.first_name
}
// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *OwnershipPersonRequest) GetLastName()(*string) {
    return m.last_name
}
// Serialize serializes information the current object
func (m *OwnershipPersonRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OwnershipPersonRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFirstName sets the first_name property value. The first_name property
func (m *OwnershipPersonRequest) SetFirstName(value *string)() {
    m.first_name = value
}
// SetLastName sets the last_name property value. The last_name property
func (m *OwnershipPersonRequest) SetLastName(value *string)() {
    m.last_name = value
}
type OwnershipPersonRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirstName()(*string)
    GetLastName()(*string)
    SetFirstName(value *string)()
    SetLastName(value *string)()
}
