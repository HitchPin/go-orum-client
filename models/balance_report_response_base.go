package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BalanceReportResponseBase struct {
    // Open and close balances for an account.
    account_balances AccountBalancesable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The end time of the report range in ISO-8601 format.
    end_time *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ledger_entry_balances property
    ledger_entry_balances []LedgerEntryBalanceable
    // The start time of the report range in ISO-8601 format.
    start_time *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewBalanceReportResponseBase instantiates a new BalanceReportResponseBase and sets the default values.
func NewBalanceReportResponseBase()(*BalanceReportResponseBase) {
    m := &BalanceReportResponseBase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBalanceReportResponseBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBalanceReportResponseBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBalanceReportResponseBase(), nil
}
// GetAccountBalances gets the account_balances property value. Open and close balances for an account.
// returns a AccountBalancesable when successful
func (m *BalanceReportResponseBase) GetAccountBalances()(AccountBalancesable) {
    return m.account_balances
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BalanceReportResponseBase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEndTime gets the end_time property value. The end time of the report range in ISO-8601 format.
// returns a *Time when successful
func (m *BalanceReportResponseBase) GetEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.end_time
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BalanceReportResponseBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_balances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountBalancesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountBalances(val.(AccountBalancesable))
        }
        return nil
    }
    res["end_time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
        }
        return nil
    }
    res["ledger_entry_balances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLedgerEntryBalanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LedgerEntryBalanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LedgerEntryBalanceable)
                }
            }
            m.SetLedgerEntryBalances(res)
        }
        return nil
    }
    res["start_time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    return res
}
// GetLedgerEntryBalances gets the ledger_entry_balances property value. The ledger_entry_balances property
// returns a []LedgerEntryBalanceable when successful
func (m *BalanceReportResponseBase) GetLedgerEntryBalances()([]LedgerEntryBalanceable) {
    return m.ledger_entry_balances
}
// GetStartTime gets the start_time property value. The start time of the report range in ISO-8601 format.
// returns a *Time when successful
func (m *BalanceReportResponseBase) GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.start_time
}
// Serialize serializes information the current object
func (m *BalanceReportResponseBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("account_balances", m.GetAccountBalances())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("end_time", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    if m.GetLedgerEntryBalances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLedgerEntryBalances()))
        for i, v := range m.GetLedgerEntryBalances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ledger_entry_balances", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("start_time", m.GetStartTime())
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
// SetAccountBalances sets the account_balances property value. Open and close balances for an account.
func (m *BalanceReportResponseBase) SetAccountBalances(value AccountBalancesable)() {
    m.account_balances = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BalanceReportResponseBase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEndTime sets the end_time property value. The end time of the report range in ISO-8601 format.
func (m *BalanceReportResponseBase) SetEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.end_time = value
}
// SetLedgerEntryBalances sets the ledger_entry_balances property value. The ledger_entry_balances property
func (m *BalanceReportResponseBase) SetLedgerEntryBalances(value []LedgerEntryBalanceable)() {
    m.ledger_entry_balances = value
}
// SetStartTime sets the start_time property value. The start time of the report range in ISO-8601 format.
func (m *BalanceReportResponseBase) SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.start_time = value
}
type BalanceReportResponseBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountBalances()(AccountBalancesable)
    GetEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLedgerEntryBalances()([]LedgerEntryBalanceable)
    GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccountBalances(value AccountBalancesable)()
    SetEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLedgerEntryBalances(value []LedgerEntryBalanceable)()
    SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
