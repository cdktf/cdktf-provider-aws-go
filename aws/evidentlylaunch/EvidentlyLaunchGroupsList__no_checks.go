//go:build no_runtime_type_checking

package evidentlylaunch

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyLaunchGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyLaunchGroupsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyLaunchGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

