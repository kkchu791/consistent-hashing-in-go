package main

type Producer struct {
	ring *Ring
}

func NewProducer(ring *ring) *Producer {
	return &Producer{
		ring: ring,
	}
}

func (p *Producer) Send(key string, message Message) error {
	// it needs to send to the right node

	// get the node

	nodeipaddress := p.ring.GetNode(key)

	// you need to send data to the ip address next

}
