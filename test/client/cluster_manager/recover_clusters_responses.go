// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// RecoverClustersReader is a Reader for the RecoverClusters structure.
type RecoverClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecoverClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRecoverClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRecoverClustersOK creates a RecoverClustersOK with default headers values
func NewRecoverClustersOK() *RecoverClustersOK {
	return &RecoverClustersOK{}
}

/*RecoverClustersOK handles this case with default header values.

RecoverClustersOK recover clusters o k
*/
type RecoverClustersOK struct {
	Payload *models.OpenpitrixRecoverClustersResponse
}

func (o *RecoverClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/recover][%d] recoverClustersOK  %+v", 200, o.Payload)
}

func (o *RecoverClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixRecoverClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
