package cfg

var MerchantConfig = &merchantConfig{}

type merchantConfig struct {
	MccCategory      mccCategoryConfig
	BusinessCategory map[string]string
}

type mccCategoryConfig struct {
	MccCategoryCombine [][3]string
}

func (*merchantConfig) ConfigName() string {
	return "merchant_config.toml"
}
