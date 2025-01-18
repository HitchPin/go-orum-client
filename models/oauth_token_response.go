package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OauthTokenResponse struct {
    // Short-lived auth (access) token that you will pass in the Authorization header of all Deliver calls.
    access_token *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time in seconds until the token expires.
    expires_in *float64
    // The type of token returned. This will always be "Bearer".
    token_type *string
}
// NewOauthTokenResponse instantiates a new OauthTokenResponse and sets the default values.
func NewOauthTokenResponse()(*OauthTokenResponse) {
    m := &OauthTokenResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOauthTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOauthTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOauthTokenResponse(), nil
}
// GetAccessToken gets the access_token property value. Short-lived auth (access) token that you will pass in the Authorization header of all Deliver calls.
// returns a *string when successful
func (m *OauthTokenResponse) GetAccessToken()(*string) {
    return m.access_token
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OauthTokenResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiresIn gets the expires_in property value. The time in seconds until the token expires.
// returns a *float64 when successful
func (m *OauthTokenResponse) GetExpiresIn()(*float64) {
    return m.expires_in
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OauthTokenResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessToken(val)
        }
        return nil
    }
    res["expires_in"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresIn(val)
        }
        return nil
    }
    res["token_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenType(val)
        }
        return nil
    }
    return res
}
// GetTokenType gets the token_type property value. The type of token returned. This will always be "Bearer".
// returns a *string when successful
func (m *OauthTokenResponse) GetTokenType()(*string) {
    return m.token_type
}
// Serialize serializes information the current object
func (m *OauthTokenResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("access_token", m.GetAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("expires_in", m.GetExpiresIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token_type", m.GetTokenType())
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
// SetAccessToken sets the access_token property value. Short-lived auth (access) token that you will pass in the Authorization header of all Deliver calls.
func (m *OauthTokenResponse) SetAccessToken(value *string)() {
    m.access_token = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OauthTokenResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiresIn sets the expires_in property value. The time in seconds until the token expires.
func (m *OauthTokenResponse) SetExpiresIn(value *float64)() {
    m.expires_in = value
}
// SetTokenType sets the token_type property value. The type of token returned. This will always be "Bearer".
func (m *OauthTokenResponse) SetTokenType(value *string)() {
    m.token_type = value
}
type OauthTokenResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessToken()(*string)
    GetExpiresIn()(*float64)
    GetTokenType()(*string)
    SetAccessToken(value *string)()
    SetExpiresIn(value *float64)()
    SetTokenType(value *string)()
}
