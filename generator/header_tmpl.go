// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package generator

var headerTmpl = `
package {{.}}
// Generated by https://github.com/hooklift/gowsdl
// Do not modify
import (
	"encoding/xml"
	"time"

	gowsdl "github.com/hooklift/gowsdl/generator"
	{{/*range .Imports*/}}
		{{/*.*/}}
	{{/*end*/}}
)

// against "unused imports"
var _ time.Time
var _ xml.Name
`
