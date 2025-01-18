package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransferGroupRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The destination property
    destination TransferGroupDestinationObjectable
    // The source property
    source TransferGroupSourceObjectable
    // The transfer_group_reference_id property
    transfer_group_reference_id *string
}
// NewTransferGroupRequest instantiates a new TransferGroupRequest and sets the default values.
func NewTransferGroupRequest()(*TransferGroupRequest) {
    m := &TransferGroupRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransferGroupRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransferGroupRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransferGroupRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransferGroupRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestination gets the destination property value. The destination property
// returns a TransferGroupDestinationObjectable when successful
func (m *TransferGroupRequest) GetDestination()(TransferGroupDestinationObjectable) {
    return m.destination
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransferGroupRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransferGroupDestinationObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestination(val.(TransferGroupDestinationObjectable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransferGroupSourceObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(TransferGroupSourceObjectable))
        }
        return nil
    }
    res["transfer_group_reference_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferGroupReferenceId(val)
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The source property
// returns a TransferGroupSourceObjectable when successful
func (m *TransferGroupRequest) GetSource()(TransferGroupSourceObjectable) {
    return m.source
}
// GetTransferGroupReferenceId gets the transfer_group_reference_id property value. The transfer_group_reference_id property
// returns a *string when successful
func (m *TransferGroupRequest) GetTransferGroupReferenceId()(*string) {
    return m.transfer_group_reference_id
}
// Serialize serializes information the current object
func (m *TransferGroupRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("destination", m.GetDestination())
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
    {
        err := writer.WriteStringValue("transfer_group_reference_id", m.GetTransferGroupReferenceId())
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
func (m *TransferGroupRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestination sets the destination property value. The destination property
func (m *TransferGroupRequest) SetDestination(value TransferGroupDestinationObjectable)() {
    m.destination = value
}
// SetSource sets the source property value. The source property
func (m *TransferGroupRequest) SetSource(value TransferGroupSourceObjectable)() {
    m.source = value
}
// SetTransferGroupReferenceId sets the transfer_group_reference_id property value. The transfer_group_reference_id property
func (m *TransferGroupRequest) SetTransferGroupReferenceId(value *string)() {
    m.transfer_group_reference_id = value
}
type TransferGroupRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestination()(TransferGroupDestinationObjectable)
    GetSource()(TransferGroupSourceObjectable)
    GetTransferGroupReferenceId()(*string)
    SetDestination(value TransferGroupDestinationObjectable)()
    SetSource(value TransferGroupSourceObjectable)()
    SetTransferGroupReferenceId(value *string)()
}
