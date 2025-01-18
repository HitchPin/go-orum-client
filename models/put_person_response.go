package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PutPersonResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The person property
    person PutPersonResponseBaseable
}
// NewPutPersonResponse instantiates a new PutPersonResponse and sets the default values.
func NewPutPersonResponse()(*PutPersonResponse) {
    m := &PutPersonResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePutPersonResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePutPersonResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPutPersonResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PutPersonResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PutPersonResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["person"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePutPersonResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerson(val.(PutPersonResponseBaseable))
        }
        return nil
    }
    return res
}
// GetPerson gets the person property value. The person property
// returns a PutPersonResponseBaseable when successful
func (m *PutPersonResponse) GetPerson()(PutPersonResponseBaseable) {
    return m.person
}
// Serialize serializes information the current object
func (m *PutPersonResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("person", m.GetPerson())
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
func (m *PutPersonResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPerson sets the person property value. The person property
func (m *PutPersonResponse) SetPerson(value PutPersonResponseBaseable)() {
    m.person = value
}
type PutPersonResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPerson()(PutPersonResponseBaseable)
    SetPerson(value PutPersonResponseBaseable)()
}
