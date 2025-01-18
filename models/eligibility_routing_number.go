package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EligibilityRoutingNumber struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Describes if the routing number is RTP or Fednow eligible or not
    eligible *bool
    // Describes if the routing number is Fednow eligible or not
    fednow *bool
    // 9-digit American Bankers Association (ABA) routing number.
    routing_number *string
    // Describes if the routing number is RTP eligible or not
    rtp *bool
}
// NewEligibilityRoutingNumber instantiates a new EligibilityRoutingNumber and sets the default values.
func NewEligibilityRoutingNumber()(*EligibilityRoutingNumber) {
    m := &EligibilityRoutingNumber{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEligibilityRoutingNumberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEligibilityRoutingNumberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEligibilityRoutingNumber(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EligibilityRoutingNumber) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEligible gets the eligible property value. Describes if the routing number is RTP or Fednow eligible or not
// returns a *bool when successful
func (m *EligibilityRoutingNumber) GetEligible()(*bool) {
    return m.eligible
}
// GetFednow gets the fednow property value. Describes if the routing number is Fednow eligible or not
// returns a *bool when successful
func (m *EligibilityRoutingNumber) GetFednow()(*bool) {
    return m.fednow
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EligibilityRoutingNumber) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["eligible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEligible(val)
        }
        return nil
    }
    res["fednow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFednow(val)
        }
        return nil
    }
    res["routing_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoutingNumber(val)
        }
        return nil
    }
    res["rtp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRtp(val)
        }
        return nil
    }
    return res
}
// GetRoutingNumber gets the routing_number property value. 9-digit American Bankers Association (ABA) routing number.
// returns a *string when successful
func (m *EligibilityRoutingNumber) GetRoutingNumber()(*string) {
    return m.routing_number
}
// GetRtp gets the rtp property value. Describes if the routing number is RTP eligible or not
// returns a *bool when successful
func (m *EligibilityRoutingNumber) GetRtp()(*bool) {
    return m.rtp
}
// Serialize serializes information the current object
func (m *EligibilityRoutingNumber) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("eligible", m.GetEligible())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fednow", m.GetFednow())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("routing_number", m.GetRoutingNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("rtp", m.GetRtp())
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
func (m *EligibilityRoutingNumber) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEligible sets the eligible property value. Describes if the routing number is RTP or Fednow eligible or not
func (m *EligibilityRoutingNumber) SetEligible(value *bool)() {
    m.eligible = value
}
// SetFednow sets the fednow property value. Describes if the routing number is Fednow eligible or not
func (m *EligibilityRoutingNumber) SetFednow(value *bool)() {
    m.fednow = value
}
// SetRoutingNumber sets the routing_number property value. 9-digit American Bankers Association (ABA) routing number.
func (m *EligibilityRoutingNumber) SetRoutingNumber(value *string)() {
    m.routing_number = value
}
// SetRtp sets the rtp property value. Describes if the routing number is RTP eligible or not
func (m *EligibilityRoutingNumber) SetRtp(value *bool)() {
    m.rtp = value
}
type EligibilityRoutingNumberable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEligible()(*bool)
    GetFednow()(*bool)
    GetRoutingNumber()(*string)
    GetRtp()(*bool)
    SetEligible(value *bool)()
    SetFednow(value *bool)()
    SetRoutingNumber(value *string)()
    SetRtp(value *bool)()
}
