package main

import (
        "fmt"
        "github.com/soh335/go-dnssd"
)

func main() {
        bc := make(chan *dnssd.BrowseReply)
        go dnssd.Browse(dnssd.DNSServiceInterfaceIndexAny, "_http._tcp", bc)

        for {
                browseReply, ok := <-bc
                if !ok {
                  fmt.Println("closed")
                  break
                }
                fmt.Println(browseReply)

                fmt.Println("start resolve")

                rc := make(chan *dnssd.ResolveReply)
                go dnssd.Resolve(
                        dnssd.DNSServiceFlagsForceMulticast,
                        browseReply.InterfaceIndex,
                        browseReply.ServiceName,
                        browseReply.RegType,
                        browseReply.ReplyDomain,
                        rc,
                )

                resolveReply, _ := <-rc
                fmt.Println(resolveReply)

                qc := make(chan *dnssd.QueryRecordReply)
                go dnssd.QueryRecord(
                        dnssd.DNSServiceFlagsForceMulticast,
                        resolveReply.InterfaceIndex,
                        resolveReply.FullName,
                        dnssd.DNSServiceType_SRV,
                        dnssd.DNSServiceClass_IN,
                        qc,
                )

                queryRecordReply, _ := <-qc
                fmt.Println(queryRecordReply)
                fmt.Println(queryRecordReply.SRV())

                gc := make(chan *dnssd.GetAddrInfoReply)
                go dnssd.GetAddrInfo(
                        dnssd.DNSServiceFlagsForceMulticast,
                        0,
                        dnssd.DNSServiceProtocol_IPv4,
                        resolveReply.HostTarget,
                        gc,
                )

                getAddrInfoReply, _ := <-gc
                fmt.Println(getAddrInfoReply)
        }
}

