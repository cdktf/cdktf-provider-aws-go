package autoscalingplansscalingplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/autoscalingplansscalingplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingplansScalingPlanScalingInstructionOutputReference interface {
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
	CustomizedLoadMetricSpecification() AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecificationOutputReference
	CustomizedLoadMetricSpecificationInput() *AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecification
	DisableDynamicScaling() interface{}
	SetDisableDynamicScaling(val interface{})
	DisableDynamicScalingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MinCapacity() *float64
	SetMinCapacity(val *float64)
	MinCapacityInput() *float64
	PredefinedLoadMetricSpecification() AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecificationOutputReference
	PredefinedLoadMetricSpecificationInput() *AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecification
	PredictiveScalingMaxCapacityBehavior() *string
	SetPredictiveScalingMaxCapacityBehavior(val *string)
	PredictiveScalingMaxCapacityBehaviorInput() *string
	PredictiveScalingMaxCapacityBuffer() *float64
	SetPredictiveScalingMaxCapacityBuffer(val *float64)
	PredictiveScalingMaxCapacityBufferInput() *float64
	PredictiveScalingMode() *string
	SetPredictiveScalingMode(val *string)
	PredictiveScalingModeInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ScalableDimension() *string
	SetScalableDimension(val *string)
	ScalableDimensionInput() *string
	ScalingPolicyUpdateBehavior() *string
	SetScalingPolicyUpdateBehavior(val *string)
	ScalingPolicyUpdateBehaviorInput() *string
	ScheduledActionBufferTime() *float64
	SetScheduledActionBufferTime(val *float64)
	ScheduledActionBufferTimeInput() *float64
	ServiceNamespace() *string
	SetServiceNamespace(val *string)
	ServiceNamespaceInput() *string
	TargetTrackingConfiguration() AutoscalingplansScalingPlanScalingInstructionTargetTrackingConfigurationList
	TargetTrackingConfigurationInput() interface{}
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
	PutCustomizedLoadMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecification)
	PutPredefinedLoadMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecification)
	PutTargetTrackingConfiguration(value interface{})
	ResetCustomizedLoadMetricSpecification()
	ResetDisableDynamicScaling()
	ResetPredefinedLoadMetricSpecification()
	ResetPredictiveScalingMaxCapacityBehavior()
	ResetPredictiveScalingMaxCapacityBuffer()
	ResetPredictiveScalingMode()
	ResetScalingPolicyUpdateBehavior()
	ResetScheduledActionBufferTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingplansScalingPlanScalingInstructionOutputReference
type jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) CustomizedLoadMetricSpecification() AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecificationOutputReference {
	var returns AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) CustomizedLoadMetricSpecificationInput() *AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecification {
	var returns *AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecification
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) DisableDynamicScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDynamicScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) DisableDynamicScalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDynamicScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) MinCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredefinedLoadMetricSpecification() AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecificationOutputReference {
	var returns AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredefinedLoadMetricSpecificationInput() *AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecification {
	var returns *AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecification
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredictiveScalingMaxCapacityBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredictiveScalingMaxCapacityBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredictiveScalingMaxCapacityBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredictiveScalingMaxCapacityBufferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredictiveScalingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PredictiveScalingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ScalableDimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ScalingPolicyUpdateBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyUpdateBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ScalingPolicyUpdateBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyUpdateBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ScheduledActionBufferTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledActionBufferTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ScheduledActionBufferTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledActionBufferTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ServiceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) TargetTrackingConfiguration() AutoscalingplansScalingPlanScalingInstructionTargetTrackingConfigurationList {
	var returns AutoscalingplansScalingPlanScalingInstructionTargetTrackingConfigurationList
	_jsii_.Get(
		j,
		"targetTrackingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) TargetTrackingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetTrackingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingplansScalingPlanScalingInstructionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutoscalingplansScalingPlanScalingInstructionOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingplansScalingPlanScalingInstructionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingplansScalingPlan.AutoscalingplansScalingPlanScalingInstructionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutoscalingplansScalingPlanScalingInstructionOutputReference_Override(a AutoscalingplansScalingPlanScalingInstructionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingplansScalingPlan.AutoscalingplansScalingPlanScalingInstructionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetDisableDynamicScaling(val interface{}) {
	if err := j.validateSetDisableDynamicScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDynamicScaling",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetMinCapacity(val *float64) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetPredictiveScalingMaxCapacityBehavior(val *string) {
	if err := j.validateSetPredictiveScalingMaxCapacityBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMaxCapacityBehavior",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetPredictiveScalingMaxCapacityBuffer(val *float64) {
	if err := j.validateSetPredictiveScalingMaxCapacityBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMaxCapacityBuffer",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetPredictiveScalingMode(val *string) {
	if err := j.validateSetPredictiveScalingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMode",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetScalableDimension(val *string) {
	if err := j.validateSetScalableDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetScalingPolicyUpdateBehavior(val *string) {
	if err := j.validateSetScalingPolicyUpdateBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingPolicyUpdateBehavior",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetScheduledActionBufferTime(val *float64) {
	if err := j.validateSetScheduledActionBufferTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledActionBufferTime",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetServiceNamespace(val *string) {
	if err := j.validateSetServiceNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PutCustomizedLoadMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionCustomizedLoadMetricSpecification) {
	if err := a.validatePutCustomizedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PutPredefinedLoadMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecification) {
	if err := a.validatePutPredefinedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) PutTargetTrackingConfiguration(value interface{}) {
	if err := a.validatePutTargetTrackingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetTrackingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetCustomizedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetDisableDynamicScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableDynamicScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetPredefinedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetPredictiveScalingMaxCapacityBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetPredictiveScalingMaxCapacityBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetPredictiveScalingMaxCapacityBuffer() {
	_jsii_.InvokeVoid(
		a,
		"resetPredictiveScalingMaxCapacityBuffer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetPredictiveScalingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetPredictiveScalingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetScalingPolicyUpdateBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingPolicyUpdateBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ResetScheduledActionBufferTime() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduledActionBufferTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

