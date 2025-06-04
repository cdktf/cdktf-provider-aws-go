// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/launchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *LaunchTemplateNetworkInterfacesConnectionTrackingSpecification
	SetInternalValue(val *LaunchTemplateNetworkInterfacesConnectionTrackingSpecification)
	TcpEstablishedTimeout() *float64
	SetTcpEstablishedTimeout(val *float64)
	TcpEstablishedTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UdpStreamTimeout() *float64
	SetUdpStreamTimeout(val *float64)
	UdpStreamTimeoutInput() *float64
	UdpTimeout() *float64
	SetUdpTimeout(val *float64)
	UdpTimeoutInput() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetTcpEstablishedTimeout()
	ResetUdpStreamTimeout()
	ResetUdpTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference
type jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) InternalValue() *LaunchTemplateNetworkInterfacesConnectionTrackingSpecification {
	var returns *LaunchTemplateNetworkInterfacesConnectionTrackingSpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) TcpEstablishedTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) TcpEstablishedTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) UdpStreamTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpStreamTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) UdpStreamTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpStreamTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) UdpTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) UdpTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpTimeoutInput",
		&returns,
	)
	return returns
}


func NewLaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewLaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference_Override(l LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetInternalValue(val *LaunchTemplateNetworkInterfacesConnectionTrackingSpecification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetTcpEstablishedTimeout(val *float64) {
	if err := j.validateSetTcpEstablishedTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpEstablishedTimeout",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetUdpStreamTimeout(val *float64) {
	if err := j.validateSetUdpStreamTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpStreamTimeout",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference)SetUdpTimeout(val *float64) {
	if err := j.validateSetUdpTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpTimeout",
		val,
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) ResetTcpEstablishedTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetTcpEstablishedTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) ResetUdpStreamTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetUdpStreamTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) ResetUdpTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetUdpTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateNetworkInterfacesConnectionTrackingSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

