Hi Gophers! Unfortunately, I don't have the time to maintain this library right now, so I can't add any features. I'll gladly accept pull requests if you're interested. Also, if you'd like to maintaining the library longer term, I'm also happy to welcome you as a contributor. If you're interested in either submitting a PR or being a longer term maintainer, please let me know on the [Gophers Slack](https://invite.slack.golangbridge.org/). My username there is `arschles`.

# functy

Often times, [vecty](https://github.com/gopherjs/vecty) code will have lots of nested function calls and the code gets hard to read and manage. This module has convenience functions for building up vecty elements in a functional pattern.

Below is an example adapted from the vecty example code for rendering Markdown:

```go
import (
    "github.com/go-functional/functy"
    "github.com/gopherjs/vecty/elem"
    "github.com/gopherjs/vecty"
)

// PageView is our main page component.
type PageView struct {
    vecty.Core
    Input string
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
    body := functy.Wrap(elem.Body)
    div := functy.Wrap(elem.Div).
        Markup(vecty.Style("float", "right"))
    textArea := functy.Wrap(elem.TextArea).
        Markup(vecty.Style("font-family", "monospace")).
        Markup(vecty.Property("rows", 14)).
        Markup(vecty.Property("cols", 70)).
        Markup(event.Input(func(e *vecty.Event) {
            p.Input = e.Target.Get("value").Strting()
            vecty.Rerender(p)
        }))
        Markup(elem.TextArea).
        Child(vecty.Text(p.Input))
    return functy.Wrap(elem.Body).Child(div.Child(textArea))
}
```

Inspired by:

- [blaze-html](https://jaspervdj.be/blaze/tutorial.html)
- [htmlgo](https://github.com/julvo/htmlgo)

And built on the shoulders of [vecty](https://github.com/gopherjs/vecty), of course!
