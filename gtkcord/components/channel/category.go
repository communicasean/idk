package channel

import (
	"sort"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/ningen/v2"
)

type sortStructure struct {
	parent   discord.Channel
	children []discord.Channel
}

func FilterChannels(s *ningen.State, chs []discord.Channel) []discord.Channel {
	filtered := make([]discord.Channel, 0, len(chs))
	u, _ := s.Me()

	hiddenCats := make(map[discord.ChannelID]bool)
	for _, ch := range chs {
		if ch.Type == discord.GuildCategory {
			p, err := s.Permissions(ch.ID, u.ID)
			if err != nil {
				continue
			}

			if !p.Has(discord.PermissionViewChannel) {
				hiddenCats[ch.ID] = true
			}
		}
	}

	for _, ch := range chs {
		p, err := s.Permissions(ch.ID, u.ID)
		if err != nil {
			continue
		}

		if !p.Has(discord.PermissionViewChannel) || hiddenCats[ch.CategoryID] {
			continue
		}

		switch ch.Type {
		case discord.DirectMessage,
			discord.GuildText,
			discord.GuildCategory,
			discord.GroupDM:

		default:
			continue
		}

		filtered = append(filtered, ch)
	}

	return filtered
}

func transformChannels(s *ningen.State, chs []discord.Channel) []*Channel {
	tree := map[discord.ChannelID]*sortStructure{}

	for _, ch := range chs {
		if ch.Type == discord.GuildCategory {
			v, ok := tree[ch.ID]
			if ok {
				v.parent = ch
			} else {
				tree[ch.ID] = &sortStructure{
					parent: ch,
				}
			}

			continue
		}

		if ch.CategoryID.IsValid() {
			v, ok := tree[ch.CategoryID]
			if ok {
				v.children = append(v.children, ch)
			} else {
				tree[ch.CategoryID] = &sortStructure{
					children: []discord.Channel{ch},
				}
			}

			continue
		}

		tree[ch.ID] = &sortStructure{
			parent: ch,
		}
	}

	list := make([]*sortStructure, 0, len(tree))

	for _, v := range tree {
		if v.children != nil {
			sort.SliceStable(v.children, func(i, j int) bool {
				return v.children[i].Position < v.children[j].Position
			})
		}

		list = append(list, v)
	}

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].parent.Position < list[j].parent.Position
	})

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].children == nil
	})

	channels := make([]*Channel, 0, len(chs))

	for i := range list {
		sch := list[i]

		if sch.parent.ID.IsValid() {
			channels = append(channels, createChannelRead(&sch.parent, s))
		}

		for i := range sch.children {
			channels = append(channels, createChannelRead(&sch.children[i], s))
		}
	}

	return channels
}
