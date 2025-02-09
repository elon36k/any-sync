package list

import (
	"github.com/anyproto/any-sync/commonspace/object/acl/aclrecordproto"
	"github.com/anyproto/any-sync/util/crypto"
)

type AclRecord struct {
	Id                string
	PrevId            string
	Timestamp         int64
	AcceptorTimestamp int64
	Data              []byte
	Identity          crypto.PubKey
	Model             interface{}
	Signature         []byte
}

type RequestRecord struct {
	RequestIdentity crypto.PubKey
	RequestMetadata []byte
	Type            RequestType
}

type AclAccountState struct {
	PubKey          crypto.PubKey
	Permissions     AclPermissions
	RequestMetadata []byte
	KeyRecordId     string
}

type RequestType int

const (
	RequestTypeRemove RequestType = iota
	RequestTypeJoin
)

type AclPermissions aclrecordproto.AclUserPermissions

func (p AclPermissions) NoPermissions() bool {
	return aclrecordproto.AclUserPermissions(p) == aclrecordproto.AclUserPermissions_None
}

func (p AclPermissions) IsOwner() bool {
	return aclrecordproto.AclUserPermissions(p) == aclrecordproto.AclUserPermissions_Owner
}

func (p AclPermissions) CanWrite() bool {
	switch aclrecordproto.AclUserPermissions(p) {
	case aclrecordproto.AclUserPermissions_Admin:
		return true
	case aclrecordproto.AclUserPermissions_Writer:
		return true
	case aclrecordproto.AclUserPermissions_Owner:
		return true
	default:
		return false
	}
}

func (p AclPermissions) CanManageAccounts() bool {
	switch aclrecordproto.AclUserPermissions(p) {
	case aclrecordproto.AclUserPermissions_Admin:
		return true
	case aclrecordproto.AclUserPermissions_Owner:
		return true
	default:
		return false
	}
}
