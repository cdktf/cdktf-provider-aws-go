package workspacesdirectory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/workspacesdirectory/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference interface {
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
	CustomSecurityGroupId() *string
	SetCustomSecurityGroupId(val *string)
	CustomSecurityGroupIdInput() *string
	DefaultOu() *string
	SetDefaultOu(val *string)
	DefaultOuInput() *string
	EnableInternetAccess() interface{}
	SetEnableInternetAccess(val interface{})
	EnableInternetAccessInput() interface{}
	EnableMaintenanceMode() interface{}
	SetEnableMaintenanceMode(val interface{})
	EnableMaintenanceModeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WorkspacesDirectoryWorkspaceCreationProperties
	SetInternalValue(val *WorkspacesDirectoryWorkspaceCreationProperties)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserEnabledAsLocalAdministrator() interface{}
	SetUserEnabledAsLocalAdministrator(val interface{})
	UserEnabledAsLocalAdministratorInput() interface{}
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
	ResetCustomSecurityGroupId()
	ResetDefaultOu()
	ResetEnableInternetAccess()
	ResetEnableMaintenanceMode()
	ResetUserEnabledAsLocalAdministrator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference
type jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) CustomSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) CustomSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) DefaultOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) DefaultOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableMaintenanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableMaintenanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) InternalValue() *WorkspacesDirectoryWorkspaceCreationProperties {
	var returns *WorkspacesDirectoryWorkspaceCreationProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) UserEnabledAsLocalAdministrator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) UserEnabledAsLocalAdministratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministratorInput",
		&returns,
	)
	return returns
}


func NewWorkspacesDirectoryWorkspaceCreationPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesDirectoryWorkspaceCreationPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.workspacesDirectory.WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesDirectoryWorkspaceCreationPropertiesOutputReference_Override(w WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.workspacesDirectory.WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetCustomSecurityGroupId(val *string) {
	if err := j.validateSetCustomSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetDefaultOu(val *string) {
	if err := j.validateSetDefaultOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOu",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetEnableInternetAccess(val interface{}) {
	if err := j.validateSetEnableInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInternetAccess",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetEnableMaintenanceMode(val interface{}) {
	if err := j.validateSetEnableMaintenanceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMaintenanceMode",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetInternalValue(val *WorkspacesDirectoryWorkspaceCreationProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)SetUserEnabledAsLocalAdministrator(val interface{}) {
	if err := j.validateSetUserEnabledAsLocalAdministratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEnabledAsLocalAdministrator",
		val,
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetCustomSecurityGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomSecurityGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetDefaultOu() {
	_jsii_.InvokeVoid(
		w,
		"resetDefaultOu",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetEnableInternetAccess() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableInternetAccess",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetEnableMaintenanceMode() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableMaintenanceMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetUserEnabledAsLocalAdministrator() {
	_jsii_.InvokeVoid(
		w,
		"resetUserEnabledAsLocalAdministrator",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

