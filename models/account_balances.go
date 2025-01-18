package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountBalances open and close balances for an account.
type AccountBalances struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Pending and available amounts.
    closing BalanceAmountsable
    // Pending and available amounts.
    opening BalanceAmountsable
}
// NewAccountBalances instantiates a new AccountBalances and sets the default values.
func NewAccountBalances()(*AccountBalances) {
    m := &AccountBalances{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccountBalancesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountBalancesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccountBalances(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccountBalances) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClosing gets the closing property value. Pending and available amounts.
// returns a BalanceAmountsable when successful
func (m *AccountBalances) GetClosing()(BalanceAmountsable) {
    return m.closing
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountBalances) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["closing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBalanceAmountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosing(val.(BalanceAmountsable))
        }
        return nil
    }
    res["opening"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBalanceAmountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpening(val.(BalanceAmountsable))
        }
        return nil
    }
    return res
}
// GetOpening gets the opening property value. Pending and available amounts.
// returns a BalanceAmountsable when successful
func (m *AccountBalances) GetOpening()(BalanceAmountsable) {
    return m.opening
}
// Serialize serializes information the current object
func (m *AccountBalances) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("closing", m.GetClosing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("opening", m.GetOpening())
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
func (m *AccountBalances) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClosing sets the closing property value. Pending and available amounts.
func (m *AccountBalances) SetClosing(value BalanceAmountsable)() {
    m.closing = value
}
// SetOpening sets the opening property value. Pending and available amounts.
func (m *AccountBalances) SetOpening(value BalanceAmountsable)() {
    m.opening = value
}
type AccountBalancesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosing()(BalanceAmountsable)
    GetOpening()(BalanceAmountsable)
    SetClosing(value BalanceAmountsable)()
    SetOpening(value BalanceAmountsable)()
}
