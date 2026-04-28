package main

import "time"

type Member struct {
	node     Node
	lastseen time.Time
}

type Registry struct {
	members map[string]Member
}

func NewRegistry() *Registry {
	return &Registry{
		members: map[string]Member{},
	}
}

func (r *Registry) Register(n Node) {
	r.members[n.name] = Member{n, time.Now()}
}

func (r *Registry) Heartbeat(name string) {
	member := r.members[name]
	member.lastseen = time.Now()
	r.members[name] = member
}

//r.CheckHealth(r)

func (reg *Registry) CheckHealth(ring *Ring) {
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
