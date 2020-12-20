package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct{
	Host string
	Port int
}

type RegistryClient interface {
	Register(addr string, port int, id string, name string, tags []string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port int) RegistryClient{
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry)Register(addr string, port int, id string,name string,  tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(cfg)
	if err != nil{
		panic(err)
	}
	check := &api.AgentServiceCheck{
		HTTP: fmt.Sprintf("http://%s:%d/health", addr, port),
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Address = addr
	registration.Tags = tags
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil{
		panic(err)
	}
	return nil
}

func (r *Registry) DeRegister(serviceId string) error{
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(cfg)
	if err != nil{
		panic(err)
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}

