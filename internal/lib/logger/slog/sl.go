package sl

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Req(req interface{}) string {
	reqJSON, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error marshaling request to JSON:", err)
		return ""
	}
	return "user request:" + string(reqJSON)
}
