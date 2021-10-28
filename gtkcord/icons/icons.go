package icons

/*
var IconTheme *gtk.IconTheme
var iconMap map[string]*gdk.Pixbuf
var iconMu sync.Mutex

func GetIcon(icon string, size int) *gdk.Pixbuf {
	var key = icon + "#" + strconv.Itoa(size)

	iconMu.Lock()
	defer iconMu.Unlock()

	if IconTheme == nil {
		v, err := semaphore.Idle(gtk.IconThemeGetDefault)
		if err != nil {
			log.Errorln("Failed to get icon theme:", err)
			return nil
		}

		IconTheme = v.(*gtk.IconTheme)
		iconMap = map[string]*gdk.Pixbuf{}
	}

	if v, ok := iconMap[key]; ok {
		p, _ := gdk.PixbufCopy(v)
		return p
	}

	v, err := semaphore.Idle(IconTheme.LoadIcon, icon, size,
		gtk.IconLookupFlags(gtk.ICON_LOOKUP_FORCE_SIZE))
	if err != nil {
		log.Errorln("Failed to get icon", icon)
		return nil
	}

	pb := v.(*gdk.Pixbuf)
	iconMap[key] = pb
	return pb
}

func GetIconUnsafe!!ILLEGAL!!(icon string, size int) *gdk.Pixbuf {
	var key = icon + "#" + strconv.Itoa(size)

	if IconTheme == nil {
		IconTheme, _ = gtk.IconThemeGetDefault()
		iconMap = map[string]*gdk.Pixbuf{}
	}

	if v, ok := iconMap[key]; ok {
		p, _ := gdk.PixbufCopy(v)
		return p
	}

	pb, err := IconTheme.LoadIcon(icon, size, gtk.IconLookupFlags(gtk.ICON_LOOKUP_FORCE_SIZE))
	if err != nil {
		log.Fatalln("Failed to load icon", icon)
	}

	iconMap[key] = pb
	return pb
}

var pngEncoder = png.Encoder{
	CompressionLevel: png.BestSpeed,
}

func Pixbuf(img image.Image) (*gdkpixbuf.Pixbuf, error) {
	if rgba, ok := img.(*image.RGBA); ok {
		return gdkpixbuf.NewPixbufFromBytes(
			glib.NewBytesWithGo(rgba.Pix), gdkpixbuf.ColorspaceRGB, true, 8,
			rgba.Rect.Dx(), rgba.Rect.Dy(), rgba.Stride,
		), nil
	}

	l, err := gdkpixbuf.NewPixbufLoaderWithType("png")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create an icon pixbuf loader")
	}

	if err := pngEncoder.Encode(gioutil.PixbufLoaderWriter(l), img); err != nil {
		return nil, errors.Wrap(err, "Failed to encode PNG")
	}

	if err := l.Close(); err != nil {
		return nil, errors.Wrap(err, "failed to close")
	}

	return l.Pixbuf(), nil
}

func PixbufIcon(img image.Image, size int) (*gdk.Pixbuf, error) {
	var buf bytes.Buffer

	if err := pngEncoder.Encode(&buf, img); err != nil {
		return nil, errors.Wrap(err, "Failed to encode PNG")
	}

	l, err := gdk.PixbufLoaderNewWithType("png")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create an icon pixbuf loader")
	}

	l.SetSize(size, size)

	p, err := l.WriteAndReturnPixbuf(buf.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "Failed to set icon to pixbuf")
	}

	return p, nil
}

func SetImage(img image.Image, gtkimg *gtk.Image) error {
	var buf bytes.Buffer

	if err := pngEncoder.Encode(&buf, img); err != nil {
		return errors.Wrap(err, "Failed to encode PNG")
	}

	l, err := gdk.PixbufLoaderNewWithType("png")
	if err != nil {
		return errors.Wrap(err, "Failed to create an icon pixbuf loader")
	}

	gtkutils.Connect(l, "area-updated", func() {
		p, err := l.GetPixbuf()
		if err != nil || p == nil {
			log.Errorln("Failed to get animation during area-prepared:", err)
			return
		}
		gtkimg.SetFromPixbuf(p)
	})

	if _, err := io.Copy(l, &buf); err != nil {
		return errors.Wrap(err, "Failed to stream to pixbuf_loader")
	}

	if err := l.Close(); err != nil {
		return errors.Wrap(err, "Failed to close pixbuf_loader")
	}

	return nil
}
*/
