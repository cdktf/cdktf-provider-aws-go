//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueCrawlerDeltaTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueCrawlerDeltaTargetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerDeltaTargetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerDeltaTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerDeltaTargetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerDeltaTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueCrawlerDeltaTargetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

