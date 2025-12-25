package core

import "fmt"

// ---------------- BaseNode ----------------
type BaseNode struct {
	ID         string
	Class      string
	Style      string
	Attributes map[string]string
	Events     map[string]string // onclick, oninput, etc.
}

func renderBaseAttrs(b BaseNode) string {
	out := ""
	if b.ID != "" {
		out += fmt.Sprintf(` id="%s"`, b.ID)
	}
	if b.Class != "" {
		out += fmt.Sprintf(` class="%s"`, b.Class)
	}
	if b.Style != "" {
		out += fmt.Sprintf(` style="%s"`, b.Style)
	}
	for k, v := range b.Attributes {
		out += fmt.Sprintf(` %s="%s"`, k, v)
	}
	for k, v := range b.Events {
		out += fmt.Sprintf(` %s="%s"`, k, v)
	}
	return out
}

// ---------------- Node Interface ----------------
type Node interface {
	Render() string
}

// ---------------- Containers ----------------
type Div struct {
	BaseNode
	Children []Node
}

func (d Div) Render() string {
	html := "<div" + renderBaseAttrs(d.BaseNode) + ">"
	for _, c := range d.Children {
		html += c.Render()
	}
	html += "</div>"
	return html
}

type Header struct {
	BaseNode
	Children []Node
}

func (h Header) Render() string {
	html := "<header" + renderBaseAttrs(h.BaseNode) + ">"
	for _, c := range h.Children {
		html += c.Render()
	}
	html += "</header>"
	return html
}

type Footer struct {
	BaseNode
	Children []Node
}

func (f Footer) Render() string {
	html := "<footer" + renderBaseAttrs(f.BaseNode) + ">"
	for _, c := range f.Children {
		html += c.Render()
	}
	html += "</footer>"
	return html
}

type Main struct {
	BaseNode
	Children []Node
}

func (m Main) Render() string {
	html := "<main" + renderBaseAttrs(m.BaseNode) + ">"
	for _, c := range m.Children {
		html += c.Render()
	}
	html += "</main>"
	return html
}

type Section struct {
	BaseNode
	Children []Node
}

func (s Section) Render() string {
	html := "<section" + renderBaseAttrs(s.BaseNode) + ">"
	for _, c := range s.Children {
		html += c.Render()
	}
	html += "</section>"
	return html
}

type Article struct {
	BaseNode
	Children []Node
}

func (a Article) Render() string {
	html := "<article" + renderBaseAttrs(a.BaseNode) + ">"
	for _, c := range a.Children {
		html += c.Render()
	}
	html += "</article>"
	return html
}

type Nav struct {
	BaseNode
	Children []Node
}

func (n Nav) Render() string {
	html := "<nav" + renderBaseAttrs(n.BaseNode) + ">"
	for _, c := range n.Children {
		html += c.Render()
	}
	html += "</nav>"
	return html
}

// ---------------- Text ----------------
type H1 struct {
	BaseNode
	Text string
}

func (h H1) Render() string { return "<h1" + renderBaseAttrs(h.BaseNode) + ">" + h.Text + "</h1>" }

type H2 struct {
	BaseNode
	Text string
}

func (h H2) Render() string { return "<h2" + renderBaseAttrs(h.BaseNode) + ">" + h.Text + "</h2>" }

type H3 struct {
	BaseNode
	Text string
}

func (h H3) Render() string { return "<h3" + renderBaseAttrs(h.BaseNode) + ">" + h.Text + "</h3>" }

type H4 struct {
	BaseNode
	Text string
}

func (h H4) Render() string { return "<h4" + renderBaseAttrs(h.BaseNode) + ">" + h.Text + "</h4>" }

type H5 struct {
	BaseNode
	Text string
}

func (h H5) Render() string { return "<h5" + renderBaseAttrs(h.BaseNode) + ">" + h.Text + "</h5>" }

type H6 struct {
	BaseNode
	Text string
}

func (h H6) Render() string { return "<h6" + renderBaseAttrs(h.BaseNode) + ">" + h.Text + "</h6>" }

type P struct {
	BaseNode
	Text string
}

func (p P) Render() string { return "<p" + renderBaseAttrs(p.BaseNode) + ">" + p.Text + "</p>" }

type Blockquote struct {
	BaseNode
	Children []Node
}

func (b Blockquote) Render() string {
	html := "<blockquote" + renderBaseAttrs(b.BaseNode) + ">"
	for _, c := range b.Children {
		html += c.Render()
	}
	html += "</blockquote>"
	return html
}

// ---------------- Links & Buttons ----------------
type A struct {
	BaseNode
	Text, Href string
}

func (a A) Render() string {
	return fmt.Sprintf(`<a href="%s"%s>%s</a>`, a.Href, renderBaseAttrs(a.BaseNode), a.Text)
}

