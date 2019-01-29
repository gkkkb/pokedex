package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gkkkb/pokedex/pkg/currentuser"
	"github.com/gkkkb/pokedex/pkg/resource"
	plog "github.com/gkkkb/piston/log"
)

// DevLog logs only on development or staging
func DevLog(v ...interface{}) {
	if os.Getenv("ENV") == "development" || os.Getenv("ENV") == "staging" {
		log.Println(v)
	}
}

// ErrLog logs errors with packen.RequestError
func ErrLog(ctx context.Context, err error, category, message string) {
	if os.Getenv("ENV") != "test" {
		res := resource.FromContext(ctx)
		plog.RequestError(fmt.Sprintf("%s", err.Error()), map[string]interface{}{
			"request_id": res.RequestID,
			"tags":       append([]string{"post"}, res.Action, category),
			"message":    message,
			"duration":   strconv.FormatFloat(time.Since(res.StartTime).Seconds(), 'f', -1, 64),
		})
	}
}

// InfoLog logs informations with packen.RequestInfo
func InfoLog(ctx context.Context, message string, tags ...string) {
	if os.Getenv("ENV") != "test" {
		res := resource.FromContext(ctx)
		user := currentuser.FromContext(ctx)
		plog.RequestInfo("", map[string]interface{}{
			"request_id": res.RequestID,
			"tags":       append([]string{"post"}, tags...),
			"message":    fmt.Sprintf("%d: %s", user.ID, message),
			"duration":   strconv.FormatFloat(time.Since(res.StartTime).Seconds(), 'f', -1, 64),
		})
	}
}

// InfoLog logs with additional informations with packen.RequestInfo
func AdditionalInfoLog(ctx context.Context, message string, mapInfo map[string]interface{}, tags ...string) {
	if os.Getenv("ENV") != "test" {
		res := resource.FromContext(ctx)
		user := currentuser.FromContext(ctx)
		mapInfo["request_id"] = res.RequestID
		mapInfo["tags"] = append([]string{"post"}, tags...)
		mapInfo["message"] = fmt.Sprintf("%d: %s", user.ID, message)
		mapInfo["duration"] = strconv.FormatFloat(time.Since(res.StartTime).Seconds(), 'f', -1, 64)
		plog.RequestInfo("", mapInfo)
	}
}

func Fatal(v ...interface{}) {
	log.Fatal(v)
}
