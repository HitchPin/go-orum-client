package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostalAddressResponse address.
type PostalAddressResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Address line 1.
    address1 *string
    // Address line 2.
    address2 *string
    // City.
    city *string
    // 2-character ISO country code.
    country *Country
    // Timestamp when the resource was created.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Orum generated unique id for the resource.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Uppercase two-character state code of the address.
    state *string
    // Type of address.
    typeEscaped *AddressType
    // Timestamp when the resource was last updated.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // 5-digit ZIP Code. Numeric values (0-9) only.
    zip5 *string
}
// NewPostalAddressResponse instantiates a new PostalAddressResponse and sets the default values.
func NewPostalAddressResponse()(*PostalAddressResponse) {
    m := &PostalAddressResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePostalAddressResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePostalAddressResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPostalAddressResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PostalAddressResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress1 gets the address1 property value. Address line 1.
// returns a *string when successful
func (m *PostalAddressResponse) GetAddress1()(*string) {
    return m.address1
}
// GetAddress2 gets the address2 property value. Address line 2.
// returns a *string when successful
func (m *PostalAddressResponse) GetAddress2()(*string) {
    return m.address2
}
// GetCity gets the city property value. City.
// returns a *string when successful
func (m *PostalAddressResponse) GetCity()(*string) {
    return m.city
}
// GetCountry gets the country property value. 2-character ISO country code.
// returns a *Country when successful
func (m *PostalAddressResponse) GetCountry()(*Country) {
    return m.country
}
// GetCreatedAt gets the created_at property value. Timestamp when the resource was created.
// returns a *Time when successful
func (m *PostalAddressResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PostalAddressResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseAddressType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AddressType))
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
// GetId gets the id property value. Orum generated unique id for the resource.
// returns a *UUID when successful
func (m *PostalAddressResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetState gets the state property value. Uppercase two-character state code of the address.
// returns a *string when successful
func (m *PostalAddressResponse) GetState()(*string) {
    return m.state
}
// GetTypeEscaped gets the type property value. Type of address.
// returns a *AddressType when successful
func (m *PostalAddressResponse) GetTypeEscaped()(*AddressType) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updated_at property value. Timestamp when the resource was last updated.
// returns a *Time when successful
func (m *PostalAddressResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetZip5 gets the zip5 property value. 5-digit ZIP Code. Numeric values (0-9) only.
// returns a *string when successful
func (m *PostalAddressResponse) GetZip5()(*string) {
    return m.zip5
}
// Serialize serializes information the current object
func (m *PostalAddressResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
func (m *PostalAddressResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress1 sets the address1 property value. Address line 1.
func (m *PostalAddressResponse) SetAddress1(value *string)() {
    m.address1 = value
}
// SetAddress2 sets the address2 property value. Address line 2.
func (m *PostalAddressResponse) SetAddress2(value *string)() {
    m.address2 = value
}
// SetCity sets the city property value. City.
func (m *PostalAddressResponse) SetCity(value *string)() {
    m.city = value
}
// SetCountry sets the country property value. 2-character ISO country code.
func (m *PostalAddressResponse) SetCountry(value *Country)() {
    m.country = value
}
// SetCreatedAt sets the created_at property value. Timestamp when the resource was created.
func (m *PostalAddressResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetId sets the id property value. Orum generated unique id for the resource.
func (m *PostalAddressResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetState sets the state property value. Uppercase two-character state code of the address.
func (m *PostalAddressResponse) SetState(value *string)() {
    m.state = value
}
// SetTypeEscaped sets the type property value. Type of address.
func (m *PostalAddressResponse) SetTypeEscaped(value *AddressType)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updated_at property value. Timestamp when the resource was last updated.
func (m *PostalAddressResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetZip5 sets the zip5 property value. 5-digit ZIP Code. Numeric values (0-9) only.
func (m *PostalAddressResponse) SetZip5(value *string)() {
    m.zip5 = value
}
type PostalAddressResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress1()(*string)
    GetAddress2()(*string)
    GetCity()(*string)
    GetCountry()(*Country)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetState()(*string)
    GetTypeEscaped()(*AddressType)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetZip5()(*string)
    SetAddress1(value *string)()
    SetAddress2(value *string)()
    SetCity(value *string)()
    SetCountry(value *Country)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetState(value *string)()
    SetTypeEscaped(value *AddressType)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetZip5(value *string)()
}
