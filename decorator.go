package main

type Vegetables struct {
	skewer ISkewer
}

func (v *Vegetables) getPrice() float32 {
	skewerPrice := v.skewer.getPrice()
	return skewerPrice + 0.2
}
func (v *Vegetables) getInfo() string {
	str := v.skewer.getInfo()
	return str + "with vegetables "
}

type Fat struct {
	skewer ISkewer
}

func (f *Fat) getPrice() float32 {
	skewerPrice := f.skewer.getPrice()
	return skewerPrice + 0.4
}
func (f *Fat) getInfo() string {
	str := f.skewer.getInfo()
	return str + "with fat "
}

type LemonSlices struct {
	skewer ISkewer
}

func (l *LemonSlices) getPrice() float32 {
	skewerPrice := l.skewer.getPrice()
	return skewerPrice + 0.1
}
func (l *LemonSlices) getInfo() string {
	str := l.skewer.getInfo()
	return str + "with lemon slices "
}
