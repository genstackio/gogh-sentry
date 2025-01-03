package gogh_sentry

import (
	"context"
	"fmt"
	"github.com/genstackio/gogh"
	"github.com/genstackio/gogh/common"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
	"time"
)

type Provider struct {
}

func convertMapStringInterfaceToMapStringString(original map[string]interface{}) (map[string]string, error) {
	converted := make(map[string]string)

	for key, value := range original {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("value for key %s is not a string", key)
		}
		converted[key] = strValue
	}

	return converted, nil
}
func wrapWithCaptureContext(ctx common.CaptureContext, fn func()) {
	sentry.WithScope(func(scope *sentry.Scope) {
		if ctx.Tags != nil {
			convertedTags, err := convertMapStringInterfaceToMapStringString(ctx.Tags)
			if nil == err {
				scope.SetTags(convertedTags)
			}
		}
		if ctx.Data != nil {
			convertedExtras, ok := ctx.Data.(map[string]interface{})
			if ok {
				scope.SetContext("extras", convertedExtras)
			}
		}
		if "" != ctx.Tag.Key {
			scope.SetTag(ctx.Tag.Key, ctx.Tag.Value.(string))
		}
		fn()
	})

}

//goland:noinspection GoUnusedParameter
func (p *Provider) AddCaptureContext(ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureError(err error, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		sentry.CaptureException(err)
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessage(message string, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		sentry.CaptureMessage(message)
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessages(messages []string, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		for i := 0; i < len(messages); i++ {
			sentry.CaptureMessage(messages[i])
		}
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureProperty(typ string, data interface{}, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		// @todo
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureData(bulkData map[string]interface{}, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		// @todo
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureEvent(event interface{}, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		// @todo
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTag(tag string, value interface{}, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		// @todo
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTags(tags map[string]interface{}, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		// @todo
	})
}

//goland:noinspection GoUnusedParameter
func (p *Provider) Error(err error, ctx common.CaptureContext) {
	wrapWithCaptureContext(ctx, func() {
		p.CaptureError(err, ctx)
	})
}

//goland:noinspection GoUnusedExportedFunction
func (p *Provider) Clean() {
	sentry.Flush(2 * time.Second)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) Wrap(h common.HandlerFn) common.HandlerFn {
	return func(ctx context.Context, payload []byte) ([]byte, error) {
		return h(ctx, payload)
	}
}

//goland:noinspection GoUnusedExportedFunction
func Create() common.Provider {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Println("gogh-sentry: unable to initialize sentry: " + err.Error())
	}
	return &Provider{}
}

//goland:noinspection GoUnusedExportedFunction
func Register() {
	err := gogh.RegisterProvider("sentry", Create)
	if err != nil {
		log.Println("gogh-sentry: unable to register provider")
	}
}
