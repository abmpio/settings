package settings

import "sort"

var (
	_settings map[string]*Setting = make(map[string]*Setting)
)

// setup settings
func Setup(settings []Setting) {
	_settings = make(map[string]*Setting)
	if len(settings) <= 0 {
		return
	}

	for i := range settings {
		_settings[settings[i].Name] = &settings[i]
	}
}

// set setting
func Set(name string, value interface{}) {
	if name == "" {
		return
	}
	_settings[name] = &Setting{
		NameValue: NameValue{
			Name:  name,
			Value: value,
		},
	}
}

// remove name
func Remove(name string) {
	delete(_settings, name)
}

func AllSettings() []Setting {
	keys := make([]string, 0)
	for key := range _settings {
		keys = append(keys, key)
	}
	// sort
	sort.Strings(keys)
	settings := make([]Setting, 0)
	for _, eachKey := range keys {
		value, ok := _settings[eachKey]
		if !ok || value == nil {
			continue
		}
		settings = append(settings, *value)
	}
	return settings
}

func GetValue(name string) interface{} {
	settingValue, ok := _settings[name]
	if !ok || settingValue == nil {
		return nil
	}
	v := settingValue.Value()
	return &v
}
