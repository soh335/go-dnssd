package dnssd

/*
#include <stdlib.h>
#include "dns_sd.h"
DNSServiceErrorType Browse(DNSServiceRef *, DNSServiceFlags, const char *, void *);
DNSServiceErrorType Resolve(DNSServiceRef *, DNSServiceFlags, uint32_t, const char *, const char *, const char *, void *);
DNSServiceErrorType QueryRecord(DNSServiceRef *, DNSServiceFlags, uint32_t, const char *, uint16_t, uint16_t, void *);
DNSServiceErrorType ServiceRegister(DNSServiceRef *, DNSServiceFlags, uint32_t, const char *, const char *, const char *, const char *, uint16_t, uint16_t, const void *, void*);
DNSServiceErrorType GetAddrInfo(DNSServiceRef *, DNSServiceFlags, uint32_t, DNSServiceProtocol, const char *, void *);
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"strings"
	"unsafe"
)

type Context struct {
	ref *C.DNSServiceRef
	c reflect.Value
}

type BrowseReply struct {
	InterfaceIndex uint32
	ServiceName    string
	RegType        string
	ReplyDomain    string
}

type ResolveReply struct {
	InterfaceIndex uint32
	FullName       string
	HostTarget     string
	Port           uint16
	TxtRecordMap   map[string]string
}

type QueryRecordReply struct {
	InterfaceIndex uint32
	FullName       string
	Rrtype         uint16
	Rrclass        uint16
	Rdlen          uint16
	Rdata          []byte
	Ttl            uint32
}

type RegisterReply struct {
	Name     string
	RegType  string
	Domain   string
}

type RecordSRV struct {
	Priority uint16
	Weight   uint16
	Port     uint16
	Host     string
}

type GetAddrInfoReply struct {
	InterfaceIndex uint32
	HostName       string
	Ip             string
	Ttl            uint32
}

func init() {
}

func Process(ctx *Context) {
	C.DNSServiceProcessResult(*(ctx.ref))
	C.DNSServiceRefDeallocate(*(ctx.ref))
	ctx.c.Close()
}

//export goBrowseReply
func goBrowseReply(interfaceIndex uint32, serviceName *C.char, regType *C.char, replyDomain *C.char, contextpt unsafe.Pointer) {
	br := &BrowseReply{interfaceIndex, C.GoString(serviceName), C.GoString(regType), C.GoString(replyDomain)}
	fn := (*(*func(*BrowseReply))(unsafe.Pointer(&contextpt)))
	fn(br)
}

func Browse(flags C.DNSServiceFlags, regType string, c chan *BrowseReply) (*Context, error) {

	var fn = func(browseReply *BrowseReply) {
		c <- browseReply
	}
	var fnptr = unsafe.Pointer(&fn)

	cregType := C.CString(regType)
	defer C.free(unsafe.Pointer(cregType))

	var ref C.DNSServiceRef
	cerr := C.Browse(&ref, flags, cregType, (*(*unsafe.Pointer)(fnptr)))

	if cerr == DNSServiceErr_NoError {
		return &Context{&ref, reflect.ValueOf(c)}, nil
	}

	return nil, createErr(cerr)
}

//export goResolveReply
func goResolveReply(interfaceIndex uint32, fullName *C.char, hostTarget *C.char, port uint16, txtLen uint16, txtRecords unsafe.Pointer, contextpt unsafe.Pointer) {

	txtRecordMap := make(map[string]string)
	count := int(C.TXTRecordGetCount((C.uint16_t)(txtLen), txtRecords))

	for i := 0; i < count; i++ {
		key := make([]byte, txtLen)
		var valueLen C.uint8_t
		var value unsafe.Pointer

		C.TXTRecordGetItemAtIndex(
			(C.uint16_t)(txtLen),
			txtRecords,
			(C.uint16_t)(i),
			256,
			(*C.char)(unsafe.Pointer(&key[0])),
			&valueLen,
			&value,
		)

		byteValue := C.GoBytes(value, C.int(valueLen))
		txtRecordMap[string(key)] = string(byteValue)
	}

	rr := &ResolveReply{interfaceIndex, C.GoString(fullName), C.GoString(hostTarget), port, txtRecordMap}
	fn := (*(*func(*ResolveReply))(unsafe.Pointer(&contextpt)))
	fn(rr)
}

func Resolve(flags C.DNSServiceFlags, interfaceIndex uint32, serviceName string, regType string, replyDomain string, c chan *ResolveReply) (*Context, error) {
	var fn = func(resolveReply *ResolveReply) {
		c <- resolveReply
	}
	var fnptr = unsafe.Pointer(&fn)

	var ref C.DNSServiceRef
	cserviceName := C.CString(serviceName)
	cregType := C.CString(regType)
	creplyDomain := C.CString(replyDomain)

	defer C.free(unsafe.Pointer(cserviceName))
	defer C.free(unsafe.Pointer(cregType))
	defer C.free(unsafe.Pointer(creplyDomain))

	cerr := C.Resolve(
		&ref,
		flags,
		(C.uint32_t)(interfaceIndex),
		cserviceName,
		cregType,
		creplyDomain,
		(*(*unsafe.Pointer)(fnptr)),
	)

	if cerr == DNSServiceErr_NoError {
		return &Context{&ref, reflect.ValueOf(c)}, nil
	}

	return nil, createErr(cerr)
}

//export goQueryRecordReply
func goQueryRecordReply(interfaceIndex uint32, fullName *C.char, rrtype uint16, rrclass uint16, rdlen uint16, rdata unsafe.Pointer, ttl uint32, contextpt unsafe.Pointer) {
	qrr := &QueryRecordReply{interfaceIndex, C.GoString(fullName), rrtype, rrclass, rdlen, C.GoBytes(rdata, C.int(rdlen)), ttl}
	fn := (*(*func(*QueryRecordReply))(unsafe.Pointer(&contextpt)))
	fn(qrr)
}

//export goServiceRegisterReply
func goServiceRegisterReply(name *C.char, regType *C.char, domain *C.char, contextpt unsafe.Pointer) {
	srr := &RegisterReply{C.GoString(name), C.GoString(regType), C.GoString(domain) }
	fn := (*(*func(*RegisterReply))(unsafe.Pointer(&contextpt)))
	fn(srr)
}

func QueryRecord(flags C.DNSServiceFlags, interfaceIndex uint32, fullName string, rrtype C.uint16_t, rrclass C.uint16_t, c chan *QueryRecordReply) (*Context, error) {

	var fn = func(queryRecordReply *QueryRecordReply) {
		c <- queryRecordReply
	}
	var fnptr = unsafe.Pointer(&fn)

	var ref C.DNSServiceRef
	cfullName := C.CString(fullName)

	defer C.free(unsafe.Pointer(cfullName))

	cerr := C.QueryRecord(
		&ref,
		flags,
		(C.uint32_t)(interfaceIndex),
		cfullName,
		rrtype,
		rrclass,
		(*(*unsafe.Pointer)(fnptr)),
	)

	if cerr == DNSServiceErr_NoError {
		return &Context{&ref, reflect.ValueOf(c)}, nil
	}

	return nil, createErr(cerr)
}

func ServiceRegister(flags C.DNSServiceFlags, interfaceIndex uint32, name string, regType string, domain string, host string, port uint16, txtRecords map[string] string, c chan *RegisterReply) (*Context, error){
	var fn = func(registerReply *RegisterReply) {
		c <- registerReply
	}
	var fnptr = unsafe.Pointer(&fn)

	var ref C.DNSServiceRef
	c_name    := C.CString(name)
	c_regType := C.CString(regType)
	c_domain  := C.CString(domain)
	c_host    := C.CString(host)

	defer C.free(unsafe.Pointer(c_name))
	defer C.free(unsafe.Pointer(c_regType))
	defer C.free(unsafe.Pointer(c_domain))
	defer C.free(unsafe.Pointer(c_host))

	// Setup text record
	var txtRecordRef C.TXTRecordRef
	C.TXTRecordCreate(
		&txtRecordRef,
		0,   // let the c library manage it's own buffer
		nil,
	)

	// Add text records
	for key, value := range txtRecords {
		c_key   := C.CString(key)
		c_value := C.CString(value)

		defer C.free(unsafe.Pointer(c_key))
		defer C.free(unsafe.Pointer(c_value))

		txtadderr := C.TXTRecordSetValue(
			&txtRecordRef,
			c_key,
			(C.uint8_t)(len(value)),
			unsafe.Pointer(c_value),
		)

		if txtadderr != DNSServiceErr_NoError {
			return nil, createErr(txtadderr)
		}
	}

	cerr := C.ServiceRegister(
		&ref,
		flags,
		(C.uint32_t)(interfaceIndex),
		c_name,
		c_regType,
		c_domain,
		c_host,
		(C.uint16_t)(port),
		C.TXTRecordGetLength(&txtRecordRef),
		C.TXTRecordGetBytesPtr(&txtRecordRef),
		(*(*unsafe.Pointer)(fnptr)),
	)

	if cerr == DNSServiceErr_NoError {
		return &Context{&ref, reflect.ValueOf(c)}, nil
	}

	return nil, createErr(cerr)
}

//export goGetAddrInfoReply
func goGetAddrInfoReply(interfaceIndex uint32, hostName *C.char, ip *C.char, ttl uint32, contextpt unsafe.Pointer) {
	gair := &GetAddrInfoReply{interfaceIndex, C.GoString(hostName), C.GoString(ip), ttl}
	fn := (*(*func(*GetAddrInfoReply))(unsafe.Pointer(&contextpt)))
	fn(gair)
}

func GetAddrInfo(flags C.DNSServiceFlags, interfaceIndex uint32, protocol C.DNSServiceProtocol, hostName string, c chan *GetAddrInfoReply) (*Context, error) {
	var fn = func(getAddrInfoReply *GetAddrInfoReply) {
		c <- getAddrInfoReply
	}
	var fnptr = unsafe.Pointer(&fn)

	var ref C.DNSServiceRef
	chostName := C.CString(hostName)

	defer C.free(unsafe.Pointer(chostName))

	cerr := C.GetAddrInfo(
		&ref,
		flags,
		(C.uint32_t)(interfaceIndex),
		protocol,
		chostName,
		(*(*unsafe.Pointer)(fnptr)),
	)

	if cerr == DNSServiceErr_NoError {
		return &Context{&ref, reflect.ValueOf(c)}, nil
	}

	return nil, createErr(cerr)
}

func (qrr *QueryRecordReply) SRV() *RecordSRV {
	reader := bytes.NewReader(qrr.Rdata)
	var priority, weight, port uint16
	var host string

	binary.Read(reader, binary.BigEndian, &priority)
	binary.Read(reader, binary.BigEndian, &weight)
	binary.Read(reader, binary.BigEndian, &port)

	hosts := make([]string, 0)

	for reader.Len() > 0 {
		var c byte
		binary.Read(reader, binary.BigEndian, &c)
		line := make([]byte, c)
		binary.Read(reader, binary.BigEndian, line)
		hosts = append(hosts, string(line))
	}

	host = strings.Join(hosts, ".")

	return &RecordSRV{priority, weight, port, host}
}
