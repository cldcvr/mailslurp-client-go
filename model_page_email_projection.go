/*
 * MailSlurp API
 *
 * MailSlurp is an API for sending and receiving emails from dynamically allocated email addresses. It's designed for developers and QA teams to test applications, process inbound emails, send templated notifications, attachments, and more.   ## Resources - [Homepage](https://www.mailslurp.com) - Get an [API KEY](https://app.mailslurp.com/sign-up/) - Generated [SDK Clients](https://www.mailslurp.com/docs/) - [Examples](https://github.com/mailslurp/examples) repository 
 *
 * API version: 6.5.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MailSlurpClient
// PageEmailProjection struct for PageEmailProjection
type PageEmailProjection struct {
	Content []EmailProjection `json:"content,omitempty"`
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
