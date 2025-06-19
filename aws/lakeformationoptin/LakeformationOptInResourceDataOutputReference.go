// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationoptin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lakeformationoptin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakeformationOptInResourceDataOutputReference interface {
	cdktf.ComplexObject
	Catalog() LakeformationOptInResourceDataCatalogList
	CatalogInput() interface{}
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
	Database() LakeformationOptInResourceDataDatabaseList
	DatabaseInput() interface{}
	DataCellsFilter() LakeformationOptInResourceDataDataCellsFilterList
	DataCellsFilterInput() interface{}
	DataLocation() LakeformationOptInResourceDataDataLocationList
	DataLocationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LfTag() LakeformationOptInResourceDataLfTagList
	LfTagExpression() LakeformationOptInResourceDataLfTagExpressionList
	LfTagExpressionInput() interface{}
	LfTagInput() interface{}
	LfTagPolicy() LakeformationOptInResourceDataLfTagPolicyList
	LfTagPolicyInput() interface{}
	Table() LakeformationOptInResourceDataTableList
	TableInput() interface{}
	TableWithColumns() LakeformationOptInResourceDataTableWithColumnsList
	TableWithColumnsInput() interface{}
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
	PutCatalog(value interface{})
	PutDatabase(value interface{})
	PutDataCellsFilter(value interface{})
	PutDataLocation(value interface{})
	PutLfTag(value interface{})
	PutLfTagExpression(value interface{})
	PutLfTagPolicy(value interface{})
	PutTable(value interface{})
	PutTableWithColumns(value interface{})
	ResetCatalog()
	ResetDatabase()
	ResetDataCellsFilter()
	ResetDataLocation()
	ResetLfTag()
	ResetLfTagExpression()
	ResetLfTagPolicy()
	ResetTable()
	ResetTableWithColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LakeformationOptInResourceDataOutputReference
type jsiiProxy_LakeformationOptInResourceDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) Catalog() LakeformationOptInResourceDataCatalogList {
	var returns LakeformationOptInResourceDataCatalogList
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) CatalogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) Database() LakeformationOptInResourceDataDatabaseList {
	var returns LakeformationOptInResourceDataDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) DataCellsFilter() LakeformationOptInResourceDataDataCellsFilterList {
	var returns LakeformationOptInResourceDataDataCellsFilterList
	_jsii_.Get(
		j,
		"dataCellsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) DataCellsFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCellsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) DataLocation() LakeformationOptInResourceDataDataLocationList {
	var returns LakeformationOptInResourceDataDataLocationList
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) DataLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) LfTag() LakeformationOptInResourceDataLfTagList {
	var returns LakeformationOptInResourceDataLfTagList
	_jsii_.Get(
		j,
		"lfTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) LfTagExpression() LakeformationOptInResourceDataLfTagExpressionList {
	var returns LakeformationOptInResourceDataLfTagExpressionList
	_jsii_.Get(
		j,
		"lfTagExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) LfTagExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) LfTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) LfTagPolicy() LakeformationOptInResourceDataLfTagPolicyList {
	var returns LakeformationOptInResourceDataLfTagPolicyList
	_jsii_.Get(
		j,
		"lfTagPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) LfTagPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) Table() LakeformationOptInResourceDataTableList {
	var returns LakeformationOptInResourceDataTableList
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) TableWithColumns() LakeformationOptInResourceDataTableWithColumnsList {
	var returns LakeformationOptInResourceDataTableWithColumnsList
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) TableWithColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLakeformationOptInResourceDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LakeformationOptInResourceDataOutputReference {
	_init_.Initialize()

	if err := validateNewLakeformationOptInResourceDataOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakeformationOptInResourceDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationOptIn.LakeformationOptInResourceDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLakeformationOptInResourceDataOutputReference_Override(l LakeformationOptInResourceDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationOptIn.LakeformationOptInResourceDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutCatalog(value interface{}) {
	if err := l.validatePutCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCatalog",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutDatabase(value interface{}) {
	if err := l.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDatabase",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutDataCellsFilter(value interface{}) {
	if err := l.validatePutDataCellsFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataCellsFilter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutDataLocation(value interface{}) {
	if err := l.validatePutDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutLfTag(value interface{}) {
	if err := l.validatePutLfTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLfTag",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutLfTagExpression(value interface{}) {
	if err := l.validatePutLfTagExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLfTagExpression",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutLfTagPolicy(value interface{}) {
	if err := l.validatePutLfTagPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLfTagPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutTable(value interface{}) {
	if err := l.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTable",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) PutTableWithColumns(value interface{}) {
	if err := l.validatePutTableWithColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalog",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		l,
		"resetDatabase",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetDataCellsFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetDataCellsFilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetDataLocation() {
	_jsii_.InvokeVoid(
		l,
		"resetDataLocation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetLfTag() {
	_jsii_.InvokeVoid(
		l,
		"resetLfTag",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetLfTagExpression() {
	_jsii_.InvokeVoid(
		l,
		"resetLfTagExpression",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetLfTagPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetLfTagPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		l,
		"resetTable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		l,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

