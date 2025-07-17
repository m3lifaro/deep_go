package homework6

import "fmt"

type Option func(*GamePerson)

func WithName(name string) Option {
	return func(p *GamePerson) {
		copy(p.name[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		manaMask := uint(mana) & 0b1111111111
		person.hm[0] = byte(manaMask >> 6)
		person.hm[1] = byte(manaMask&0b111111<<2 | uint(person.hm[1])&0b00000011)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		healthMask := uint(health) & 0b1111111111
		person.hm[1] = byte(uint(person.hm[1])&0b11111100 | (healthMask >> 8))
		person.hm[2] = byte(healthMask & 0b11111111)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		var p uint16
		var resp = uint16(respect << 12)
		var other = person.attrs & 0xFFF
		p = resp | other
		person.attrs = p
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		var strgth = uint16(strength << 8)
		person.attrs = person.attrs&0xF000 | strgth | person.attrs&0xFF
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		var exp = uint16(experience << 4)
		person.attrs = person.attrs&0xFF00 | exp | person.attrs&0xF
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		var lvl = uint16(level & 0xF)
		person.attrs = person.attrs&0xFFF0 | lvl
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		bMap := person.params & 0xF
		house := bMap | 0b0100
		person.params = person.params | bMap | house
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		bMap := person.params & 0xF
		gun := bMap | 0b0010
		person.params = person.params | bMap | gun
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		bMap := person.params & 0xF
		family := bMap | 0b0001
		person.params = person.params | bMap | family
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		pType := uint8(personType << 4)
		person.params = pType | person.params&0xF
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	x      int32
	y      int32
	z      int32
	gold   uint32
	hm     [3]byte
	params uint8
	attrs  uint16
	name   [42]byte
}

func NewGamePerson(options ...Option) GamePerson {
	var g = GamePerson{}
	for _, option := range options {
		option(&g)
	}
	return g
}

func (p *GamePerson) Name() string {
	n := 0
	for n < len(p.name) && p.name[n] != 0 {
		n++
	}
	return string(p.name[:n])
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	var leading = uint(p.hm[0]) & 0b1111
	var trailing = (uint(p.hm[1]) >> 2) & 0b111111
	var res = leading<<6 | trailing
	fmt.Printf("Двоичное представление: %b\n", leading)
	fmt.Printf("Двоичное представление: %b\n", trailing)
	fmt.Printf("Двоичное представление: %b\n", res)
	return int(res)
}

func (p *GamePerson) Health() int {
	var leading = uint(p.hm[1]) & 0b11
	var trailing = uint(p.hm[2]) & 0b11111111
	var res = leading<<8 | trailing
	return int(res)
}

func (p *GamePerson) Respect() int {
	return int(p.attrs & 0xF000 >> 12)
}

func (p *GamePerson) Strength() int {
	return int(p.attrs & 0xF00 >> 8)
}

func (p *GamePerson) Experience() int {
	return int(p.attrs & 0xF0 >> 4)
}

func (p *GamePerson) Level() int {
	return int(p.attrs & 0xF)
}

func (p *GamePerson) HasHouse() bool {
	return (p.params & 0b0100) == 0b0100
}

func (p *GamePerson) HasGun() bool {
	return (p.params & 0b0010) == 0b0010
}

func (p *GamePerson) HasFamilty() bool {
	return (p.params & 0b0001) == 0b0001
}

func (p *GamePerson) Type() int {
	return int(p.params & 0xF0 >> 4)
}
