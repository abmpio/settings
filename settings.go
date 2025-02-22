package settings

import (
	"sort"
	"sync"

	"github.com/abmpio/libx/lang"
)

var (
	_settings map[string]map[string]*Setting = map[string]map[string]*Setting{}
	_rwLock   sync.RWMutex
)

func ensureAppExist(appName string) map[string]*Setting {
	_rwLock.Lock()
	defer _rwLock.Unlock()
	appSettings, ok := _settings[appName]
	if !ok {
		appSettings = make(map[string]*Setting)
		_settings[appName] = appSettings
	}
	return appSettings
}

// setup settings
func Setup(settings []*Setting) {
	_settings = map[string]map[string]*Setting{}
	if len(settings) <= 0 {
		return
	}

	for _, eachSetting := range settings {
		appSettings := ensureAppExist(eachSetting.AppName)
		appSettings[eachSetting.Name] = eachSetting
	}
}

// set setting
func Set(appName string, name string, value interface{}) {
	if name == "" {
		return
	}
	appSettings := ensureAppExist(appName)
	appSettings[name] = &Setting{
		NameValue: lang.NameValue{
			Name:  name,
			Value: value,
		},
	}
}

// remove name
func Remove(appName string, name string) {
	appSettings := ensureAppExist(appName)
	delete(appSettings, name)
}

func AllSettings(appName string) []Setting {
	appSettings := ensureAppExist(appName)

	keys := make([]string, 0)
	for key := range appSettings {
		keys = append(keys, key)
	}
	// sort
	sort.Strings(keys)
	settings := make([]Setting, 0)
	for _, eachKey := range keys {
		value, ok := appSettings[eachKey]
		if !ok || value == nil {
			continue
		}
		settings = append(settings, *value)
	}
	return settings
}

// get Setting by name
func GetSetting(appName string, name string) *Setting {
	appSettings := ensureAppExist(appName)
	settingValue, ok := appSettings[name]
	if !ok || settingValue == nil {
		return nil
	}
	return settingValue
}

// get Setting by name
func GetSettingInAll(name string) *Setting {
	_rwLock.RLock()
	defer _rwLock.RUnlock()

	for _, eachSettings := range _settings {
		if eachSettings == nil {
			continue
		}
		settingValue, ok := eachSettings[name]
		if !ok {
			continue
		}
		return settingValue
	}
	return nil
}

// get value by name and appName
func GetValue(appName string, name string) interface{} {
	appSettings := ensureAppExist(appName)
	settingValue, ok := appSettings[name]
	if !ok || settingValue == nil {
		return nil
	}
	v := settingValue.Value()
	return &v
}

func GetValueAsString(appName string, name string) string {
	appSettings := ensureAppExist(appName)
	settingValue, ok := appSettings[name]
	if !ok || settingValue == nil {
		return ""
	}
	return settingValue.ValueAsString()
}

// find setting in all app
func GetValueAsStringInAll(name string) string {
	_rwLock.RLock()
	defer _rwLock.RUnlock()

	for _, eachSettings := range _settings {
		if eachSettings == nil {
			continue
		}
		settingValue, ok := eachSettings[name]
		if !ok {
			continue
		}
		return settingValue.ValueAsString()
	}
	return ""
}