type Button struct {
	BaseNode
	Text    string
	OnClick string
}

func (b Button) Render() string {
	if b.OnClick != "" {
		b.Events = map[string]string{"onclick": b.OnClick}
	}
	return "<button" + renderBaseAttrs(b.BaseNode) + ">" + b.Text + "</button>"
}

// ---------------- Forms ----------------
type Form struct {
	BaseNode
	Children       []Node
	Method, Action string
}

func (f Form) Render() string {
	html := fmt.Sprintf(`<form method="%s" action="%s"%s>`, f.Method, f.Action, renderBaseAttrs(f.BaseNode))
	for _, c := range f.Children {
		html += c.Render()
	}
	html += "</form>"
	return html
}

type Input struct {
	BaseNode
	Type, Name, Value, Placeholder string
	Checked, Disabled              bool
}

func (i Input) Render() string {
	html := fmt.Sprintf(`<input type="%s" name="%s" value="%s" placeholder="%s"%s`, i.Type, i.Name, i.Value, i.Placeholder, renderBaseAttrs(i.BaseNode))
	if i.Checked {
		html += " checked"
	}
	if i.Disabled {
		html += " disabled"
	}
	html += ">"
	return html
}

type Textarea struct {
	BaseNode
	Name, Placeholder string
	Rows, Cols        int
	Text              string
}

func (t Textarea) Render() string {
	html := fmt.Sprintf(`<textarea name="%s" rows="%d" cols="%d"%s>%s</textarea>`, t.Name, t.Rows, t.Cols, renderBaseAttrs(t.BaseNode), t.Text)
	return html
}

type Select struct {
	BaseNode
	Name    string
	Options []Option
}

func (s Select) Render() string {
	html := fmt.Sprintf(`<select name="%s"%s>`, s.Name, renderBaseAttrs(s.BaseNode))
	for _, o := range s.Options {
		html += o.Render()
	}
	html += "</select>"
	return html
}

type Option struct {
	BaseNode
	Text, Value string
	Selected    bool
}

func (o Option) Render() string {
	sel := ""
	if o.Selected {
		sel = " selected"
	}
	return fmt.Sprintf(`<option value="%s"%s%s>%s</option>`, o.Value, sel, renderBaseAttrs(o.BaseNode), o.Text)
}

// ---------------- Media ----------------
type Img struct {
	BaseNode
	Src, Alt string
}

func (i Img) Render() string {
	return fmt.Sprintf(`<img src="%s" alt="%s"%s>`, i.Src, i.Alt, renderBaseAttrs(i.BaseNode))
}

type Audio struct {
	BaseNode
	Controls bool
	Src      string
}

func (a Audio) Render() string {
	html := "<audio"
	if a.Controls {
		html += " controls"
	}
	html += renderBaseAttrs(a.BaseNode)
	if a.Src != "" {
		html += fmt.Sprintf(` src="%s"`, a.Src)
	}
	html += ">audio</audio>"
	return html
}

type Video struct {
	BaseNode
	Controls bool
	Src      string
}

func (v Video) Render() string {
	html := "<video"
	if v.Controls {
		html += " controls"
	}
	html += renderBaseAttrs(v.BaseNode)
	if v.Src != "" {
		html += fmt.Sprintf(` src="%s"`, v.Src)
	}
	html += ">video</video>"
	return html
}

type Canvas struct {
	BaseNode
	Width, Height int
	Text          string
}

func (c Canvas) Render() string {
	return fmt.Sprintf(`<canvas width="%d" height="%d"%s>%s</canvas>`, c.Width, c.Height, renderBaseAttrs(c.BaseNode), c.Text)
}

type Iframe struct {
	BaseNode
	Src           string
	Width, Height int
}

func (i Iframe) Render() string {
	return fmt.Sprintf(`<iframe src="%s" width="%d" height="%d"%s></iframe>`, i.Src, i.Width, i.Height, renderBaseAttrs(i.BaseNode))
}

// ---------------- Scripts & Styles ----------------
type Script struct {
	BaseNode
	Src, Code string
}

func (s Script) Render() string {
	if s.Src != "" {
		return fmt.Sprintf(`<script src="%s"%s></script>`, s.Src, renderBaseAttrs(s.BaseNode))
	}
	return fmt.Sprintf(`<script%s>%s</script>`, renderBaseAttrs(s.BaseNode), s.Code)
}

type Style struct {
	BaseNode
	Code string
}

func (s Style) Render() string {
	return fmt.Sprintf(`<style%s>%s</style>`, renderBaseAttrs(s.BaseNode), s.Code)
}

// ---------------- Helper ----------------
func Render(node Node) string { return node.Render() }
