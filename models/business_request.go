package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BusinessRequest struct {
    // The account_holder_name property
    account_holder_name *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Incorporated address information for the business. Note: This field is required for payouts, deposits, or account-to-account transfers. Orum requires a physical address for all registered businesses; aligned with standards mandated by FinCEN and the FDIC. A physical address is a tangible building address, such as a home, office, or office suite. A PO Box is a mailing address, but not a physical address.
    addresses []BusinessPostalAddressRequestable
    // The business_name property
    business_name *string
    // List of contact information for the business. 'email', 'phone', or 'website' are required. A maximum of three persons or businesses can have the same phone number.
    contacts []ContactRequestBusinessable
    // The customer_reference_id property
    customer_reference_id *string
    // The entity_type property
    entity_type *BusinessEntityType
    // The incorporation_date property
    incorporation_date *string
    // The legal_name property
    legal_name *string
    // The metadata property
    metadata Metadataable
    // The tax_id property
    tax_id *string
    // The tax_id_type property
    tax_id_type *TaxIdType
}
// NewBusinessRequest instantiates a new BusinessRequest and sets the default values.
func NewBusinessRequest()(*BusinessRequest) {
    m := &BusinessRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBusinessRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBusinessRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessRequest(), nil
}
// GetAccountHolderName gets the account_holder_name property value. The account_holder_name property
// returns a *string when successful
func (m *BusinessRequest) GetAccountHolderName()(*string) {
    return m.account_holder_name
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BusinessRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddresses gets the addresses property value. Incorporated address information for the business. Note: This field is required for payouts, deposits, or account-to-account transfers. Orum requires a physical address for all registered businesses; aligned with standards mandated by FinCEN and the FDIC. A physical address is a tangible building address, such as a home, office, or office suite. A PO Box is a mailing address, but not a physical address.
// returns a []BusinessPostalAddressRequestable when successful
func (m *BusinessRequest) GetAddresses()([]BusinessPostalAddressRequestable) {
    return m.addresses
}
// GetBusinessName gets the business_name property value. The business_name property
// returns a *string when successful
func (m *BusinessRequest) GetBusinessName()(*string) {
    return m.business_name
}
// GetContacts gets the contacts property value. List of contact information for the business. 'email', 'phone', or 'website' are required. A maximum of three persons or businesses can have the same phone number.
// returns a []ContactRequestBusinessable when successful
func (m *BusinessRequest) GetContacts()([]ContactRequestBusinessable) {
    return m.contacts
}
// GetCustomerReferenceId gets the customer_reference_id property value. The customer_reference_id property
// returns a *string when successful
func (m *BusinessRequest) GetCustomerReferenceId()(*string) {
    return m.customer_reference_id
}
// GetEntityType gets the entity_type property value. The entity_type property
// returns a *BusinessEntityType when successful
func (m *BusinessRequest) GetEntityType()(*BusinessEntityType) {
    return m.entity_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BusinessRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBusinessPostalAddressRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessPostalAddressRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BusinessPostalAddressRequestable)
                }
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["business_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessName(val)
        }
        return nil
    }
    res["contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContactRequestBusinessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContactRequestBusinessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContactRequestBusinessable)
                }
            }
            m.SetContacts(res)
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
    res["entity_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBusinessEntityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityType(val.(*BusinessEntityType))
        }
        return nil
    }
    res["incorporation_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncorporationDate(val)
        }
        return nil
    }
    res["legal_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegalName(val)
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
    res["tax_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxId(val)
        }
        return nil
    }
    res["tax_id_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaxIdType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxIdType(val.(*TaxIdType))
        }
        return nil
    }
    return res
}
// GetIncorporationDate gets the incorporation_date property value. The incorporation_date property
// returns a *string when successful
func (m *BusinessRequest) GetIncorporationDate()(*string) {
    return m.incorporation_date
}
// GetLegalName gets the legal_name property value. The legal_name property
// returns a *string when successful
func (m *BusinessRequest) GetLegalName()(*string) {
    return m.legal_name
}
// GetMetadata gets the metadata property value. The metadata property
// returns a Metadataable when successful
func (m *BusinessRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetTaxId gets the tax_id property value. The tax_id property
// returns a *string when successful
func (m *BusinessRequest) GetTaxId()(*string) {
    return m.tax_id
}
// GetTaxIdType gets the tax_id_type property value. The tax_id_type property
// returns a *TaxIdType when successful
func (m *BusinessRequest) GetTaxIdType()(*TaxIdType) {
    return m.tax_id_type
}
// Serialize serializes information the current object
func (m *BusinessRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("account_holder_name", m.GetAccountHolderName())
        if err != nil {
            return err
        }
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("business_name", m.GetBusinessName())
        if err != nil {
            return err
        }
    }
    if m.GetContacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("contacts", cast)
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
    if m.GetEntityType() != nil {
        cast := (*m.GetEntityType()).String()
        err := writer.WriteStringValue("entity_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("incorporation_date", m.GetIncorporationDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("legal_name", m.GetLegalName())
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
        err := writer.WriteStringValue("tax_id", m.GetTaxId())
        if err != nil {
            return err
        }
    }
    if m.GetTaxIdType() != nil {
        cast := (*m.GetTaxIdType()).String()
        err := writer.WriteStringValue("tax_id_type", &cast)
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
func (m *BusinessRequest) SetAccountHolderName(value *string)() {
    m.account_holder_name = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddresses sets the addresses property value. Incorporated address information for the business. Note: This field is required for payouts, deposits, or account-to-account transfers. Orum requires a physical address for all registered businesses; aligned with standards mandated by FinCEN and the FDIC. A physical address is a tangible building address, such as a home, office, or office suite. A PO Box is a mailing address, but not a physical address.
func (m *BusinessRequest) SetAddresses(value []BusinessPostalAddressRequestable)() {
    m.addresses = value
}
// SetBusinessName sets the business_name property value. The business_name property
func (m *BusinessRequest) SetBusinessName(value *string)() {
    m.business_name = value
}
// SetContacts sets the contacts property value. List of contact information for the business. 'email', 'phone', or 'website' are required. A maximum of three persons or businesses can have the same phone number.
func (m *BusinessRequest) SetContacts(value []ContactRequestBusinessable)() {
    m.contacts = value
}
// SetCustomerReferenceId sets the customer_reference_id property value. The customer_reference_id property
func (m *BusinessRequest) SetCustomerReferenceId(value *string)() {
    m.customer_reference_id = value
}
// SetEntityType sets the entity_type property value. The entity_type property
func (m *BusinessRequest) SetEntityType(value *BusinessEntityType)() {
    m.entity_type = value
}
// SetIncorporationDate sets the incorporation_date property value. The incorporation_date property
func (m *BusinessRequest) SetIncorporationDate(value *string)() {
    m.incorporation_date = value
}
// SetLegalName sets the legal_name property value. The legal_name property
func (m *BusinessRequest) SetLegalName(value *string)() {
    m.legal_name = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *BusinessRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetTaxId sets the tax_id property value. The tax_id property
func (m *BusinessRequest) SetTaxId(value *string)() {
    m.tax_id = value
}
// SetTaxIdType sets the tax_id_type property value. The tax_id_type property
func (m *BusinessRequest) SetTaxIdType(value *TaxIdType)() {
    m.tax_id_type = value
}
type BusinessRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(*string)
    GetAddresses()([]BusinessPostalAddressRequestable)
    GetBusinessName()(*string)
    GetContacts()([]ContactRequestBusinessable)
    GetCustomerReferenceId()(*string)
    GetEntityType()(*BusinessEntityType)
    GetIncorporationDate()(*string)
    GetLegalName()(*string)
    GetMetadata()(Metadataable)
    GetTaxId()(*string)
    GetTaxIdType()(*TaxIdType)
    SetAccountHolderName(value *string)()
    SetAddresses(value []BusinessPostalAddressRequestable)()
    SetBusinessName(value *string)()
    SetContacts(value []ContactRequestBusinessable)()
    SetCustomerReferenceId(value *string)()
    SetEntityType(value *BusinessEntityType)()
    SetIncorporationDate(value *string)()
    SetLegalName(value *string)()
    SetMetadata(value Metadataable)()
    SetTaxId(value *string)()
    SetTaxIdType(value *TaxIdType)()
}
