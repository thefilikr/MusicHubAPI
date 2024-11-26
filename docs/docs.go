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
        "/song/create": {
            "post": {
                "description": "Create a new song with all required fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Create a new song",
                "parameters": [
                    {
                        "description": "Song object",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Song"
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
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create song",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song/delete": {
            "delete": {
                "description": "Deletes information about the song by name and group name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Delete information about a song.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the group",
                        "name": "group",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The name of the song",
                        "name": "song",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song successfully deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song/edit": {
            "put": {
                "description": "Changes the information about the song by name and group name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Change the information about the song.",
                "parameters": [
                    {
                        "description": "Song object",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song successfully changed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song/info": {
            "get": {
                "description": "Returns information about the song by name and group name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Get information about the song.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the group",
                        "name": "group",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The name of the song",
                        "name": "song",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of verses to return",
                        "name": "countVerses",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "numPage",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song object",
                        "schema": {
                            "$ref": "#/definitions/handler.Song"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Returns songs by filters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "parameters": [
                    {
                        "description": "Song object",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Array of Song objects",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Song": {
            "description": "Details about a song",
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "The Beatles"
                },
                "link": {
                    "type": "string",
                    "example": "https://example.com/song"
                },
                "release_date": {
                    "type": "string",
                    "example": "1968-08-26"
                },
                "song": {
                    "type": "string",
                    "example": "Hey Jude"
                },
                "text": {
                    "type": "array",
                    "items": {
                        "type": "string"
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
