// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslaunchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawslaunchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsLaunchTemplateNetworkInterfacesOutputReference interface {
	cdktf.ComplexObject
	AssociateCarrierIpAddress() *string
	AssociatePublicIpAddress() *string
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
	ConnectionTrackingSpecification() DataAwsLaunchTemplateNetworkInterfacesConnectionTrackingSpecificationList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteOnTermination() *string
	Description() *string
	DeviceIndex() *float64
	// Experimental.
	Fqn() *string
	InterfaceType() *string
	InternalValue() *DataAwsLaunchTemplateNetworkInterfaces
	SetInternalValue(val *DataAwsLaunchTemplateNetworkInterfaces)
	Ipv4AddressCount() *float64
	Ipv4Addresses() *[]*string
	Ipv4PrefixCount() *float64
	Ipv4Prefixes() *[]*string
	Ipv6AddressCount() *float64
	Ipv6Addresses() *[]*string
	Ipv6PrefixCount() *float64
	Ipv6Prefixes() *[]*string
	NetworkCardIndex() *float64
	NetworkInterfaceId() *string
	PrimaryIpv6() *string
	PrivateIpAddress() *string
	SecurityGroups() *[]*string
	SubnetId() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsLaunchTemplateNetworkInterfacesOutputReference
type jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) AssociateCarrierIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateCarrierIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) AssociatePublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) ConnectionTrackingSpecification() DataAwsLaunchTemplateNetworkInterfacesConnectionTrackingSpecificationList {
	var returns DataAwsLaunchTemplateNetworkInterfacesConnectionTrackingSpecificationList
	_jsii_.Get(
		j,
		"connectionTrackingSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) DeleteOnTermination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) InternalValue() *DataAwsLaunchTemplateNetworkInterfaces {
	var returns *DataAwsLaunchTemplateNetworkInterfaces
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv4AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv4PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv4Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv6PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Ipv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) PrimaryIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsLaunchTemplateNetworkInterfacesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsLaunchTemplateNetworkInterfacesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsLaunchTemplateNetworkInterfacesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplateNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsLaunchTemplateNetworkInterfacesOutputReference_Override(d DataAwsLaunchTemplateNetworkInterfacesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplateNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference)SetInternalValue(val *DataAwsLaunchTemplateNetworkInterfaces) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateNetworkInterfacesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

