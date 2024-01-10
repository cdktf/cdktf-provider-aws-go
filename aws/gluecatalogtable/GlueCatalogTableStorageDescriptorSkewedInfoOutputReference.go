// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/gluecatalogtable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueCatalogTableStorageDescriptorSkewedInfoOutputReference interface {
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
	InternalValue() *GlueCatalogTableStorageDescriptorSkewedInfo
	SetInternalValue(val *GlueCatalogTableStorageDescriptorSkewedInfo)
	SkewedColumnNames() *[]*string
	SetSkewedColumnNames(val *[]*string)
	SkewedColumnNamesInput() *[]*string
	SkewedColumnValueLocationMaps() *map[string]*string
	SetSkewedColumnValueLocationMaps(val *map[string]*string)
	SkewedColumnValueLocationMapsInput() *map[string]*string
	SkewedColumnValues() *[]*string
	SetSkewedColumnValues(val *[]*string)
	SkewedColumnValuesInput() *[]*string
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
	ResetSkewedColumnNames()
	ResetSkewedColumnValueLocationMaps()
	ResetSkewedColumnValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableStorageDescriptorSkewedInfoOutputReference
type jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) InternalValue() *GlueCatalogTableStorageDescriptorSkewedInfo {
	var returns *GlueCatalogTableStorageDescriptorSkewedInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMaps() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMapsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) SkewedColumnValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableStorageDescriptorSkewedInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GlueCatalogTableStorageDescriptorSkewedInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableStorageDescriptorSkewedInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTableStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueCatalogTableStorageDescriptorSkewedInfoOutputReference_Override(g GlueCatalogTableStorageDescriptorSkewedInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTableStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetInternalValue(val *GlueCatalogTableStorageDescriptorSkewedInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetSkewedColumnNames(val *[]*string) {
	if err := j.validateSetSkewedColumnNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnNames",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetSkewedColumnValueLocationMaps(val *map[string]*string) {
	if err := j.validateSetSkewedColumnValueLocationMapsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnValueLocationMaps",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetSkewedColumnValues(val *[]*string) {
	if err := j.validateSetSkewedColumnValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnValues",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnNames() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValueLocationMaps() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValueLocationMaps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValues() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSkewedInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

