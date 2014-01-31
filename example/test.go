package main

import (
	"fmt"
	"github.com/soh335/go-dnssd"
)

func main() {
	ctx, err := RegisterService()
	if err != nil {
		panic(err)
	}
	defer ctx.Release()
	Discover()
}

func RegisterService() (*dnssd.Context, error) {
	txtRecords := map[string]string{
		"path": "/path-to-page.html",
	}

	rc := make(chan *dnssd.RegisterReply)
	ctx, err := dnssd.ServiceRegister(
		dnssd.DNSServiceFlagsSuppressUnusable,
		0, // most applications will pass 0
		"My Server",
		"_http._tcp.",
		"",   // empty string ends up as local domain
		"",   // most applications do not specify a host
		3000, // port the service is running on
		txtRecords,
		rc,
	)

	go dnssd.Process(ctx)

	if err != nil {
		return nil, err
	}

	registerReply, _ := <-rc
	fmt.Println("Register Reply: ", registerReply)
	return ctx, nil
}

func Discover() {
	bc := make(chan *dnssd.BrowseReply)
	ctx, err := dnssd.Browse(dnssd.DNSServiceInterfaceIndexAny, "_http._tcp", bc)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer ctx.Release()
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

		defer rctx.Release()
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

		defer qctx.Release()
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

		defer gctx.Release()
		go dnssd.Process(gctx)

		getAddrInfoReply, _ := <-gc
		fmt.Println(getAddrInfoReply)
	}
}
