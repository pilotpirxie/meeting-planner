package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]any{"error": message})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return http.ErrMissingFile
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func ToJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func ToJSONPretty(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(b)
}

var validate = validator.New()

type RequestOptions struct {
	Body    any
	Params  any
	Query   any
	Headers any
}

func ParseRequest(r *http.Request, opts RequestOptions) error {
	if opts.Body != nil {
		if r.Body == nil {
			return fmt.Errorf("Missing request body")
		}
		if err := json.NewDecoder(r.Body).Decode(opts.Body); err != nil {
			return fmt.Errorf("Invalid JSON body: %w", err)
		}
		defer r.Body.Close()
		if err := validate.Struct(opts.Body); err != nil {
			return fmt.Errorf("Body validation failed: %w", err)
		}
	}

	if opts.Params != nil {
		if err := parsePathParams(r, opts.Params); err != nil {
			return fmt.Errorf("Invalid path params: %w", err)
		}
		if err := validate.Struct(opts.Params); err != nil {
			return fmt.Errorf("Path params validation failed: %w", err)
		}
	}

	if opts.Query != nil {
		if err := parseQueryParams(r, opts.Query); err != nil {
			return fmt.Errorf("Invalid query params: %w", err)
		}
		if err := validate.Struct(opts.Query); err != nil {
			return fmt.Errorf("Query params validation failed: %w", err)
		}
	}

	if opts.Headers != nil {
		if err := parseHeaders(r, opts.Headers); err != nil {
			return fmt.Errorf("Invalid headers: %w", err)
		}
		if err := validate.Struct(opts.Headers); err != nil {
			return fmt.Errorf("Headers validation failed: %w", err)
		}
	}

	return nil
}

func parsePathParams(r *http.Request, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("ParamsScheme must be a pointer to struct")
	}

	rv = rv.Elem()
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)

		tagName := fieldType.Tag.Get("param")
		if tagName == "" {
			tagName = strings.ToLower(fieldType.Name)
		}

		paramValue := r.PathValue(tagName)
		if paramValue == "" {
			continue
		}

		if err := setFieldValue(field, paramValue); err != nil {
			return fmt.Errorf("field %s: %w", fieldType.Name, err)
		}
	}

	return nil
}

func parseQueryParams(r *http.Request, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("QueryScheme must be a pointer to struct")
	}

	rv = rv.Elem()
	rt := rv.Type()
	query := r.URL.Query()

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)

		tagName := fieldType.Tag.Get("query")
		if tagName == "" {
			tagName = strings.ToLower(fieldType.Name)
		}

		queryValue := query.Get(tagName)
		if queryValue == "" {
			continue
		}

		if err := setFieldValue(field, queryValue); err != nil {
			return fmt.Errorf("Field %s: %w", fieldType.Name, err)
		}
	}

	return nil
}

func parseHeaders(r *http.Request, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("HeadersScheme must be a pointer to struct")
	}

	rv = rv.Elem()
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)

		tagName := fieldType.Tag.Get("header")
		if tagName == "-" {
			continue
		}
		if tagName == "" {
			tagName = http.CanonicalHeaderKey(fieldType.Name)
		}

		headerValue := r.Header.Get(tagName)
		if headerValue == "" {
			continue
		}

		if err := setFieldValue(field, headerValue); err != nil {
			return fmt.Errorf("Field %s: %w", fieldType.Name, err)
		}
	}

	return nil
}

func setFieldValue(field reflect.Value, value string) error {
	if !field.CanSet() {
		return fmt.Errorf("Field cannot be set")
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("Cannot parse as int: %w", err)
		}
		field.SetInt(intVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("Cannot parse as uint: %w", err)
		}
		field.SetUint(uintVal)
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("Cannot parse as float: %w", err)
		}
		field.SetFloat(floatVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("Cannot parse as bool: %w", err)
		}
		field.SetBool(boolVal)
	case reflect.Ptr:
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		return setFieldValue(field.Elem(), value)
	default:
		return fmt.Errorf("Unsupported field type: %s", field.Kind())
	}

	return nil
}
