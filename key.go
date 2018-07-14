package arcane

type Key string

func (k *Key) String() string {
	return "Key: :" + string(*k)
}

func (k *Key) arcaneType() {}
func (k *Key) atomic()     {}
