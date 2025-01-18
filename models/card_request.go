package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CardRequest struct {
    // The account_holder_name property
    account_holder_name CardRequest_CardRequest_account_holder_nameable
    // The account_reference_id property
    account_reference_id *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address property
    address CardRequest_CardRequest_addressable
    // The customer_reference_id property
    customer_reference_id *string
    // The customer_resource_type property
    customer_resource_type *EndCustomerResourceType
    // The expiration_month property
    expiration_month *int32
    // The expiration_year property
    expiration_year *int32
    // The metadata property
    metadata Metadataable
    // The card's number
    number *string
}
// CardRequest_CardRequest_account_holder_name composed type wrapper for classes CardAccountHolderNameBusinessable, CardAccountHolderNamePersonable
type CardRequest_CardRequest_account_holder_name struct {
    // Composed type representation for type CardAccountHolderNameBusinessable
    cardAccountHolderNameBusiness CardAccountHolderNameBusinessable
    // Composed type representation for type CardAccountHolderNamePersonable
    cardAccountHolderNamePerson CardAccountHolderNamePersonable
}
// NewCardRequest_CardRequest_account_holder_name instantiates a new CardRequest_CardRequest_account_holder_name and sets the default values.
func NewCardRequest_CardRequest_account_holder_name()(*CardRequest_CardRequest_account_holder_name) {
    m := &CardRequest_CardRequest_account_holder_name{
    }
    return m
}
// CreateCardRequest_CardRequest_account_holder_nameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCardRequest_CardRequest_account_holder_nameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCardRequest_CardRequest_account_holder_name()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetCardAccountHolderNameBusiness gets the CardAccountHolderNameBusiness property value. Composed type representation for type CardAccountHolderNameBusinessable
// returns a CardAccountHolderNameBusinessable when successful
func (m *CardRequest_CardRequest_account_holder_name) GetCardAccountHolderNameBusiness()(CardAccountHolderNameBusinessable) {
    return m.cardAccountHolderNameBusiness
}
// GetCardAccountHolderNamePerson gets the CardAccountHolderNamePerson property value. Composed type representation for type CardAccountHolderNamePersonable
// returns a CardAccountHolderNamePersonable when successful
func (m *CardRequest_CardRequest_account_holder_name) GetCardAccountHolderNamePerson()(CardAccountHolderNamePersonable) {
    return m.cardAccountHolderNamePerson
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CardRequest_CardRequest_account_holder_name) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetCardAccountHolderNameBusiness() != nil {
        return m.GetCardAccountHolderNameBusiness().GetFieldDeserializers()
    } else if m.GetCardAccountHolderNamePerson() != nil {
        return m.GetCardAccountHolderNamePerson().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CardRequest_CardRequest_account_holder_name) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *CardRequest_CardRequest_account_holder_name) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCardAccountHolderNameBusiness() != nil {
        err := writer.WriteObjectValue("", m.GetCardAccountHolderNameBusiness())
        if err != nil {
            return err
        }
    } else if m.GetCardAccountHolderNamePerson() != nil {
        err := writer.WriteObjectValue("", m.GetCardAccountHolderNamePerson())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCardAccountHolderNameBusiness sets the CardAccountHolderNameBusiness property value. Composed type representation for type CardAccountHolderNameBusinessable
func (m *CardRequest_CardRequest_account_holder_name) SetCardAccountHolderNameBusiness(value CardAccountHolderNameBusinessable)() {
    m.cardAccountHolderNameBusiness = value
}
// SetCardAccountHolderNamePerson sets the CardAccountHolderNamePerson property value. Composed type representation for type CardAccountHolderNamePersonable
func (m *CardRequest_CardRequest_account_holder_name) SetCardAccountHolderNamePerson(value CardAccountHolderNamePersonable)() {
    m.cardAccountHolderNamePerson = value
}
// CardRequest_CardRequest_address composed type wrapper for classes CardFullAddressable, CardZipOnlyAddressable
type CardRequest_CardRequest_address struct {
    // Composed type representation for type CardFullAddressable
    cardFullAddress CardFullAddressable
    // Composed type representation for type CardZipOnlyAddressable
    cardZipOnlyAddress CardZipOnlyAddressable
}
// NewCardRequest_CardRequest_address instantiates a new CardRequest_CardRequest_address and sets the default values.
func NewCardRequest_CardRequest_address()(*CardRequest_CardRequest_address) {
    m := &CardRequest_CardRequest_address{
    }
    return m
}
// CreateCardRequest_CardRequest_addressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCardRequest_CardRequest_addressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCardRequest_CardRequest_address()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateCardFullAddressFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(CardFullAddressable); ok {
                result.SetCardFullAddress(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateCardZipOnlyAddressFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(CardZipOnlyAddressable); ok {
                result.SetCardZipOnlyAddress(cast)
            }
        }
    }
    return result, nil
}
// GetCardFullAddress gets the CardFullAddress property value. Composed type representation for type CardFullAddressable
// returns a CardFullAddressable when successful
func (m *CardRequest_CardRequest_address) GetCardFullAddress()(CardFullAddressable) {
    return m.cardFullAddress
}
// GetCardZipOnlyAddress gets the CardZipOnlyAddress property value. Composed type representation for type CardZipOnlyAddressable
// returns a CardZipOnlyAddressable when successful
func (m *CardRequest_CardRequest_address) GetCardZipOnlyAddress()(CardZipOnlyAddressable) {
    return m.cardZipOnlyAddress
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CardRequest_CardRequest_address) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CardRequest_CardRequest_address) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *CardRequest_CardRequest_address) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCardFullAddress() != nil {
        err := writer.WriteObjectValue("", m.GetCardFullAddress())
        if err != nil {
            return err
        }
    } else if m.GetCardZipOnlyAddress() != nil {
        err := writer.WriteObjectValue("", m.GetCardZipOnlyAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCardFullAddress sets the CardFullAddress property value. Composed type representation for type CardFullAddressable
func (m *CardRequest_CardRequest_address) SetCardFullAddress(value CardFullAddressable)() {
    m.cardFullAddress = value
}
// SetCardZipOnlyAddress sets the CardZipOnlyAddress property value. Composed type representation for type CardZipOnlyAddressable
func (m *CardRequest_CardRequest_address) SetCardZipOnlyAddress(value CardZipOnlyAddressable)() {
    m.cardZipOnlyAddress = value
}
type CardRequest_CardRequest_account_holder_nameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCardAccountHolderNameBusiness()(CardAccountHolderNameBusinessable)
    GetCardAccountHolderNamePerson()(CardAccountHolderNamePersonable)
    SetCardAccountHolderNameBusiness(value CardAccountHolderNameBusinessable)()
    SetCardAccountHolderNamePerson(value CardAccountHolderNamePersonable)()
}
type CardRequest_CardRequest_addressable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCardFullAddress()(CardFullAddressable)
    GetCardZipOnlyAddress()(CardZipOnlyAddressable)
    SetCardFullAddress(value CardFullAddressable)()
    SetCardZipOnlyAddress(value CardZipOnlyAddressable)()
}
// NewCardRequest instantiates a new CardRequest and sets the default values.
func NewCardRequest()(*CardRequest) {
    m := &CardRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCardRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCardRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCardRequest(), nil
}
// GetAccountHolderName gets the account_holder_name property value. The account_holder_name property
// returns a CardRequest_CardRequest_account_holder_nameable when successful
func (m *CardRequest) GetAccountHolderName()(CardRequest_CardRequest_account_holder_nameable) {
    return m.account_holder_name
}
// GetAccountReferenceId gets the account_reference_id property value. The account_reference_id property
// returns a *string when successful
func (m *CardRequest) GetAccountReferenceId()(*string) {
    return m.account_reference_id
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CardRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The address property
// returns a CardRequest_CardRequest_addressable when successful
func (m *CardRequest) GetAddress()(CardRequest_CardRequest_addressable) {
    return m.address
}
// GetCustomerReferenceId gets the customer_reference_id property value. The customer_reference_id property
// returns a *string when successful
func (m *CardRequest) GetCustomerReferenceId()(*string) {
    return m.customer_reference_id
}
// GetCustomerResourceType gets the customer_resource_type property value. The customer_resource_type property
// returns a *EndCustomerResourceType when successful
func (m *CardRequest) GetCustomerResourceType()(*EndCustomerResourceType) {
    return m.customer_resource_type
}
// GetExpirationMonth gets the expiration_month property value. The expiration_month property
// returns a *int32 when successful
func (m *CardRequest) GetExpirationMonth()(*int32) {
    return m.expiration_month
}
// GetExpirationYear gets the expiration_year property value. The expiration_year property
// returns a *int32 when successful
func (m *CardRequest) GetExpirationYear()(*int32) {
    return m.expiration_year
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CardRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_holder_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCardRequest_CardRequest_account_holder_nameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountHolderName(val.(CardRequest_CardRequest_account_holder_nameable))
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
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCardRequest_CardRequest_addressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(CardRequest_CardRequest_addressable))
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
    res["expiration_month"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationMonth(val)
        }
        return nil
    }
    res["expiration_year"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationYear(val)
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
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. The metadata property
// returns a Metadataable when successful
func (m *CardRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetNumber gets the number property value. The card's number
// returns a *string when successful
func (m *CardRequest) GetNumber()(*string) {
    return m.number
}
// Serialize serializes information the current object
func (m *CardRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("account_holder_name", m.GetAccountHolderName())
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
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
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
        err := writer.WriteInt32Value("expiration_month", m.GetExpirationMonth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("expiration_year", m.GetExpirationYear())
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
        err := writer.WriteStringValue("number", m.GetNumber())
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
func (m *CardRequest) SetAccountHolderName(value CardRequest_CardRequest_account_holder_nameable)() {
    m.account_holder_name = value
}
// SetAccountReferenceId sets the account_reference_id property value. The account_reference_id property
func (m *CardRequest) SetAccountReferenceId(value *string)() {
    m.account_reference_id = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CardRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The address property
func (m *CardRequest) SetAddress(value CardRequest_CardRequest_addressable)() {
    m.address = value
}
// SetCustomerReferenceId sets the customer_reference_id property value. The customer_reference_id property
func (m *CardRequest) SetCustomerReferenceId(value *string)() {
    m.customer_reference_id = value
}
// SetCustomerResourceType sets the customer_resource_type property value. The customer_resource_type property
func (m *CardRequest) SetCustomerResourceType(value *EndCustomerResourceType)() {
    m.customer_resource_type = value
}
// SetExpirationMonth sets the expiration_month property value. The expiration_month property
func (m *CardRequest) SetExpirationMonth(value *int32)() {
    m.expiration_month = value
}
// SetExpirationYear sets the expiration_year property value. The expiration_year property
func (m *CardRequest) SetExpirationYear(value *int32)() {
    m.expiration_year = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *CardRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetNumber sets the number property value. The card's number
func (m *CardRequest) SetNumber(value *string)() {
    m.number = value
}
type CardRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(CardRequest_CardRequest_account_holder_nameable)
    GetAccountReferenceId()(*string)
    GetAddress()(CardRequest_CardRequest_addressable)
    GetCustomerReferenceId()(*string)
    GetCustomerResourceType()(*EndCustomerResourceType)
    GetExpirationMonth()(*int32)
    GetExpirationYear()(*int32)
    GetMetadata()(Metadataable)
    GetNumber()(*string)
    SetAccountHolderName(value CardRequest_CardRequest_account_holder_nameable)()
    SetAccountReferenceId(value *string)()
    SetAddress(value CardRequest_CardRequest_addressable)()
    SetCustomerReferenceId(value *string)()
    SetCustomerResourceType(value *EndCustomerResourceType)()
    SetExpirationMonth(value *int32)()
    SetExpirationYear(value *int32)()
    SetMetadata(value Metadataable)()
    SetNumber(value *string)()
}
