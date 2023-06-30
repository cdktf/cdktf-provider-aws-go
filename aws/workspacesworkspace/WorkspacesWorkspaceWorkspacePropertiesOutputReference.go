package workspacesworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/workspacesworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspacesWorkspaceWorkspacePropertiesOutputReference interface {
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
	ComputeTypeName() *string
	SetComputeTypeName(val *string)
	ComputeTypeNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WorkspacesWorkspaceWorkspaceProperties
	SetInternalValue(val *WorkspacesWorkspaceWorkspaceProperties)
	RootVolumeSizeGib() *float64
	SetRootVolumeSizeGib(val *float64)
	RootVolumeSizeGibInput() *float64
	RunningMode() *string
	SetRunningMode(val *string)
	RunningModeAutoStopTimeoutInMinutes() *float64
	SetRunningModeAutoStopTimeoutInMinutes(val *float64)
	RunningModeAutoStopTimeoutInMinutesInput() *float64
	RunningModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserVolumeSizeGib() *float64
	SetUserVolumeSizeGib(val *float64)
	UserVolumeSizeGibInput() *float64
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
	ResetComputeTypeName()
	ResetRootVolumeSizeGib()
	ResetRunningMode()
	ResetRunningModeAutoStopTimeoutInMinutes()
	ResetUserVolumeSizeGib()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesWorkspaceWorkspacePropertiesOutputReference
type jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ComputeTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ComputeTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) InternalValue() *WorkspacesWorkspaceWorkspaceProperties {
	var returns *WorkspacesWorkspaceWorkspaceProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RootVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RootVolumeSizeGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningModeAutoStopTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningModeAutoStopTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningModeAutoStopTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningModeAutoStopTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) UserVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) UserVolumeSizeGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userVolumeSizeGibInput",
		&returns,
	)
	return returns
}


func NewWorkspacesWorkspaceWorkspacePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkspacesWorkspaceWorkspacePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesWorkspaceWorkspacePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.workspacesWorkspace.WorkspacesWorkspaceWorkspacePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesWorkspaceWorkspacePropertiesOutputReference_Override(w WorkspacesWorkspaceWorkspacePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.workspacesWorkspace.WorkspacesWorkspaceWorkspacePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetComputeTypeName(val *string) {
	if err := j.validateSetComputeTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeTypeName",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetInternalValue(val *WorkspacesWorkspaceWorkspaceProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetRootVolumeSizeGib(val *float64) {
	if err := j.validateSetRootVolumeSizeGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootVolumeSizeGib",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetRunningMode(val *string) {
	if err := j.validateSetRunningModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runningMode",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetRunningModeAutoStopTimeoutInMinutes(val *float64) {
	if err := j.validateSetRunningModeAutoStopTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runningModeAutoStopTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference)SetUserVolumeSizeGib(val *float64) {
	if err := j.validateSetUserVolumeSizeGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userVolumeSizeGib",
		val,
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetComputeTypeName() {
	_jsii_.InvokeVoid(
		w,
		"resetComputeTypeName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetRootVolumeSizeGib() {
	_jsii_.InvokeVoid(
		w,
		"resetRootVolumeSizeGib",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetRunningMode() {
	_jsii_.InvokeVoid(
		w,
		"resetRunningMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetRunningModeAutoStopTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetRunningModeAutoStopTimeoutInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetUserVolumeSizeGib() {
	_jsii_.InvokeVoid(
		w,
		"resetUserVolumeSizeGib",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

