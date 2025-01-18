package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransferRequest struct {
    // Banks can display additional information to help the end user understand why they received funds. This information will appear on both the source and destination bank account statements. The field supports 10 alphanumeric characters for ACH and 140 for RTP.
    account_statement_descriptor *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The amount property
    amount *int32
    // The currency property
    currency *TransferCurrency
    // The destination property
    destination DestinationTransferPartyable
    // The metadata property
    metadata Metadataable
    // The source property
    source SourceTransferPartyable
    // The speed property
    speed *TransferSpeed
    // The transfer_group_id property
    transfer_group_id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The transfer_reference_id property
    transfer_reference_id *string
}
// NewTransferRequest instantiates a new TransferRequest and sets the default values.
func NewTransferRequest()(*TransferRequest) {
    m := &TransferRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransferRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransferRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransferRequest(), nil
}
// GetAccountStatementDescriptor gets the account_statement_descriptor property value. Banks can display additional information to help the end user understand why they received funds. This information will appear on both the source and destination bank account statements. The field supports 10 alphanumeric characters for ACH and 140 for RTP.
// returns a *string when successful
func (m *TransferRequest) GetAccountStatementDescriptor()(*string) {
    return m.account_statement_descriptor
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransferRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAmount gets the amount property value. The amount property
// returns a *int32 when successful
func (m *TransferRequest) GetAmount()(*int32) {
    return m.amount
}
// GetCurrency gets the currency property value. The currency property
// returns a *TransferCurrency when successful
func (m *TransferRequest) GetCurrency()(*TransferCurrency) {
    return m.currency
}
// GetDestination gets the destination property value. The destination property
// returns a DestinationTransferPartyable when successful
func (m *TransferRequest) GetDestination()(DestinationTransferPartyable) {
    return m.destination
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransferRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_statement_descriptor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountStatementDescriptor(val)
        }
        return nil
    }
    res["amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmount(val)
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
    res["destination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDestinationTransferPartyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestination(val.(DestinationTransferPartyable))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(Metadataable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSourceTransferPartyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(SourceTransferPartyable))
        }
        return nil
    }
    res["speed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTransferSpeed)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeed(val.(*TransferSpeed))
        }
        return nil
    }
    res["transfer_group_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferGroupId(val)
        }
        return nil
    }
    res["transfer_reference_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferReferenceId(val)
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. The metadata property
// returns a Metadataable when successful
func (m *TransferRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetSource gets the source property value. The source property
// returns a SourceTransferPartyable when successful
func (m *TransferRequest) GetSource()(SourceTransferPartyable) {
    return m.source
}
// GetSpeed gets the speed property value. The speed property
// returns a *TransferSpeed when successful
func (m *TransferRequest) GetSpeed()(*TransferSpeed) {
    return m.speed
}
// GetTransferGroupId gets the transfer_group_id property value. The transfer_group_id property
// returns a *UUID when successful
func (m *TransferRequest) GetTransferGroupId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.transfer_group_id
}
// GetTransferReferenceId gets the transfer_reference_id property value. The transfer_reference_id property
// returns a *string when successful
func (m *TransferRequest) GetTransferReferenceId()(*string) {
    return m.transfer_reference_id
}
// Serialize serializes information the current object
func (m *TransferRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("account_statement_descriptor", m.GetAccountStatementDescriptor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("amount", m.GetAmount())
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
        err := writer.WriteObjectValue("destination", m.GetDestination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    if m.GetSpeed() != nil {
        cast := (*m.GetSpeed()).String()
        err := writer.WriteStringValue("speed", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("transfer_group_id", m.GetTransferGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("transfer_reference_id", m.GetTransferReferenceId())
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
// SetAccountStatementDescriptor sets the account_statement_descriptor property value. Banks can display additional information to help the end user understand why they received funds. This information will appear on both the source and destination bank account statements. The field supports 10 alphanumeric characters for ACH and 140 for RTP.
func (m *TransferRequest) SetAccountStatementDescriptor(value *string)() {
    m.account_statement_descriptor = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TransferRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAmount sets the amount property value. The amount property
func (m *TransferRequest) SetAmount(value *int32)() {
    m.amount = value
}
// SetCurrency sets the currency property value. The currency property
func (m *TransferRequest) SetCurrency(value *TransferCurrency)() {
    m.currency = value
}
// SetDestination sets the destination property value. The destination property
func (m *TransferRequest) SetDestination(value DestinationTransferPartyable)() {
    m.destination = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *TransferRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetSource sets the source property value. The source property
func (m *TransferRequest) SetSource(value SourceTransferPartyable)() {
    m.source = value
}
// SetSpeed sets the speed property value. The speed property
func (m *TransferRequest) SetSpeed(value *TransferSpeed)() {
    m.speed = value
}
// SetTransferGroupId sets the transfer_group_id property value. The transfer_group_id property
func (m *TransferRequest) SetTransferGroupId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.transfer_group_id = value
}
// SetTransferReferenceId sets the transfer_reference_id property value. The transfer_reference_id property
func (m *TransferRequest) SetTransferReferenceId(value *string)() {
    m.transfer_reference_id = value
}
type TransferRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountStatementDescriptor()(*string)
    GetAmount()(*int32)
    GetCurrency()(*TransferCurrency)
    GetDestination()(DestinationTransferPartyable)
    GetMetadata()(Metadataable)
    GetSource()(SourceTransferPartyable)
    GetSpeed()(*TransferSpeed)
    GetTransferGroupId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetTransferReferenceId()(*string)
    SetAccountStatementDescriptor(value *string)()
    SetAmount(value *int32)()
    SetCurrency(value *TransferCurrency)()
    SetDestination(value DestinationTransferPartyable)()
    SetMetadata(value Metadataable)()
    SetSource(value SourceTransferPartyable)()
    SetSpeed(value *TransferSpeed)()
    SetTransferGroupId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetTransferReferenceId(value *string)()
}
