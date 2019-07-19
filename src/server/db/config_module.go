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
	}
	return str
}
