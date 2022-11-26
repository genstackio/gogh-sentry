package gogh_sentry

import (
	"context"
	"github.com/genstackio/gogh"
	"github.com/genstackio/gogh/common"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
)

type Provider struct {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) AddCaptureContext(ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureError(err error, ctx common.CaptureContext) {
	sentry.CaptureException(err)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessage(message string, ctx common.CaptureContext) {
	sentry.CaptureMessage(message)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessages(messages []string, ctx common.CaptureContext) {
	for i := 0; i < len(messages); i++ {
		sentry.CaptureMessage(messages[i])
	}
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureProperty(typ string, data interface{}, ctx common.CaptureContext) {
	// @todo
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureData(bulkData map[string]interface{}, ctx common.CaptureContext) {
	// @todo
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureEvent(event interface{}, ctx common.CaptureContext) {
	// @todo
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTag(tag string, value interface{}, ctx common.CaptureContext) {
	// @todo
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTags(tags map[string]interface{}, ctx common.CaptureContext) {
	// @todo
}

//goland:noinspection GoUnusedParameter
func (p *Provider) Error(err error, ctx common.CaptureContext) {
	p.CaptureError(err, ctx)
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
