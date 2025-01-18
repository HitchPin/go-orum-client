package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TransferStatusReason for failed transfers - details on why the transfer is in a failed state.
type TransferStatusReason struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The destination property
    destination ReasonCodesable
    // The source property
    source ReasonCodesable
}
// NewTransferStatusReason instantiates a new TransferStatusReason and sets the default values.
func NewTransferStatusReason()(*TransferStatusReason) {
    m := &TransferStatusReason{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTransferStatusReasonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransferStatusReasonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransferStatusReason(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransferStatusReason) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestination gets the destination property value. The destination property
// returns a ReasonCodesable when successful
func (m *TransferStatusReason) GetDestination()(ReasonCodesable) {
    return m.destination
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransferStatusReason) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReasonCodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestination(val.(ReasonCodesable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReasonCodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(ReasonCodesable))
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The source property
// returns a ReasonCodesable when successful
func (m *TransferStatusReason) GetSource()(ReasonCodesable) {
    return m.source
}
// Serialize serializes information the current object
func (m *TransferStatusReason) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TransferStatusReason) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestination sets the destination property value. The destination property
func (m *TransferStatusReason) SetDestination(value ReasonCodesable)() {
    m.destination = value
}
// SetSource sets the source property value. The source property
func (m *TransferStatusReason) SetSource(value ReasonCodesable)() {
    m.source = value
}
type TransferStatusReasonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestination()(ReasonCodesable)
    GetSource()(ReasonCodesable)
    SetDestination(value ReasonCodesable)()
    SetSource(value ReasonCodesable)()
}
