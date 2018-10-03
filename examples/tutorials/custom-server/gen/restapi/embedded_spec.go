// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Greeting Server",
    "version": "1.0.0"
  },
  "paths": {
    "/hello": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getGreeting",
        "parameters": [
          {
            "type": "string",
            "description": "defaults to World if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a greeting",
            "schema": {
              "description": "contains the actual greeting as plain text",
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
}
