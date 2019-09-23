/*
 * MailSlurp API
 *
 * For documentation see [developer guide](https://www.mailslurp.com/developers). [Create an account](https://app.mailslurp.com) in the MailSlurp Dashboard to [view your API Key](https://app). For all bugs, feature requests, or help please [see support](https://www.mailslurp.com/support/).
 *
 * API version: 0.0.1-alpha
 * Contact: contact@mailslurp.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mailslurp

// Options for creating a domain to use with MailSlurp. You must have ownership access to this domain in order to verify it.
type CreateDomainOptions struct {
	// The top level domain you wish to use with MailSlurp
	Domain string `json:"domain,omitempty"`
}
