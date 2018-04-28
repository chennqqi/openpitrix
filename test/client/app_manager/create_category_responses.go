// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// CreateCategoryReader is a Reader for the CreateCategory structure.
type CreateCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCategoryOK creates a CreateCategoryOK with default headers values
func NewCreateCategoryOK() *CreateCategoryOK {
	return &CreateCategoryOK{}
}

/*CreateCategoryOK handles this case with default header values.

CreateCategoryOK create category o k
*/
type CreateCategoryOK struct {
	Payload *models.OpenpitrixCreateCategoryResponse
}

func (o *CreateCategoryOK) Error() string {
	return fmt.Sprintf("[POST /v1/categories][%d] createCategoryOK  %+v", 200, o.Payload)
}

func (o *CreateCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixCreateCategoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
