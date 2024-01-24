// Package stockitem_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package stockitem_api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse = []struct {
	Message *string `json:"message,omitempty"`
}

// StockItem defines model for StockItem.
type StockItem struct {
	CreatedAt time.Time          `json:"createdAt"`
	Id        openapi_types.UUID `json:"id"`
	Name      string             `json:"name"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

// BadRequest defines model for BadRequest.
type BadRequest = BadRequestResponse

// Created defines model for Created.
type Created = StockItem

// OK defines model for OK.
type OK = StockItem

// PostStockItemJSONBody defines parameters for PostStockItem.
type PostStockItemJSONBody struct {
	Name string `json:"name"`
}

// PutStockItemJSONBody defines parameters for PutStockItem.
type PutStockItemJSONBody struct {
	Name string `json:"name"`
}

// PostStockItemJSONRequestBody defines body for PostStockItem for application/json ContentType.
type PostStockItemJSONRequestBody PostStockItemJSONBody

// PutStockItemJSONRequestBody defines body for PutStockItem for application/json ContentType.
type PutStockItemJSONRequestBody PutStockItemJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create Stock Item
	// (POST /stock/items)
	PostStockItem(ctx echo.Context) error
	// Update Stock Item
	// (PUT /stock/items/{id})
	PutStockItem(ctx echo.Context, id openapi_types.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) PostStockItem(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostStockItem(ctx)
	return err
}

// PutStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) PutStockItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutStockItem(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/stock/items", wrapper.PostStockItem)
	router.PUT(baseURL+"/stock/items/:id", wrapper.PutStockItem)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xUzW7bPBB8FWG/78hGSn8uuiVFD0YKJEjQU5ADK25spuZPliujhsB3L0jZlhSrTQIU",
	"LdCbuBzNLmeG7KBxxjuLlgPUHRAG72zAvDiX6hofWwycVo2zjDZ/Su/XupGsnS0fgrOpFpoVGpm+/ie8",
	"hxr+Kwfqst8N5UB5vesEMUYBCkND2idGqFPjYt85CvhIKBnVbxvihl3zbcFo5nrvm0UBC8tIVq5vkDZI",
	"n4gcJeopfg8qelTRw6KAy4s/M/DlBaTiDjv17SBy3YFmNHnbk/NIrHuTDYYglxnBW49QQ2DSdplJdxX3",
	"9QGbbMWuIInkNq2HyY6Im17Is3z2e0dGMtSgJOMb1gZBPO0nQKsJtm21moNZaebmFdB69bqWUQDhY6sp",
	"pesWcrvMLkbjj3nvjjRJHNreuzyQ5nXay6oUSZbi7GoBAjZIoXfr9KQ6qdKszqOVXkMN73JJgJe8ysqV",
	"If1fDoa5/gbOBbUYekEmpRywhYIarlzgwaD+qBj43Kntq6I59XVefQEbudZK9sN1YOT3z2iXvIL6tKoE",
	"GG0P6/hU90w5L+0AY2oxF0Yv1Nvq9Gd354ArRzf6fVU9jx89e1HAh5f8MvdS5EvZGiNpO+tW2h87XXZa",
	"xSx3O+P2lxzBX7rdTsz2kqRBRgpQ33agE0lK2D7hdZ/2qb5i5Poz1zDe/ZuBeoHb6cX9m1k6zkKMMf4I",
	"AAD//+shVxLLBwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
