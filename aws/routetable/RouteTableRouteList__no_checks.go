//go:build no_runtime_type_checking

package routetable

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteTableRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouteTableRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouteTableRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouteTableRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouteTableRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouteTableRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouteTableRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

