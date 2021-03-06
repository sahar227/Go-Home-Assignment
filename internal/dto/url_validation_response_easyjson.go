// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD69cde9dDecodeValidatorInternalDto(in *jlexer.Lexer, out *URLValidationResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "location":
			out.Location = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD69cde9dEncodeValidatorInternalDto(out *jwriter.Writer, in URLValidationResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"location\":"
		out.RawString(prefix[1:])
		out.String(string(in.Location))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URLValidationResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD69cde9dEncodeValidatorInternalDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URLValidationResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD69cde9dEncodeValidatorInternalDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URLValidationResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD69cde9dDecodeValidatorInternalDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URLValidationResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD69cde9dDecodeValidatorInternalDto(l, v)
}
