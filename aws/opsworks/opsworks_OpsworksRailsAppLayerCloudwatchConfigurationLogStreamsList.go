package opsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/opsworks/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList
type jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList {
	_init_.Initialize()

	if err := validateNewOpsworksRailsAppLayerCloudwatchConfigurationLogStreamsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworks.OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList_Override(o OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworks.OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) Get(index *float64) OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksRailsAppLayerCloudwatchConfigurationLogStreamsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
