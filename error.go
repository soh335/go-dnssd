package dnssd

/*
#include "dns_sd.h"
*/
import "C"

import (
        "errors"
)

func createErr(e C.DNSServiceErrorType) error {
        switch e {
        case DNSServiceErr_NoError:
                return nil
        case DNSServiceErr_Unknown:
                return errors.New("unknown")
        case DNSServiceErr_NoSuchName:
                return errors.New("no such name")
        case DNSServiceErr_NoMemory:
                return errors.New("no memory")
        case DNSServiceErr_BadParam:
                return errors.New("bad param")
        case DNSServiceErr_BadReference:
                return errors.New("bad reference")
        case DNSServiceErr_BadState:
                return errors.New("bad state")
        case DNSServiceErr_BadFlags:
                return errors.New("bad flags")
        case DNSServiceErr_Unsupported:
                return errors.New("unsupported")
        case DNSServiceErr_NotInitialized:
                return errors.New("not initialized")
        case DNSServiceErr_AlreadyRegistered:
                return errors.New("already registerd")
        case DNSServiceErr_NameConflict:
                return errors.New("name conflict")
        case DNSServiceErr_Invalid:
                return errors.New("invalid")
        case DNSServiceErr_Firewall:
                return errors.New("firewall")
        case DNSServiceErr_Incompatible:
                return errors.New("incompatible")
        case DNSServiceErr_BadInterfaceIndex:
                return errors.New("bad interface index")
        case DNSServiceErr_Refused:
                return errors.New("refused")
        case DNSServiceErr_NoSuchRecord:
                return errors.New("no such record")
        case DNSServiceErr_NoAuth:
                return errors.New("no auth")
        case DNSServiceErr_NoSuchKey:
                return errors.New("no such key")
        case DNSServiceErr_NATTraversal:
                return errors.New("NAT traversal")
        case DNSServiceErr_DoubleNAT:
                return errors.New("double NAT")
        case DNSServiceErr_BadTime:
                return errors.New("bad time")
        case DNSServiceErr_BadSig:
                return errors.New("bad sig")
        case DNSServiceErr_BadKey:
                return errors.New("bad key")
        case DNSServiceErr_Transient:
                return errors.New("transient")
        case DNSServiceErr_ServiceNotRunning:
                return errors.New("service not running")
        case DNSServiceErr_NATPortMappingUnsupported:
                return errors.New("NAT port mapping unsported")
        case DNSServiceErr_NATPortMappingDisabled:
                return errors.New("NAT port mapping disabled")
        case DNSServiceErr_NoRouter:
                return errors.New("no router")
        case DNSServiceErr_PollingMode:
                return errors.New("polling mode")
        case DNSServiceErr_Timeout:
                return errors.New("timeout")
        default:
                return nil
        }
}
