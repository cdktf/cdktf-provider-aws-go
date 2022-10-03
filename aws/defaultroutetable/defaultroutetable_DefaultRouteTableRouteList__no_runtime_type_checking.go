//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package defaultroutetable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultRouteTableRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DefaultRouteTableRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DefaultRouteTableRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DefaultRouteTableRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DefaultRouteTableRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DefaultRouteTableRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDefaultRouteTableRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

