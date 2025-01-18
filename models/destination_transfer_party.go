package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DestinationTransferParty information about the transfer funds destination, which will be credited.
type DestinationTransferParty struct {
    // Unique reference ID for the account being credited.
    account_reference_id *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique reference ID for the customer (person or business) associated with the external account being credited.
    customer_reference_id *string
    // The name that will appear on the bank account statement of the account being credited. The field supports 16 alphanumeric characters for ACH, and 140 for RTP.
    statement_display_name *string
}
// NewDestinationTransferParty instantiates a new DestinationTransferParty and sets the default values.
func NewDestinationTransferParty()(*DestinationTransferParty) {
    m := &DestinationTransferParty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDestinationTransferPartyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDestinationTransferPartyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDestinationTransferParty(), nil
}
// GetAccountReferenceId gets the account_reference_id property value. Unique reference ID for the account being credited.
// returns a *string when successful
func (m *DestinationTransferParty) GetAccountReferenceId()(*string) {
    return m.account_reference_id
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DestinationTransferParty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomerReferenceId gets the customer_reference_id property value. Unique reference ID for the customer (person or business) associated with the external account being credited.
// returns a *string when successful
func (m *DestinationTransferParty) GetCustomerReferenceId()(*string) {
    return m.customer_reference_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DestinationTransferParty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_reference_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountReferenceId(val)
        }
        return nil
    }
    res["customer_reference_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerReferenceId(val)
        }
        return nil
    }
    res["statement_display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatementDisplayName(val)
        }
        return nil
    }
    return res
}
// GetStatementDisplayName gets the statement_display_name property value. The name that will appear on the bank account statement of the account being credited. The field supports 16 alphanumeric characters for ACH, and 140 for RTP.
// returns a *string when successful
func (m *DestinationTransferParty) GetStatementDisplayName()(*string) {
    return m.statement_display_name
}
// Serialize serializes information the current object
func (m *DestinationTransferParty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("account_reference_id", m.GetAccountReferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customer_reference_id", m.GetCustomerReferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("statement_display_name", m.GetStatementDisplayName())
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
// SetAccountReferenceId sets the account_reference_id property value. Unique reference ID for the account being credited.
func (m *DestinationTransferParty) SetAccountReferenceId(value *string)() {
    m.account_reference_id = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DestinationTransferParty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomerReferenceId sets the customer_reference_id property value. Unique reference ID for the customer (person or business) associated with the external account being credited.
func (m *DestinationTransferParty) SetCustomerReferenceId(value *string)() {
    m.customer_reference_id = value
}
// SetStatementDisplayName sets the statement_display_name property value. The name that will appear on the bank account statement of the account being credited. The field supports 16 alphanumeric characters for ACH, and 140 for RTP.
func (m *DestinationTransferParty) SetStatementDisplayName(value *string)() {
    m.statement_display_name = value
}
type DestinationTransferPartyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountReferenceId()(*string)
    GetCustomerReferenceId()(*string)
    GetStatementDisplayName()(*string)
    SetAccountReferenceId(value *string)()
    SetCustomerReferenceId(value *string)()
    SetStatementDisplayName(value *string)()
}
