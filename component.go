package main

type ISkewer interface {
	getPrice() float32
	getInfo() string
}

type EmptySkewer struct {
}

func (e *EmptySkewer) getPrice() float32 {
	return 0
}

func (e *EmptySkewer) getInfo() string {
	return "Skewer "
}

type ChickenSkewer struct {
}

func (c *ChickenSkewer) getPrice() float32 {
	return 2.1
}

func (c *ChickenSkewer) getInfo() string {
	return "Chicken "
}

type LambSkewer struct {
}

func (l *LambSkewer) getPrice() float32 {
	return 2.3
}

func (l *LambSkewer) getInfo() string {
	return "Lamb "
}

type BeefSkewer struct {
}

func (b *BeefSkewer) getPrice() float32 {
	return 2.2
}

func (b *BeefSkewer) getInfo() string {
	return "Beef "
}
