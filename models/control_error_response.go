package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ControlErrorResponse struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attempts_remaining property
    attempts_remaining *int32
    // The error_code property
    error_code *string
    // The message property
    message *string
}
// NewControlErrorResponse instantiates a new ControlErrorResponse and sets the default values.
func NewControlErrorResponse()(*ControlErrorResponse) {
    m := &ControlErrorResponse{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateControlErrorResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateControlErrorResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewControlErrorResponse(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *ControlErrorResponse) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ControlErrorResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttemptsRemaining gets the attempts_remaining property value. The attempts_remaining property
// returns a *int32 when successful
func (m *ControlErrorResponse) GetAttemptsRemaining()(*int32) {
    return m.attempts_remaining
}
// GetErrorCode gets the error_code property value. The error_code property
// returns a *string when successful
func (m *ControlErrorResponse) GetErrorCode()(*string) {
    return m.error_code
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ControlErrorResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attempts_remaining"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttemptsRemaining(val)
        }
        return nil
    }
    res["error_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *ControlErrorResponse) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *ControlErrorResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("attempts_remaining", m.GetAttemptsRemaining())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error_code", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *ControlErrorResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttemptsRemaining sets the attempts_remaining property value. The attempts_remaining property
func (m *ControlErrorResponse) SetAttemptsRemaining(value *int32)() {
    m.attempts_remaining = value
}
// SetErrorCode sets the error_code property value. The error_code property
func (m *ControlErrorResponse) SetErrorCode(value *string)() {
    m.error_code = value
}
// SetMessage sets the message property value. The message property
func (m *ControlErrorResponse) SetMessage(value *string)() {
    m.message = value
}
type ControlErrorResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttemptsRemaining()(*int32)
    GetErrorCode()(*string)
    GetMessage()(*string)
    SetAttemptsRemaining(value *int32)()
    SetErrorCode(value *string)()
    SetMessage(value *string)()
}
