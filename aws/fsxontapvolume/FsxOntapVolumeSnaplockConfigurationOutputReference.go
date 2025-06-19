// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/fsxontapvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOntapVolumeSnaplockConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuditLogVolume() interface{}
	SetAuditLogVolume(val interface{})
	AuditLogVolumeInput() interface{}
	AutocommitPeriod() FsxOntapVolumeSnaplockConfigurationAutocommitPeriodOutputReference
	AutocommitPeriodInput() *FsxOntapVolumeSnaplockConfigurationAutocommitPeriod
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
	InternalValue() *FsxOntapVolumeSnaplockConfiguration
	SetInternalValue(val *FsxOntapVolumeSnaplockConfiguration)
	PrivilegedDelete() *string
	SetPrivilegedDelete(val *string)
	PrivilegedDeleteInput() *string
	RetentionPeriod() FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference
	RetentionPeriodInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriod
	SnaplockType() *string
	SetSnaplockType(val *string)
	SnaplockTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeAppendModeEnabled() interface{}
	SetVolumeAppendModeEnabled(val interface{})
	VolumeAppendModeEnabledInput() interface{}
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
	PutAutocommitPeriod(value *FsxOntapVolumeSnaplockConfigurationAutocommitPeriod)
	PutRetentionPeriod(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriod)
	ResetAuditLogVolume()
	ResetAutocommitPeriod()
	ResetPrivilegedDelete()
	ResetRetentionPeriod()
	ResetVolumeAppendModeEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxOntapVolumeSnaplockConfigurationOutputReference
type jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) AuditLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) AuditLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) AutocommitPeriod() FsxOntapVolumeSnaplockConfigurationAutocommitPeriodOutputReference {
	var returns FsxOntapVolumeSnaplockConfigurationAutocommitPeriodOutputReference
	_jsii_.Get(
		j,
		"autocommitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) AutocommitPeriodInput() *FsxOntapVolumeSnaplockConfigurationAutocommitPeriod {
	var returns *FsxOntapVolumeSnaplockConfigurationAutocommitPeriod
	_jsii_.Get(
		j,
		"autocommitPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) InternalValue() *FsxOntapVolumeSnaplockConfiguration {
	var returns *FsxOntapVolumeSnaplockConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) PrivilegedDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) PrivilegedDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) RetentionPeriod() FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference {
	var returns FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) RetentionPeriodInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriod {
	var returns *FsxOntapVolumeSnaplockConfigurationRetentionPeriod
	_jsii_.Get(
		j,
		"retentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) SnaplockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) SnaplockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) VolumeAppendModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) VolumeAppendModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabledInput",
		&returns,
	)
	return returns
}


func NewFsxOntapVolumeSnaplockConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FsxOntapVolumeSnaplockConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxOntapVolumeSnaplockConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolumeSnaplockConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxOntapVolumeSnaplockConfigurationOutputReference_Override(f FsxOntapVolumeSnaplockConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolumeSnaplockConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetAuditLogVolume(val interface{}) {
	if err := j.validateSetAuditLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogVolume",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetInternalValue(val *FsxOntapVolumeSnaplockConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetPrivilegedDelete(val *string) {
	if err := j.validateSetPrivilegedDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedDelete",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetSnaplockType(val *string) {
	if err := j.validateSetSnaplockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snaplockType",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference)SetVolumeAppendModeEnabled(val interface{}) {
	if err := j.validateSetVolumeAppendModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAppendModeEnabled",
		val,
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) PutAutocommitPeriod(value *FsxOntapVolumeSnaplockConfigurationAutocommitPeriod) {
	if err := f.validatePutAutocommitPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAutocommitPeriod",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) PutRetentionPeriod(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriod) {
	if err := f.validatePutRetentionPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putRetentionPeriod",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ResetAuditLogVolume() {
	_jsii_.InvokeVoid(
		f,
		"resetAuditLogVolume",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ResetAutocommitPeriod() {
	_jsii_.InvokeVoid(
		f,
		"resetAutocommitPeriod",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ResetPrivilegedDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetPrivilegedDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ResetRetentionPeriod() {
	_jsii_.InvokeVoid(
		f,
		"resetRetentionPeriod",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ResetVolumeAppendModeEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeAppendModeEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

