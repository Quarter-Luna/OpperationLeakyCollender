package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			// fmt.Println(port, "closed") // records what ports are closed
		}
		return
	}
	conn.Close()
	fmt.Println(port, "open") // records if a port is open
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(port)
	}
}

func GetOutIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	var ips []string
	fmt.Println("Enter ips of computers ending with a '0'")
	for scn.Scan() {
		line := scn.Text()
		if line == "0" {
			break
		}
		ips = append(ips, line)
	}

	if len(ips) > 0 {
		for _, ip := range ips {
			var f int
			var l int
			var t int
			ps := &PortScanner{
				ip:   ip,
				lock: semaphore.NewWeighted(Ulimit()),
			}
			fmt.Println("enter ports and time for", ip)
			fmt.Println("Start port: ")
			fmt.Scanln(&f)
			fmt.Println("Last Port: ")
			fmt.Scanln(&l)
			fmt.Println("Time in milli: ")
			fmt.Scanln(&t)
			fmt.Println("Open ports for", ip, ":")

			ps.Start(f, l, time.Duration(t)*time.Millisecond)
			fmt.Println()
		}
	}
	/*
		ps := &PortScanner{
			ip:   "127.0.0.1",
			lock: semaphore.NewWeighted(Ulimit()),
		}

		ps.Start(1, 1024, 500*time.Millisecond)
		// fmt.Println(GetOutIP())
	*/
}
