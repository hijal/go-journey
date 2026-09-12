package main

import (
	"fmt"
	"maps"
	"slices"
)

type Server struct {
	Host string
	Env  string
	Up   bool
}

type DeployPlan struct {
	allowedEnvs map[string]struct{}
	targets     []*Server
}

func (d *DeployPlan) Allow(envs ...string) {
	if d.allowedEnvs == nil {
		d.allowedEnvs = make(map[string]struct{}, len(envs))
	}

	for _, e := range envs {
		d.allowedEnvs[e] = struct{}{}
	}
}

func (d *DeployPlan) Add(s *Server) {
	d.targets = append(d.targets, s)
}

func (d *DeployPlan) Run() []string {
	eligible := []string{}

	for _, s := range d.targets {
		_, envOk := d.allowedEnvs[s.Env]
		if envOk && s.Up {
			eligible = append(eligible, s.Host)
		}
	}
	return eligible
}

func main() {
	plan := new(DeployPlan)
	plan.Allow("prod", "staging")

	servers := []*Server{
		{Host: "web-01", Env: "prod", Up: true},
		{Host: "web-02", Env: "prod", Up: false}, // down: skipped
		{Host: "dev-01", Env: "dev", Up: true},   // env not allowed: skipped
		{Host: "stg-01", Env: "staging", Up: true},
	}

	for _, s := range servers {
		plan.Add(s)
	}

	picked := plan.Run()
	slices.Sort(picked)
	fmt.Println("deploy targets:", picked)

	done := make(map[string]struct{})

	for _, h := range picked {
		done[h] = struct{}{}
	}

	host := "stg-01"

	if _, ok := done[host]; ok {
		fmt.Printf("%s already deployed, skipping\n", host)
	}

	fmt.Println("all done:", slices.Sorted(maps.Keys(done)))
}
