package window

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
	"github.com/diamondburned/gtkcord3/gtkcord/components/animations"
)

const LoadingTitle = "Connecting to Discord — gtkcord3"

var previousLoadingChild gtk.Widgetter

// NowLoading fades the internal stack view to show a spinning circle.
func NowLoading() {
	// Use a spinner:
	s := animations.NewSpinner(75)

	previousLoadingChild = Window.header

	// Use a custom header instead of the actual Header:
	h := gtk.NewHeaderBar()
	h.SetTitle(LoadingTitle)
	h.SetShowCloseButton(true)
	h.ShowAll()

	// Set the loading animation:
	stackSet(Window.main, "loading", s)
	SetHeader(h)
}
