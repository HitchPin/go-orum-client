package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Balance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Available amount of the balance in integral cents (example: 100 = $1).
    available *int64
    // Currency code in ISO 4217 format. Only USD is supported.
    currency *TransferCurrency
    // Pending amount of the balance in integral cents (example: 100 = $1).
    pending *int64
    // Timestamp when the resource was last updated.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewBalance instantiates a new Balance and sets the default values.
func NewBalance()(*Balance) {
    m := &Balance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBalanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBalanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBalance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Balance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailable gets the available property value. Available amount of the balance in integral cents (example: 100 = $1).
// returns a *int64 when successful
func (m *Balance) GetAvailable()(*int64) {
    return m.available
}
// GetCurrency gets the currency property value. Currency code in ISO 4217 format. Only USD is supported.
// returns a *TransferCurrency when successful
func (m *Balance) GetCurrency()(*TransferCurrency) {
    return m.currency
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Balance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["available"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailable(val)
        }
        return nil
    }
    res["currency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransferCurrency)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val.(*TransferCurrency))
        }
        return nil
    }
    res["pending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPending(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetPending gets the pending property value. Pending amount of the balance in integral cents (example: 100 = $1).
// returns a *int64 when successful
func (m *Balance) GetPending()(*int64) {
    return m.pending
}
// GetUpdatedAt gets the updated_at property value. Timestamp when the resource was last updated.
// returns a *Time when successful
func (m *Balance) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *Balance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("available", m.GetAvailable())
        if err != nil {
            return err
        }
    }
    if m.GetCurrency() != nil {
        cast := (*m.GetCurrency()).String()
        err := writer.WriteStringValue("currency", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("pending", m.GetPending())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
func (m *Balance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailable sets the available property value. Available amount of the balance in integral cents (example: 100 = $1).
func (m *Balance) SetAvailable(value *int64)() {
    m.available = value
}
// SetCurrency sets the currency property value. Currency code in ISO 4217 format. Only USD is supported.
func (m *Balance) SetCurrency(value *TransferCurrency)() {
    m.currency = value
}
// SetPending sets the pending property value. Pending amount of the balance in integral cents (example: 100 = $1).
func (m *Balance) SetPending(value *int64)() {
    m.pending = value
}
// SetUpdatedAt sets the updated_at property value. Timestamp when the resource was last updated.
func (m *Balance) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type Balanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailable()(*int64)
    GetCurrency()(*TransferCurrency)
    GetPending()(*int64)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAvailable(value *int64)()
    SetCurrency(value *TransferCurrency)()
    SetPending(value *int64)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
