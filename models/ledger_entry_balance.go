package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LedgerEntryBalance ledger entries with type and debit/credit balances.
type LedgerEntryBalance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Amount in integral cents.
    credit *int64
    // Amount in integral cents.
    debit *int64
    // Type of ledger entry.
    typeEscaped *LedgerEntryType
}
// NewLedgerEntryBalance instantiates a new LedgerEntryBalance and sets the default values.
func NewLedgerEntryBalance()(*LedgerEntryBalance) {
    m := &LedgerEntryBalance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLedgerEntryBalanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLedgerEntryBalanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLedgerEntryBalance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LedgerEntryBalance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCredit gets the credit property value. Amount in integral cents.
// returns a *int64 when successful
func (m *LedgerEntryBalance) GetCredit()(*int64) {
    return m.credit
}
// GetDebit gets the debit property value. Amount in integral cents.
// returns a *int64 when successful
func (m *LedgerEntryBalance) GetDebit()(*int64) {
    return m.debit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LedgerEntryBalance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredit(val)
        }
        return nil
    }
    res["debit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebit(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLedgerEntryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*LedgerEntryType))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. Type of ledger entry.
// returns a *LedgerEntryType when successful
func (m *LedgerEntryBalance) GetTypeEscaped()(*LedgerEntryType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *LedgerEntryBalance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("credit", m.GetCredit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("debit", m.GetDebit())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *LedgerEntryBalance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCredit sets the credit property value. Amount in integral cents.
func (m *LedgerEntryBalance) SetCredit(value *int64)() {
    m.credit = value
}
// SetDebit sets the debit property value. Amount in integral cents.
func (m *LedgerEntryBalance) SetDebit(value *int64)() {
    m.debit = value
}
// SetTypeEscaped sets the type property value. Type of ledger entry.
func (m *LedgerEntryBalance) SetTypeEscaped(value *LedgerEntryType)() {
    m.typeEscaped = value
}
type LedgerEntryBalanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCredit()(*int64)
    GetDebit()(*int64)
    GetTypeEscaped()(*LedgerEntryType)
    SetCredit(value *int64)()
    SetDebit(value *int64)()
    SetTypeEscaped(value *LedgerEntryType)()
}
