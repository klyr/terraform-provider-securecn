// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostCdConnectionsRuleReader is a Reader for the PostCdConnectionsRule structure.
type PostCdConnectionsRuleReader struct {
	Formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCdConnectionsRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCdConnectionsRuleCreated()
		if err := result.readResponse(response, consumer, o.Formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCdConnectionsRuleCreated creates a PostCdConnectionsRuleCreated with default headers values
func NewPostCdConnectionsRuleCreated() *PostCdConnectionsRuleCreated {
	return &PostCdConnectionsRuleCreated{}
}

/*PostCdConnectionsRuleCreated handles this case with default header values.

created.
*/
type PostCdConnectionsRuleCreated struct {
	Payload *CdConnectionRule
}

func (o *PostCdConnectionsRuleCreated) Error() string {
	return fmt.Sprintf("[POST /cd/connectionsRule][%d] postCdConnectionsRuleCreated  %+v", 201, o.Payload)
}

func (o *PostCdConnectionsRuleCreated) GetPayload() *CdConnectionRule {
	return o.Payload
}

func (o *PostCdConnectionsRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CdConnectionRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
