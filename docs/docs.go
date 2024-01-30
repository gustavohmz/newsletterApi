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
        "/api/v1/newsletters": {
            "get": {
                "description": "Retrieves a list of newsletters with optional search and pagination parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "newsletters"
                ],
                "summary": "Get a list of newsletters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the newsletter to search for",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page for pagination",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Allows an admin user to update an existing newsletter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "newsletters"
                ],
                "summary": "Update an existing newsletter",
                "parameters": [
                    {
                        "description": "Update newsletter details",
                        "name": "updateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateNewsletterRequest"
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
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Allows an admin user to create a new newsletter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "newsletters"
                ],
                "summary": "Create a new newsletter",
                "parameters": [
                    {
                        "description": "Newsletter details",
                        "name": "newsletter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Newsletter"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/newsletters/send/{newsletterID}": {
            "post": {
                "description": "Allows an admin user to send a newsletter to a list of subscribers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "newsletters"
                ],
                "summary": "Send newsletter to subscribers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the newsletter to be sent",
                        "name": "newsletterID",
                        "in": "path",
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
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/newsletters/{id}": {
            "delete": {
                "description": "Allows an admin user to delete a newsletter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "newsletters"
                ],
                "summary": "Delete a newsletter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the newsletter to delete",
                        "name": "id",
                        "in": "path",
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
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/subscribe/{email}/{category}": {
            "post": {
                "description": "Allows a user to subscribe to the newsletter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscribers"
                ],
                "summary": "Subscribe to the newsletter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address to subscribe",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Category to subscribe to",
                        "name": "category",
                        "in": "path",
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
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/subscribers": {
            "get": {
                "description": "Retrieves a list of subscribers with optional search and pagination parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscribers"
                ],
                "summary": "Get a list of subscribers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address of the subscriber to search for",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Category of the subscriber to search for",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page for pagination",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/subscribers/{email}/{category}": {
            "get": {
                "description": "Get details of a subscriber by email address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscribers"
                ],
                "summary": "Get subscriber by email and category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address to get details",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Subscriber"
                        }
                    },
                    "404": {
                        "description": "Subscriber not found",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/unsubscribe/{email}/{category}": {
            "delete": {
                "description": "Allows a user to unsubscribe from the newsletter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscribers"
                ],
                "summary": "Unsubscribe from the newsletter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address to unsubscribe",
                        "name": "email",
                        "in": "path",
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
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Attachment": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "domain.Newsletter": {
            "type": "object",
            "properties": {
                "attachments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Attachment"
                    }
                },
                "category": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": ""
                },
                "name": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        },
        "domain.Subscriber": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "subscription_date": {
                    "type": "string"
                }
            }
        },
        "request.Attachment": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "request.UpdateNewsletterRequest": {
            "type": "object",
            "properties": {
                "attachments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.Attachment"
                    }
                },
                "category": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        },
        "v1.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
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
