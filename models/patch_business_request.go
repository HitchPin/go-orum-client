package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PatchBusinessRequest struct {
    // The account_holder_name property
    account_holder_name *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Incorporated address information for the business. "type" is required if you are updating other address fields. Note: This field is required for payouts, deposits, or account-to-account transfers. Orum requires a physical address for all registered businesses; aligned with standards mandated by FinCEN and the FDIC. A physical address is a tangible building address, such as a home, office, or office suite. A PO Box is a mailing address, but not a physical address.
    addresses []BusinessPatchPostalAddressRequestable
    // The business_name property
    business_name *string
    // List of contact information for the business. 'email', 'phone', or 'website' are required. A maximum of three persons or businesses can have the same phone number.
    contacts []PatchContactRequestBusinessable
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
// NewPatchBusinessRequest instantiates a new PatchBusinessRequest and sets the default values.
func NewPatchBusinessRequest()(*PatchBusinessRequest) {
    m := &PatchBusinessRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatchBusinessRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePatchBusinessRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatchBusinessRequest(), nil
}
// GetAccountHolderName gets the account_holder_name property value. The account_holder_name property
// returns a *string when successful
func (m *PatchBusinessRequest) GetAccountHolderName()(*string) {
    return m.account_holder_name
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PatchBusinessRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddresses gets the addresses property value. Incorporated address information for the business. "type" is required if you are updating other address fields. Note: This field is required for payouts, deposits, or account-to-account transfers. Orum requires a physical address for all registered businesses; aligned with standards mandated by FinCEN and the FDIC. A physical address is a tangible building address, such as a home, office, or office suite. A PO Box is a mailing address, but not a physical address.
// returns a []BusinessPatchPostalAddressRequestable when successful
func (m *PatchBusinessRequest) GetAddresses()([]BusinessPatchPostalAddressRequestable) {
    return m.addresses
}
// GetBusinessName gets the business_name property value. The business_name property
// returns a *string when successful
func (m *PatchBusinessRequest) GetBusinessName()(*string) {
    return m.business_name
}
// GetContacts gets the contacts property value. List of contact information for the business. 'email', 'phone', or 'website' are required. A maximum of three persons or businesses can have the same phone number.
// returns a []PatchContactRequestBusinessable when successful
func (m *PatchBusinessRequest) GetContacts()([]PatchContactRequestBusinessable) {
    return m.contacts
}
// GetEntityType gets the entity_type property value. The entity_type property
// returns a *BusinessEntityType when successful
func (m *PatchBusinessRequest) GetEntityType()(*BusinessEntityType) {
    return m.entity_type
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PatchBusinessRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateBusinessPatchPostalAddressRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessPatchPostalAddressRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BusinessPatchPostalAddressRequestable)
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
        val, err := n.GetCollectionOfObjectValues(CreatePatchContactRequestBusinessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PatchContactRequestBusinessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PatchContactRequestBusinessable)
                }
            }
            m.SetContacts(res)
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
func (m *PatchBusinessRequest) GetIncorporationDate()(*string) {
    return m.incorporation_date
}
// GetLegalName gets the legal_name property value. The legal_name property
// returns a *string when successful
func (m *PatchBusinessRequest) GetLegalName()(*string) {
    return m.legal_name
}
// GetMetadata gets the metadata property value. The metadata property
// returns a Metadataable when successful
func (m *PatchBusinessRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetTaxId gets the tax_id property value. The tax_id property
// returns a *string when successful
func (m *PatchBusinessRequest) GetTaxId()(*string) {
    return m.tax_id
}
// GetTaxIdType gets the tax_id_type property value. The tax_id_type property
// returns a *TaxIdType when successful
func (m *PatchBusinessRequest) GetTaxIdType()(*TaxIdType) {
    return m.tax_id_type
}
// Serialize serializes information the current object
func (m *PatchBusinessRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PatchBusinessRequest) SetAccountHolderName(value *string)() {
    m.account_holder_name = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PatchBusinessRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddresses sets the addresses property value. Incorporated address information for the business. "type" is required if you are updating other address fields. Note: This field is required for payouts, deposits, or account-to-account transfers. Orum requires a physical address for all registered businesses; aligned with standards mandated by FinCEN and the FDIC. A physical address is a tangible building address, such as a home, office, or office suite. A PO Box is a mailing address, but not a physical address.
func (m *PatchBusinessRequest) SetAddresses(value []BusinessPatchPostalAddressRequestable)() {
    m.addresses = value
}
// SetBusinessName sets the business_name property value. The business_name property
func (m *PatchBusinessRequest) SetBusinessName(value *string)() {
    m.business_name = value
}
// SetContacts sets the contacts property value. List of contact information for the business. 'email', 'phone', or 'website' are required. A maximum of three persons or businesses can have the same phone number.
func (m *PatchBusinessRequest) SetContacts(value []PatchContactRequestBusinessable)() {
    m.contacts = value
}
// SetEntityType sets the entity_type property value. The entity_type property
func (m *PatchBusinessRequest) SetEntityType(value *BusinessEntityType)() {
    m.entity_type = value
}
// SetIncorporationDate sets the incorporation_date property value. The incorporation_date property
func (m *PatchBusinessRequest) SetIncorporationDate(value *string)() {
    m.incorporation_date = value
}
// SetLegalName sets the legal_name property value. The legal_name property
func (m *PatchBusinessRequest) SetLegalName(value *string)() {
    m.legal_name = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *PatchBusinessRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetTaxId sets the tax_id property value. The tax_id property
func (m *PatchBusinessRequest) SetTaxId(value *string)() {
    m.tax_id = value
}
// SetTaxIdType sets the tax_id_type property value. The tax_id_type property
func (m *PatchBusinessRequest) SetTaxIdType(value *TaxIdType)() {
    m.tax_id_type = value
}
type PatchBusinessRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(*string)
    GetAddresses()([]BusinessPatchPostalAddressRequestable)
    GetBusinessName()(*string)
    GetContacts()([]PatchContactRequestBusinessable)
    GetEntityType()(*BusinessEntityType)
    GetIncorporationDate()(*string)
    GetLegalName()(*string)
    GetMetadata()(Metadataable)
    GetTaxId()(*string)
    GetTaxIdType()(*TaxIdType)
    SetAccountHolderName(value *string)()
    SetAddresses(value []BusinessPatchPostalAddressRequestable)()
    SetBusinessName(value *string)()
    SetContacts(value []PatchContactRequestBusinessable)()
    SetEntityType(value *BusinessEntityType)()
    SetIncorporationDate(value *string)()
    SetLegalName(value *string)()
    SetMetadata(value Metadataable)()
    SetTaxId(value *string)()
    SetTaxIdType(value *TaxIdType)()
}
