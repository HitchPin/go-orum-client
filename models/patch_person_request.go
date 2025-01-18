package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PatchPersonRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Address information for the person. "type" is required if you are updating other address fields. Note: This field is required for payouts, deposits, or account-to-account transfers.
    addresses []PersonPatchPostalAddressRequestable
    // Optional list of contact information for the person. If object is present, either phone and/or email contact types are accepted.  A maximum of three persons or businesses can have the same phone number.
    contacts []PatchContactRequestPersonable
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
// NewPatchPersonRequest instantiates a new PatchPersonRequest and sets the default values.
func NewPatchPersonRequest()(*PatchPersonRequest) {
    m := &PatchPersonRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatchPersonRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePatchPersonRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatchPersonRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PatchPersonRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddresses gets the addresses property value. Address information for the person. "type" is required if you are updating other address fields. Note: This field is required for payouts, deposits, or account-to-account transfers.
// returns a []PersonPatchPostalAddressRequestable when successful
func (m *PatchPersonRequest) GetAddresses()([]PersonPatchPostalAddressRequestable) {
    return m.addresses
}
// GetContacts gets the contacts property value. Optional list of contact information for the person. If object is present, either phone and/or email contact types are accepted.  A maximum of three persons or businesses can have the same phone number.
// returns a []PatchContactRequestPersonable when successful
func (m *PatchPersonRequest) GetContacts()([]PatchContactRequestPersonable) {
    return m.contacts
}
// GetDateOfBirth gets the date_of_birth property value. The date_of_birth property
// returns a *string when successful
func (m *PatchPersonRequest) GetDateOfBirth()(*string) {
    return m.date_of_birth
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PatchPersonRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonPatchPostalAddressRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonPatchPostalAddressRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PersonPatchPostalAddressRequestable)
                }
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePatchContactRequestPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PatchContactRequestPersonable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PatchContactRequestPersonable)
                }
            }
            m.SetContacts(res)
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
func (m *PatchPersonRequest) GetFirstName()(*string) {
    return m.first_name
}
// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *PatchPersonRequest) GetLastName()(*string) {
    return m.last_name
}
// GetMetadata gets the metadata property value. The metadata property
// returns a Metadataable when successful
func (m *PatchPersonRequest) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *PatchPersonRequest) GetMiddleName()(*string) {
    return m.middle_name
}
// GetSocialSecurityNumber gets the social_security_number property value. The social_security_number property
// returns a *string when successful
func (m *PatchPersonRequest) GetSocialSecurityNumber()(*string) {
    return m.social_security_number
}
// Serialize serializes information the current object
func (m *PatchPersonRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PatchPersonRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddresses sets the addresses property value. Address information for the person. "type" is required if you are updating other address fields. Note: This field is required for payouts, deposits, or account-to-account transfers.
func (m *PatchPersonRequest) SetAddresses(value []PersonPatchPostalAddressRequestable)() {
    m.addresses = value
}
// SetContacts sets the contacts property value. Optional list of contact information for the person. If object is present, either phone and/or email contact types are accepted.  A maximum of three persons or businesses can have the same phone number.
func (m *PatchPersonRequest) SetContacts(value []PatchContactRequestPersonable)() {
    m.contacts = value
}
// SetDateOfBirth sets the date_of_birth property value. The date_of_birth property
func (m *PatchPersonRequest) SetDateOfBirth(value *string)() {
    m.date_of_birth = value
}
// SetFirstName sets the first_name property value. The first_name property
func (m *PatchPersonRequest) SetFirstName(value *string)() {
    m.first_name = value
}
// SetLastName sets the last_name property value. The last_name property
func (m *PatchPersonRequest) SetLastName(value *string)() {
    m.last_name = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *PatchPersonRequest) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetMiddleName sets the middle_name property value. The middle_name property
func (m *PatchPersonRequest) SetMiddleName(value *string)() {
    m.middle_name = value
}
// SetSocialSecurityNumber sets the social_security_number property value. The social_security_number property
func (m *PatchPersonRequest) SetSocialSecurityNumber(value *string)() {
    m.social_security_number = value
}
type PatchPersonRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddresses()([]PersonPatchPostalAddressRequestable)
    GetContacts()([]PatchContactRequestPersonable)
    GetDateOfBirth()(*string)
    GetFirstName()(*string)
    GetLastName()(*string)
    GetMetadata()(Metadataable)
    GetMiddleName()(*string)
    GetSocialSecurityNumber()(*string)
    SetAddresses(value []PersonPatchPostalAddressRequestable)()
    SetContacts(value []PatchContactRequestPersonable)()
    SetDateOfBirth(value *string)()
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetMetadata(value Metadataable)()
    SetMiddleName(value *string)()
    SetSocialSecurityNumber(value *string)()
}
