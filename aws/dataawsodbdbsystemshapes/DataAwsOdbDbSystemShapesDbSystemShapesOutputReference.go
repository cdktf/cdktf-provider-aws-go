// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsodbdbsystemshapes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsodbdbsystemshapes/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsOdbDbSystemShapesDbSystemShapesOutputReference interface {
	cdktf.ComplexObject
	AvailableCoreCount() *float64
	AvailableCoreCountPerNode() *float64
	AvailableDataStorageInTbs() *float64
	AvailableDataStoragePerServerInTbs() *float64
	AvailableDbNodePerNodeInGbs() *float64
	AvailableDbNodeStorageInGbs() *float64
	AvailableMemoryInGbs() *float64
	AvailableMemoryPerNodeInGbs() *float64
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
	CoreCountIncrement() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsOdbDbSystemShapesDbSystemShapes
	SetInternalValue(val *DataAwsOdbDbSystemShapesDbSystemShapes)
	MaximumNodeCount() *float64
	MaxStorageCount() *float64
	MinCoreCountPerNode() *float64
	MinDataStorageInTbs() *float64
	MinDbNodeStoragePerNodeInGbs() *float64
	MinimumCoreCount() *float64
	MinimumNodeCount() *float64
	MinMemoryPerNodeInGbs() *float64
	MinStorageCount() *float64
	Name() *string
	RuntimeMinimumCoreCount() *float64
	ShapeFamily() *string
	ShapeType() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsOdbDbSystemShapesDbSystemShapesOutputReference
type jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableDataStoragePerServerInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDataStoragePerServerInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableDbNodePerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDbNodePerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableDbNodeStorageInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDbNodeStorageInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableMemoryInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) AvailableMemoryPerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryPerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) CoreCountIncrement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreCountIncrement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) InternalValue() *DataAwsOdbDbSystemShapesDbSystemShapes {
	var returns *DataAwsOdbDbSystemShapesDbSystemShapes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MaximumNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MaxStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MinCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MinDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MinDbNodeStoragePerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDbNodeStoragePerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MinimumCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MinimumNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MinMemoryPerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryPerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) MinStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) RuntimeMinimumCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runtimeMinimumCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) ShapeFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) ShapeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsOdbDbSystemShapesDbSystemShapesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsOdbDbSystemShapesDbSystemShapesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsOdbDbSystemShapesDbSystemShapesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOdbDbSystemShapes.DataAwsOdbDbSystemShapesDbSystemShapesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsOdbDbSystemShapesDbSystemShapesOutputReference_Override(d DataAwsOdbDbSystemShapesDbSystemShapesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOdbDbSystemShapes.DataAwsOdbDbSystemShapesDbSystemShapesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference)SetInternalValue(val *DataAwsOdbDbSystemShapesDbSystemShapes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapesDbSystemShapesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

