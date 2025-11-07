// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/vpnconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference interface {
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
	InternalValue() *VpnConnectionTunnel2LogOptionsCloudwatchLogOptions
	SetInternalValue(val *VpnConnectionTunnel2LogOptionsCloudwatchLogOptions)
	LogEnabled() interface{}
	SetLogEnabled(val interface{})
	LogEnabledInput() interface{}
	LogGroupArn() *string
	SetLogGroupArn(val *string)
	LogGroupArnInput() *string
	LogOutputFormat() *string
	SetLogOutputFormat(val *string)
	LogOutputFormatInput() *string
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
	ResetLogEnabled()
	ResetLogGroupArn()
	ResetLogOutputFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference
type jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) InternalValue() *VpnConnectionTunnel2LogOptionsCloudwatchLogOptions {
	var returns *VpnConnectionTunnel2LogOptionsCloudwatchLogOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) LogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) LogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) LogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) LogOutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) LogOutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logOutputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewVpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpnConnection.VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference_Override(v VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpnConnection.VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetInternalValue(val *VpnConnectionTunnel2LogOptionsCloudwatchLogOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetLogEnabled(val interface{}) {
	if err := j.validateSetLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetLogGroupArn(val *string) {
	if err := j.validateSetLogGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupArn",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetLogOutputFormat(val *string) {
	if err := j.validateSetLogOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logOutputFormat",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) ResetLogEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetLogEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) ResetLogGroupArn() {
	_jsii_.InvokeVoid(
		v,
		"resetLogGroupArn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) ResetLogOutputFormat() {
	_jsii_.InvokeVoid(
		v,
		"resetLogOutputFormat",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2LogOptionsCloudwatchLogOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

