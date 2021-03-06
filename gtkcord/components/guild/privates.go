package guild

import (
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
	"github.com/diamondburned/gtkcord3/gtkcord/components/channel"
	"github.com/diamondburned/gtkcord3/gtkcord/gtkutils"
	"github.com/diamondburned/gtkcord3/internal/log"
	"github.com/diamondburned/ningen/v2"
)

type DMButton struct {
	*gtk.ListBoxRow
	Unread *UnreadStrip

	OnClick func()

	class string
}

// thread-safe
func NewPMButton(s *ningen.State) (dm *DMButton) {
	r := gtk.NewListBoxRow()
	r.SetHAlign(gtk.AlignFill)
	r.SetVAlign(gtk.AlignCenter)
	r.SetActivatable(true)
	marginate(r)
	r.Show()
	gtkutils.InjectCSS(r, "dmbutton", "")

	i := gtk.NewImageFromIconName("system-users-symbolic", 0)
	i.SetSizeRequest(IconSize, IconSize)
	i.SetPixelSize(IconSize / 3 * 2)
	i.SetHAlign(gtk.AlignCenter)
	i.SetVAlign(gtk.AlignCenter)
	i.Show()

	ov := NewUnreadStrip(i)
	r.Add(ov)

	dm = &DMButton{
		ListBoxRow: r,
		Unread:     ov,
	}

	name := "Private Messages"
	BindName(r, ov, &name)

	// Initialize the read state.
	dm.resetRead(s)

	return
}

func (dm *DMButton) setUnread(unread bool) {
	if unread {
		dm.Unread.SetPinged()
	} else {
		dm.Unread.SetRead()
	}
}

func (dm *DMButton) inactive() {
	dm.Unread.SetActive(false)
}

func (dm *DMButton) resetRead(s *ningen.State) {
	go func() {
		// Find and detect any unread channels:
		chs, err := s.PrivateChannels()
		if err != nil {
			log.Errorln("failed to get private channels for DMButton:", err)
			return
		}

		channel.ScanUnreadDMs(s, chs, func(ch *discord.Channel) {
			glib.IdleAdd(func() { dm.setUnread(true) })
		})
	}()
}
