// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationseteventdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sesv2configurationseteventdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference interface {
	cdktf.ComplexObject
	CloudWatchDestination() Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationOutputReference
	CloudWatchDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestination
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventBridgeDestination() Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestinationOutputReference
	EventBridgeDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestination
	// Experimental.
	Fqn() *string
	InternalValue() *Sesv2ConfigurationSetEventDestinationEventDestination
	SetInternalValue(val *Sesv2ConfigurationSetEventDestinationEventDestination)
	KinesisFirehoseDestination() Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationOutputReference
	KinesisFirehoseDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination
	MatchingEventTypes() *[]*string
	SetMatchingEventTypes(val *[]*string)
	MatchingEventTypesInput() *[]*string
	PinpointDestination() Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestinationOutputReference
	PinpointDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestination
	SnsDestination() Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestinationOutputReference
	SnsDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestination
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutCloudWatchDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestination)
	PutEventBridgeDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestination)
	PutKinesisFirehoseDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination)
	PutPinpointDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestination)
	PutSnsDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestination)
	ResetCloudWatchDestination()
	ResetEnabled()
	ResetEventBridgeDestination()
	ResetKinesisFirehoseDestination()
	ResetPinpointDestination()
	ResetSnsDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference
type jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) CloudWatchDestination() Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationOutputReference {
	var returns Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationOutputReference
	_jsii_.Get(
		j,
		"cloudWatchDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) CloudWatchDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestination {
	var returns *Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestination
	_jsii_.Get(
		j,
		"cloudWatchDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) EventBridgeDestination() Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestinationOutputReference {
	var returns Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestinationOutputReference
	_jsii_.Get(
		j,
		"eventBridgeDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) EventBridgeDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestination {
	var returns *Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestination
	_jsii_.Get(
		j,
		"eventBridgeDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) InternalValue() *Sesv2ConfigurationSetEventDestinationEventDestination {
	var returns *Sesv2ConfigurationSetEventDestinationEventDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) KinesisFirehoseDestination() Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationOutputReference {
	var returns Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) KinesisFirehoseDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination {
	var returns *Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination
	_jsii_.Get(
		j,
		"kinesisFirehoseDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) MatchingEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) MatchingEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) PinpointDestination() Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestinationOutputReference {
	var returns Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestinationOutputReference
	_jsii_.Get(
		j,
		"pinpointDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) PinpointDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestination {
	var returns *Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestination
	_jsii_.Get(
		j,
		"pinpointDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) SnsDestination() Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestinationOutputReference {
	var returns Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestinationOutputReference
	_jsii_.Get(
		j,
		"snsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) SnsDestinationInput() *Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestination {
	var returns *Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestination
	_jsii_.Get(
		j,
		"snsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesv2ConfigurationSetEventDestinationEventDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewSesv2ConfigurationSetEventDestinationEventDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sesv2ConfigurationSetEventDestination.Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesv2ConfigurationSetEventDestinationEventDestinationOutputReference_Override(s Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sesv2ConfigurationSetEventDestination.Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference)SetInternalValue(val *Sesv2ConfigurationSetEventDestinationEventDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference)SetMatchingEventTypes(val *[]*string) {
	if err := j.validateSetMatchingEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingEventTypes",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) PutCloudWatchDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationCloudWatchDestination) {
	if err := s.validatePutCloudWatchDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudWatchDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) PutEventBridgeDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationEventBridgeDestination) {
	if err := s.validatePutEventBridgeDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEventBridgeDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) PutKinesisFirehoseDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination) {
	if err := s.validatePutKinesisFirehoseDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKinesisFirehoseDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) PutPinpointDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationPinpointDestination) {
	if err := s.validatePutPinpointDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPinpointDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) PutSnsDestination(value *Sesv2ConfigurationSetEventDestinationEventDestinationSnsDestination) {
	if err := s.validatePutSnsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSnsDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ResetCloudWatchDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudWatchDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ResetEventBridgeDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetEventBridgeDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ResetKinesisFirehoseDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetKinesisFirehoseDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ResetPinpointDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetPinpointDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ResetSnsDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetSnsDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2ConfigurationSetEventDestinationEventDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

