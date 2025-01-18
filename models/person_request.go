package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PersonRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Address information for the person. Note: This field is required for payouts, deposits, or account-to-account transfers.
    addresses []PersonPostalAddressRequestable
    // Optional list of contact information for the person. If object is present, either phone and/or email contact types are accepted.  A maximum of three persons or businesses can have the same phone number.
    contacts []ContactRequestPersonable
    // The customer_reference_id property
    customer_reference_id *string
    // The date_of_birth property
    date_of_birth *string
    // The first_name property
    first_name *string
    // The last_name property
    last_name *string
    // The metadata property
    metadata Metadataable
    // The middle_name property
    middle_name *string
    // The social_security_number property
    social_security_number *string
}
// NewPersonRequest instantiates a new PersonRequest and sets the default values.
func NewPersonRequest()(*PersonRequest) {
    m := &PersonRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePersonRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePersonRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PersonRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddresses gets the addresses property value. Address information for the person. Note: This field is required for payouts, deposits, or account-to-account transfers.
// returns a []PersonPostalAddressRequestable when successful
func (m *PersonRequest) GetAddresses()([]PersonPostalAddressRequestable) {
    return m.addresses
}
// GetContacts gets the contacts property value. Optional list of contact information for the person. If object is present, either phone and/or email contact types are accepted.  A maximum of three persons or businesses can have the same phone number.
// returns a []ContactRequestPersonable when successful
func (m *PersonRequest) GetContacts()([]ContactRequestPersonable) {
    return m.contacts
}
// GetCustomerReferenceId gets the customer_reference_id property value. The customer_reference_id property
// returns a *string when successful
func (m *PersonRequest) GetCustomerReferenceId()(*string) {
    return m.customer_reference_id
}
// GetDateOfBirth gets the date_of_birth property value. The date_of_birth property
// returns a *string when successful
func (m *PersonRequest) GetDateOfBirth()(*string) {
    return m.date_of_birth
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PersonRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonPostalAddressRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonPostalAddressRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PersonPostalAddressRequestable)
                }
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContactRequestPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContactRequestPersonable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContactRequestPersonable)
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
    res["date_of_birth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateOfBirth(val)
        }
        return nil
    }
    res["first_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["last_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
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
    res["middle_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["social_security_number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSocialSecurityNumber(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *PersonRequest) GetFirstName()(*string) {
    return m.first_name
}
// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *PersonRequest) GetLastName()(*string) {
    return m.last_name
}
// GetMetadata gets the metadata property value. The metadata property
// returns a Metadataable when successful
func (m *PersonRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *PersonRequest) GetMiddleName()(*string) {
    return m.middle_name
}
// GetSocialSecurityNumber gets the social_security_number property value. The social_security_number property
// returns a *string when successful
func (m *PersonRequest) GetSocialSecurityNumber()(*string) {
    return m.social_security_number
}
// Serialize serializes information the current object
func (m *PersonRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("date_of_birth", m.GetDateOfBirth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("first_name", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_name", m.GetLastName())
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
        err := writer.WriteStringValue("middle_name", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("social_security_number", m.GetSocialSecurityNumber())
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
func (m *PersonRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddresses sets the addresses property value. Address information for the person. Note: This field is required for payouts, deposits, or account-to-account transfers.
func (m *PersonRequest) SetAddresses(value []PersonPostalAddressRequestable)() {
    m.addresses = value
}
// SetContacts sets the contacts property value. Optional list of contact information for the person. If object is present, either phone and/or email contact types are accepted.  A maximum of three persons or businesses can have the same phone number.
func (m *PersonRequest) SetContacts(value []ContactRequestPersonable)() {
    m.contacts = value
}
// SetCustomerReferenceId sets the customer_reference_id property value. The customer_reference_id property
func (m *PersonRequest) SetCustomerReferenceId(value *string)() {
    m.customer_reference_id = value
}
// SetDateOfBirth sets the date_of_birth property value. The date_of_birth property
func (m *PersonRequest) SetDateOfBirth(value *string)() {
    m.date_of_birth = value
}
// SetFirstName sets the first_name property value. The first_name property
func (m *PersonRequest) SetFirstName(value *string)() {
    m.first_name = value
}
// SetLastName sets the last_name property value. The last_name property
func (m *PersonRequest) SetLastName(value *string)() {
    m.last_name = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *PersonRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetMiddleName sets the middle_name property value. The middle_name property
func (m *PersonRequest) SetMiddleName(value *string)() {
    m.middle_name = value
}
// SetSocialSecurityNumber sets the social_security_number property value. The social_security_number property
func (m *PersonRequest) SetSocialSecurityNumber(value *string)() {
    m.social_security_number = value
}
type PersonRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddresses()([]PersonPostalAddressRequestable)
    GetContacts()([]ContactRequestPersonable)
    GetCustomerReferenceId()(*string)
    GetDateOfBirth()(*string)
    GetFirstName()(*string)
    GetLastName()(*string)
    GetMetadata()(Metadataable)
    GetMiddleName()(*string)
    GetSocialSecurityNumber()(*string)
    SetAddresses(value []PersonPostalAddressRequestable)()
    SetContacts(value []ContactRequestPersonable)()
    SetCustomerReferenceId(value *string)()
    SetDateOfBirth(value *string)()
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetMetadata(value Metadataable)()
    SetMiddleName(value *string)()
    SetSocialSecurityNumber(value *string)()
}
