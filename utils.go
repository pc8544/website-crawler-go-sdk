package websitecrawler

import "encoding/json"

// mapToStruct converts a generic map[string]interface{} into a typed struct.
func mapToStruct(m map[string]interface{}, target interface{}) error {
    data, err := json.Marshal(m)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, target)
}
