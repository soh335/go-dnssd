#include "dns_sd.h"
#include <stdlib.h>
#include <stdio.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include "_cgo_export.h"

void BrowseReply
(
    DNSServiceRef sdRef,
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    DNSServiceErrorType errorCode,
    const char *serviceName,
    const char *regType,
    const char *replyDomain,
    void *context
)
{
    goBrowseReply(
        interfaceIndex,
        (char *)serviceName,
        (char *)regType,
        (char *)replyDomain,
        context
    );
}

DNSServiceErrorType Browse
(
    DNSServiceFlags flags,
    const char* regType,
    void *context
)
{
    DNSServiceRef BrowseRef;
    DNSServiceErrorType err = DNSServiceBrowse(&BrowseRef, 0, flags, regType, NULL, BrowseReply, context);

    if ( err == kDNSServiceErr_NoError ) {
        DNSServiceProcessResult(BrowseRef);
        DNSServiceRefDeallocate(BrowseRef);
    }

    return err;
}

void ResolveReply
(
    DNSServiceRef sdRef,
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    DNSServiceErrorType errorCode,
    const char *fullName,
    const char *hostTarget,
    uint16_t port,
    uint16_t txtLen,
    const unsigned char *txtRecords,
    void *context
)
{
    goResolveReply(
        interfaceIndex,
        (char *)fullName,
        (char *)hostTarget,
        port,
        txtLen,
        (void *)txtRecords,
        context
    );
}

DNSServiceErrorType Resolve
(
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    const char *serviceName,
    const char *regType,
    const char *replyDomain,
    void *context
)
{
    DNSServiceRef ResolveRef;
    DNSServiceErrorType err = DNSServiceResolve(&ResolveRef, flags, interfaceIndex, serviceName, regType, replyDomain, ResolveReply, context);

    if ( err == kDNSServiceErr_NoError ) {
        DNSServiceProcessResult(ResolveRef);
        DNSServiceRefDeallocate(ResolveRef);
    }

    return err;
}

void QueryRecordReply
(
    DNSServiceRef sdRef,
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    DNSServiceErrorType errorCode,
    const char *fullName,
    uint16_t rrtype,
    uint16_t rrclass,
    uint16_t rdlen,
    const void *rdata,
    uint32_t ttl,
    void *context
)
{
    goQueryRecordReply(
            interfaceIndex,
            (char *)fullName,
            rrtype,
            rrclass,
            rdlen,
            (void *)rdata,
            ttl,
            context
    );
}

DNSServiceErrorType QueryRecord
(
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    const char *fullName,
    uint16_t rrtype,
    uint16_t rrclass,
    void *context
)
{
    DNSServiceRef QueryRecordRef;
    DNSServiceErrorType err = DNSServiceQueryRecord(&QueryRecordRef, flags, interfaceIndex, fullName, rrtype, rrclass, QueryRecordReply, context);

    if ( err == kDNSServiceErr_NoError ) {
        DNSServiceProcessResult(QueryRecordRef);
        DNSServiceRefDeallocate(QueryRecordRef);
    }

    return err;
}

void GetAddrInfoReply
(
    DNSServiceRef sdRef,
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    DNSServiceErrorType errorCode,
    const char *hostName,
    const struct sockaddr *address,
    uint32_t ttl,
    void *context
)
{
    char ip[INET6_ADDRSTRLEN];
    switch( address->sa_family ) {
        case AF_INET:
            inet_ntop(AF_INET, &(((struct sockaddr_in *)(address))->sin_addr), ip, INET6_ADDRSTRLEN);
            break;
        case AF_INET6:
            inet_ntop(AF_INET6, &(((struct sockaddr_in6 *)(address))->sin6_addr), ip, INET6_ADDRSTRLEN);
            break;
        default:
            break;
    }

    goGetAddrInfoReply(
            interfaceIndex,
            (char *)hostName,
            ip,
            ttl,
            context
    );
}

DNSServiceErrorType GetAddrInfo
(
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    DNSServiceProtocol protocol,
    const char *hostName,
    void *context
)
{
    DNSServiceRef GetAddrInfoRef;
    DNSServiceErrorType err = DNSServiceGetAddrInfo(&GetAddrInfoRef, flags, interfaceIndex, protocol, hostName, GetAddrInfoReply, context);

    if ( err == kDNSServiceErr_NoError ) {
        DNSServiceProcessResult(GetAddrInfoRef);
        DNSServiceRefDeallocate(GetAddrInfoRef);
    }

    return err;
}
