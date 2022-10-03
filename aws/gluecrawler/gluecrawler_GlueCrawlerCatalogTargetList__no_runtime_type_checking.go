//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package gluecrawler

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueCrawlerCatalogTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueCrawlerCatalogTargetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerCatalogTargetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerCatalogTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerCatalogTargetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueCrawlerCatalogTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueCrawlerCatalogTargetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

