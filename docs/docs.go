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
        "/passenger/attributes": {
            "get": {
                "description": "Get passenger information filtered by specific attributes.",
                "produces": [
                    "application/json"
                ],
                "operationId": "getPassengerInfoByAttributes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Passenger ID",
                        "name": "passengerId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Comma-separated list of attribute names",
                        "name": "attributes",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Passenger information",
                        "schema": {
                            "$ref": "#/definitions/model.PassengerInfoDTO"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/passengers/v1/info": {
            "get": {
                "description": "Get a list of all passenger information.",
                "produces": [
                    "application/json"
                ],
                "operationId": "getAllPassengerInfo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.PassengerInfo"
                            }
                        }
                    }
                }
            }
        },
        "/passengers/v1/info/fares/histogram": {
            "get": {
                "description": "Get a histogram of fare prices with percentiles.",
                "produces": [
                    "text/html"
                ],
                "operationId": "getFaresHistogram",
                "responses": {
                    "200": {
                        "description": "html",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/passengers/v1/info/{passengerId}": {
            "get": {
                "description": "Get Passenger Info by passenger ID",
                "produces": [
                    "application/json"
                ],
                "operationId": "GetPassengerInfo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Passenger ID",
                        "name": "passengerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Passenger Info",
                        "schema": {
                            "$ref": "#/definitions/model.PassengerInfo"
                        }
                    },
                    "404": {
                        "description": "Passenger not found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "model.PassengerInfo": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "number"
                },
                "cabin": {
                    "type": "string"
                },
                "embarked": {
                    "type": "string"
                },
                "fare": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "pClass": {
                    "type": "integer"
                },
                "parch": {
                    "type": "integer"
                },
                "passenger_id": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "sib_sb": {
                    "type": "integer"
                },
                "survived": {
                    "type": "integer"
                },
                "ticket": {
                    "type": "string"
                }
            }
        },
        "model.PassengerInfoDTO": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "number"
                },
                "cabin": {
                    "type": "string"
                },
                "embarked": {
                    "type": "string"
                },
                "fare": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "pClass": {
                    "type": "integer"
                },
                "parch": {
                    "type": "integer"
                },
                "passengerId": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "sibSb": {
                    "type": "integer"
                },
                "survived": {
                    "type": "integer"
                },
                "ticket": {
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
