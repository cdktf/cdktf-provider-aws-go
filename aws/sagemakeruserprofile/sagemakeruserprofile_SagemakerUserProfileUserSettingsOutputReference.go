package sagemakeruserprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/sagemakeruserprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerUserProfileUserSettingsOutputReference interface {
	cdktf.ComplexObject
	CanvasAppSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference
	CanvasAppSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettings
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
	InternalValue() *SagemakerUserProfileUserSettings
	SetInternalValue(val *SagemakerUserProfileUserSettings)
	JupyterServerAppSettings() SagemakerUserProfileUserSettingsJupyterServerAppSettingsOutputReference
	JupyterServerAppSettingsInput() *SagemakerUserProfileUserSettingsJupyterServerAppSettings
	KernelGatewayAppSettings() SagemakerUserProfileUserSettingsKernelGatewayAppSettingsOutputReference
	KernelGatewayAppSettingsInput() *SagemakerUserProfileUserSettingsKernelGatewayAppSettings
	RSessionAppSettings() SagemakerUserProfileUserSettingsRSessionAppSettingsOutputReference
	RSessionAppSettingsInput() *SagemakerUserProfileUserSettingsRSessionAppSettings
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SharingSettings() SagemakerUserProfileUserSettingsSharingSettingsOutputReference
	SharingSettingsInput() *SagemakerUserProfileUserSettingsSharingSettings
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
	PutJupyterServerAppSettings(value *SagemakerUserProfileUserSettingsJupyterServerAppSettings)
	PutKernelGatewayAppSettings(value *SagemakerUserProfileUserSettingsKernelGatewayAppSettings)
	PutRSessionAppSettings(value *SagemakerUserProfileUserSettingsRSessionAppSettings)
	PutSharingSettings(value *SagemakerUserProfileUserSettingsSharingSettings)
	PutTensorBoardAppSettings(value *SagemakerUserProfileUserSettingsTensorBoardAppSettings)
	ResetCanvasAppSettings()
	ResetJupyterServerAppSettings()
	ResetKernelGatewayAppSettings()
	ResetRSessionAppSettings()
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

// The jsii proxy struct for SagemakerUserProfileUserSettingsOutputReference
type jsiiProxy_SagemakerUserProfileUserSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (s *jsiiProxy_SagemakerUserProfileUserSettingsOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetCanvasAppSettings",
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

