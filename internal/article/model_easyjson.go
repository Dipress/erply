// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package article

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

func easyjsonC80ae7adDecodeGithubComRomanyxErplyInternalArticle(in *jlexer.Lexer, out *Model) {
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
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonC80ae7adEncodeGithubComRomanyxErplyInternalArticle(out *jwriter.Writer, in Model) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Model) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComRomanyxErplyInternalArticle(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Model) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComRomanyxErplyInternalArticle(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Model) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComRomanyxErplyInternalArticle(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Model) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComRomanyxErplyInternalArticle(l, v)
}
