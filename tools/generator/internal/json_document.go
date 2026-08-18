package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// jsonObject keeps the key order it parsed, config.json is committed.
type jsonObject struct {
	keys   []string
	values map[string]any
}

func newJSONObject() *jsonObject {
	return &jsonObject{values: make(map[string]any)}
}

func (o *jsonObject) Get(key string) (any, bool) {
	value, ok := o.values[key]
	return value, ok
}

func (o *jsonObject) Set(key string, value any) {
	if _, ok := o.values[key]; !ok {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

func (o *jsonObject) Keys() []string {
	return o.keys
}

func (o *jsonObject) Values() []any {
	var values = make([]any, 0, len(o.keys))
	for _, key := range o.keys {
		values = append(values, o.values[key])
	}
	return values
}

// Numbers keep their source spelling.
func decodeJSONDocument(data []byte) (any, error) {
	var dec = json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	value, err := decodeJSONValue(dec)
	if err != nil {
		return nil, err
	}
	if _, err = dec.Token(); !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("unexpected data after the top level value")
	}
	return value, nil
}

func decodeJSONValue(dec *json.Decoder) (any, error) {
	var token, err = dec.Token()
	if err != nil {
		return nil, err
	}

	delim, ok := token.(json.Delim)
	if !ok {
		return token, nil
	}

	switch delim {
	case '{':
		var obj = newJSONObject()
		for dec.More() {
			key, err := dec.Token()
			if err != nil {
				return nil, err
			}
			name, ok := key.(string)
			if !ok {
				return nil, fmt.Errorf("object key %v is not a string", key)
			}
			value, err := decodeJSONValue(dec)
			if err != nil {
				return nil, err
			}
			obj.Set(name, value)
		}
		_, err = dec.Token()
		return obj, err
	case '[':
		var items = []any{}
		for dec.More() {
			item, err := decodeJSONValue(dec)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		_, err = dec.Token()
		return items, err
	}

	return nil, fmt.Errorf("unexpected %v", delim)
}

// The format config.json is committed in: two space indents, no trailing
// newline, no escaping of & and <.
func encodeJSONDocument(value any) ([]byte, error) {
	var buf bytes.Buffer
	if err := writeJSONValue(&buf, value, ""); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

const jsonIndent = "  "

func writeJSONValue(buf *bytes.Buffer, value any, indent string) error {
	switch v := value.(type) {
	case *jsonObject:
		return writeJSONObject(buf, v, indent)
	case []any:
		return writeJSONArray(buf, v, indent)
	case string:
		return writeJSONString(buf, v)
	case json.Number:
		buf.WriteString(v.String())
	case bool:
		_, _ = fmt.Fprintf(buf, "%t", v)
	case nil:
		buf.WriteString("null")
	default:
		return fmt.Errorf("cannot encode %T", value)
	}
	return nil
}

func writeJSONObject(buf *bytes.Buffer, obj *jsonObject, indent string) error {
	if len(obj.keys) == 0 {
		buf.WriteString("{}")
		return nil
	}

	var inner = indent + jsonIndent
	buf.WriteString("{\n")
	for idx, key := range obj.keys {
		if idx > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteString(inner)
		if err := writeJSONString(buf, key); err != nil {
			return err
		}
		buf.WriteString(": ")
		if err := writeJSONValue(buf, obj.values[key], inner); err != nil {
			return err
		}
	}
	buf.WriteString("\n" + indent + "}")
	return nil
}

func writeJSONArray(buf *bytes.Buffer, items []any, indent string) error {
	if len(items) == 0 {
		buf.WriteString("[]")
		return nil
	}

	var inner = indent + jsonIndent
	buf.WriteString("[\n")
	for idx, item := range items {
		if idx > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteString(inner)
		if err := writeJSONValue(buf, item, inner); err != nil {
			return err
		}
	}
	buf.WriteString("\n" + indent + "]")
	return nil
}

func writeJSONString(buf *bytes.Buffer, value string) error {
	var quoted bytes.Buffer
	var enc = json.NewEncoder(&quoted)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(value); err != nil {
		return err
	}
	buf.Write(bytes.TrimRight(quoted.Bytes(), "\n"))
	return nil
}
