package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransferGroupResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The transfer_group property
    transfer_group TransferGroupable
}
// NewTransferGroupResponse instantiates a new TransferGroupResponse and sets the default values.
func NewTransferGroupResponse()(*TransferGroupResponse) {
    m := &TransferGroupResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransferGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransferGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransferGroupResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransferGroupResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransferGroupResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["transfer_group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransferGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferGroup(val.(TransferGroupable))
        }
        return nil
    }
    return res
}
// GetTransferGroup gets the transfer_group property value. The transfer_group property
// returns a TransferGroupable when successful
func (m *TransferGroupResponse) GetTransferGroup()(TransferGroupable) {
    return m.transfer_group
}
// Serialize serializes information the current object
func (m *TransferGroupResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("transfer_group", m.GetTransferGroup())
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
func (m *TransferGroupResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTransferGroup sets the transfer_group property value. The transfer_group property
func (m *TransferGroupResponse) SetTransferGroup(value TransferGroupable)() {
    m.transfer_group = value
}
type TransferGroupResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTransferGroup()(TransferGroupable)
    SetTransferGroup(value TransferGroupable)()
}
