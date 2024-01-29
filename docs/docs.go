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
                    },
                    {
                        "description": "Request body containing recipients",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SendNewsletterRequest"
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
            }
        },
        "/api/v1/subscribe/{email}": {
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
        "/api/v1/subscribers/{email}": {
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
                "summary": "Get subscriber by email",
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
        "/api/v1/unsubscribe/{email}": {
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
                }
            }
        },
        "domain.Category": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "SpecialOffers",
                "Memberships",
                "MonthlyPromotions",
                "NewProducts"
            ]
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
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Category"
                    }
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
                }
            }
        },
        "domain.Subscriber": {
            "type": "object",
            "properties": {
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
        "v1.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "v1.SendNewsletterRequest": {
            "type": "object",
            "properties": {
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "email": {
                                "type": "string"
                            }
                        }
                    }
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
