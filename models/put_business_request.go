package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PutBusinessRequest struct {
    BusinessRequest
}
// NewPutBusinessRequest instantiates a new PutBusinessRequest and sets the default values.
func NewPutBusinessRequest()(*PutBusinessRequest) {
    m := &PutBusinessRequest{
        BusinessRequest: *NewBusinessRequest(),
    }
    return m
}
// CreatePutBusinessRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePutBusinessRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPutBusinessRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PutBusinessRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BusinessRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PutBusinessRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BusinessRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type PutBusinessRequestable interface {
    BusinessRequestable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
