// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-12 22:49:58.824653 +0800 CST m=+0.053212055

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "商品中心",
        "title": "mms",
        "contact": {
            "name": "richard sun",
            "email": "cugriver@163.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "0.0.1"
    },
    "host": "localhost:13147",
    "basePath": "/",
    "paths": {
        "/v1/categories": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "创建分类",
                "parameters": [
                    {
                        "description": "CreateCategory",
                        "name": "{}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.RequestCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.HttpOk"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.HttpError"
                        }
                    }
                }
            }
        },
        "/v1/merchandises/version": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "显示当前服务的版本和代码版本号",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.ResponseVersion"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/route.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "route.HttpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "route.HttpOk": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "route.RequestCategory": {
            "type": "object",
            "properties": {
                "categoryName": {
                    "type": "string"
                },
                "categoryStatus": {
                    "type": "integer"
                },
                "childCategories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "parentCategories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "route.ResponseVersion": {
            "type": "object",
            "properties": {
                "info": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{ Schemes: []string{"http","https"}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}