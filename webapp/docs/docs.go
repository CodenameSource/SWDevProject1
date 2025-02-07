// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/addItem": {
            "post": {
                "description": "Add a new item to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new item",
                "parameters": [
                    {
                        "description": "New item details",
                        "name": "newItem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backend.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/getItems": {
            "get": {
                "description": "Get a list of items",
                "produces": [
                    "application/json"
                ],
                "summary": "Get items",
                "responses": {
                    "200": {
                        "description": "List of items",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/backend.Item"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/removeItem": {
            "delete": {
                "description": "Remove an item from the system",
                "produces": [
                    "application/json"
                ],
                "summary": "Remove an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "URL of the item",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/updatePrice": {
            "get": {
                "description": "Update the price of an item",
                "produces": [
                    "application/json"
                ],
                "summary": "Update item price",
                "parameters": [
                    {
                        "type": "string",
                        "description": "URL of the item",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/updatePrices": {
            "get": {
                "description": "Update the prices of all items",
                "produces": [
                    "application/json"
                ],
                "summary": "Update prices of all items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "backend.Item": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
