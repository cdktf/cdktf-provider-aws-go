// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsodbdbservers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsodbdbservers/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsOdbDbServersDbServersOutputReference interface {
	cdktf.ComplexObject
	AutonomousVirtualMachineIds() *[]*string
	AutonomousVmClusterIds() *[]*string
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
	ComputeModel() *string
	CpuCoreCount() *float64
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DbNodeStorageSizeInGbs() *float64
	DbServerPatchingDetails() DataAwsOdbDbServersDbServersDbServerPatchingDetailsList
	DisplayName() *string
	ExadataInfrastructureId() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataAwsOdbDbServersDbServers
	SetInternalValue(val *DataAwsOdbDbServersDbServers)
	MaxCpuCount() *float64
	MaxDbNodeStorageInGbs() *float64
	MaxMemoryInGbs() *float64
	MemorySizeInGbs() *float64
	Ocid() *string
	OciResourceAnchorName() *string
	Shape() *string
	Status() *string
	StatusReason() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmClusterIds() *[]*string
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

// The jsii proxy struct for DataAwsOdbDbServersDbServersOutputReference
type jsiiProxy_DataAwsOdbDbServersDbServersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) AutonomousVirtualMachineIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autonomousVirtualMachineIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) AutonomousVmClusterIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autonomousVmClusterIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) DbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) DbServerPatchingDetails() DataAwsOdbDbServersDbServersDbServerPatchingDetailsList {
	var returns DataAwsOdbDbServersDbServersDbServerPatchingDetailsList
	_jsii_.Get(
		j,
		"dbServerPatchingDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) ExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) InternalValue() *DataAwsOdbDbServersDbServers {
	var returns *DataAwsOdbDbServersDbServers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) MaxCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) MaxDbNodeStorageInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDbNodeStorageInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) MaxMemoryInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) VmClusterIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmClusterIds",
		&returns,
	)
	return returns
}


func NewDataAwsOdbDbServersDbServersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsOdbDbServersDbServersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsOdbDbServersDbServersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOdbDbServersDbServersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOdbDbServers.DataAwsOdbDbServersDbServersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsOdbDbServersDbServersOutputReference_Override(d DataAwsOdbDbServersDbServersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOdbDbServers.DataAwsOdbDbServersDbServersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference)SetInternalValue(val *DataAwsOdbDbServersDbServers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsOdbDbServersDbServersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

