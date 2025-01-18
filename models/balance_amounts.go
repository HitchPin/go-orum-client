package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BalanceAmounts pending and available amounts.
type BalanceAmounts struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Amount in integral cents.
    available_amount *int64
    // Amount in integral cents.
    pending_amount *int64
}
// NewBalanceAmounts instantiates a new BalanceAmounts and sets the default values.
func NewBalanceAmounts()(*BalanceAmounts) {
    m := &BalanceAmounts{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBalanceAmountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBalanceAmountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBalanceAmounts(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BalanceAmounts) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableAmount gets the available_amount property value. Amount in integral cents.
// returns a *int64 when successful
func (m *BalanceAmounts) GetAvailableAmount()(*int64) {
    return m.available_amount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BalanceAmounts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["available_amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableAmount(val)
        }
        return nil
    }
    res["pending_amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingAmount(val)
        }
        return nil
    }
    return res
}
// GetPendingAmount gets the pending_amount property value. Amount in integral cents.
// returns a *int64 when successful
func (m *BalanceAmounts) GetPendingAmount()(*int64) {
    return m.pending_amount
}
// Serialize serializes information the current object
func (m *BalanceAmounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("available_amount", m.GetAvailableAmount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("pending_amount", m.GetPendingAmount())
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
func (m *BalanceAmounts) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableAmount sets the available_amount property value. Amount in integral cents.
func (m *BalanceAmounts) SetAvailableAmount(value *int64)() {
    m.available_amount = value
}
// SetPendingAmount sets the pending_amount property value. Amount in integral cents.
func (m *BalanceAmounts) SetPendingAmount(value *int64)() {
    m.pending_amount = value
}
type BalanceAmountsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableAmount()(*int64)
    GetPendingAmount()(*int64)
    SetAvailableAmount(value *int64)()
    SetPendingAmount(value *int64)()
}
