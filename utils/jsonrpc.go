
package utils

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func JsonRPCRequest(url string, request map[string]interface{}) (map[string]interface{}, error) {
    jsonData, err := json.Marshal(request)
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result, nil
}
