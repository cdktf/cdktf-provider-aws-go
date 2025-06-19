// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/verifiedaccessendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VerifiedaccessEndpointRdsOptionsOutputReference interface {
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
	InternalValue() *VerifiedaccessEndpointRdsOptions
	SetInternalValue(val *VerifiedaccessEndpointRdsOptions)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RdsDbClusterArn() *string
	SetRdsDbClusterArn(val *string)
	RdsDbClusterArnInput() *string
	RdsDbInstanceArn() *string
	SetRdsDbInstanceArn(val *string)
	RdsDbInstanceArnInput() *string
	RdsDbProxyArn() *string
	SetRdsDbProxyArn(val *string)
	RdsDbProxyArnInput() *string
	RdsEndpoint() *string
	SetRdsEndpoint(val *string)
	RdsEndpointInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetPort()
	ResetProtocol()
	ResetRdsDbClusterArn()
	ResetRdsDbInstanceArn()
	ResetRdsDbProxyArn()
	ResetRdsEndpoint()
	ResetSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VerifiedaccessEndpointRdsOptionsOutputReference
type jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) InternalValue() *VerifiedaccessEndpointRdsOptions {
	var returns *VerifiedaccessEndpointRdsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsDbClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsDbClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsDbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsDbInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsDbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsDbProxyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) RdsEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVerifiedaccessEndpointRdsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VerifiedaccessEndpointRdsOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewVerifiedaccessEndpointRdsOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpointRdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVerifiedaccessEndpointRdsOptionsOutputReference_Override(v VerifiedaccessEndpointRdsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessEndpoint.VerifiedaccessEndpointRdsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetInternalValue(val *VerifiedaccessEndpointRdsOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetRdsDbClusterArn(val *string) {
	if err := j.validateSetRdsDbClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbClusterArn",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetRdsDbInstanceArn(val *string) {
	if err := j.validateSetRdsDbInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbInstanceArn",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetRdsDbProxyArn(val *string) {
	if err := j.validateSetRdsDbProxyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbProxyArn",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetRdsEndpoint(val *string) {
	if err := j.validateSetRdsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsEndpoint",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		v,
		"resetPort",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ResetRdsDbClusterArn() {
	_jsii_.InvokeVoid(
		v,
		"resetRdsDbClusterArn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ResetRdsDbInstanceArn() {
	_jsii_.InvokeVoid(
		v,
		"resetRdsDbInstanceArn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ResetRdsDbProxyArn() {
	_jsii_.InvokeVoid(
		v,
		"resetRdsDbProxyArn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ResetRdsEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetRdsEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		v,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessEndpointRdsOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

