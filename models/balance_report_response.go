package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BalanceReportResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The balance_statement_summary_report property
    balance_statement_summary_report BalanceReportResponseBaseable
}
// NewBalanceReportResponse instantiates a new BalanceReportResponse and sets the default values.
func NewBalanceReportResponse()(*BalanceReportResponse) {
    m := &BalanceReportResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBalanceReportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBalanceReportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBalanceReportResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BalanceReportResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBalanceStatementSummaryReport gets the balance_statement_summary_report property value. The balance_statement_summary_report property
// returns a BalanceReportResponseBaseable when successful
func (m *BalanceReportResponse) GetBalanceStatementSummaryReport()(BalanceReportResponseBaseable) {
    return m.balance_statement_summary_report
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BalanceReportResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["balance_statement_summary_report"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBalanceReportResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalanceStatementSummaryReport(val.(BalanceReportResponseBaseable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BalanceReportResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("balance_statement_summary_report", m.GetBalanceStatementSummaryReport())
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
func (m *BalanceReportResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBalanceStatementSummaryReport sets the balance_statement_summary_report property value. The balance_statement_summary_report property
func (m *BalanceReportResponse) SetBalanceStatementSummaryReport(value BalanceReportResponseBaseable)() {
    m.balance_statement_summary_report = value
}
type BalanceReportResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBalanceStatementSummaryReport()(BalanceReportResponseBaseable)
    SetBalanceStatementSummaryReport(value BalanceReportResponseBaseable)()
}
