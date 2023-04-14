package iotindexingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/iotindexingconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList interface {
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
	Get(index *float64) IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList
type jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewIotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList {
	_init_.Initialize()

	if err := validateNewIotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList{}

	_jsii_.Create(
		"@cdktf/provider-aws.iotIndexingConfiguration.IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewIotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList_Override(i IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.iotIndexingConfiguration.IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		i,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (i *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) Get(index *float64) IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldOutputReference {
	if err := i.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldOutputReference

	_jsii_.Invoke(
		i,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotIndexingConfigurationThingGroupIndexingConfigurationCustomFieldList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

