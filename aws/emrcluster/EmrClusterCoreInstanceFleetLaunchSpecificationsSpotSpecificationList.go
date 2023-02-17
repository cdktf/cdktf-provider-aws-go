package emrcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/emrcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList interface {
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
	Get(index *float64) EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList
type jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList {
	_init_.Initialize()

	if err := validateNewEmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList_Override(e EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrCluster.EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) Get(index *float64) EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecificationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

