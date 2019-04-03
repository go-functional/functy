package functy

import "github.com/gopherjs/vecty"

// Wrapper is the top level struct for any vecty element. Don't create one
// of these structs directly. Instead, call the Wrap function to bring
// elements from the vecty world into the functy world
type Wrapper struct {
	markupOrChild []vecty.MarkupOrChild
	fn            func(...vecty.MarkupOrChild) *vecty.HTML
}

// Markup adds some markup to the vecty element in the wrapper
func (w Wrapper) Markup(appliers ...vecty.Applyer) Wrapper {
	w.markupOrChild = append(w.markupOrChild, vecty.Markup(appliers...))
	return w
}

// Child adds a child to the element inside this wrapper
func (w Wrapper) Child(fn func(...vecty.MarkupOrChild) *vecty.HTML) Wrapper {
	parentFn := w.fn
	childFn := fn
	return Wrapper{
		fn: func(args ...vecty.MarkupOrChild) *vecty.HTML {
			return parentFn(childFn(args...))
		},
	}
}

// PeerFn wraps fn into a puts w and the
func (w Wrapper) PeerFn(
	parent Wrapper,
	fn func(...vecty.MarkupOrChild) *vecty.HTML) Wrapper {

}

// Apply converts the wrapper into a vecty.HTML so you can use it in your
// "regularly scheduled" code
func (w Wrapper) Apply() *vecty.HTML {
	return w.fn(w.markupOrChild...)
}
