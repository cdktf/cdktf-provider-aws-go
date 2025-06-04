// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/workspacesdirectory/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference interface {
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
	DeviceTypeAndroid() *string
	SetDeviceTypeAndroid(val *string)
	DeviceTypeAndroidInput() *string
	DeviceTypeChromeos() *string
	SetDeviceTypeChromeos(val *string)
	DeviceTypeChromeosInput() *string
	DeviceTypeIos() *string
	SetDeviceTypeIos(val *string)
	DeviceTypeIosInput() *string
	DeviceTypeLinux() *string
	SetDeviceTypeLinux(val *string)
	DeviceTypeLinuxInput() *string
	DeviceTypeOsx() *string
	SetDeviceTypeOsx(val *string)
	DeviceTypeOsxInput() *string
	DeviceTypeWeb() *string
	SetDeviceTypeWeb(val *string)
	DeviceTypeWebInput() *string
	DeviceTypeWindows() *string
	SetDeviceTypeWindows(val *string)
	DeviceTypeWindowsInput() *string
	DeviceTypeZeroclient() *string
	SetDeviceTypeZeroclient(val *string)
	DeviceTypeZeroclientInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *WorkspacesDirectoryWorkspaceAccessProperties
	SetInternalValue(val *WorkspacesDirectoryWorkspaceAccessProperties)
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
	ResetDeviceTypeAndroid()
	ResetDeviceTypeChromeos()
	ResetDeviceTypeIos()
	ResetDeviceTypeLinux()
	ResetDeviceTypeOsx()
	ResetDeviceTypeWeb()
	ResetDeviceTypeWindows()
	ResetDeviceTypeZeroclient()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference
type jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeAndroid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeAndroid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeAndroidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeAndroidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeChromeos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeChromeos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeChromeosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeChromeosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeIos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeIosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeLinux() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeLinuxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeOsx() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeOsx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeOsxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeOsxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWeb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWebInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWindows() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWindowsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeZeroclient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeZeroclient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeZeroclientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeZeroclientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) InternalValue() *WorkspacesDirectoryWorkspaceAccessProperties {
	var returns *WorkspacesDirectoryWorkspaceAccessProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspacesDirectoryWorkspaceAccessPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesDirectoryWorkspaceAccessPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.workspacesDirectory.WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesDirectoryWorkspaceAccessPropertiesOutputReference_Override(w WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.workspacesDirectory.WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeAndroid(val *string) {
	if err := j.validateSetDeviceTypeAndroidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeAndroid",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeChromeos(val *string) {
	if err := j.validateSetDeviceTypeChromeosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeChromeos",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeIos(val *string) {
	if err := j.validateSetDeviceTypeIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeIos",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeLinux(val *string) {
	if err := j.validateSetDeviceTypeLinuxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeLinux",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeOsx(val *string) {
	if err := j.validateSetDeviceTypeOsxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeOsx",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeWeb(val *string) {
	if err := j.validateSetDeviceTypeWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeWeb",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeWindows(val *string) {
	if err := j.validateSetDeviceTypeWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeWindows",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetDeviceTypeZeroclient(val *string) {
	if err := j.validateSetDeviceTypeZeroclientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeZeroclient",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetInternalValue(val *WorkspacesDirectoryWorkspaceAccessProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeAndroid() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeAndroid",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeChromeos() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeChromeos",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeIos() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeIos",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeLinux() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeLinux",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeOsx() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeOsx",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeWeb() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeWeb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeWindows() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeWindows",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeZeroclient() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeZeroclient",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

