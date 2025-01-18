package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonPostalAddressRequest address.
type PersonPostalAddressRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address1 property
    address1 *string
    // The address2 property
    address2 *string
    // The city property
    city *string
    // The country property
    country *Country
    // The state property
    state *string
    // The type property
    typeEscaped *AddressTypePerson
    // The zip5 property
    zip5 *string
}
// NewPersonPostalAddressRequest instantiates a new PersonPostalAddressRequest and sets the default values.
func NewPersonPostalAddressRequest()(*PersonPostalAddressRequest) {
    m := &PersonPostalAddressRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePersonPostalAddressRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePersonPostalAddressRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonPostalAddressRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PersonPostalAddressRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress1 gets the address1 property value. The address1 property
// returns a *string when successful
func (m *PersonPostalAddressRequest) GetAddress1()(*string) {
    return m.address1
}
// GetAddress2 gets the address2 property value. The address2 property
// returns a *string when successful
func (m *PersonPostalAddressRequest) GetAddress2()(*string) {
    return m.address2
}
// GetCity gets the city property value. The city property
// returns a *string when successful
func (m *PersonPostalAddressRequest) GetCity()(*string) {
    return m.city
}
// GetCountry gets the country property value. The country property
// returns a *Country when successful
func (m *PersonPostalAddressRequest) GetCountry()(*Country) {
    return m.country
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PersonPostalAddressRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress1(val)
        }
        return nil
    }
    res["address2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress2(val)
        }
        return nil
    }
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCountry)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val.(*Country))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAddressTypePerson)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AddressTypePerson))
        }
        return nil
    }
    res["zip5"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZip5(val)
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *PersonPostalAddressRequest) GetState()(*string) {
    return m.state
}
// GetTypeEscaped gets the type property value. The type property
// returns a *AddressTypePerson when successful
func (m *PersonPostalAddressRequest) GetTypeEscaped()(*AddressTypePerson) {
    return m.typeEscaped
}
// GetZip5 gets the zip5 property value. The zip5 property
// returns a *string when successful
func (m *PersonPostalAddressRequest) GetZip5()(*string) {
    return m.zip5
}
// Serialize serializes information the current object
func (m *PersonPostalAddressRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address1", m.GetAddress1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("address2", m.GetAddress2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    if m.GetCountry() != nil {
        cast := (*m.GetCountry()).String()
        err := writer.WriteStringValue("country", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("zip5", m.GetZip5())
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
func (m *PersonPostalAddressRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress1 sets the address1 property value. The address1 property
func (m *PersonPostalAddressRequest) SetAddress1(value *string)() {
    m.address1 = value
}
// SetAddress2 sets the address2 property value. The address2 property
func (m *PersonPostalAddressRequest) SetAddress2(value *string)() {
    m.address2 = value
}
// SetCity sets the city property value. The city property
func (m *PersonPostalAddressRequest) SetCity(value *string)() {
    m.city = value
}
// SetCountry sets the country property value. The country property
func (m *PersonPostalAddressRequest) SetCountry(value *Country)() {
    m.country = value
}
// SetState sets the state property value. The state property
func (m *PersonPostalAddressRequest) SetState(value *string)() {
    m.state = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *PersonPostalAddressRequest) SetTypeEscaped(value *AddressTypePerson)() {
    m.typeEscaped = value
}
// SetZip5 sets the zip5 property value. The zip5 property
func (m *PersonPostalAddressRequest) SetZip5(value *string)() {
    m.zip5 = value
}
type PersonPostalAddressRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress1()(*string)
    GetAddress2()(*string)
    GetCity()(*string)
    GetCountry()(*Country)
    GetState()(*string)
    GetTypeEscaped()(*AddressTypePerson)
    GetZip5()(*string)
    SetAddress1(value *string)()
    SetAddress2(value *string)()
    SetCity(value *string)()
    SetCountry(value *Country)()
    SetState(value *string)()
    SetTypeEscaped(value *AddressTypePerson)()
    SetZip5(value *string)()
}
