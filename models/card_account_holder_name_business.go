package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CardAccountHolderNameBusiness business name of account holder. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
type CardAccountHolderNameBusiness struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
    company_name *string
}
// NewCardAccountHolderNameBusiness instantiates a new CardAccountHolderNameBusiness and sets the default values.
func NewCardAccountHolderNameBusiness()(*CardAccountHolderNameBusiness) {
    m := &CardAccountHolderNameBusiness{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCardAccountHolderNameBusinessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCardAccountHolderNameBusinessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCardAccountHolderNameBusiness(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CardAccountHolderNameBusiness) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompanyName gets the company_name property value. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
// returns a *string when successful
func (m *CardAccountHolderNameBusiness) GetCompanyName()(*string) {
    return m.company_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CardAccountHolderNameBusiness) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["company_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CardAccountHolderNameBusiness) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("company_name", m.GetCompanyName())
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
func (m *CardAccountHolderNameBusiness) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompanyName sets the company_name property value. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
func (m *CardAccountHolderNameBusiness) SetCompanyName(value *string)() {
    m.company_name = value
}
type CardAccountHolderNameBusinessable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompanyName()(*string)
    SetCompanyName(value *string)()
}
