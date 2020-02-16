/*
 * MailSlurp API
 *
 * ## Introduction  [MailSlurp](https://www.mailslurp.com) is an Email API for developers and QA testers. It let's users: - create emails addresses on demand - receive emails and attachments in code - send templated HTML emails  ## About  This page contains the REST API documentation for MailSlurp. All requests require API Key authentication passed as an `x-api-key` header.  Create an account to [get your free API Key](https://app.mailslurp.com/sign-up/).  ## Resources - 🔑 [Get API Key](https://app.mailslurp.com/sign-up/)                    - 🎓 [Developer Portal](https://www.mailslurp.com/docs/)           - 📦 [Library SDKs](https://www.mailslurp.com/docs/) - ✍️ [Code Examples](https://www.mailslurp.com/examples) - ⚠️ [Report an issue](https://drift.me/mailslurp)  ## Explore  
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// PageWebhookProjection struct for PageWebhookProjection
type PageWebhookProjection struct {
	Content []WebhookProjection `json:"content,omitempty"`
	Empty bool `json:"empty,omitempty"`
	First bool `json:"first,omitempty"`
	Last bool `json:"last,omitempty"`
	Number int32 `json:"number,omitempty"`
	NumberOfElements int32 `json:"numberOfElements,omitempty"`
	Pageable Pageable `json:"pageable,omitempty"`
	Size int32 `json:"size,omitempty"`
	Sort Sort `json:"sort,omitempty"`
	TotalElements int64 `json:"totalElements,omitempty"`
	TotalPages int32 `json:"totalPages,omitempty"`
}
