// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package request

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

func easyjsonE7952480DecodeGithubComBsmOpenrtbNativeRequest(in *jlexer.Lexer, out *Title) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "len":
			out.Length = int(in.Int())
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
			}
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
func easyjsonE7952480EncodeGithubComBsmOpenrtbNativeRequest(out *jwriter.Writer, in Title) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"len\":")
	out.Int(int(in.Length))
	if len(in.Ext) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ext\":")
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Title) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE7952480EncodeGithubComBsmOpenrtbNativeRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Title) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE7952480EncodeGithubComBsmOpenrtbNativeRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Title) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE7952480DecodeGithubComBsmOpenrtbNativeRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Title) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE7952480DecodeGithubComBsmOpenrtbNativeRequest(l, v)
}
