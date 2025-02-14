// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtableexport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dynamodbtableexport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamodbTableExportIncrementalExportSpecificationOutputReference interface {
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
	ExportFromTime() *string
	SetExportFromTime(val *string)
	ExportFromTimeInput() *string
	ExportToTime() *string
	SetExportToTime(val *string)
	ExportToTimeInput() *string
	ExportViewType() *string
	SetExportViewType(val *string)
	ExportViewTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DynamodbTableExportIncrementalExportSpecification
	SetInternalValue(val *DynamodbTableExportIncrementalExportSpecification)
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
	ResetExportFromTime()
	ResetExportToTime()
	ResetExportViewType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbTableExportIncrementalExportSpecificationOutputReference
type jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ExportFromTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFromTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ExportFromTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFromTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ExportToTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ExportToTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ExportViewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportViewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ExportViewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportViewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) InternalValue() *DynamodbTableExportIncrementalExportSpecification {
	var returns *DynamodbTableExportIncrementalExportSpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDynamodbTableExportIncrementalExportSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DynamodbTableExportIncrementalExportSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbTableExportIncrementalExportSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExportIncrementalExportSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDynamodbTableExportIncrementalExportSpecificationOutputReference_Override(d DynamodbTableExportIncrementalExportSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExportIncrementalExportSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetExportFromTime(val *string) {
	if err := j.validateSetExportFromTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportFromTime",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetExportToTime(val *string) {
	if err := j.validateSetExportToTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToTime",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetExportViewType(val *string) {
	if err := j.validateSetExportViewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportViewType",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetInternalValue(val *DynamodbTableExportIncrementalExportSpecification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ResetExportFromTime() {
	_jsii_.InvokeVoid(
		d,
		"resetExportFromTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ResetExportToTime() {
	_jsii_.InvokeVoid(
		d,
		"resetExportToTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ResetExportViewType() {
	_jsii_.InvokeVoid(
		d,
		"resetExportViewType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbTableExportIncrementalExportSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

