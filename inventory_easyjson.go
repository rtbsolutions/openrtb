// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package openrtb

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

func easyjson6f8bf452DecodeGithubComBsmOpenrtb(in *jlexer.Lexer, out *Site) {
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
		case "page":
			out.Page = string(in.String())
		case "ref":
			out.Ref = string(in.String())
		case "search":
			out.Search = string(in.String())
		case "mobile":
			out.Mobile = int(in.Int())
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "domain":
			out.Domain = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Cat = nil
			} else {
				in.Delim('[')
				if out.Cat == nil {
					if !in.IsDelim(']') {
						out.Cat = make([]string, 0, 4)
					} else {
						out.Cat = []string{}
					}
				} else {
					out.Cat = (out.Cat)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Cat = append(out.Cat, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sectioncat":
			if in.IsNull() {
				in.Skip()
				out.SectionCat = nil
			} else {
				in.Delim('[')
				if out.SectionCat == nil {
					if !in.IsDelim(']') {
						out.SectionCat = make([]string, 0, 4)
					} else {
						out.SectionCat = []string{}
					}
				} else {
					out.SectionCat = (out.SectionCat)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.SectionCat = append(out.SectionCat, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pagecat":
			if in.IsNull() {
				in.Skip()
				out.PageCat = nil
			} else {
				in.Delim('[')
				if out.PageCat == nil {
					if !in.IsDelim(']') {
						out.PageCat = make([]string, 0, 4)
					} else {
						out.PageCat = []string{}
					}
				} else {
					out.PageCat = (out.PageCat)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.PageCat = append(out.PageCat, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "privacypolicy":
			if in.IsNull() {
				in.Skip()
				out.PrivacyPolicy = nil
			} else {
				if out.PrivacyPolicy == nil {
					out.PrivacyPolicy = new(int)
				}
				*out.PrivacyPolicy = int(in.Int())
			}
		case "publisher":
			if in.IsNull() {
				in.Skip()
				out.Publisher = nil
			} else {
				if out.Publisher == nil {
					out.Publisher = new(Publisher)
				}
				easyjson6f8bf452DecodeGithubComBsmOpenrtb1(in, &*out.Publisher)
			}
		case "content":
			if in.IsNull() {
				in.Skip()
				out.Content = nil
			} else {
				if out.Content == nil {
					out.Content = new(Content)
				}
				(*out.Content).UnmarshalEasyJSON(in)
			}
		case "keywords":
			out.Keywords = string(in.String())
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
func easyjson6f8bf452EncodeGithubComBsmOpenrtb(out *jwriter.Writer, in Site) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Page != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"page\":")
		out.String(string(in.Page))
	}
	if in.Ref != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ref\":")
		out.String(string(in.Ref))
	}
	if in.Search != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"search\":")
		out.String(string(in.Search))
	}
	if in.Mobile != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"mobile\":")
		out.Int(int(in.Mobile))
	}
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Domain != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"domain\":")
		out.String(string(in.Domain))
	}
	if len(in.Cat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cat\":")
		if in.Cat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Cat {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	if len(in.SectionCat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"sectioncat\":")
		if in.SectionCat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.SectionCat {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.PageCat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pagecat\":")
		if in.PageCat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.PageCat {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.PrivacyPolicy != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"privacypolicy\":")
		if in.PrivacyPolicy == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.PrivacyPolicy))
		}
	}
	if in.Publisher != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"publisher\":")
		if in.Publisher == nil {
			out.RawString("null")
		} else {
			easyjson6f8bf452EncodeGithubComBsmOpenrtb1(out, *in.Publisher)
		}
	}
	if in.Content != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"content\":")
		if in.Content == nil {
			out.RawString("null")
		} else {
			(*in.Content).MarshalEasyJSON(out)
		}
	}
	if in.Keywords != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"keywords\":")
		out.String(string(in.Keywords))
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
func (v Site) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6f8bf452EncodeGithubComBsmOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Site) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6f8bf452EncodeGithubComBsmOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Site) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6f8bf452DecodeGithubComBsmOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Site) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6f8bf452DecodeGithubComBsmOpenrtb(l, v)
}
func easyjson6f8bf452DecodeGithubComBsmOpenrtb1(in *jlexer.Lexer, out *Publisher) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Cat = nil
			} else {
				in.Delim('[')
				if out.Cat == nil {
					if !in.IsDelim(']') {
						out.Cat = make([]string, 0, 4)
					} else {
						out.Cat = []string{}
					}
				} else {
					out.Cat = (out.Cat)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Cat = append(out.Cat, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "domain":
			out.Domain = string(in.String())
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
func easyjson6f8bf452EncodeGithubComBsmOpenrtb1(out *jwriter.Writer, in Publisher) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if len(in.Cat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cat\":")
		if in.Cat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Cat {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if in.Domain != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"domain\":")
		out.String(string(in.Domain))
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
func easyjson6f8bf452DecodeGithubComBsmOpenrtb2(in *jlexer.Lexer, out *Inventory) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "domain":
			out.Domain = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Cat = nil
			} else {
				in.Delim('[')
				if out.Cat == nil {
					if !in.IsDelim(']') {
						out.Cat = make([]string, 0, 4)
					} else {
						out.Cat = []string{}
					}
				} else {
					out.Cat = (out.Cat)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Cat = append(out.Cat, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sectioncat":
			if in.IsNull() {
				in.Skip()
				out.SectionCat = nil
			} else {
				in.Delim('[')
				if out.SectionCat == nil {
					if !in.IsDelim(']') {
						out.SectionCat = make([]string, 0, 4)
					} else {
						out.SectionCat = []string{}
					}
				} else {
					out.SectionCat = (out.SectionCat)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.SectionCat = append(out.SectionCat, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pagecat":
			if in.IsNull() {
				in.Skip()
				out.PageCat = nil
			} else {
				in.Delim('[')
				if out.PageCat == nil {
					if !in.IsDelim(']') {
						out.PageCat = make([]string, 0, 4)
					} else {
						out.PageCat = []string{}
					}
				} else {
					out.PageCat = (out.PageCat)[:0]
				}
				for !in.IsDelim(']') {
					var v15 string
					v15 = string(in.String())
					out.PageCat = append(out.PageCat, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "privacypolicy":
			if in.IsNull() {
				in.Skip()
				out.PrivacyPolicy = nil
			} else {
				if out.PrivacyPolicy == nil {
					out.PrivacyPolicy = new(int)
				}
				*out.PrivacyPolicy = int(in.Int())
			}
		case "publisher":
			if in.IsNull() {
				in.Skip()
				out.Publisher = nil
			} else {
				if out.Publisher == nil {
					out.Publisher = new(Publisher)
				}
				easyjson6f8bf452DecodeGithubComBsmOpenrtb1(in, &*out.Publisher)
			}
		case "content":
			if in.IsNull() {
				in.Skip()
				out.Content = nil
			} else {
				if out.Content == nil {
					out.Content = new(Content)
				}
				(*out.Content).UnmarshalEasyJSON(in)
			}
		case "keywords":
			out.Keywords = string(in.String())
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
func easyjson6f8bf452EncodeGithubComBsmOpenrtb2(out *jwriter.Writer, in Inventory) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Domain != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"domain\":")
		out.String(string(in.Domain))
	}
	if len(in.Cat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cat\":")
		if in.Cat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.Cat {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	if len(in.SectionCat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"sectioncat\":")
		if in.SectionCat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.SectionCat {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	if len(in.PageCat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pagecat\":")
		if in.PageCat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.PageCat {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	if in.PrivacyPolicy != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"privacypolicy\":")
		if in.PrivacyPolicy == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.PrivacyPolicy))
		}
	}
	if in.Publisher != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"publisher\":")
		if in.Publisher == nil {
			out.RawString("null")
		} else {
			easyjson6f8bf452EncodeGithubComBsmOpenrtb1(out, *in.Publisher)
		}
	}
	if in.Content != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"content\":")
		if in.Content == nil {
			out.RawString("null")
		} else {
			(*in.Content).MarshalEasyJSON(out)
		}
	}
	if in.Keywords != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"keywords\":")
		out.String(string(in.Keywords))
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
func (v Inventory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6f8bf452EncodeGithubComBsmOpenrtb2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Inventory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6f8bf452EncodeGithubComBsmOpenrtb2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Inventory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6f8bf452DecodeGithubComBsmOpenrtb2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Inventory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6f8bf452DecodeGithubComBsmOpenrtb2(l, v)
}
func easyjson6f8bf452DecodeGithubComBsmOpenrtb3(in *jlexer.Lexer, out *App) {
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
		case "bundle":
			out.Bundle = string(in.String())
		case "storeurl":
			out.StoreURL = string(in.String())
		case "ver":
			out.Ver = string(in.String())
		case "paid":
			out.Paid = int(in.Int())
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "domain":
			out.Domain = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Cat = nil
			} else {
				in.Delim('[')
				if out.Cat == nil {
					if !in.IsDelim(']') {
						out.Cat = make([]string, 0, 4)
					} else {
						out.Cat = []string{}
					}
				} else {
					out.Cat = (out.Cat)[:0]
				}
				for !in.IsDelim(']') {
					var v22 string
					v22 = string(in.String())
					out.Cat = append(out.Cat, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sectioncat":
			if in.IsNull() {
				in.Skip()
				out.SectionCat = nil
			} else {
				in.Delim('[')
				if out.SectionCat == nil {
					if !in.IsDelim(']') {
						out.SectionCat = make([]string, 0, 4)
					} else {
						out.SectionCat = []string{}
					}
				} else {
					out.SectionCat = (out.SectionCat)[:0]
				}
				for !in.IsDelim(']') {
					var v23 string
					v23 = string(in.String())
					out.SectionCat = append(out.SectionCat, v23)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pagecat":
			if in.IsNull() {
				in.Skip()
				out.PageCat = nil
			} else {
				in.Delim('[')
				if out.PageCat == nil {
					if !in.IsDelim(']') {
						out.PageCat = make([]string, 0, 4)
					} else {
						out.PageCat = []string{}
					}
				} else {
					out.PageCat = (out.PageCat)[:0]
				}
				for !in.IsDelim(']') {
					var v24 string
					v24 = string(in.String())
					out.PageCat = append(out.PageCat, v24)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "privacypolicy":
			if in.IsNull() {
				in.Skip()
				out.PrivacyPolicy = nil
			} else {
				if out.PrivacyPolicy == nil {
					out.PrivacyPolicy = new(int)
				}
				*out.PrivacyPolicy = int(in.Int())
			}
		case "publisher":
			if in.IsNull() {
				in.Skip()
				out.Publisher = nil
			} else {
				if out.Publisher == nil {
					out.Publisher = new(Publisher)
				}
				easyjson6f8bf452DecodeGithubComBsmOpenrtb1(in, &*out.Publisher)
			}
		case "content":
			if in.IsNull() {
				in.Skip()
				out.Content = nil
			} else {
				if out.Content == nil {
					out.Content = new(Content)
				}
				(*out.Content).UnmarshalEasyJSON(in)
			}
		case "keywords":
			out.Keywords = string(in.String())
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
func easyjson6f8bf452EncodeGithubComBsmOpenrtb3(out *jwriter.Writer, in App) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Bundle != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bundle\":")
		out.String(string(in.Bundle))
	}
	if in.StoreURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"storeurl\":")
		out.String(string(in.StoreURL))
	}
	if in.Ver != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ver\":")
		out.String(string(in.Ver))
	}
	if in.Paid != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"paid\":")
		out.Int(int(in.Paid))
	}
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Domain != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"domain\":")
		out.String(string(in.Domain))
	}
	if len(in.Cat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cat\":")
		if in.Cat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v25, v26 := range in.Cat {
				if v25 > 0 {
					out.RawByte(',')
				}
				out.String(string(v26))
			}
			out.RawByte(']')
		}
	}
	if len(in.SectionCat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"sectioncat\":")
		if in.SectionCat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v27, v28 := range in.SectionCat {
				if v27 > 0 {
					out.RawByte(',')
				}
				out.String(string(v28))
			}
			out.RawByte(']')
		}
	}
	if len(in.PageCat) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"pagecat\":")
		if in.PageCat == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v29, v30 := range in.PageCat {
				if v29 > 0 {
					out.RawByte(',')
				}
				out.String(string(v30))
			}
			out.RawByte(']')
		}
	}
	if in.PrivacyPolicy != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"privacypolicy\":")
		if in.PrivacyPolicy == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.PrivacyPolicy))
		}
	}
	if in.Publisher != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"publisher\":")
		if in.Publisher == nil {
			out.RawString("null")
		} else {
			easyjson6f8bf452EncodeGithubComBsmOpenrtb1(out, *in.Publisher)
		}
	}
	if in.Content != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"content\":")
		if in.Content == nil {
			out.RawString("null")
		} else {
			(*in.Content).MarshalEasyJSON(out)
		}
	}
	if in.Keywords != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"keywords\":")
		out.String(string(in.Keywords))
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
func (v App) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6f8bf452EncodeGithubComBsmOpenrtb3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v App) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6f8bf452EncodeGithubComBsmOpenrtb3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *App) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6f8bf452DecodeGithubComBsmOpenrtb3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *App) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6f8bf452DecodeGithubComBsmOpenrtb3(l, v)
}
