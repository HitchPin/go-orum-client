package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VerifyAccountsResponse struct {
    // An array of verify accounts.
    accounts []VerifyAccountResponseBaseable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewVerifyAccountsResponse instantiates a new VerifyAccountsResponse and sets the default values.
func NewVerifyAccountsResponse()(*VerifyAccountsResponse) {
    m := &VerifyAccountsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVerifyAccountsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVerifyAccountsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifyAccountsResponse(), nil
}
// GetAccounts gets the accounts property value. An array of verify accounts.
// returns a []VerifyAccountResponseBaseable when successful
func (m *VerifyAccountsResponse) GetAccounts()([]VerifyAccountResponseBaseable) {
    return m.accounts
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VerifyAccountsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VerifyAccountsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVerifyAccountResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VerifyAccountResponseBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VerifyAccountResponseBaseable)
                }
            }
            m.SetAccounts(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *VerifyAccountsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("accounts", cast)
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
// SetAccounts sets the accounts property value. An array of verify accounts.
func (m *VerifyAccountsResponse) SetAccounts(value []VerifyAccountResponseBaseable)() {
    m.accounts = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifyAccountsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type VerifyAccountsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()([]VerifyAccountResponseBaseable)
    SetAccounts(value []VerifyAccountResponseBaseable)()
}
