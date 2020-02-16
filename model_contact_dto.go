/*
 * MailSlurp API
 *
 * ## Introduction  [MailSlurp](https://www.mailslurp.com) is an Email API for developers and QA testers. It let's users: - create emails addresses on demand - receive emails and attachments in code - send templated HTML emails  ## About  This page contains the REST API documentation for MailSlurp. All requests require API Key authentication passed as an `x-api-key` header.  Create an account to [get your free API Key](https://app.mailslurp.com/sign-up/).  ## Resources - 🔑 [Get API Key](https://app.mailslurp.com/sign-up/)                    - 🎓 [Developer Portal](https://www.mailslurp.com/docs/)           - 📦 [Library SDKs](https://www.mailslurp.com/docs/) - ✍️ [Code Examples](https://www.mailslurp.com/examples) - ⚠️ [Report an issue](https://drift.me/mailslurp)  ## Explore  
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp
// ContactDto struct for ContactDto
type ContactDto struct {
	Company string `json:"company,omitempty"`
	EmailAddresses []string `json:"emailAddresses"`
	FirstName string `json:"firstName,omitempty"`
	Id string `json:"id"`
	LastName string `json:"lastName,omitempty"`
	MetaData JsonNode `json:"metaData,omitempty"`
	OptOut bool `json:"optOut,omitempty"`
	Tags []string `json:"tags"`
}
