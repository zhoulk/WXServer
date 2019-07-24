package db

import (
	"bytes"
	"encoding/json"
)

// CostOfCloth ...
func (m *Module) CostOfCloth(t int32, level int32) string {
	for _, config := range m.clothConfigs {
		if config.Type == t && config.Level == level {
			return config.Cost
		}
	}
	return "[0]"
}

// GetConfigStr ...
func (m *Module) GetConfigStr(t int32) string {
	var str = "{}"
	if t == 1 {
		bs, err := json.Marshal(m.clothConfigs)
		if err == nil {
			str = bytes.NewBuffer(bs).String()
		}
	} else if t == 2 {
		bs, err := json.Marshal(m.sceneConfigs)
		if err == nil {
			str = bytes.NewBuffer(bs).String()
		}
	} else if t == 3 {
		bs, err := json.Marshal(m.levelConfigs)
		if err == nil {
			str = bytes.NewBuffer(bs).String()
		}
	} else if t == 4 {
		bs, err := json.Marshal(m.signConfigs)
		if err == nil {
			str = bytes.NewBuffer(bs).String()
		}
	}
	return str
}
