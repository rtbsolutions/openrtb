// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package response

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

func easyjson16eb09bcDecodeGithubComBidbossOpenrtbNativeResponse(in *jlexer.Lexer, out *Link) {
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
		case "url":
			out.URL = string(in.String())
		case "clicktrackers":
			if in.IsNull() {
				in.Skip()
				out.ClickTrackers = nil
			} else {
				in.Delim('[')
				if out.ClickTrackers == nil {
					if !in.IsDelim(']') {
						out.ClickTrackers = make([]string, 0, 4)
					} else {
						out.ClickTrackers = []string{}
					}
				} else {
					out.ClickTrackers = (out.ClickTrackers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.ClickTrackers = append(out.ClickTrackers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "fallback":
			out.FallbackURL = string(in.String())
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
func easyjson16eb09bcEncodeGithubComBidbossOpenrtbNativeResponse(out *jwriter.Writer, in Link) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	if len(in.ClickTrackers) != 0 {
		const prefix string = ",\"clicktrackers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.ClickTrackers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.FallbackURL != "" {
		const prefix string = ",\"fallback\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FallbackURL))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Link) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16eb09bcEncodeGithubComBidbossOpenrtbNativeResponse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Link) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16eb09bcEncodeGithubComBidbossOpenrtbNativeResponse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Link) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16eb09bcDecodeGithubComBidbossOpenrtbNativeResponse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Link) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16eb09bcDecodeGithubComBidbossOpenrtbNativeResponse(l, v)
}
