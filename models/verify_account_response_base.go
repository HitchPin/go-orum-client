package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VerifyAccountResponseBase struct {
    // Name of account holder. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
    account_holder_name *string
    // Account number for US bank account. 17 digits maximum.
    account_number *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Status of account control.
    control_status *ControlStatus
    // Timestamp when the resource was created.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Status of account debit.
    debit_status *DebitStatus
    // Status reason for failed and blocked accounts.
    debit_status_reason *DebitStatusReason
    // Email address to notify once the statement code is sent to the account to verify account control.
    email *string
    // Estimated date that the account verification will be complete. Time will always be midnight and should be ignored.
    estimated_verification_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Orum generated unique id for the resource.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Status of account ownership.
    ownership_status *OwnershipStatus
    // Ownership details for a person
    person OwnershipPersonable
    // 9-digit American Bankers Association (ABA) routing number.
    routing_number *string
    // Status reason for failed, invalid, and closed accounts.
    status_reason *StatusReason
    // Timestamp when the resource was last updated.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Status of account verification.
    verification_status *VerificationStatus
}
// NewVerifyAccountResponseBase instantiates a new VerifyAccountResponseBase and sets the default values.
func NewVerifyAccountResponseBase()(*VerifyAccountResponseBase) {
    m := &VerifyAccountResponseBase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVerifyAccountResponseBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVerifyAccountResponseBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifyAccountResponseBase(), nil
}
// GetAccountHolderName gets the account_holder_name property value. Name of account holder. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
// returns a *string when successful
func (m *VerifyAccountResponseBase) GetAccountHolderName()(*string) {
    return m.account_holder_name
}
// GetAccountNumber gets the account_number property value. Account number for US bank account. 17 digits maximum.
// returns a *string when successful
func (m *VerifyAccountResponseBase) GetAccountNumber()(*string) {
    return m.account_number
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VerifyAccountResponseBase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetControlStatus gets the control_status property value. Status of account control.
// returns a *ControlStatus when successful
func (m *VerifyAccountResponseBase) GetControlStatus()(*ControlStatus) {
    return m.control_status
}
// GetCreatedAt gets the created_at property value. Timestamp when the resource was created.
// returns a *Time when successful
func (m *VerifyAccountResponseBase) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDebitStatus gets the debit_status property value. Status of account debit.
// returns a *DebitStatus when successful
func (m *VerifyAccountResponseBase) GetDebitStatus()(*DebitStatus) {
    return m.debit_status
}
// GetDebitStatusReason gets the debit_status_reason property value. Status reason for failed and blocked accounts.
// returns a *DebitStatusReason when successful
func (m *VerifyAccountResponseBase) GetDebitStatusReason()(*DebitStatusReason) {
    return m.debit_status_reason
}
// GetEmail gets the email property value. Email address to notify once the statement code is sent to the account to verify account control.
// returns a *string when successful
func (m *VerifyAccountResponseBase) GetEmail()(*string) {
    return m.email
}
// GetEstimatedVerificationDate gets the estimated_verification_date property value. Estimated date that the account verification will be complete. Time will always be midnight and should be ignored.
// returns a *Time when successful
func (m *VerifyAccountResponseBase) GetEstimatedVerificationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.estimated_verification_date
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VerifyAccountResponseBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["control_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseControlStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlStatus(val.(*ControlStatus))
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
    res["debit_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDebitStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebitStatus(val.(*DebitStatus))
        }
        return nil
    }
    res["debit_status_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDebitStatusReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebitStatusReason(val.(*DebitStatusReason))
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
    res["estimated_verification_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEstimatedVerificationDate(val)
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
    res["ownership_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnershipStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnershipStatus(val.(*OwnershipStatus))
        }
        return nil
    }
    res["person"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOwnershipPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerson(val.(OwnershipPersonable))
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
    res["status_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatusReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusReason(val.(*StatusReason))
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
    res["verification_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVerificationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerificationStatus(val.(*VerificationStatus))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Orum generated unique id for the resource.
// returns a *UUID when successful
func (m *VerifyAccountResponseBase) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetOwnershipStatus gets the ownership_status property value. Status of account ownership.
// returns a *OwnershipStatus when successful
func (m *VerifyAccountResponseBase) GetOwnershipStatus()(*OwnershipStatus) {
    return m.ownership_status
}
// GetPerson gets the person property value. Ownership details for a person
// returns a OwnershipPersonable when successful
func (m *VerifyAccountResponseBase) GetPerson()(OwnershipPersonable) {
    return m.person
}
// GetRoutingNumber gets the routing_number property value. 9-digit American Bankers Association (ABA) routing number.
// returns a *string when successful
func (m *VerifyAccountResponseBase) GetRoutingNumber()(*string) {
    return m.routing_number
}
// GetStatusReason gets the status_reason property value. Status reason for failed, invalid, and closed accounts.
// returns a *StatusReason when successful
func (m *VerifyAccountResponseBase) GetStatusReason()(*StatusReason) {
    return m.status_reason
}
// GetUpdatedAt gets the updated_at property value. Timestamp when the resource was last updated.
// returns a *Time when successful
func (m *VerifyAccountResponseBase) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetVerificationStatus gets the verification_status property value. Status of account verification.
// returns a *VerificationStatus when successful
func (m *VerifyAccountResponseBase) GetVerificationStatus()(*VerificationStatus) {
    return m.verification_status
}
// Serialize serializes information the current object
func (m *VerifyAccountResponseBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetControlStatus() != nil {
        cast := (*m.GetControlStatus()).String()
        err := writer.WriteStringValue("control_status", &cast)
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
    if m.GetDebitStatus() != nil {
        cast := (*m.GetDebitStatus()).String()
        err := writer.WriteStringValue("debit_status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDebitStatusReason() != nil {
        cast := (*m.GetDebitStatusReason()).String()
        err := writer.WriteStringValue("debit_status_reason", &cast)
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
        err := writer.WriteTimeValue("estimated_verification_date", m.GetEstimatedVerificationDate())
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
    if m.GetOwnershipStatus() != nil {
        cast := (*m.GetOwnershipStatus()).String()
        err := writer.WriteStringValue("ownership_status", &cast)
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
    if m.GetStatusReason() != nil {
        cast := (*m.GetStatusReason()).String()
        err := writer.WriteStringValue("status_reason", &cast)
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
    if m.GetVerificationStatus() != nil {
        cast := (*m.GetVerificationStatus()).String()
        err := writer.WriteStringValue("verification_status", &cast)
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
// SetAccountHolderName sets the account_holder_name property value. Name of account holder. Accepts alphanumeric characters and hyphens, dashes, periods, apostrophes, spaces, and diacritics.
func (m *VerifyAccountResponseBase) SetAccountHolderName(value *string)() {
    m.account_holder_name = value
}
// SetAccountNumber sets the account_number property value. Account number for US bank account. 17 digits maximum.
func (m *VerifyAccountResponseBase) SetAccountNumber(value *string)() {
    m.account_number = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifyAccountResponseBase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetControlStatus sets the control_status property value. Status of account control.
func (m *VerifyAccountResponseBase) SetControlStatus(value *ControlStatus)() {
    m.control_status = value
}
// SetCreatedAt sets the created_at property value. Timestamp when the resource was created.
func (m *VerifyAccountResponseBase) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDebitStatus sets the debit_status property value. Status of account debit.
func (m *VerifyAccountResponseBase) SetDebitStatus(value *DebitStatus)() {
    m.debit_status = value
}
// SetDebitStatusReason sets the debit_status_reason property value. Status reason for failed and blocked accounts.
func (m *VerifyAccountResponseBase) SetDebitStatusReason(value *DebitStatusReason)() {
    m.debit_status_reason = value
}
// SetEmail sets the email property value. Email address to notify once the statement code is sent to the account to verify account control.
func (m *VerifyAccountResponseBase) SetEmail(value *string)() {
    m.email = value
}
// SetEstimatedVerificationDate sets the estimated_verification_date property value. Estimated date that the account verification will be complete. Time will always be midnight and should be ignored.
func (m *VerifyAccountResponseBase) SetEstimatedVerificationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.estimated_verification_date = value
}
// SetId sets the id property value. Orum generated unique id for the resource.
func (m *VerifyAccountResponseBase) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetOwnershipStatus sets the ownership_status property value. Status of account ownership.
func (m *VerifyAccountResponseBase) SetOwnershipStatus(value *OwnershipStatus)() {
    m.ownership_status = value
}
// SetPerson sets the person property value. Ownership details for a person
func (m *VerifyAccountResponseBase) SetPerson(value OwnershipPersonable)() {
    m.person = value
}
// SetRoutingNumber sets the routing_number property value. 9-digit American Bankers Association (ABA) routing number.
func (m *VerifyAccountResponseBase) SetRoutingNumber(value *string)() {
    m.routing_number = value
}
// SetStatusReason sets the status_reason property value. Status reason for failed, invalid, and closed accounts.
func (m *VerifyAccountResponseBase) SetStatusReason(value *StatusReason)() {
    m.status_reason = value
}
// SetUpdatedAt sets the updated_at property value. Timestamp when the resource was last updated.
func (m *VerifyAccountResponseBase) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetVerificationStatus sets the verification_status property value. Status of account verification.
func (m *VerifyAccountResponseBase) SetVerificationStatus(value *VerificationStatus)() {
    m.verification_status = value
}
type VerifyAccountResponseBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountHolderName()(*string)
    GetAccountNumber()(*string)
    GetControlStatus()(*ControlStatus)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDebitStatus()(*DebitStatus)
    GetDebitStatusReason()(*DebitStatusReason)
    GetEmail()(*string)
    GetEstimatedVerificationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetOwnershipStatus()(*OwnershipStatus)
    GetPerson()(OwnershipPersonable)
    GetRoutingNumber()(*string)
    GetStatusReason()(*StatusReason)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVerificationStatus()(*VerificationStatus)
    SetAccountHolderName(value *string)()
    SetAccountNumber(value *string)()
    SetControlStatus(value *ControlStatus)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDebitStatus(value *DebitStatus)()
    SetDebitStatusReason(value *DebitStatusReason)()
    SetEmail(value *string)()
    SetEstimatedVerificationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetOwnershipStatus(value *OwnershipStatus)()
    SetPerson(value OwnershipPersonable)()
    SetRoutingNumber(value *string)()
    SetStatusReason(value *StatusReason)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVerificationStatus(value *VerificationStatus)()
}
