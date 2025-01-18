package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VerifyAccountRequest struct {
    // The account_holder_name property
    account_holder_name *string
    // The account_number property
    account_number *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email property
    email *string
    // The person property
    person OwnershipPersonRequestable
    // The routing_number property
    routing_number *string
    // The type property
    typeEscaped *AccountType
}
// NewVerifyAccountRequest instantiates a new VerifyAccountRequest and sets the default values.
func NewVerifyAccountRequest()(*VerifyAccountRequest) {
    m := &VerifyAccountRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVerifyAccountRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVerifyAccountRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifyAccountRequest(), nil
}
// GetAccountHolderName gets the account_holder_name property value. The account_holder_name property
// returns a *string when successful
func (m *VerifyAccountRequest) GetAccountHolderName()(*string) {
    return m.account_holder_name
}
// GetAccountNumber gets the account_number property value. The account_number property
// returns a *string when successful
func (m *VerifyAccountRequest) GetAccountNumber()(*string) {
    return m.account_number
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VerifyAccountRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *VerifyAccountRequest) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VerifyAccountRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["person"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOwnershipPersonRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerson(val.(OwnershipPersonRequestable))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AccountType))
        }
        return nil
    }
    return res
}
// GetPerson gets the person property value. The person property
// returns a OwnershipPersonRequestable when successful
func (m *VerifyAccountRequest) GetPerson()(OwnershipPersonRequestable) {
    return m.person
}
// GetRoutingNumber gets the routing_number property value. The routing_number property
// returns a *string when successful
func (m *VerifyAccountRequest) GetRoutingNumber()(*string) {
    return m.routing_number
}
// GetTypeEscaped gets the type property value. The type property
// returns a *AccountType when successful
func (m *VerifyAccountRequest) GetTypeEscaped()(*AccountType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *VerifyAccountRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("person", m.GetPerson())
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *VerifyAccountRequest) SetAccountHolderName(value *string)() {
    m.account_holder_name = value
}
// SetAccountNumber sets the account_number property value. The account_number property
func (m *VerifyAccountRequest) SetAccountNumber(value *string)() {
    m.account_number = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifyAccountRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email property
func (m *VerifyAccountRequest) SetEmail(value *string)() {
    m.email = value
}
// SetPerson sets the person property value. The person property
func (m *VerifyAccountRequest) SetPerson(value OwnershipPersonRequestable)() {
    m.person = value
}
// SetRoutingNumber sets the routing_number property value. The routing_number property
func (m *VerifyAccountRequest) SetRoutingNumber(value *string)() {
    m.routing_number = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *VerifyAccountRequest) SetTypeEscaped(value *AccountType)() {
    m.typeEscaped = value
}
type VerifyAccountRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(*string)
    GetAccountNumber()(*string)
    GetEmail()(*string)
    GetPerson()(OwnershipPersonRequestable)
    GetRoutingNumber()(*string)
    GetTypeEscaped()(*AccountType)
    SetAccountHolderName(value *string)()
    SetAccountNumber(value *string)()
    SetEmail(value *string)()
    SetPerson(value OwnershipPersonRequestable)()
    SetRoutingNumber(value *string)()
    SetTypeEscaped(value *AccountType)()
}
