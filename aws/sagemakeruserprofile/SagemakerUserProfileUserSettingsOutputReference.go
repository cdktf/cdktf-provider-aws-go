// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakeruserprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerUserProfileUserSettingsOutputReference interface {
	cdktf.ComplexObject
	AutoMountHomeEfs() *string
	SetAutoMountHomeEfs(val *string)
	AutoMountHomeEfsInput() *string
	CanvasAppSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference
	CanvasAppSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettings
	CodeEditorAppSettings() SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference
	CodeEditorAppSettingsInput() *SagemakerUserProfileUserSettingsCodeEditorAppSettings
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
	CustomFileSystemConfig() SagemakerUserProfileUserSettingsCustomFileSystemConfigList
	CustomFileSystemConfigInput() interface{}
	CustomPosixUserConfig() SagemakerUserProfileUserSettingsCustomPosixUserConfigOutputReference
	CustomPosixUserConfigInput() *SagemakerUserProfileUserSettingsCustomPosixUserConfig
	DefaultLandingUri() *string
	SetDefaultLandingUri(val *string)
	DefaultLandingUriInput() *string
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerUserProfileUserSettings
	SetInternalValue(val *SagemakerUserProfileUserSettings)
	JupyterLabAppSettings() SagemakerUserProfileUserSettingsJupyterLabAppSettingsOutputReference
	JupyterLabAppSettingsInput() *SagemakerUserProfileUserSettingsJupyterLabAppSettings
	JupyterServerAppSettings() SagemakerUserProfileUserSettingsJupyterServerAppSettingsOutputReference
	JupyterServerAppSettingsInput() *SagemakerUserProfileUserSettingsJupyterServerAppSettings
	KernelGatewayAppSettings() SagemakerUserProfileUserSettingsKernelGatewayAppSettingsOutputReference
	KernelGatewayAppSettingsInput() *SagemakerUserProfileUserSettingsKernelGatewayAppSettings
	RSessionAppSettings() SagemakerUserProfileUserSettingsRSessionAppSettingsOutputReference
	RSessionAppSettingsInput() *SagemakerUserProfileUserSettingsRSessionAppSettings
	RStudioServerProAppSettings() SagemakerUserProfileUserSettingsRStudioServerProAppSettingsOutputReference
	RStudioServerProAppSettingsInput() *SagemakerUserProfileUserSettingsRStudioServerProAppSettings
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SharingSettings() SagemakerUserProfileUserSettingsSharingSettingsOutputReference
	SharingSettingsInput() *SagemakerUserProfileUserSettingsSharingSettings
	SpaceStorageSettings() SagemakerUserProfileUserSettingsSpaceStorageSettingsOutputReference
	SpaceStorageSettingsInput() *SagemakerUserProfileUserSettingsSpaceStorageSettings
	StudioWebPortal() *string
	SetStudioWebPortal(val *string)
	StudioWebPortalInput() *string
	StudioWebPortalSettings() SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference
	StudioWebPortalSettingsInput() *SagemakerUserProfileUserSettingsStudioWebPortalSettings
	TensorBoardAppSettings() SagemakerUserProfileUserSettingsTensorBoardAppSettingsOutputReference
	TensorBoardAppSettingsInput() *SagemakerUserProfileUserSettingsTensorBoardAppSettings
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
	PutCanvasAppSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettings)
	PutCodeEditorAppSettings(value *SagemakerUserProfileUserSettingsCodeEditorAppSettings)
	PutCustomFileSystemConfig(value interface{})
	PutCustomPosixUserConfig(value *SagemakerUserProfileUserSettingsCustomPosixUserConfig)
	PutJupyterLabAppSettings(value *SagemakerUserProfileUserSettingsJupyterLabAppSettings)
	PutJupyterServerAppSettings(value *SagemakerUserProfileUserSettingsJupyterServerAppSettings)
	PutKernelGatewayAppSettings(value *SagemakerUserProfileUserSettingsKernelGatewayAppSettings)
	PutRSessionAppSettings(value *SagemakerUserProfileUserSettingsRSessionAppSettings)
	PutRStudioServerProAppSettings(value *SagemakerUserProfileUserSettingsRStudioServerProAppSettings)
	PutSharingSettings(value *SagemakerUserProfileUserSettingsSharingSettings)
	PutSpaceStorageSettings(value *SagemakerUserProfileUserSettingsSpaceStorageSettings)
	PutStudioWebPortalSettings(value *SagemakerUserProfileUserSettingsStudioWebPortalSettings)
	PutTensorBoardAppSettings(value *SagemakerUserProfileUserSettingsTensorBoardAppSettings)
	ResetAutoMountHomeEfs()
	ResetCanvasAppSettings()
	ResetCodeEditorAppSettings()
	ResetCustomFileSystemConfig()
	ResetCustomPosixUserConfig()
	ResetDefaultLandingUri()
	ResetJupyterLabAppSettings()
	ResetJupyterServerAppSettings()
	ResetKernelGatewayAppSettings()
	ResetRSessionAppSettings()
	ResetRStudioServerProAppSettings()
	ResetSecurityGroups()
	ResetSharingSettings()
	ResetSpaceStorageSettings()
	ResetStudioWebPortal()
	ResetStudioWebPortalSettings()
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

