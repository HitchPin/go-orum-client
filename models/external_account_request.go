package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExternalAccountRequest struct {
    // The account_holder_name property
    account_holder_name *string
    // The account_number property
    account_number *string
    // The account_reference_id property
    account_reference_id *string
    // The account_type property
    account_type *AccountType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The customer_reference_id property
    customer_reference_id *string
    // The customer_resource_type property
    customer_resource_type *EndCustomerResourceType
    // The metadata property
    metadata Metadataable
    // The routing_number property
    routing_number *string
}
// NewExternalAccountRequest instantiates a new ExternalAccountRequest and sets the default values.
func NewExternalAccountRequest()(*ExternalAccountRequest) {
    m := &ExternalAccountRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalAccountRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalAccountRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalAccountRequest(), nil
}
// GetAccountHolderName gets the account_holder_name property value. The account_holder_name property
// returns a *string when successful
func (m *ExternalAccountRequest) GetAccountHolderName()(*string) {
    return m.account_holder_name
}
// GetAccountNumber gets the account_number property value. The account_number property
// returns a *string when successful
func (m *ExternalAccountRequest) GetAccountNumber()(*string) {
    return m.account_number
}
// GetAccountReferenceId gets the account_reference_id property value. The account_reference_id property
// returns a *string when successful
func (m *ExternalAccountRequest) GetAccountReferenceId()(*string) {
    return m.account_reference_id
}
// GetAccountType gets the account_type property value. The account_type property
// returns a *AccountType when successful
func (m *ExternalAccountRequest) GetAccountType()(*AccountType) {
    return m.account_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalAccountRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomerReferenceId gets the customer_reference_id property value. The customer_reference_id property
// returns a *string when successful
func (m *ExternalAccountRequest) GetCustomerReferenceId()(*string) {
    return m.customer_reference_id
}
// GetCustomerResourceType gets the customer_resource_type property value. The customer_resource_type property
// returns a *EndCustomerResourceType when successful
func (m *ExternalAccountRequest) GetCustomerResourceType()(*EndCustomerResourceType) {
    return m.customer_resource_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalAccountRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["account_reference_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountReferenceId(val)
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
    res["customer_reference_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerReferenceId(val)
        }
        return nil
    }
    res["customer_resource_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndCustomerResourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerResourceType(val.(*EndCustomerResourceType))
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
func (m *ExternalAccountRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetRoutingNumber gets the routing_number property value. The routing_number property
// returns a *string when successful
func (m *ExternalAccountRequest) GetRoutingNumber()(*string) {
    return m.routing_number
}
// Serialize serializes information the current object
func (m *ExternalAccountRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("account_reference_id", m.GetAccountReferenceId())
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
        err := writer.WriteStringValue("customer_reference_id", m.GetCustomerReferenceId())
        if err != nil {
            return err
        }
    }
    if m.GetCustomerResourceType() != nil {
        cast := (*m.GetCustomerResourceType()).String()
        err := writer.WriteStringValue("customer_resource_type", &cast)
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
func (m *ExternalAccountRequest) SetAccountHolderName(value *string)() {
    m.account_holder_name = value
}
// SetAccountNumber sets the account_number property value. The account_number property
func (m *ExternalAccountRequest) SetAccountNumber(value *string)() {
    m.account_number = value
}
// SetAccountReferenceId sets the account_reference_id property value. The account_reference_id property
func (m *ExternalAccountRequest) SetAccountReferenceId(value *string)() {
    m.account_reference_id = value
}
// SetAccountType sets the account_type property value. The account_type property
func (m *ExternalAccountRequest) SetAccountType(value *AccountType)() {
    m.account_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExternalAccountRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomerReferenceId sets the customer_reference_id property value. The customer_reference_id property
func (m *ExternalAccountRequest) SetCustomerReferenceId(value *string)() {
    m.customer_reference_id = value
}
// SetCustomerResourceType sets the customer_resource_type property value. The customer_resource_type property
func (m *ExternalAccountRequest) SetCustomerResourceType(value *EndCustomerResourceType)() {
    m.customer_resource_type = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *ExternalAccountRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetRoutingNumber sets the routing_number property value. The routing_number property
func (m *ExternalAccountRequest) SetRoutingNumber(value *string)() {
    m.routing_number = value
}
type ExternalAccountRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(*string)
    GetAccountNumber()(*string)
    GetAccountReferenceId()(*string)
    GetAccountType()(*AccountType)
    GetCustomerReferenceId()(*string)
    GetCustomerResourceType()(*EndCustomerResourceType)
    GetMetadata()(Metadataable)
    GetRoutingNumber()(*string)
    SetAccountHolderName(value *string)()
    SetAccountNumber(value *string)()
    SetAccountReferenceId(value *string)()
    SetAccountType(value *AccountType)()
    SetCustomerReferenceId(value *string)()
    SetCustomerResourceType(value *EndCustomerResourceType)()
    SetMetadata(value Metadataable)()
    SetRoutingNumber(value *string)()
}
