package main

import (
	"flag"
	"log"
	"os/exec"
	"time"
)

var (
	UserName   string
	Password   string
	DomainName string
)

func init() {
	flag.StringVar(&Password, "password", "test123", "DDNS Password")
	flag.StringVar(&DomainName, "domain", "example.org", "The domain name to update")
}

func UpdateDomainName4() error {
	cmd := exec.Command("curl", "-4", "https://"+DomainName+":"+Password+"@dyn.dns.he.net/nic/update?hostname="+DomainName)
	return cmd.Run()
}
func UpdateDomainName6() error {
	cmd := exec.Command("curl", "-6", "https://"+DomainName+":"+Password+"@dyn.dns.he.net/nic/update?hostname="+DomainName)
	return cmd.Run()
}

func main() {
	var err error
	flag.Parse()
	t := time.NewTicker(time.Minute * 5)
	for _ = range t.C {
		err = UpdateDomainName4()
		if err != nil {
			log.Println(err)
		}
		err = UpdateDomainName6()
		if err != nil {
			log.Println(err)
		}
	}
}
