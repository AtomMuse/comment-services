// Code generated by swaggo/swag. DO NOT EDIT.

package doc

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
        "/api-comments/comments": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Create a new comment",
                "operationId": "CreateComment",
                "parameters": [
                    {
                        "description": "Comment data to create",
                        "name": "requestExhibition",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestCreateComment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api-comments/comments/exhibitions/{id}": {
            "get": {
                "description": "GetCommentsByExhibitionID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "GetCommentsByExhibitionID",
                "operationId": "GetCommentsByExhibitionID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exhibition ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api-comments/comments/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a new comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Update a new comment",
                "operationId": "UpdateComment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Comment data to update",
                        "name": "requestUpdateExhibition",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestUpdateComment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "DeleteCommentByID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "DeleteCommentByID",
                "operationId": "DeleteCommentByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/comments/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "GetCommentByID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "GetCommentByID",
                "operationId": "GetCommentByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.RequestCreateComment": {
            "type": "object",
            "required": [
                "commentMessage",
                "exhibitionId"
            ],
            "properties": {
                "commentMessage": {
                    "type": "string"
                },
                "createDateAt": {
                    "type": "integer"
                },
                "exhibitionId": {
                    "type": "string"
                },
                "updateDateAt": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                }
            }
        },
        "model.RequestUpdateComment": {
            "type": "object",
            "required": [
                "commentMessage"
            ],
            "properties": {
                "commentMessage": {
                    "type": "string"
                },
                "updateDateAt": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "profile": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Comment Service API",
	Description:      "Comment Service สำหรับขอจัดการเกี่ยวกับ Workshop Manager ทั้งการสร้าง แก้ไข ลบ Workshop Manager",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
