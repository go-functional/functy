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

// ChildFn returns a copy of w with the child inside of it
func (w Wrapper) ChildFn(fn func(...vecty.MarkupOrChild) *vecty.HTML) Wrapper {
	parentFn := w.fn
	childFn := fn
	return Wrapper{
		fn: func(args ...vecty.MarkupOrChild) *vecty.HTML {
			return parentFn(childFn(args...))
		},
	}
}

// Peer returns a new Wrapper with children as first-level children
// of parent
func Peer(parent Wrapper, children ...Wrapper) Wrapper {
	for _, child := range children {
		parent.markupOrChild = append(parent.markupOrChild, child)
	}
	return parent
}

// Apply converts the wrapper into a vecty.HTML so you can use it in your
// "regularly scheduled" code
func (w Wrapper) Apply() *vecty.HTML {
	return w.fn(w.markupOrChild...)
}
