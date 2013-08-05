package dnssd

/*
#include "dns_sd.h"
*/
import "C"

const (
	// interface
	DNSServiceInterfaceIndexAny       = C.kDNSServiceInterfaceIndexAny
	DNSServiceInterfaceIndexLocalOnly = C.kDNSServiceInterfaceIndexLocalOnly
	DNSServiceInterfaceIndexUnicast   = C.kDNSServiceInterfaceIndexUnicast
	DNSServiceInterfaceIndexP2P       = C.kDNSServiceInterfaceIndexP2P

	// flags
	DNSServiceFlagsMoreComing             = C.kDNSServiceFlagsMoreComing
	DNSServiceFlagsAdd                    = C.kDNSServiceFlagsAdd
	DNSServiceFlagsDefault                = C.kDNSServiceFlagsAdd
	DNSServiceFlagsShared                 = C.kDNSServiceFlagsShared
	DNSServiceFlagsUnique                 = C.kDNSServiceFlagsUnique
	DNSServiceFlagsBrowseDomains          = C.kDNSServiceFlagsBrowseDomains
	DNSServiceFlagsRegistrationDomains    = C.kDNSServiceFlagsRegistrationDomains
	DNSServiceFlagsLongLivedQuery         = C.kDNSServiceFlagsLongLivedQuery
	DNSServiceFlagsAllowRemoteQuery       = C.kDNSServiceFlagsAllowRemoteQuery
	DNSServiceFlagsForceMulticast         = C.kDNSServiceFlagsForceMulticast
	DNSServiceFlagsForce                  = C.kDNSServiceFlagsForce
	DNSServiceFlagsReturnIntermediates    = C.kDNSServiceFlagsReturnIntermediates
	DNSServiceFlagsNonBrowsable           = C.kDNSServiceFlagsNonBrowsable
	DNSServiceFlagsShareConnection        = C.kDNSServiceFlagsShareConnection
	DNSServiceFlagsSuppressUnusable       = C.kDNSServiceFlagsSuppressUnusable
	DNSServiceFlagsTimeout                = C.kDNSServiceFlagsTimeout
	DNSServiceFlagsIncludeP2P             = C.kDNSServiceFlagsIncludeP2P
	DNSServiceFlagsWakeOnResolve          = C.kDNSServiceFlagsWakeOnResolve
	DNSServiceFlagsBackgroundTrafficClass = C.kDNSServiceFlagsBackgroundTrafficClass
	DNSServiceFlagsIncludeAWDL            = C.kDNSServiceFlagsIncludeAWDL

	// protocol
	DNSServiceProtocol_IPv4 = C.kDNSServiceProtocol_IPv4
	DNSServiceProtocol_IPv6 = C.kDNSServiceProtocol_IPv6

	DNSServiceProtocol_UDP = C.kDNSServiceProtocol_UDP
	DNSServiceProtocol_TCP = C.kDNSServiceProtocol_TCP

	// class
	DNSServiceClass_IN = C.kDNSServiceClass_IN

	// type

	DNSServiceType_A          = C.kDNSServiceType_A
	DNSServiceType_NS         = C.kDNSServiceType_NS
	DNSServiceType_MD         = C.kDNSServiceType_MD
	DNSServiceType_MF         = C.kDNSServiceType_MF
	DNSServiceType_CNAME      = C.kDNSServiceType_CNAME
	DNSServiceType_SOA        = C.kDNSServiceType_SOA
	DNSServiceType_MB         = C.kDNSServiceType_MB
	DNSServiceType_MG         = C.kDNSServiceType_MG
	DNSServiceType_MR         = C.kDNSServiceType_MR
	DNSServiceType_NULL       = C.kDNSServiceType_NULL
	DNSServiceType_WKS        = C.kDNSServiceType_WKS
	DNSServiceType_PTR        = C.kDNSServiceType_PTR
	DNSServiceType_HINFO      = C.kDNSServiceType_HINFO
	DNSServiceType_MINFO      = C.kDNSServiceType_MINFO
	DNSServiceType_MX         = C.kDNSServiceType_MX
	DNSServiceType_TXT        = C.kDNSServiceType_TXT
	DNSServiceType_RP         = C.kDNSServiceType_RP
	DNSServiceType_AFSDB      = C.kDNSServiceType_AFSDB
	DNSServiceType_X25        = C.kDNSServiceType_X25
	DNSServiceType_ISDN       = C.kDNSServiceType_ISDN
	DNSServiceType_RT         = C.kDNSServiceType_RT
	DNSServiceType_NSAP       = C.kDNSServiceType_NSAP
	DNSServiceType_NSAP_PTR   = C.kDNSServiceType_NSAP_PTR
	DNSServiceType_SIG        = C.kDNSServiceType_SIG
	DNSServiceType_KEY        = C.kDNSServiceType_KEY
	DNSServiceType_PX         = C.kDNSServiceType_PX
	DNSServiceType_GPOS       = C.kDNSServiceType_GPOS
	DNSServiceType_AAAA       = C.kDNSServiceType_AAAA
	DNSServiceType_LOC        = C.kDNSServiceType_LOC
	DNSServiceType_NXT        = C.kDNSServiceType_NXT
	DNSServiceType_EID        = C.kDNSServiceType_EID
	DNSServiceType_NIMLOC     = C.kDNSServiceType_NIMLOC
	DNSServiceType_SRV        = C.kDNSServiceType_SRV
	DNSServiceType_ATMA       = C.kDNSServiceType_ATMA
	DNSServiceType_NAPTR      = C.kDNSServiceType_NAPTR
	DNSServiceType_KX         = C.kDNSServiceType_KX
	DNSServiceType_CERT       = C.kDNSServiceType_CERT
	DNSServiceType_A6         = C.kDNSServiceType_A6
	DNSServiceType_DNAME      = C.kDNSServiceType_DNAME
	DNSServiceType_SINK       = C.kDNSServiceType_SINK
	DNSServiceType_OPT        = C.kDNSServiceType_OPT
	DNSServiceType_APL        = C.kDNSServiceType_APL
	DNSServiceType_DS         = C.kDNSServiceType_DS
	DNSServiceType_SSHFP      = C.kDNSServiceType_SSHFP
	DNSServiceType_IPSECKEY   = C.kDNSServiceType_IPSECKEY
	DNSServiceType_RRSIG      = C.kDNSServiceType_RRSIG
	DNSServiceType_NSEC       = C.kDNSServiceType_NSEC
	DNSServiceType_DNSKEY     = C.kDNSServiceType_DNSKEY
	DNSServiceType_DHCID      = C.kDNSServiceType_DHCID
	DNSServiceType_NSEC3      = C.kDNSServiceType_NSEC3
	DNSServiceType_NSEC3PARAM = C.kDNSServiceType_NSEC3PARAM

	DNSServiceType_HIP = C.kDNSServiceType_HIP

	DNSServiceType_SPF    = C.kDNSServiceType_SPF
	DNSServiceType_UINFO  = C.kDNSServiceType_UINFO
	DNSServiceType_UID    = C.kDNSServiceType_UID
	DNSServiceType_GID    = C.kDNSServiceType_GID
	DNSServiceType_UNSPEC = C.kDNSServiceType_UNSPEC

	DNSServiceType_TKEY  = C.kDNSServiceType_TKEY
	DNSServiceType_TSIG  = C.kDNSServiceType_TSIG
	DNSServiceType_IXFR  = C.kDNSServiceType_IXFR
	DNSServiceType_AXFR  = C.kDNSServiceType_AXFR
	DNSServiceType_MAILB = C.kDNSServiceType_MAILB
	DNSServiceType_MAILA = C.kDNSServiceType_MAILA
	DNSServiceType_ANY   = C.kDNSServiceType_ANY

	// error
	DNSServiceErr_NoError                   = C.kDNSServiceErr_NoError
	DNSServiceErr_Unknown                   = C.kDNSServiceErr_Unknown
	DNSServiceErr_NoSuchName                = C.kDNSServiceErr_NoSuchName
	DNSServiceErr_NoMemory                  = C.kDNSServiceErr_NoMemory
	DNSServiceErr_BadParam                  = C.kDNSServiceErr_BadParam
	DNSServiceErr_BadReference              = C.kDNSServiceErr_BadReference
	DNSServiceErr_BadState                  = C.kDNSServiceErr_BadState
	DNSServiceErr_BadFlags                  = C.kDNSServiceErr_BadFlags
	DNSServiceErr_Unsupported               = C.kDNSServiceErr_Unsupported
	DNSServiceErr_NotInitialized            = C.kDNSServiceErr_NotInitialized
	DNSServiceErr_AlreadyRegistered         = C.kDNSServiceErr_AlreadyRegistered
	DNSServiceErr_NameConflict              = C.kDNSServiceErr_NameConflict
	DNSServiceErr_Invalid                   = C.kDNSServiceErr_Invalid
	DNSServiceErr_Firewall                  = C.kDNSServiceErr_Firewall
	DNSServiceErr_Incompatible              = C.kDNSServiceErr_Incompatible
	DNSServiceErr_BadInterfaceIndex         = C.kDNSServiceErr_BadInterfaceIndex
	DNSServiceErr_Refused                   = C.kDNSServiceErr_Refused
	DNSServiceErr_NoSuchRecord              = C.kDNSServiceErr_NoSuchRecord
	DNSServiceErr_NoAuth                    = C.kDNSServiceErr_NoAuth
	DNSServiceErr_NoSuchKey                 = C.kDNSServiceErr_NoSuchKey
	DNSServiceErr_NATTraversal              = C.kDNSServiceErr_NATTraversal
	DNSServiceErr_DoubleNAT                 = C.kDNSServiceErr_DoubleNAT
	DNSServiceErr_BadTime                   = C.kDNSServiceErr_BadTime
	DNSServiceErr_BadSig                    = C.kDNSServiceErr_BadSig
	DNSServiceErr_BadKey                    = C.kDNSServiceErr_BadKey
	DNSServiceErr_Transient                 = C.kDNSServiceErr_Transient
	DNSServiceErr_ServiceNotRunning         = C.kDNSServiceErr_ServiceNotRunning
	DNSServiceErr_NATPortMappingUnsupported = C.kDNSServiceErr_NATPortMappingUnsupported
	DNSServiceErr_NATPortMappingDisabled    = C.kDNSServiceErr_NATPortMappingDisabled
	DNSServiceErr_NoRouter                  = C.kDNSServiceErr_NoRouter
	DNSServiceErr_PollingMode               = C.kDNSServiceErr_PollingMode
	DNSServiceErr_Timeout                   = C.kDNSServiceErr_Timeout
)
