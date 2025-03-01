package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalAccountResponseBase a single object of External Account.
type ExternalAccountResponseBase struct {
    // Name of account holder.
    account_holder_name *string
    // Account number for US bank account. 17 digits maximum.
    account_number *string
    // Unique reference id for the external account. Generated by you.
    account_reference_id *string
    // Type of bank account - checking or savings.
    account_type *AccountType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Timestamp when the resource was created.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique customer_reference_id that you passed when creating the associated customer (business or person) resource.
    customer_reference_id *string
    // Type of customer resource - business, person, or enterprise.
    customer_resource_type *EndCustomerResourceType
    // Orum generated unique id for the resource.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Additional data you would like to provide on the resource. The field supports valid JSON of up to 5 key-value pairs with a maximum of 20 characters for the key and 50 characters for the value. Do not include any sensitive information.
    metadata Metadataable
    // 9-digit American Bankers Association (ABA) routing number.
    routing_number *string
    // Status of the external account.
    status *ExternalAccountStatus
    // Timestamp when the resource was last updated.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewExternalAccountResponseBase instantiates a new ExternalAccountResponseBase and sets the default values.
func NewExternalAccountResponseBase()(*ExternalAccountResponseBase) {
    m := &ExternalAccountResponseBase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExternalAccountResponseBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalAccountResponseBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalAccountResponseBase(), nil
}
// GetAccountHolderName gets the account_holder_name property value. Name of account holder.
// returns a *string when successful
func (m *ExternalAccountResponseBase) GetAccountHolderName()(*string) {
    return m.account_holder_name
}
// GetAccountNumber gets the account_number property value. Account number for US bank account. 17 digits maximum.
// returns a *string when successful
func (m *ExternalAccountResponseBase) GetAccountNumber()(*string) {
    return m.account_number
}
// GetAccountReferenceId gets the account_reference_id property value. Unique reference id for the external account. Generated by you.
// returns a *string when successful
func (m *ExternalAccountResponseBase) GetAccountReferenceId()(*string) {
    return m.account_reference_id
}
// GetAccountType gets the account_type property value. Type of bank account - checking or savings.
// returns a *AccountType when successful
func (m *ExternalAccountResponseBase) GetAccountType()(*AccountType) {
    return m.account_type
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExternalAccountResponseBase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. Timestamp when the resource was created.
// returns a *Time when successful
func (m *ExternalAccountResponseBase) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCustomerReferenceId gets the customer_reference_id property value. Unique customer_reference_id that you passed when creating the associated customer (business or person) resource.
// returns a *string when successful
func (m *ExternalAccountResponseBase) GetCustomerReferenceId()(*string) {
    return m.customer_reference_id
}
// GetCustomerResourceType gets the customer_resource_type property value. Type of customer resource - business, person, or enterprise.
// returns a *EndCustomerResourceType when successful
func (m *ExternalAccountResponseBase) GetCustomerResourceType()(*EndCustomerResourceType) {
    return m.customer_resource_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalAccountResponseBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAccountStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ExternalAccountStatus))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Orum generated unique id for the resource.
// returns a *UUID when successful
func (m *ExternalAccountResponseBase) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetMetadata gets the metadata property value. Additional data you would like to provide on the resource. The field supports valid JSON of up to 5 key-value pairs with a maximum of 20 characters for the key and 50 characters for the value. Do not include any sensitive information.
// returns a Metadataable when successful
func (m *ExternalAccountResponseBase) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetRoutingNumber gets the routing_number property value. 9-digit American Bankers Association (ABA) routing number.
// returns a *string when successful
func (m *ExternalAccountResponseBase) GetRoutingNumber()(*string) {
    return m.routing_number
}
// GetStatus gets the status property value. Status of the external account.
// returns a *ExternalAccountStatus when successful
func (m *ExternalAccountResponseBase) GetStatus()(*ExternalAccountStatus) {
    return m.status
}
// GetUpdatedAt gets the updated_at property value. Timestamp when the resource was last updated.
// returns a *Time when successful
func (m *ExternalAccountResponseBase) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *ExternalAccountResponseBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteUUIDValue("id", m.GetId())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
// SetAccountHolderName sets the account_holder_name property value. Name of account holder.
func (m *ExternalAccountResponseBase) SetAccountHolderName(value *string)() {
    m.account_holder_name = value
}
// SetAccountNumber sets the account_number property value. Account number for US bank account. 17 digits maximum.
func (m *ExternalAccountResponseBase) SetAccountNumber(value *string)() {
    m.account_number = value
}
// SetAccountReferenceId sets the account_reference_id property value. Unique reference id for the external account. Generated by you.
func (m *ExternalAccountResponseBase) SetAccountReferenceId(value *string)() {
    m.account_reference_id = value
}
// SetAccountType sets the account_type property value. Type of bank account - checking or savings.
func (m *ExternalAccountResponseBase) SetAccountType(value *AccountType)() {
    m.account_type = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExternalAccountResponseBase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. Timestamp when the resource was created.
func (m *ExternalAccountResponseBase) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCustomerReferenceId sets the customer_reference_id property value. Unique customer_reference_id that you passed when creating the associated customer (business or person) resource.
func (m *ExternalAccountResponseBase) SetCustomerReferenceId(value *string)() {
    m.customer_reference_id = value
}
// SetCustomerResourceType sets the customer_resource_type property value. Type of customer resource - business, person, or enterprise.
func (m *ExternalAccountResponseBase) SetCustomerResourceType(value *EndCustomerResourceType)() {
    m.customer_resource_type = value
}
// SetId sets the id property value. Orum generated unique id for the resource.
func (m *ExternalAccountResponseBase) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetMetadata sets the metadata property value. Additional data you would like to provide on the resource. The field supports valid JSON of up to 5 key-value pairs with a maximum of 20 characters for the key and 50 characters for the value. Do not include any sensitive information.
func (m *ExternalAccountResponseBase) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetRoutingNumber sets the routing_number property value. 9-digit American Bankers Association (ABA) routing number.
func (m *ExternalAccountResponseBase) SetRoutingNumber(value *string)() {
    m.routing_number = value
}
// SetStatus sets the status property value. Status of the external account.
func (m *ExternalAccountResponseBase) SetStatus(value *ExternalAccountStatus)() {
    m.status = value
}
// SetUpdatedAt sets the updated_at property value. Timestamp when the resource was last updated.
func (m *ExternalAccountResponseBase) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type ExternalAccountResponseBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(*string)
    GetAccountNumber()(*string)
    GetAccountReferenceId()(*string)
    GetAccountType()(*AccountType)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerReferenceId()(*string)
    GetCustomerResourceType()(*EndCustomerResourceType)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetMetadata()(Metadataable)
    GetRoutingNumber()(*string)
    GetStatus()(*ExternalAccountStatus)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccountHolderName(value *string)()
    SetAccountNumber(value *string)()
    SetAccountReferenceId(value *string)()
    SetAccountType(value *AccountType)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerReferenceId(value *string)()
    SetCustomerResourceType(value *EndCustomerResourceType)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetMetadata(value Metadataable)()
    SetRoutingNumber(value *string)()
    SetStatus(value *ExternalAccountStatus)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
