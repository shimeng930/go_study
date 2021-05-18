package cfg

var includeConfigs = map[string]includeableConfig{}

func registerIncludeConfigs() {
	includeConfigs[MerchantConfig.ConfigName()] = MerchantConfig
}

func getIncludeConfig(configName string) interface{} {
	return includeConfigs[configName]
}

type includeableConfig interface {
	ConfigName() string
}
