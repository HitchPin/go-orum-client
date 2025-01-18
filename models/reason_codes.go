package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ReasonCodes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reason code sent by the network (ex R01).
    network_reason_code *string
    // The message sent by the network.
    network_reason_code_message *string
    // The payment rail used in the transfer. Either RTP or ACH.
    network_reason_code_rail_type *string
    // A word-based code created by Orum to describe the reason for the failure. Orum reason codes are rail-agnostic.
    reason_code *string
    // A human-readable description of the reason code.
    reason_code_message *string
}
// NewReasonCodes instantiates a new ReasonCodes and sets the default values.
func NewReasonCodes()(*ReasonCodes) {
    m := &ReasonCodes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReasonCodesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReasonCodesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReasonCodes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReasonCodes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReasonCodes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["network_reason_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkReasonCode(val)
        }
        return nil
    }
    res["network_reason_code_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkReasonCodeMessage(val)
        }
        return nil
    }
    res["network_reason_code_rail_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkReasonCodeRailType(val)
        }
        return nil
    }
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
// GetNetworkReasonCode gets the network_reason_code property value. The reason code sent by the network (ex R01).
// returns a *string when successful
func (m *ReasonCodes) GetNetworkReasonCode()(*string) {
    return m.network_reason_code
}
// GetNetworkReasonCodeMessage gets the network_reason_code_message property value. The message sent by the network.
// returns a *string when successful
func (m *ReasonCodes) GetNetworkReasonCodeMessage()(*string) {
    return m.network_reason_code_message
}
// GetNetworkReasonCodeRailType gets the network_reason_code_rail_type property value. The payment rail used in the transfer. Either RTP or ACH.
// returns a *string when successful
func (m *ReasonCodes) GetNetworkReasonCodeRailType()(*string) {
    return m.network_reason_code_rail_type
}
// GetReasonCode gets the reason_code property value. A word-based code created by Orum to describe the reason for the failure. Orum reason codes are rail-agnostic.
// returns a *string when successful
func (m *ReasonCodes) GetReasonCode()(*string) {
    return m.reason_code
}
// GetReasonCodeMessage gets the reason_code_message property value. A human-readable description of the reason code.
// returns a *string when successful
func (m *ReasonCodes) GetReasonCodeMessage()(*string) {
    return m.reason_code_message
}
// Serialize serializes information the current object
func (m *ReasonCodes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("network_reason_code", m.GetNetworkReasonCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network_reason_code_message", m.GetNetworkReasonCodeMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network_reason_code_rail_type", m.GetNetworkReasonCodeRailType())
        if err != nil {
            return err
        }
    }
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
func (m *ReasonCodes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNetworkReasonCode sets the network_reason_code property value. The reason code sent by the network (ex R01).
func (m *ReasonCodes) SetNetworkReasonCode(value *string)() {
    m.network_reason_code = value
}
// SetNetworkReasonCodeMessage sets the network_reason_code_message property value. The message sent by the network.
func (m *ReasonCodes) SetNetworkReasonCodeMessage(value *string)() {
    m.network_reason_code_message = value
}
// SetNetworkReasonCodeRailType sets the network_reason_code_rail_type property value. The payment rail used in the transfer. Either RTP or ACH.
func (m *ReasonCodes) SetNetworkReasonCodeRailType(value *string)() {
    m.network_reason_code_rail_type = value
}
// SetReasonCode sets the reason_code property value. A word-based code created by Orum to describe the reason for the failure. Orum reason codes are rail-agnostic.
func (m *ReasonCodes) SetReasonCode(value *string)() {
    m.reason_code = value
}
// SetReasonCodeMessage sets the reason_code_message property value. A human-readable description of the reason code.
func (m *ReasonCodes) SetReasonCodeMessage(value *string)() {
    m.reason_code_message = value
}
type ReasonCodesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetworkReasonCode()(*string)
    GetNetworkReasonCodeMessage()(*string)
    GetNetworkReasonCodeRailType()(*string)
    GetReasonCode()(*string)
    GetReasonCodeMessage()(*string)
    SetNetworkReasonCode(value *string)()
    SetNetworkReasonCodeMessage(value *string)()
    SetNetworkReasonCodeRailType(value *string)()
    SetReasonCode(value *string)()
    SetReasonCodeMessage(value *string)()
}
