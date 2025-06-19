// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupreportplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/backupreportplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupReportPlanReportSettingOutputReference interface {
	cdktf.ComplexObject
	Accounts() *[]*string
	SetAccounts(val *[]*string)
	AccountsInput() *[]*string
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
	FrameworkArns() *[]*string
	SetFrameworkArns(val *[]*string)
	FrameworkArnsInput() *[]*string
	InternalValue() *BackupReportPlanReportSetting
	SetInternalValue(val *BackupReportPlanReportSetting)
	NumberOfFrameworks() *float64
	SetNumberOfFrameworks(val *float64)
	NumberOfFrameworksInput() *float64
	OrganizationUnits() *[]*string
	SetOrganizationUnits(val *[]*string)
	OrganizationUnitsInput() *[]*string
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	ReportTemplate() *string
	SetReportTemplate(val *string)
	ReportTemplateInput() *string
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
	ResetAccounts()
	ResetFrameworkArns()
	ResetNumberOfFrameworks()
	ResetOrganizationUnits()
	ResetRegions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupReportPlanReportSettingOutputReference
type jsiiProxy_BackupReportPlanReportSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) FrameworkArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frameworkArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) FrameworkArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frameworkArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) InternalValue() *BackupReportPlanReportSetting {
	var returns *BackupReportPlanReportSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) NumberOfFrameworks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFrameworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) NumberOfFrameworksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfFrameworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) OrganizationUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) OrganizationUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) ReportTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) ReportTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupReportPlanReportSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupReportPlanReportSettingOutputReference {
	_init_.Initialize()

	if err := validateNewBackupReportPlanReportSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupReportPlanReportSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.backupReportPlan.BackupReportPlanReportSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupReportPlanReportSettingOutputReference_Override(b BackupReportPlanReportSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.backupReportPlan.BackupReportPlanReportSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetFrameworkArns(val *[]*string) {
	if err := j.validateSetFrameworkArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameworkArns",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetInternalValue(val *BackupReportPlanReportSetting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetNumberOfFrameworks(val *float64) {
	if err := j.validateSetNumberOfFrameworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfFrameworks",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetOrganizationUnits(val *[]*string) {
	if err := j.validateSetOrganizationUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationUnits",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetReportTemplate(val *string) {
	if err := j.validateSetReportTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportTemplate",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupReportPlanReportSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		b,
		"resetAccounts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) ResetFrameworkArns() {
	_jsii_.InvokeVoid(
		b,
		"resetFrameworkArns",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) ResetNumberOfFrameworks() {
	_jsii_.InvokeVoid(
		b,
		"resetNumberOfFrameworks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) ResetOrganizationUnits() {
	_jsii_.InvokeVoid(
		b,
		"resetOrganizationUnits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		b,
		"resetRegions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupReportPlanReportSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

