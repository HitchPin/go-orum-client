package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalAccountsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of External Accounts.
    external_accounts []ExternalAccountResponseBaseable
}
// NewExternalAccountsResponse instantiates a new ExternalAccountsResponse and sets the default values.
func NewExternalAccountsResponse()(*ExternalAccountsResponse) {
    m := &ExternalAccountsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalAccountsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalAccountsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalAccountsResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalAccountsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExternalAccounts gets the external_accounts property value. List of External Accounts.
// returns a []ExternalAccountResponseBaseable when successful
func (m *ExternalAccountsResponse) GetExternalAccounts()([]ExternalAccountResponseBaseable) {
    return m.external_accounts
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalAccountsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["external_accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalAccountResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalAccountResponseBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalAccountResponseBaseable)
                }
            }
            m.SetExternalAccounts(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExternalAccountsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExternalAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExternalAccounts()))
        for i, v := range m.GetExternalAccounts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("external_accounts", cast)
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
func (m *ExternalAccountsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExternalAccounts sets the external_accounts property value. List of External Accounts.
func (m *ExternalAccountsResponse) SetExternalAccounts(value []ExternalAccountResponseBaseable)() {
    m.external_accounts = value
}
type ExternalAccountsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalAccounts()([]ExternalAccountResponseBaseable)
    SetExternalAccounts(value []ExternalAccountResponseBaseable)()
}
