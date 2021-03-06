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

func easyjson3c9ce8c3DecodeGithubComBsmOpenrtbNativeRequest(in *jlexer.Lexer, out *Video) {
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
		case "mimes":
			if in.IsNull() {
				in.Skip()
				out.Mimes = nil
			} else {
				in.Delim('[')
				if out.Mimes == nil {
					if !in.IsDelim(']') {
						out.Mimes = make([]string, 0, 4)
					} else {
						out.Mimes = []string{}
					}
				} else {
					out.Mimes = (out.Mimes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Mimes = append(out.Mimes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "minduration":
			out.MinDuration = int(in.Int())
		case "maxduration":
			out.MaxDuration = int(in.Int())
		case "protocols":
			if in.IsNull() {
				in.Skip()
				out.Protocols = nil
			} else {
				in.Delim('[')
				if out.Protocols == nil {
					if !in.IsDelim(']') {
						out.Protocols = make([]int, 0, 8)
					} else {
						out.Protocols = []int{}
					}
				} else {
					out.Protocols = (out.Protocols)[:0]
				}
				for !in.IsDelim(']') {
					var v2 int
					v2 = int(in.Int())
					out.Protocols = append(out.Protocols, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson3c9ce8c3EncodeGithubComBsmOpenrtbNativeRequest(out *jwriter.Writer, in Video) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Mimes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"mimes\":")
		if in.Mimes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Mimes {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if in.MinDuration != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"minduration\":")
		out.Int(int(in.MinDuration))
	}
	if in.MaxDuration != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"maxduration\":")
		out.Int(int(in.MaxDuration))
	}
	if len(in.Protocols) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"protocols\":")
		if in.Protocols == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Protocols {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
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
func (v Video) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9ce8c3EncodeGithubComBsmOpenrtbNativeRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Video) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9ce8c3EncodeGithubComBsmOpenrtbNativeRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Video) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9ce8c3DecodeGithubComBsmOpenrtbNativeRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Video) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9ce8c3DecodeGithubComBsmOpenrtbNativeRequest(l, v)
}
