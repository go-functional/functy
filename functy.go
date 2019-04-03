package functy

import (
	"github.com/gopherjs/vecty"
)

// Wrap some standard vecty HTML so that you can manipulate it in the
// functional style. For example, here's how you'd build a <form>:
//
//	form := Wrap(elem.Form). // don't _call_ elem.Form() here!
//		Markup(vecty.Class("sample-form")). // add class="sample-form"
//		Child(elem.Input). // add an <input> element into this form
//		Markup(vecty.Class("edit")). // add class="edit"
//		Markup(prop.Value("edit something here!")). // value="edit something here!"
//		Event(event.Input(inputCallback)) // call the (go) function called inputCallback()
func Wrap(fn func(...vecty.MarkupOrChild) *vecty.HTML) Wrapper {
	return Wrapper{
		fn: fn,
	}
}
