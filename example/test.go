package main

import (
	"fmt"
	"github.com/soh335/go-dnssd"
)

func main() {
	bc := make(chan *dnssd.BrowseReply)
	ctx, err := dnssd.Browse(dnssd.DNSServiceInterfaceIndexAny, "_http._tcp", bc)

	if err != nil {
		fmt.Println(err)
		return
	}

	go dnssd.Process(ctx)

	for {
		browseReply, ok := <-bc
		if !ok {
			fmt.Println("closed")
			break
		}
		fmt.Println(browseReply)

		fmt.Println("start resolve")

		rc := make(chan *dnssd.ResolveReply)
		rctx, err := dnssd.Resolve(
			dnssd.DNSServiceFlagsForceMulticast,
			browseReply.InterfaceIndex,
			browseReply.ServiceName,
			browseReply.RegType,
			browseReply.ReplyDomain,
			rc,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		go dnssd.Process(rctx)

		resolveReply, _ := <-rc
		fmt.Println(resolveReply)

		qc := make(chan *dnssd.QueryRecordReply)
		qctx, err := dnssd.QueryRecord(
			dnssd.DNSServiceFlagsForceMulticast,
			resolveReply.InterfaceIndex,
			resolveReply.FullName,
			dnssd.DNSServiceType_SRV,
			dnssd.DNSServiceClass_IN,
			qc,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		go dnssd.Process(qctx)

		queryRecordReply, _ := <-qc
		fmt.Println(queryRecordReply)
		fmt.Println(queryRecordReply.SRV())

		gc := make(chan *dnssd.GetAddrInfoReply)
		gctx, err := dnssd.GetAddrInfo(
			dnssd.DNSServiceFlagsForceMulticast,
			0,
			dnssd.DNSServiceProtocol_IPv4,
			resolveReply.HostTarget,
			gc,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		go dnssd.Process(gctx)

		getAddrInfoReply, _ := <-gc
		fmt.Println(getAddrInfoReply)
	}
}
