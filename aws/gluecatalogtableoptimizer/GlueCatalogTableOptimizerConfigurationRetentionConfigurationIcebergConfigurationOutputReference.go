// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtableoptimizer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/gluecatalogtableoptimizer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference interface {
	cdktf.ComplexObject
	CleanExpiredFiles() interface{}
	SetCleanExpiredFiles(val interface{})
	CleanExpiredFilesInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NumberOfSnapshotsToRetain() *float64
	SetNumberOfSnapshotsToRetain(val *float64)
	NumberOfSnapshotsToRetainInput() *float64
	SnapshotRetentionPeriodInDays() *float64
	SetSnapshotRetentionPeriodInDays(val *float64)
	SnapshotRetentionPeriodInDaysInput() *float64
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
	ResetCleanExpiredFiles()
	ResetNumberOfSnapshotsToRetain()
	ResetSnapshotRetentionPeriodInDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference
type jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) CleanExpiredFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanExpiredFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) CleanExpiredFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanExpiredFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) NumberOfSnapshotsToRetain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSnapshotsToRetain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) NumberOfSnapshotsToRetainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSnapshotsToRetainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) SnapshotRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) SnapshotRetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTableOptimizer.GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference_Override(g GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTableOptimizer.GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetCleanExpiredFiles(val interface{}) {
	if err := j.validateSetCleanExpiredFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanExpiredFiles",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetNumberOfSnapshotsToRetain(val *float64) {
	if err := j.validateSetNumberOfSnapshotsToRetainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfSnapshotsToRetain",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetSnapshotRetentionPeriodInDays(val *float64) {
	if err := j.validateSetSnapshotRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) ResetCleanExpiredFiles() {
	_jsii_.InvokeVoid(
		g,
		"resetCleanExpiredFiles",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) ResetNumberOfSnapshotsToRetain() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfSnapshotsToRetain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) ResetSnapshotRetentionPeriodInDays() {
	_jsii_.InvokeVoid(
		g,
		"resetSnapshotRetentionPeriodInDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOptimizerConfigurationRetentionConfigurationIcebergConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

