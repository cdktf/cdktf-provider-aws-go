// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dmsreplicationconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsReplicationConfigComputeConfigOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
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
	DnsNameServers() *string
	SetDnsNameServers(val *string)
	DnsNameServersInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DmsReplicationConfigComputeConfig
	SetInternalValue(val *DmsReplicationConfigComputeConfig)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	MaxCapacityUnits() *float64
	SetMaxCapacityUnits(val *float64)
	MaxCapacityUnitsInput() *float64
	MinCapacityUnits() *float64
	SetMinCapacityUnits(val *float64)
	MinCapacityUnitsInput() *float64
	MultiAz() interface{}
	SetMultiAz(val interface{})
	MultiAzInput() interface{}
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	ReplicationSubnetGroupId() *string
	SetReplicationSubnetGroupId(val *string)
	ReplicationSubnetGroupIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
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
	ResetAvailabilityZone()
	ResetDnsNameServers()
	ResetKmsKeyId()
	ResetMaxCapacityUnits()
	ResetMinCapacityUnits()
	ResetMultiAz()
	ResetPreferredMaintenanceWindow()
	ResetVpcSecurityGroupIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsReplicationConfigComputeConfigOutputReference
type jsiiProxy_DmsReplicationConfigComputeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) DnsNameServers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) DnsNameServersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) InternalValue() *DmsReplicationConfigComputeConfig {
	var returns *DmsReplicationConfigComputeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) MaxCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) MaxCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) MinCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) MinCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ReplicationSubnetGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ReplicationSubnetGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


func NewDmsReplicationConfigComputeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsReplicationConfigComputeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDmsReplicationConfigComputeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsReplicationConfigComputeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfigComputeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsReplicationConfigComputeConfigOutputReference_Override(d DmsReplicationConfigComputeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsReplicationConfig.DmsReplicationConfigComputeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetDnsNameServers(val *string) {
	if err := j.validateSetDnsNameServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsNameServers",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetInternalValue(val *DmsReplicationConfigComputeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetMaxCapacityUnits(val *float64) {
	if err := j.validateSetMaxCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetMinCapacityUnits(val *float64) {
	if err := j.validateSetMinCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetReplicationSubnetGroupId(val *string) {
	if err := j.validateSetReplicationSubnetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSubnetGroupId",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetDnsNameServers() {
	_jsii_.InvokeVoid(
		d,
		"resetDnsNameServers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetMaxCapacityUnits() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxCapacityUnits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetMinCapacityUnits() {
	_jsii_.InvokeVoid(
		d,
		"resetMinCapacityUnits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetMultiAz() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationConfigComputeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

