// Code generated by qtc from "basepage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line examples/basicserver/templates/basepage.qtpl:3
package templates

//line examples/basicserver/templates/basepage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line examples/basicserver/templates/basepage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line examples/basicserver/templates/basepage.qtpl:4
type Page interface {
//line examples/basicserver/templates/basepage.qtpl:4
	Title() string
//line examples/basicserver/templates/basepage.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line examples/basicserver/templates/basepage.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line examples/basicserver/templates/basepage.qtpl:4
	Body() string
//line examples/basicserver/templates/basepage.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line examples/basicserver/templates/basepage.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line examples/basicserver/templates/basepage.qtpl:4
}

// Page prints a page implementing Page interface.

//line examples/basicserver/templates/basepage.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line examples/basicserver/templates/basepage.qtpl:12
	qw422016.N().S(`
<html>
	<head>
		<title>`)
//line examples/basicserver/templates/basepage.qtpl:15
	p.StreamTitle(qw422016)
//line examples/basicserver/templates/basepage.qtpl:15
	qw422016.N().S(`</title>
	</head>
	<body>
		<div>
			<a href="/">return to main page</a>
		</div>
		`)
//line examples/basicserver/templates/basepage.qtpl:21
	p.StreamBody(qw422016)
//line examples/basicserver/templates/basepage.qtpl:21
	qw422016.N().S(`
	</body>
</html>
`)
//line examples/basicserver/templates/basepage.qtpl:24
}

//line examples/basicserver/templates/basepage.qtpl:24
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line examples/basicserver/templates/basepage.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line examples/basicserver/templates/basepage.qtpl:24
	StreamPageTemplate(qw422016, p)
//line examples/basicserver/templates/basepage.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line examples/basicserver/templates/basepage.qtpl:24
}

//line examples/basicserver/templates/basepage.qtpl:24
func PageTemplate(p Page) string {
//line examples/basicserver/templates/basepage.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
//line examples/basicserver/templates/basepage.qtpl:24
	WritePageTemplate(qb422016, p)
//line examples/basicserver/templates/basepage.qtpl:24
	qs422016 := string(qb422016.B)
//line examples/basicserver/templates/basepage.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
//line examples/basicserver/templates/basepage.qtpl:24
	return qs422016
//line examples/basicserver/templates/basepage.qtpl:24
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line examples/basicserver/templates/basepage.qtpl:29
type BasePage struct{}

//line examples/basicserver/templates/basepage.qtpl:30
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line examples/basicserver/templates/basepage.qtpl:30
	qw422016.N().S(`This is a base title`)
//line examples/basicserver/templates/basepage.qtpl:30
}

//line examples/basicserver/templates/basepage.qtpl:30
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line examples/basicserver/templates/basepage.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line examples/basicserver/templates/basepage.qtpl:30
	p.StreamTitle(qw422016)
//line examples/basicserver/templates/basepage.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line examples/basicserver/templates/basepage.qtpl:30
}

//line examples/basicserver/templates/basepage.qtpl:30
func (p *BasePage) Title() string {
//line examples/basicserver/templates/basepage.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
//line examples/basicserver/templates/basepage.qtpl:30
	p.WriteTitle(qb422016)
//line examples/basicserver/templates/basepage.qtpl:30
	qs422016 := string(qb422016.B)
//line examples/basicserver/templates/basepage.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
//line examples/basicserver/templates/basepage.qtpl:30
	return qs422016
//line examples/basicserver/templates/basepage.qtpl:30
}

//line examples/basicserver/templates/basepage.qtpl:31
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line examples/basicserver/templates/basepage.qtpl:31
	qw422016.N().S(`This is a base body`)
//line examples/basicserver/templates/basepage.qtpl:31
}

//line examples/basicserver/templates/basepage.qtpl:31
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line examples/basicserver/templates/basepage.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line examples/basicserver/templates/basepage.qtpl:31
	p.StreamBody(qw422016)
//line examples/basicserver/templates/basepage.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line examples/basicserver/templates/basepage.qtpl:31
}

//line examples/basicserver/templates/basepage.qtpl:31
func (p *BasePage) Body() string {
//line examples/basicserver/templates/basepage.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
//line examples/basicserver/templates/basepage.qtpl:31
	p.WriteBody(qb422016)
//line examples/basicserver/templates/basepage.qtpl:31
	qs422016 := string(qb422016.B)
//line examples/basicserver/templates/basepage.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
//line examples/basicserver/templates/basepage.qtpl:31
	return qs422016
//line examples/basicserver/templates/basepage.qtpl:31
}
