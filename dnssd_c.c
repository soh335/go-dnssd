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
    DNSServiceRef *ref,
    DNSServiceFlags flags,
    const char* regType,
    void *context
)
{
    DNSServiceErrorType err = DNSServiceBrowse(ref, 0, flags, regType, NULL, BrowseReply, context);
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
    DNSServiceRef *ref,
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    const char *serviceName,
    const char *regType,
    const char *replyDomain,
    void *context
)
{
    DNSServiceErrorType err = DNSServiceResolve(ref, flags, interfaceIndex, serviceName, regType, replyDomain, ResolveReply, context);
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

void RegisterReply
(
    DNSServiceRef        sdRef,
    DNSServiceFlags      flags,
    DNSServiceErrorType  errorCode,
    const char           *name,
    const char           *regType,
    const char           *domain,
    void                 *context
)
{
    goServiceRegisterReply(
        (char *)name,
        (char *)regType,
        (char *)domain,
        context
    );
}

DNSServiceErrorType QueryRecord
(
    DNSServiceRef *ref,
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    const char *fullName,
    uint16_t rrtype,
    uint16_t rrclass,
    void *context
)
{
    DNSServiceErrorType err = DNSServiceQueryRecord(ref, flags, interfaceIndex, fullName, rrtype, rrclass, QueryRecordReply, context);
    return err;
}

DNSServiceErrorType ServiceRegister
(
    DNSServiceRef    *ref,
    DNSServiceFlags  flags,
    uint32_t         interfaceIndex,
    const char       *name,       /* may be NULL */
    const char       *regType,
    const char       *domain,     /* may be NULL */
    const char       *host,       /* may be NULL */
    uint16_t         port,        /* In network byte order */
    uint16_t         txtLen,
    const void       *txtRecord,  /* may be NULL */
    void             *context     /* may be NULL */
)
{
    DNSServiceErrorType err = DNSServiceRegister(ref, flags, interfaceIndex, name, regType, domain, host, htons(port), txtLen, txtRecord, RegisterReply, context);
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
    DNSServiceRef *ref,
    DNSServiceFlags flags,
    uint32_t interfaceIndex,
    DNSServiceProtocol protocol,
    const char *hostName,
    void *context
)
{
    DNSServiceErrorType err = DNSServiceGetAddrInfo(ref, flags, interfaceIndex, protocol, hostName, GetAddrInfoReply, context);
    return err;
}
