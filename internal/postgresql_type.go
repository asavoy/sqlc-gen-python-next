package python

import (
	"log"

	"github.com/sqlc-dev/plugin-sdk-go/plugin"
	"github.com/sqlc-dev/plugin-sdk-go/sdk"
)

func postgresType(req *plugin.GenerateRequest, col *plugin.Column) pyType {
	return postgresTypeWithConfig(req, col, Config{})
}

func postgresTypeWithConfig(req *plugin.GenerateRequest, col *plugin.Column, conf Config) pyType {
	columnType := sdk.DataType(col.Type)

	switch columnType {
	case "serial", "serial4", "pg_catalog.serial4", "bigserial", "serial8", "pg_catalog.serial8", "smallserial", "serial2", "pg_catalog.serial2", "integer", "int", "int4", "pg_catalog.int4", "bigint", "int8", "pg_catalog.int8", "smallint", "int2", "pg_catalog.int2":
		return pyType{Name: "int", Module: ""}
	case "float", "double precision", "float8", "pg_catalog.float8", "real", "float4", "pg_catalog.float4":
		return pyType{Name: "float", Module: ""}
	case "numeric", "pg_catalog.numeric", "money":
		return pyType{Name: "Decimal", Module: "decimal"}
	case "boolean", "bool", "pg_catalog.bool":
		return pyType{Name: "bool", Module: ""}
	case "json", "jsonb":
		return pyType{Name: "Any", Module: "typing"}
	case "bytea", "blob", "pg_catalog.bytea":
		return pyType{Name: "memoryview", Module: ""}
	case "date":
		return pyType{Name: "date", Module: "datetime"}
	case "pg_catalog.time", "pg_catalog.timetz":
		return pyType{Name: "time", Module: "datetime"}
	case "pg_catalog.timestamptz", "timestamptz":
		// For timezone-aware timestamps, use pydantic.AwareDatetime if using Pydantic models
		if conf.EmitPydanticModels {
			return pyType{Name: "AwareDatetime", Module: "pydantic"}
		}
		return pyType{Name: "datetime", Module: "datetime"}
	case "pg_catalog.timestamp":
		// Plain timestamp without timezone
		return pyType{Name: "datetime", Module: "datetime"}
	case "interval", "pg_catalog.interval":
		return pyType{Name: "timedelta", Module: "datetime"}
	case "text", "pg_catalog.varchar", "pg_catalog.bpchar", "string", "citext":
		return pyType{Name: "str", Module: ""}
	case "uuid":
		return pyType{Name: "UUID", Module: "uuid"}
	case "inet", "cidr", "macaddr", "macaddr8":
		// psycopg2 does have support for ipaddress objects, but it is not enabled by default
		//
		// https://www.psycopg.org/docs/extras.html#adapt-network
		return pyType{Name: "str", Module: ""}
	case "ltree", "lquery", "ltxtquery":
		return pyType{Name: "str", Module: ""}
	default:
		for _, schema := range req.Catalog.Schemas {
			if schema.Name == "pg_catalog" || schema.Name == "information_schema" {
				continue
			}
			for _, enum := range schema.Enums {
				if columnType == enum.Name {
					enumName := modelName(enum.Name, req.Settings)
					if schema.Name != req.Catalog.DefaultSchema {
						enumName = modelName(schema.Name+"_"+enum.Name, req.Settings)
					}
					return pyType{Name: enumName, Module: "models"}
				}
			}
		}
		log.Printf("unknown PostgreSQL type: %s\n", columnType)
		return pyType{Name: "Any", Module: "typing"}
	}
}
