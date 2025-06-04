// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuprestoretestingplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/backuprestoretestingplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupRestoreTestingPlanRecoveryPointSelectionOutputReference interface {
	cdktf.ComplexObject
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
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
	ExcludeVaults() *[]*string
	SetExcludeVaults(val *[]*string)
	ExcludeVaultsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeVaults() *[]*string
	SetIncludeVaults(val *[]*string)
	IncludeVaultsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RecoveryPointTypes() *[]*string
	SetRecoveryPointTypes(val *[]*string)
	RecoveryPointTypesInput() *[]*string
	SelectionWindowDays() *float64
	SetSelectionWindowDays(val *float64)
	SelectionWindowDaysInput() *float64
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
	ResetExcludeVaults()
	ResetSelectionWindowDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupRestoreTestingPlanRecoveryPointSelectionOutputReference
type jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ExcludeVaults() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ExcludeVaultsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) IncludeVaults() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) IncludeVaultsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeVaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) RecoveryPointTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryPointTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) RecoveryPointTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryPointTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) SelectionWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selectionWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) SelectionWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"selectionWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupRestoreTestingPlanRecoveryPointSelectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupRestoreTestingPlanRecoveryPointSelectionOutputReference {
	_init_.Initialize()

	if err := validateNewBackupRestoreTestingPlanRecoveryPointSelectionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.backupRestoreTestingPlan.BackupRestoreTestingPlanRecoveryPointSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupRestoreTestingPlanRecoveryPointSelectionOutputReference_Override(b BackupRestoreTestingPlanRecoveryPointSelectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.backupRestoreTestingPlan.BackupRestoreTestingPlanRecoveryPointSelectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetExcludeVaults(val *[]*string) {
	if err := j.validateSetExcludeVaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeVaults",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetIncludeVaults(val *[]*string) {
	if err := j.validateSetIncludeVaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVaults",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetRecoveryPointTypes(val *[]*string) {
	if err := j.validateSetRecoveryPointTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointTypes",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetSelectionWindowDays(val *float64) {
	if err := j.validateSetSelectionWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectionWindowDays",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ResetExcludeVaults() {
	_jsii_.InvokeVoid(
		b,
		"resetExcludeVaults",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ResetSelectionWindowDays() {
	_jsii_.InvokeVoid(
		b,
		"resetSelectionWindowDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BackupRestoreTestingPlanRecoveryPointSelectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

