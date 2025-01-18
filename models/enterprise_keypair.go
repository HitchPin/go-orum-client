package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseKeypair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *string
    // The enterprise_name property
    enterprise_name *string
    // The id property
    id *string
    // The public_key property
    public_key *string
    // The updated_at property
    updated_at *string
}
// NewEnterpriseKeypair instantiates a new EnterpriseKeypair and sets the default values.
func NewEnterpriseKeypair()(*EnterpriseKeypair) {
    m := &EnterpriseKeypair{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseKeypairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseKeypairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseKeypair(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseKeypair) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *EnterpriseKeypair) GetCreatedAt()(*string) {
    return m.created_at
}
// GetEnterpriseName gets the enterprise_name property value. The enterprise_name property
// returns a *string when successful
func (m *EnterpriseKeypair) GetEnterpriseName()(*string) {
    return m.enterprise_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseKeypair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["enterprise_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["public_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKey(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *EnterpriseKeypair) GetId()(*string) {
    return m.id
}
// GetPublicKey gets the public_key property value. The public_key property
// returns a *string when successful
func (m *EnterpriseKeypair) GetPublicKey()(*string) {
    return m.public_key
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *EnterpriseKeypair) GetUpdatedAt()(*string) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *EnterpriseKeypair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enterprise_name", m.GetEnterpriseName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("public_key", m.GetPublicKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
func (m *EnterpriseKeypair) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *EnterpriseKeypair) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetEnterpriseName sets the enterprise_name property value. The enterprise_name property
func (m *EnterpriseKeypair) SetEnterpriseName(value *string)() {
    m.enterprise_name = value
}
// SetId sets the id property value. The id property
func (m *EnterpriseKeypair) SetId(value *string)() {
    m.id = value
}
// SetPublicKey sets the public_key property value. The public_key property
func (m *EnterpriseKeypair) SetPublicKey(value *string)() {
    m.public_key = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *EnterpriseKeypair) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
type EnterpriseKeypairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetEnterpriseName()(*string)
    GetId()(*string)
    GetPublicKey()(*string)
    GetUpdatedAt()(*string)
    SetCreatedAt(value *string)()
    SetEnterpriseName(value *string)()
    SetId(value *string)()
    SetPublicKey(value *string)()
    SetUpdatedAt(value *string)()
}
