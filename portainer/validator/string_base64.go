package validator

import (
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"golang.org/x/net/context"
)

type stringBase64Validator struct {
}

// Description returns a plain text description
func (v stringBase64Validator) Description(ctx context.Context) string {
	return fmt.Sprintf("string must be Base64-encoded")
}

// MarkdownDescription returns a markdown formatted description
func (v stringBase64Validator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("string must be Base64-encoded")
}

func (v stringBase64Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	str := req.ConfigValue.ValueString()
	_, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid String value",
			fmt.Sprintf("String must must be Base64-encoded. Either provide an encoded value, or use the `base64encode` function"),
		)

		return
	}
}

// StringBase64 returns a validator which ensures that any configured
// attribute value is StringBase64-encoded. Null (unconfigured) and unknown (known after apply)
// values are skipped.
func StringBase64() validator.String {
	return stringBase64Validator{}
}
