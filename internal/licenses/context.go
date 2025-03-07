package licenses

import (
	"context"
)

type licenseScannerKey struct{}

func SetContextLicenseScanner(ctx context.Context, s Scanner) context.Context {
	return context.WithValue(ctx, licenseScannerKey{}, s)
}

func ContextLicenseScanner(ctx context.Context) (Scanner, error) {
	if s, ok := ctx.Value(licenseScannerKey{}).(Scanner); ok {
		return s, nil
	}
	return NewDefaultScanner()
}
