package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetKeypairResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enterprise_keypair property
    enterprise_keypair EnterpriseKeypairable
}
// NewGetKeypairResponse instantiates a new GetKeypairResponse and sets the default values.
func NewGetKeypairResponse()(*GetKeypairResponse) {
    m := &GetKeypairResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetKeypairResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetKeypairResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetKeypairResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetKeypairResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnterpriseKeypair gets the enterprise_keypair property value. The enterprise_keypair property
// returns a EnterpriseKeypairable when successful
func (m *GetKeypairResponse) GetEnterpriseKeypair()(EnterpriseKeypairable) {
    return m.enterprise_keypair
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetKeypairResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enterprise_keypair"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseKeypairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseKeypair(val.(EnterpriseKeypairable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GetKeypairResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("enterprise_keypair", m.GetEnterpriseKeypair())
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
func (m *GetKeypairResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnterpriseKeypair sets the enterprise_keypair property value. The enterprise_keypair property
func (m *GetKeypairResponse) SetEnterpriseKeypair(value EnterpriseKeypairable)() {
    m.enterprise_keypair = value
}
type GetKeypairResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterpriseKeypair()(EnterpriseKeypairable)
    SetEnterpriseKeypair(value EnterpriseKeypairable)()
}
