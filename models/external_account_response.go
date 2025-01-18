package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalAccountResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A single object of External Account.
    external_account ExternalAccountResponseBaseable
}
// NewExternalAccountResponse instantiates a new ExternalAccountResponse and sets the default values.
func NewExternalAccountResponse()(*ExternalAccountResponse) {
    m := &ExternalAccountResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalAccountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalAccountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalAccountResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalAccountResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExternalAccount gets the external_account property value. A single object of External Account.
// returns a ExternalAccountResponseBaseable when successful
func (m *ExternalAccountResponse) GetExternalAccount()(ExternalAccountResponseBaseable) {
    return m.external_account
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalAccountResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["external_account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalAccountResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAccount(val.(ExternalAccountResponseBaseable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExternalAccountResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("external_account", m.GetExternalAccount())
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
func (m *ExternalAccountResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExternalAccount sets the external_account property value. A single object of External Account.
func (m *ExternalAccountResponse) SetExternalAccount(value ExternalAccountResponseBaseable)() {
    m.external_account = value
}
type ExternalAccountResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalAccount()(ExternalAccountResponseBaseable)
    SetExternalAccount(value ExternalAccountResponseBaseable)()
}
