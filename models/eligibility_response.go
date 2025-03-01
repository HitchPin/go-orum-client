package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EligibilityResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The routing_numbers property
    routing_numbers []EligibilityRoutingNumberable
}
// NewEligibilityResponse instantiates a new EligibilityResponse and sets the default values.
func NewEligibilityResponse()(*EligibilityResponse) {
    m := &EligibilityResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEligibilityResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEligibilityResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEligibilityResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EligibilityResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EligibilityResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["routing_numbers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEligibilityRoutingNumberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EligibilityRoutingNumberable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EligibilityRoutingNumberable)
                }
            }
            m.SetRoutingNumbers(res)
        }
        return nil
    }
    return res
}
// GetRoutingNumbers gets the routing_numbers property value. The routing_numbers property
// returns a []EligibilityRoutingNumberable when successful
func (m *EligibilityResponse) GetRoutingNumbers()([]EligibilityRoutingNumberable) {
    return m.routing_numbers
}
// Serialize serializes information the current object
func (m *EligibilityResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRoutingNumbers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoutingNumbers()))
        for i, v := range m.GetRoutingNumbers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("routing_numbers", cast)
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
func (m *EligibilityResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRoutingNumbers sets the routing_numbers property value. The routing_numbers property
func (m *EligibilityResponse) SetRoutingNumbers(value []EligibilityRoutingNumberable)() {
    m.routing_numbers = value
}
type EligibilityResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoutingNumbers()([]EligibilityRoutingNumberable)
    SetRoutingNumbers(value []EligibilityRoutingNumberable)()
}