// The jsii proxy struct for SagemakerUserProfileUserSettingsOutputReference
type jsiiProxy_SagemakerUserProfileUserSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) AutoMountHomeEfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) AutoMountHomeEfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CanvasAppSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference
	_jsii_.Get(
		j,
		"canvasAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CanvasAppSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettings
	_jsii_.Get(
		j,
		"canvasAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CodeEditorAppSettings() SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCodeEditorAppSettingsOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CodeEditorAppSettingsInput() *SagemakerUserProfileUserSettingsCodeEditorAppSettings {
	var returns *SagemakerUserProfileUserSettingsCodeEditorAppSettings
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CustomFileSystemConfig() SagemakerUserProfileUserSettingsCustomFileSystemConfigList {
	var returns SagemakerUserProfileUserSettingsCustomFileSystemConfigList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CustomPosixUserConfig() SagemakerUserProfileUserSettingsCustomPosixUserConfigOutputReference {
	var returns SagemakerUserProfileUserSettingsCustomPosixUserConfigOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) CustomPosixUserConfigInput() *SagemakerUserProfileUserSettingsCustomPosixUserConfig {
	var returns *SagemakerUserProfileUserSettingsCustomPosixUserConfig
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) DefaultLandingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) DefaultLandingUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) InternalValue() *SagemakerUserProfileUserSettings {
	var returns *SagemakerUserProfileUserSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) JupyterLabAppSettings() SagemakerUserProfileUserSettingsJupyterLabAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsJupyterLabAppSettingsOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) JupyterLabAppSettingsInput() *SagemakerUserProfileUserSettingsJupyterLabAppSettings {
	var returns *SagemakerUserProfileUserSettingsJupyterLabAppSettings
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) JupyterServerAppSettings() SagemakerUserProfileUserSettingsJupyterServerAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsJupyterServerAppSettingsOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) JupyterServerAppSettingsInput() *SagemakerUserProfileUserSettingsJupyterServerAppSettings {
	var returns *SagemakerUserProfileUserSettingsJupyterServerAppSettings
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) KernelGatewayAppSettings() SagemakerUserProfileUserSettingsKernelGatewayAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsKernelGatewayAppSettingsOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) KernelGatewayAppSettingsInput() *SagemakerUserProfileUserSettingsKernelGatewayAppSettings {
	var returns *SagemakerUserProfileUserSettingsKernelGatewayAppSettings
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) RSessionAppSettings() SagemakerUserProfileUserSettingsRSessionAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsRSessionAppSettingsOutputReference
	_jsii_.Get(
		j,
		"rSessionAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) RSessionAppSettingsInput() *SagemakerUserProfileUserSettingsRSessionAppSettings {
	var returns *SagemakerUserProfileUserSettingsRSessionAppSettings
	_jsii_.Get(
		j,
		"rSessionAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) RStudioServerProAppSettings() SagemakerUserProfileUserSettingsRStudioServerProAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsRStudioServerProAppSettingsOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) RStudioServerProAppSettingsInput() *SagemakerUserProfileUserSettingsRStudioServerProAppSettings {
	var returns *SagemakerUserProfileUserSettingsRStudioServerProAppSettings
	_jsii_.Get(
		j,
		"rStudioServerProAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) SharingSettings() SagemakerUserProfileUserSettingsSharingSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsSharingSettingsOutputReference
	_jsii_.Get(
		j,
		"sharingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) SharingSettingsInput() *SagemakerUserProfileUserSettingsSharingSettings {
	var returns *SagemakerUserProfileUserSettingsSharingSettings
	_jsii_.Get(
		j,
		"sharingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) SpaceStorageSettings() SagemakerUserProfileUserSettingsSpaceStorageSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsSpaceStorageSettingsOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) SpaceStorageSettingsInput() *SagemakerUserProfileUserSettingsSpaceStorageSettings {
	var returns *SagemakerUserProfileUserSettingsSpaceStorageSettings
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) StudioWebPortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) StudioWebPortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) StudioWebPortalSettings() SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsStudioWebPortalSettingsOutputReference
	_jsii_.Get(
		j,
		"studioWebPortalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) StudioWebPortalSettingsInput() *SagemakerUserProfileUserSettingsStudioWebPortalSettings {
	var returns *SagemakerUserProfileUserSettingsStudioWebPortalSettings
	_jsii_.Get(
		j,
		"studioWebPortalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) TensorBoardAppSettings() SagemakerUserProfileUserSettingsTensorBoardAppSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsTensorBoardAppSettingsOutputReference
	_jsii_.Get(
		j,
		"tensorBoardAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) TensorBoardAppSettingsInput() *SagemakerUserProfileUserSettingsTensorBoardAppSettings {
	var returns *SagemakerUserProfileUserSettingsTensorBoardAppSettings
	_jsii_.Get(
		j,
		"tensorBoardAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerUserProfileUserSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerUserProfileUserSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerUserProfileUserSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerUserProfileUserSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerUserProfileUserSettingsOutputReference_Override(s SagemakerUserProfileUserSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetAutoMountHomeEfs(val *string) {
	if err := j.validateSetAutoMountHomeEfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMountHomeEfs",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetDefaultLandingUri(val *string) {
	if err := j.validateSetDefaultLandingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLandingUri",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetInternalValue(val *SagemakerUserProfileUserSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetStudioWebPortal(val *string) {
	if err := j.validateSetStudioWebPortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioWebPortal",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutCanvasAppSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettings) {
	if err := s.validatePutCanvasAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCanvasAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutCodeEditorAppSettings(value *SagemakerUserProfileUserSettingsCodeEditorAppSettings) {
	if err := s.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := s.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutCustomPosixUserConfig(value *SagemakerUserProfileUserSettingsCustomPosixUserConfig) {
	if err := s.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutJupyterLabAppSettings(value *SagemakerUserProfileUserSettingsJupyterLabAppSettings) {
	if err := s.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutJupyterServerAppSettings(value *SagemakerUserProfileUserSettingsJupyterServerAppSettings) {
	if err := s.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutKernelGatewayAppSettings(value *SagemakerUserProfileUserSettingsKernelGatewayAppSettings) {
	if err := s.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutRSessionAppSettings(value *SagemakerUserProfileUserSettingsRSessionAppSettings) {
	if err := s.validatePutRSessionAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRSessionAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutRStudioServerProAppSettings(value *SagemakerUserProfileUserSettingsRStudioServerProAppSettings) {
	if err := s.validatePutRStudioServerProAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRStudioServerProAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutSharingSettings(value *SagemakerUserProfileUserSettingsSharingSettings) {
	if err := s.validatePutSharingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSharingSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutSpaceStorageSettings(value *SagemakerUserProfileUserSettingsSpaceStorageSettings) {
	if err := s.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutStudioWebPortalSettings(value *SagemakerUserProfileUserSettingsStudioWebPortalSettings) {
	if err := s.validatePutStudioWebPortalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStudioWebPortalSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) PutTensorBoardAppSettings(value *SagemakerUserProfileUserSettingsTensorBoardAppSettings) {
	if err := s.validatePutTensorBoardAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTensorBoardAppSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetAutoMountHomeEfs() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoMountHomeEfs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetCanvasAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetDefaultLandingUri() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultLandingUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetRSessionAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetRSessionAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetRStudioServerProAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetRStudioServerProAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetSharingSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetSharingSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetStudioWebPortal() {
	_jsii_.InvokeVoid(
		s,
		"resetStudioWebPortal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetStudioWebPortalSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetStudioWebPortalSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetTensorBoardAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetTensorBoardAppSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

