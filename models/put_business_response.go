package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PutBusinessResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The business property
    business PutBusinessResponseBaseable
}
// NewPutBusinessResponse instantiates a new PutBusinessResponse and sets the default values.
func NewPutBusinessResponse()(*PutBusinessResponse) {
    m := &PutBusinessResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePutBusinessResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePutBusinessResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPutBusinessResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PutBusinessResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBusiness gets the business property value. The business property
// returns a PutBusinessResponseBaseable when successful
func (m *PutBusinessResponse) GetBusiness()(PutBusinessResponseBaseable) {
    return m.business
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PutBusinessResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["business"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePutBusinessResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusiness(val.(PutBusinessResponseBaseable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PutBusinessResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("business", m.GetBusiness())
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
func (m *PutBusinessResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBusiness sets the business property value. The business property
func (m *PutBusinessResponse) SetBusiness(value PutBusinessResponseBaseable)() {
    m.business = value
}
type PutBusinessResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusiness()(PutBusinessResponseBaseable)
    SetBusiness(value PutBusinessResponseBaseable)()
}
