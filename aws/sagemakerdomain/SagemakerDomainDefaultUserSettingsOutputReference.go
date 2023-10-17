// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/sagemakerdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDomainDefaultUserSettingsOutputReference interface {
	cdktf.ComplexObject
	CanvasAppSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference
	CanvasAppSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettings
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
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerDomainDefaultUserSettings
	SetInternalValue(val *SagemakerDomainDefaultUserSettings)
	JupyterServerAppSettings() SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference
	JupyterServerAppSettingsInput() *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings
	KernelGatewayAppSettings() SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference
	KernelGatewayAppSettingsInput() *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings
	RSessionAppSettings() SagemakerDomainDefaultUserSettingsRSessionAppSettingsOutputReference
	RSessionAppSettingsInput() *SagemakerDomainDefaultUserSettingsRSessionAppSettings
	RStudioServerProAppSettings() SagemakerDomainDefaultUserSettingsRStudioServerProAppSettingsOutputReference
	RStudioServerProAppSettingsInput() *SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SharingSettings() SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference
	SharingSettingsInput() *SagemakerDomainDefaultUserSettingsSharingSettings
	TensorBoardAppSettings() SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsOutputReference
	TensorBoardAppSettingsInput() *SagemakerDomainDefaultUserSettingsTensorBoardAppSettings
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
	PutCanvasAppSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettings)
	PutJupyterServerAppSettings(value *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings)
	PutKernelGatewayAppSettings(value *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings)
	PutRSessionAppSettings(value *SagemakerDomainDefaultUserSettingsRSessionAppSettings)
	PutRStudioServerProAppSettings(value *SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings)
	PutSharingSettings(value *SagemakerDomainDefaultUserSettingsSharingSettings)
	PutTensorBoardAppSettings(value *SagemakerDomainDefaultUserSettingsTensorBoardAppSettings)
	ResetCanvasAppSettings()
	ResetJupyterServerAppSettings()
	ResetKernelGatewayAppSettings()
	ResetRSessionAppSettings()
	ResetRStudioServerProAppSettings()
	ResetSecurityGroups()
	ResetSharingSettings()
	ResetTensorBoardAppSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDefaultUserSettingsOutputReference
type jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) CanvasAppSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference
	_jsii_.Get(
		j,
		"canvasAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) CanvasAppSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettings {
	var returns *SagemakerDomainDefaultUserSettingsCanvasAppSettings
	_jsii_.Get(
		j,
		"canvasAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) InternalValue() *SagemakerDomainDefaultUserSettings {
	var returns *SagemakerDomainDefaultUserSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) JupyterServerAppSettings() SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) JupyterServerAppSettingsInput() *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings {
	var returns *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) KernelGatewayAppSettings() SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) KernelGatewayAppSettingsInput() *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings {
	var returns *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) RSessionAppSettings() SagemakerDomainDefaultUserSettingsRSessionAppSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsRSessionAppSettingsOutputReference
	_jsii_.Get(
		j,
		"rSessionAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) RSessionAppSettingsInput() *SagemakerDomainDefaultUserSettingsRSessionAppSettings {
	var returns *SagemakerDomainDefaultUserSettingsRSessionAppSettings
	_jsii_.Get(
		j,
		"rSessionAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) RStudioServerProAppSettings() SagemakerDomainDefaultUserSettingsRStudioServerProAppSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsRStudioServerProAppSettingsOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) RStudioServerProAppSettingsInput() *SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings {
	var returns *SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings
	_jsii_.Get(
		j,
		"rStudioServerProAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) SharingSettings() SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference
	_jsii_.Get(
		j,
		"sharingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) SharingSettingsInput() *SagemakerDomainDefaultUserSettingsSharingSettings {
	var returns *SagemakerDomainDefaultUserSettingsSharingSettings
	_jsii_.Get(
		j,
		"sharingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) TensorBoardAppSettings() SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsOutputReference
	_jsii_.Get(
		j,
		"tensorBoardAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) TensorBoardAppSettingsInput() *SagemakerDomainDefaultUserSettingsTensorBoardAppSettings {
	var returns *SagemakerDomainDefaultUserSettingsTensorBoardAppSettings
	_jsii_.Get(
		j,
		"tensorBoardAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultUserSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDefaultUserSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultUserSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultUserSettingsOutputReference_Override(s SagemakerDomainDefaultUserSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference)SetInternalValue(val *SagemakerDomainDefaultUserSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) PutCanvasAppSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettings) {
	if err := s.validatePutCanvasAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCanvasAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) PutJupyterServerAppSettings(value *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings) {
	if err := s.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) PutKernelGatewayAppSettings(value *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings) {
	if err := s.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) PutRSessionAppSettings(value *SagemakerDomainDefaultUserSettingsRSessionAppSettings) {
	if err := s.validatePutRSessionAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRSessionAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) PutRStudioServerProAppSettings(value *SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings) {
	if err := s.validatePutRStudioServerProAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRStudioServerProAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) PutSharingSettings(value *SagemakerDomainDefaultUserSettingsSharingSettings) {
	if err := s.validatePutSharingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSharingSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) PutTensorBoardAppSettings(value *SagemakerDomainDefaultUserSettingsTensorBoardAppSettings) {
	if err := s.validatePutTensorBoardAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTensorBoardAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetCanvasAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetRSessionAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetRSessionAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetRStudioServerProAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetRStudioServerProAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetSharingSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetSharingSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ResetTensorBoardAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetTensorBoardAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

