package helper

import (
	"encoding/json"

	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

func HomeContentToMap(hc *model.HomeContent) (map[string]interface{}, error) {
	var value interface{}
	if hc.IsJson {
		err := json.Unmarshal([]byte(hc.Value), &value)
		if err != nil {
			return nil, err
		}
	} else {
		value = hc.Value
	}
	return map[string]interface{}{
		"section":    hc.Section,
		"key":        hc.Key,
		"value":      value,
		"is_json":    hc.IsJson,
		"updated_at": hc.UpdatedAt,
	}, nil
}

func HomeContentSliceToMapSlice(hcs []*model.HomeContent) ([]map[string]interface{}, error) {
	var maps []map[string]interface{}
	for _, p := range hcs {
		mapHomeContent, err := HomeContentToMap(p)
		if err != nil {
			return nil, err
		}
		maps = append(maps, mapHomeContent)
	}
	return maps, nil
}
