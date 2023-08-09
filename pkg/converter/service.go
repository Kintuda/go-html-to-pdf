package converter

import "context"

type HtmlConversionProvider interface {
	SendSMS(ctx context.Context, content string) (string, error)
}
