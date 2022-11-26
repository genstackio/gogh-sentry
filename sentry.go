package gogh_sentry

import (
	"context"
	"github.com/genstackio/gogh"
	"github.com/genstackio/gogh/common"
	"log"
)

type Provider struct {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) AddCaptureContext(ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureError(err error, ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessage(message string, ctx common.CaptureContext) {
	log.Println("BLABLA", message)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessages(messages []string, ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureProperty(typ string, data interface{}, ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureData(bulkData map[string]interface{}, ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureEvent(event interface{}, ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTag(tag string, value interface{}, ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTags(tags map[string]interface{}, ctx common.CaptureContext) {
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
	return &Provider{}
}

//goland:noinspection GoUnusedExportedFunction
func Register() {
	err := gogh.RegisterProvider("sentry", Create)
	if err != nil {
		log.Println("gogh-sentry: unable to register provider")
	}
}
