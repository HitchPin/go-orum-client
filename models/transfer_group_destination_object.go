package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransferGroupDestinationObject information about the ultimate destination of the transfer group.
type TransferGroupDestinationObject struct {
    // Unique id of the account on the Enterprise platform
    account_reference_id *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique id of the customer on the Enterprise platform
    customer_reference_id *string
}
// NewTransferGroupDestinationObject instantiates a new TransferGroupDestinationObject and sets the default values.
func NewTransferGroupDestinationObject()(*TransferGroupDestinationObject) {
    m := &TransferGroupDestinationObject{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransferGroupDestinationObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransferGroupDestinationObjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransferGroupDestinationObject(), nil
}
// GetAccountReferenceId gets the account_reference_id property value. Unique id of the account on the Enterprise platform
// returns a *string when successful
func (m *TransferGroupDestinationObject) GetAccountReferenceId()(*string) {
    return m.account_reference_id
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransferGroupDestinationObject) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomerReferenceId gets the customer_reference_id property value. Unique id of the customer on the Enterprise platform
// returns a *string when successful
func (m *TransferGroupDestinationObject) GetCustomerReferenceId()(*string) {
    return m.customer_reference_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransferGroupDestinationObject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *TransferGroupDestinationObject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountReferenceId sets the account_reference_id property value. Unique id of the account on the Enterprise platform
func (m *TransferGroupDestinationObject) SetAccountReferenceId(value *string)() {
    m.account_reference_id = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TransferGroupDestinationObject) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomerReferenceId sets the customer_reference_id property value. Unique id of the customer on the Enterprise platform
func (m *TransferGroupDestinationObject) SetCustomerReferenceId(value *string)() {
    m.customer_reference_id = value
}
type TransferGroupDestinationObjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountReferenceId()(*string)
    GetCustomerReferenceId()(*string)
    SetAccountReferenceId(value *string)()
    SetCustomerReferenceId(value *string)()
}
