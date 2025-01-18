package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BalanceResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The balance property
    balance Balanceable
}
// NewBalanceResponse instantiates a new BalanceResponse and sets the default values.
func NewBalanceResponse()(*BalanceResponse) {
    m := &BalanceResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBalanceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBalanceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBalanceResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BalanceResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBalance gets the balance property value. The balance property
// returns a Balanceable when successful
func (m *BalanceResponse) GetBalance()(Balanceable) {
    return m.balance
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BalanceResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["balance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBalanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalance(val.(Balanceable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BalanceResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("balance", m.GetBalance())
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
func (m *BalanceResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBalance sets the balance property value. The balance property
func (m *BalanceResponse) SetBalance(value Balanceable)() {
    m.balance = value
}
type BalanceResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBalance()(Balanceable)
    SetBalance(value Balanceable)()
}
