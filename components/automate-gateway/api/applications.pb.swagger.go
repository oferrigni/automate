package api

func init() {
	Swagger.Add("applications", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/applications/applications.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/beta/applications/service-groups": {
      "get": {
        "operationId": "GetServiceGroups",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServiceGroups"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "pagination.page",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "sorting.field",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sorting.order",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ASC",
              "DESC"
            ],
            "default": "ASC"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/service-groups/{service_group_id}": {
      "get": {
        "operationId": "GetServicesBySG",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesBySGRes"
            }
          }
        },
        "parameters": [
          {
            "name": "service_group_id",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.page",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "sorting.field",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sorting.order",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ASC",
              "DESC"
            ],
            "default": "ASC"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/service_groups_health_counts": {
      "get": {
        "operationId": "GetServiceGroupsHealthCounts",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsHealthCounts"
            }
          }
        },
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/services": {
      "get": {
        "operationId": "GetServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesRes"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "pagination.page",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "sorting.field",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sorting.order",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ASC",
              "DESC"
            ],
            "default": "ASC"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/version": {
      "get": {
        "operationId": "GetVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/versionVersionInfo"
            }
          }
        },
        "tags": [
          "ApplicationsService"
        ]
      }
    }
  },
  "definitions": {
    "applicationsHealthCounts": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32"
        },
        "ok": {
          "type": "integer",
          "format": "int32"
        },
        "warning": {
          "type": "integer",
          "format": "int32"
        },
        "critical": {
          "type": "integer",
          "format": "int32"
        },
        "unknown": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "applicationsHealthStatus": {
      "type": "string",
      "enum": [
        "OK",
        "WARNING",
        "CRITICAL",
        "UNKNOWN"
      ],
      "default": "OK",
      "title": "The HealthStatus enum matches the habitat implementation for health-check status:\n=\u003e https://www.habitat.sh/docs/reference/#health-check"
    },
    "applicationsService": {
      "type": "object",
      "properties": {
        "supervisor_id": {
          "type": "string"
        },
        "release": {
          "type": "string"
        },
        "group": {
          "type": "string"
        },
        "health_check": {
          "$ref": "#/definitions/applicationsHealthStatus"
        },
        "status": {
          "$ref": "#/definitions/applicationsServiceStatus"
        },
        "application": {
          "type": "string"
        },
        "environment": {
          "type": "string"
        },
        "fqdn": {
          "type": "string"
        }
      }
    },
    "applicationsServiceGroup": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "release": {
          "type": "string",
          "title": "This field is the full package identifier combined in a single string like:\nExample: core/redis/0.1.0/8743278934278923"
        },
        "status": {
          "$ref": "#/definitions/applicationsHealthStatus"
        },
        "health_percentage": {
          "type": "integer",
          "format": "int32",
          "title": "The health_percentage can be a number between 0-100"
        },
        "services_health_counts": {
          "$ref": "#/definitions/applicationsHealthCounts"
        },
        "id": {
          "type": "string"
        }
      },
      "title": "A service group message is the representation of one single service group that\nis internally generated by aggregating all the services"
    },
    "applicationsServiceGroups": {
      "type": "object",
      "properties": {
        "service_groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/applicationsServiceGroup"
          }
        }
      }
    },
    "applicationsServiceStatus": {
      "type": "string",
      "enum": [
        "RUNNING",
        "INITIALIZING",
        "DEPLOYING",
        "DOWN"
      ],
      "default": "RUNNING",
      "title": "The ServiceStatus enum describes the status of the service\n@afiune have we defined these states somewhere?"
    },
    "applicationsServicesBySGRes": {
      "type": "object",
      "properties": {
        "group": {
          "type": "string"
        },
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/applicationsService"
          }
        }
      }
    },
    "applicationsServicesRes": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/applicationsService"
          }
        }
      }
    },
    "queryPagination": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "querySortOrder": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    },
    "querySorting": {
      "type": "object",
      "properties": {
        "field": {
          "type": "string"
        },
        "order": {
          "$ref": "#/definitions/querySortOrder"
        }
      }
    },
    "versionVersionInfo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "sha": {
          "type": "string"
        },
        "built": {
          "type": "string"
        }
      }
    }
  }
}
`)
}
