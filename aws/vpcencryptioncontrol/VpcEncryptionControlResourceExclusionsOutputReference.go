// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcencryptioncontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/vpcencryptioncontrol/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcEncryptionControlResourceExclusionsOutputReference interface {
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
	EgressOnlyInternetGateway() VpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference
	ElasticFileSystem() VpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *VpcEncryptionControlResourceExclusions
	SetInternalValue(val *VpcEncryptionControlResourceExclusions)
	InternetGateway() VpcEncryptionControlResourceExclusionsInternetGatewayOutputReference
	Lambda() VpcEncryptionControlResourceExclusionsLambdaOutputReference
	NatGateway() VpcEncryptionControlResourceExclusionsNatGatewayOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualPrivateGateway() VpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference
	VpcLattice() VpcEncryptionControlResourceExclusionsVpcLatticeOutputReference
	VpcPeering() VpcEncryptionControlResourceExclusionsVpcPeeringOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpcEncryptionControlResourceExclusionsOutputReference
type jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) EgressOnlyInternetGateway() VpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference {
	var returns VpcEncryptionControlResourceExclusionsEgressOnlyInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"egressOnlyInternetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) ElasticFileSystem() VpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference {
	var returns VpcEncryptionControlResourceExclusionsElasticFileSystemOutputReference
	_jsii_.Get(
		j,
		"elasticFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) InternalValue() *VpcEncryptionControlResourceExclusions {
	var returns *VpcEncryptionControlResourceExclusions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) InternetGateway() VpcEncryptionControlResourceExclusionsInternetGatewayOutputReference {
	var returns VpcEncryptionControlResourceExclusionsInternetGatewayOutputReference
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) Lambda() VpcEncryptionControlResourceExclusionsLambdaOutputReference {
	var returns VpcEncryptionControlResourceExclusionsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) NatGateway() VpcEncryptionControlResourceExclusionsNatGatewayOutputReference {
	var returns VpcEncryptionControlResourceExclusionsNatGatewayOutputReference
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) VirtualPrivateGateway() VpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference {
	var returns VpcEncryptionControlResourceExclusionsVirtualPrivateGatewayOutputReference
	_jsii_.Get(
		j,
		"virtualPrivateGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) VpcLattice() VpcEncryptionControlResourceExclusionsVpcLatticeOutputReference {
	var returns VpcEncryptionControlResourceExclusionsVpcLatticeOutputReference
	_jsii_.Get(
		j,
		"vpcLattice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) VpcPeering() VpcEncryptionControlResourceExclusionsVpcPeeringOutputReference {
	var returns VpcEncryptionControlResourceExclusionsVpcPeeringOutputReference
	_jsii_.Get(
		j,
		"vpcPeering",
		&returns,
	)
	return returns
}


func NewVpcEncryptionControlResourceExclusionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpcEncryptionControlResourceExclusionsOutputReference {
	_init_.Initialize()

	if err := validateNewVpcEncryptionControlResourceExclusionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcEncryptionControl.VpcEncryptionControlResourceExclusionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpcEncryptionControlResourceExclusionsOutputReference_Override(v VpcEncryptionControlResourceExclusionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcEncryptionControl.VpcEncryptionControlResourceExclusionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference)SetInternalValue(val *VpcEncryptionControlResourceExclusions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpcEncryptionControlResourceExclusionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

