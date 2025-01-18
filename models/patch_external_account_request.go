package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PatchExternalAccountRequest struct {
    // The account_holder_name property
    account_holder_name *string
    // The account_number property
    account_number *string
    // The account_type property
    account_type *AccountType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The metadata property
    metadata Metadataable
    // The routing_number property
    routing_number *string
}
// NewPatchExternalAccountRequest instantiates a new PatchExternalAccountRequest and sets the default values.
func NewPatchExternalAccountRequest()(*PatchExternalAccountRequest) {
    m := &PatchExternalAccountRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatchExternalAccountRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePatchExternalAccountRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatchExternalAccountRequest(), nil
}
// GetAccountHolderName gets the account_holder_name property value. The account_holder_name property
// returns a *string when successful
func (m *PatchExternalAccountRequest) GetAccountHolderName()(*string) {
    return m.account_holder_name
}
// GetAccountNumber gets the account_number property value. The account_number property
// returns a *string when successful
func (m *PatchExternalAccountRequest) GetAccountNumber()(*string) {
    return m.account_number
}
// GetAccountType gets the account_type property value. The account_type property
// returns a *AccountType when successful
func (m *PatchExternalAccountRequest) GetAccountType()(*AccountType) {
    return m.account_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PatchExternalAccountRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PatchExternalAccountRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_holder_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountHolderName(val)
        }
        return nil
    }
    res["account_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountNumber(val)
        }
        return nil
    }
    res["account_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountType(val.(*AccountType))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(Metadataable))
        }
        return nil
    }
    res["routing_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoutingNumber(val)
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. The metadata property
// returns a Metadataable when successful
func (m *PatchExternalAccountRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetRoutingNumber gets the routing_number property value. The routing_number property
// returns a *string when successful
func (m *PatchExternalAccountRequest) GetRoutingNumber()(*string) {
    return m.routing_number
}
// Serialize serializes information the current object
func (m *PatchExternalAccountRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("account_holder_name", m.GetAccountHolderName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("account_number", m.GetAccountNumber())
        if err != nil {
            return err
        }
    }
    if m.GetAccountType() != nil {
        cast := (*m.GetAccountType()).String()
        err := writer.WriteStringValue("account_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("routing_number", m.GetRoutingNumber())
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
// SetAccountHolderName sets the account_holder_name property value. The account_holder_name property
func (m *PatchExternalAccountRequest) SetAccountHolderName(value *string)() {
    m.account_holder_name = value
}
// SetAccountNumber sets the account_number property value. The account_number property
func (m *PatchExternalAccountRequest) SetAccountNumber(value *string)() {
    m.account_number = value
}
// SetAccountType sets the account_type property value. The account_type property
func (m *PatchExternalAccountRequest) SetAccountType(value *AccountType)() {
    m.account_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PatchExternalAccountRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *PatchExternalAccountRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetRoutingNumber sets the routing_number property value. The routing_number property
func (m *PatchExternalAccountRequest) SetRoutingNumber(value *string)() {
    m.routing_number = value
}
type PatchExternalAccountRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(*string)
    GetAccountNumber()(*string)
    GetAccountType()(*AccountType)
    GetMetadata()(Metadataable)
    GetRoutingNumber()(*string)
    SetAccountHolderName(value *string)()
    SetAccountNumber(value *string)()
    SetAccountType(value *AccountType)()
    SetMetadata(value Metadataable)()
    SetRoutingNumber(value *string)()
}
