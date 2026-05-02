package main

import (
	"net"
	"strings"
	"sync"
	"time"
)

type Member struct {
	node     Node
	lastseen time.Time
}

type Registry struct {
	members map[string]Member
	address string
	mu      sync.Mutex
}

func NewRegistry(address string) *Registry {
	return &Registry{
		members: map[string]Member{},
		address: address,
	}
}

func (r *Registry) Register(n Node) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.members[n.name] = Member{n, time.Now()}
}

func (r *Registry) Heartbeat(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	member := r.members[name]
	member.lastseen = time.Now()
	r.members[name] = member
}

//r.CheckHealth(r)

func (reg *Registry) CheckHealth(ring *Ring) {
	reg.mu.Lock()
	defer reg.mu.Unlock()

	for nodename, member := range reg.members {
		if time.Since(member.lastseen) > 10*time.Second {
			// remove from the member map
			delete(reg.members, nodename)
			ring.RemoveNode(member.node)
		}
	}
}

func (reg *Registry) StartHealthChecker(ring *Ring) {
	for {
		time.Sleep(5 * time.Second)
		reg.CheckHealth(ring)

	}
}

func (reg *Registry) ListenHeartBeat() {
	ln, err := net.Listen("tcp", reg.address)
	handleError(err)

	for {
		conn, err := ln.Accept()

		handleError(err)
		go handleConnection(reg, conn)
	}
}

func handleConnection(r *Registry, c net.Conn) {

	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	handleError(err)

	nodename := strings.TrimSpace(string(buf[:n]))
	r.Heartbeat(nodename)

}
