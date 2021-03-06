package api

func init() {
	Swagger.Add("legacy_legacy", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/legacy/legacy.proto",
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
    "/events/data-collector": {
      "get": {
        "summary": "This is used by chef-server, it requests a GET /data-collector/v0 to check\nAutomate's status.\nWe proxy /data-collector/v0 to /api/v0/events/data-collector, so this is\nwhere we need to respond.\nSince this is for legacy-support only, we don't bother much about having\ngoogle.protobuf.Empty as argument.",
        "operationId": "Status",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/legacyStatusResponse"
            }
          }
        },
        "tags": [
          "LegacyDataCollector"
        ]
      }
    }
  },
  "definitions": {
    "legacyStatusResponse": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    }
  }
}
`)
}
