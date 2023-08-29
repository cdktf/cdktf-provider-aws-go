// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/backupplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupPlanRuleOutputReference interface {
	cdktf.ComplexObject
	CompletionWindow() *float64
	SetCompletionWindow(val *float64)
	CompletionWindowInput() *float64
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
	CopyAction() BackupPlanRuleCopyActionList
	CopyActionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableContinuousBackup() interface{}
	SetEnableContinuousBackup(val interface{})
	EnableContinuousBackupInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lifecycle() BackupPlanRuleLifecycleOutputReference
	LifecycleInput() *BackupPlanRuleLifecycle
	RecoveryPointTags() *map[string]*string
	SetRecoveryPointTags(val *map[string]*string)
	RecoveryPointTagsInput() *map[string]*string
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	StartWindow() *float64
	SetStartWindow(val *float64)
	StartWindowInput() *float64
	TargetVaultName() *string
	SetTargetVaultName(val *string)
	TargetVaultNameInput() *string
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
	PutCopyAction(value interface{})
	PutLifecycle(value *BackupPlanRuleLifecycle)
	ResetCompletionWindow()
	ResetCopyAction()
	ResetEnableContinuousBackup()
	ResetLifecycle()
	ResetRecoveryPointTags()
	ResetSchedule()
	ResetStartWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPlanRuleOutputReference
type jsiiProxy_BackupPlanRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) CompletionWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) CompletionWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) CopyAction() BackupPlanRuleCopyActionList {
	var returns BackupPlanRuleCopyActionList
	_jsii_.Get(
		j,
		"copyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) CopyActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) EnableContinuousBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableContinuousBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) EnableContinuousBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableContinuousBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) Lifecycle() BackupPlanRuleLifecycleOutputReference {
	var returns BackupPlanRuleLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) LifecycleInput() *BackupPlanRuleLifecycle {
	var returns *BackupPlanRuleLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) RecoveryPointTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"recoveryPointTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) RecoveryPointTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"recoveryPointTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) StartWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) StartWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) TargetVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) TargetVaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupPlanRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupPlanRuleOutputReference {
	_init_.Initialize()

	if err := validateNewBackupPlanRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupPlanRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.backupPlan.BackupPlanRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupPlanRuleOutputReference_Override(b BackupPlanRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.backupPlan.BackupPlanRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetCompletionWindow(val *float64) {
	if err := j.validateSetCompletionWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completionWindow",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetEnableContinuousBackup(val interface{}) {
	if err := j.validateSetEnableContinuousBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableContinuousBackup",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetRecoveryPointTags(val *map[string]*string) {
	if err := j.validateSetRecoveryPointTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointTags",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetStartWindow(val *float64) {
	if err := j.validateSetStartWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startWindow",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetTargetVaultName(val *string) {
	if err := j.validateSetTargetVaultNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVaultName",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) PutCopyAction(value interface{}) {
	if err := b.validatePutCopyActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCopyAction",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) PutLifecycle(value *BackupPlanRuleLifecycle) {
	if err := b.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLifecycle",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ResetCompletionWindow() {
	_jsii_.InvokeVoid(
		b,
		"resetCompletionWindow",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ResetCopyAction() {
	_jsii_.InvokeVoid(
		b,
		"resetCopyAction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ResetEnableContinuousBackup() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableContinuousBackup",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		b,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ResetRecoveryPointTags() {
	_jsii_.InvokeVoid(
		b,
		"resetRecoveryPointTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		b,
		"resetSchedule",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) ResetStartWindow() {
	_jsii_.InvokeVoid(
		b,
		"resetStartWindow",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BackupPlanRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

