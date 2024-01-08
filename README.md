# Gotmpl

This is a modifiaction layer ontop of the `html/template` library provided by the go standard library. Originally pulled from the [extemplate](https://github.com/dannyvankooten/extemplate) library created by Danny van Kooten.

This extension layer allows for an inheritence based model, where templates can extend other templates and inject custom content inside them through a `component` system.

This is less "go-like", but suites my own personal development preference when building websites, coming from NextJS and React.

The modifications made by Jake Landers fix a few bugs found in the original version, and define some default functions that make development with the templates a bit easier.

## How to Use

```go
import "github.com/sapphirenw/gotmpl"

var init() {
    var XT *gotmpl.Extemplate
}

func main() {
    XT = gotmpl.New()
    ...
}
```

## Example

`base.html` defines blocks of content like:

```html
<div class="p-4 safe-area w-full mt-[80px] z-0">
    {{block "content" .}}{{end}}
</div>
```

Which allows the template to override these blocks like:

```html
{{extends "base.html"}}

{{define "title"}}<title>W. Notepad - Search</title>{{end}}

{{define "content"}}

<div class="pb-[100px]">
    {{if ne .Category ""}}
    <p class="text-2xl md:text-3xl font-bold my-4 md:my-8">{{.Category}}</p>
    {{end}}
    {{if ne .Keyword ""}}
    <p class="text-2xl md:text-3xl font-bold my-4 md:my-8"># {{.Keyword}}</p>
    {{end}}
    {{if ne .SearchText ""}}
    <p class="text-2xl md:text-3xl font-bold my-4 md:my-8">Results for "{{.SearchText}}"</p>
    {{end}}

    {{template "articles.html" .}}
</div>

{{end}}
```

Then the base is rendered with the content defined inside the template file in the necessary locations.
