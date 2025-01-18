package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PutBusinessResponseBase struct {
    BusinessResponseBase
}
// NewPutBusinessResponseBase instantiates a new PutBusinessResponseBase and sets the default values.
func NewPutBusinessResponseBase()(*PutBusinessResponseBase) {
    m := &PutBusinessResponseBase{
        BusinessResponseBase: *NewBusinessResponseBase(),
    }
    return m
}
// CreatePutBusinessResponseBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePutBusinessResponseBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPutBusinessResponseBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PutBusinessResponseBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BusinessResponseBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PutBusinessResponseBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BusinessResponseBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type PutBusinessResponseBaseable interface {
    BusinessResponseBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
