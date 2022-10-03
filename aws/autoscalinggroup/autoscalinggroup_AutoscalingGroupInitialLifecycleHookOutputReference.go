package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupInitialLifecycleHookOutputReference interface {
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
	DefaultResult() *string
	SetDefaultResult(val *string)
	DefaultResultInput() *string
	// Experimental.
	Fqn() *string
	HeartbeatTimeout() *float64
	SetHeartbeatTimeout(val *float64)
	HeartbeatTimeoutInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifecycleTransition() *string
	SetLifecycleTransition(val *string)
	LifecycleTransitionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NotificationMetadata() *string
	SetNotificationMetadata(val *string)
	NotificationMetadataInput() *string
	NotificationTargetArn() *string
	SetNotificationTargetArn(val *string)
	NotificationTargetArnInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetDefaultResult()
	ResetHeartbeatTimeout()
	ResetNotificationMetadata()
	ResetNotificationTargetArn()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupInitialLifecycleHookOutputReference
type jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) DefaultResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) DefaultResultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) HeartbeatTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) HeartbeatTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) LifecycleTransition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) LifecycleTransitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) NotificationMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) NotificationMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) NotificationTargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) NotificationTargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupInitialLifecycleHookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutoscalingGroupInitialLifecycleHookOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupInitialLifecycleHookOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupInitialLifecycleHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutoscalingGroupInitialLifecycleHookOutputReference_Override(a AutoscalingGroupInitialLifecycleHookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupInitialLifecycleHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetDefaultResult(val *string) {
	if err := j.validateSetDefaultResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultResult",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetHeartbeatTimeout(val *float64) {
	if err := j.validateSetHeartbeatTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatTimeout",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetLifecycleTransition(val *string) {
	if err := j.validateSetLifecycleTransitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleTransition",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetNotificationMetadata(val *string) {
	if err := j.validateSetNotificationMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationMetadata",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetNotificationTargetArn(val *string) {
	if err := j.validateSetNotificationTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTargetArn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ResetDefaultResult() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResult",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ResetHeartbeatTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetHeartbeatTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ResetNotificationMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ResetNotificationTargetArn() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationTargetArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

