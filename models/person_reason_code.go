package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PersonReasonCode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A word-based code that describes the verification failure.
    reason_code *string
    // A human-readable description of the reason code.
    reason_code_message *string
}
// NewPersonReasonCode instantiates a new PersonReasonCode and sets the default values.
func NewPersonReasonCode()(*PersonReasonCode) {
    m := &PersonReasonCode{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePersonReasonCodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePersonReasonCodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonReasonCode(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PersonReasonCode) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PersonReasonCode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reason_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReasonCode(val)
        }
        return nil
    }
    res["reason_code_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReasonCodeMessage(val)
        }
        return nil
    }
    return res
}
// GetReasonCode gets the reason_code property value. A word-based code that describes the verification failure.
// returns a *string when successful
func (m *PersonReasonCode) GetReasonCode()(*string) {
    return m.reason_code
}
// GetReasonCodeMessage gets the reason_code_message property value. A human-readable description of the reason code.
// returns a *string when successful
func (m *PersonReasonCode) GetReasonCodeMessage()(*string) {
    return m.reason_code_message
}
// Serialize serializes information the current object
func (m *PersonReasonCode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reason_code", m.GetReasonCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason_code_message", m.GetReasonCodeMessage())
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
func (m *PersonReasonCode) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReasonCode sets the reason_code property value. A word-based code that describes the verification failure.
func (m *PersonReasonCode) SetReasonCode(value *string)() {
    m.reason_code = value
}
// SetReasonCodeMessage sets the reason_code_message property value. A human-readable description of the reason code.
func (m *PersonReasonCode) SetReasonCodeMessage(value *string)() {
    m.reason_code_message = value
}
type PersonReasonCodeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReasonCode()(*string)
    GetReasonCodeMessage()(*string)
    SetReasonCode(value *string)()
    SetReasonCodeMessage(value *string)()
}
