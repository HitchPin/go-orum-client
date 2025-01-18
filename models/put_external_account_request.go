package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PutExternalAccountRequest struct {
    ExternalAccountRequest
}
// NewPutExternalAccountRequest instantiates a new PutExternalAccountRequest and sets the default values.
func NewPutExternalAccountRequest()(*PutExternalAccountRequest) {
    m := &PutExternalAccountRequest{
        ExternalAccountRequest: *NewExternalAccountRequest(),
    }
    return m
}
// CreatePutExternalAccountRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePutExternalAccountRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPutExternalAccountRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PutExternalAccountRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExternalAccountRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PutExternalAccountRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExternalAccountRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type PutExternalAccountRequestable interface {
    ExternalAccountRequestable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
