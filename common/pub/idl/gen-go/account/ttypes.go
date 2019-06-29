// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package account

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/base"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = base.GoUnusedProtection__
var GoUnusedProtection__ int

// Attributes:
//  - Account
//  - Password
//  - Sign
type LogInReq struct {
	Account  string `thrift:"Account,1" json:"Account"`
	Password string `thrift:"Password,2" json:"Password"`
	Sign     string `thrift:"Sign,3" json:"Sign"`
}

func NewLogInReq() *LogInReq {
	return &LogInReq{}
}

func (p *LogInReq) GetAccount() string {
	return p.Account
}

func (p *LogInReq) GetPassword() string {
	return p.Password
}

func (p *LogInReq) GetSign() string {
	return p.Sign
}
func (p *LogInReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LogInReq) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Account = v
	}
	return nil
}

func (p *LogInReq) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Password = v
	}
	return nil
}

func (p *LogInReq) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Sign = v
	}
	return nil
}

func (p *LogInReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("LogInReq"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LogInReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Account", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Account: ", p), err)
	}
	if err := oprot.WriteString(string(p.Account)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Account (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Account: ", p), err)
	}
	return err
}

func (p *LogInReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Password", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Password: ", p), err)
	}
	if err := oprot.WriteString(string(p.Password)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Password (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Password: ", p), err)
	}
	return err
}

func (p *LogInReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Sign", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:Sign: ", p), err)
	}
	if err := oprot.WriteString(string(p.Sign)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Sign (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:Sign: ", p), err)
	}
	return err
}

func (p *LogInReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LogInReq(%+v)", *p)
}

// Attributes:
//  - UID
//  - Status
//  - LoginStatus
//  - ErrInfo
type LogInRes struct {
	UID         *int64 `thrift:"Uid,1" json:"Uid,omitempty"`
	Status      *int32 `thrift:"Status,2" json:"Status,omitempty"`
	LoginStatus *int32 `thrift:"LoginStatus,3" json:"LoginStatus,omitempty"`
	// unused fields # 4 to 99
	ErrInfo *base.ErrorInfo `thrift:"ErrInfo,100" json:"ErrInfo,omitempty"`
}

func NewLogInRes() *LogInRes {
	return &LogInRes{}
}

var LogInRes_UID_DEFAULT int64

func (p *LogInRes) GetUID() int64 {
	if !p.IsSetUID() {
		return LogInRes_UID_DEFAULT
	}
	return *p.UID
}

var LogInRes_Status_DEFAULT int32

func (p *LogInRes) GetStatus() int32 {
	if !p.IsSetStatus() {
		return LogInRes_Status_DEFAULT
	}
	return *p.Status
}

var LogInRes_LoginStatus_DEFAULT int32

func (p *LogInRes) GetLoginStatus() int32 {
	if !p.IsSetLoginStatus() {
		return LogInRes_LoginStatus_DEFAULT
	}
	return *p.LoginStatus
}

var LogInRes_ErrInfo_DEFAULT *base.ErrorInfo

func (p *LogInRes) GetErrInfo() *base.ErrorInfo {
	if !p.IsSetErrInfo() {
		return LogInRes_ErrInfo_DEFAULT
	}
	return p.ErrInfo
}
func (p *LogInRes) IsSetUID() bool {
	return p.UID != nil
}

func (p *LogInRes) IsSetStatus() bool {
	return p.Status != nil
}

func (p *LogInRes) IsSetLoginStatus() bool {
	return p.LoginStatus != nil
}

func (p *LogInRes) IsSetErrInfo() bool {
	return p.ErrInfo != nil
}

func (p *LogInRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 100:
			if err := p.readField100(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LogInRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.UID = &v
	}
	return nil
}

func (p *LogInRes) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Status = &v
	}
	return nil
}

func (p *LogInRes) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.LoginStatus = &v
	}
	return nil
}

func (p *LogInRes) readField100(iprot thrift.TProtocol) error {
	p.ErrInfo = &base.ErrorInfo{}
	if err := p.ErrInfo.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ErrInfo), err)
	}
	return nil
}

func (p *LogInRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("LogInRes"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField100(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LogInRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetUID() {
		if err := oprot.WriteFieldBegin("Uid", thrift.I64, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Uid: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.UID)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Uid (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Uid: ", p), err)
		}
	}
	return err
}

func (p *LogInRes) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetStatus() {
		if err := oprot.WriteFieldBegin("Status", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Status: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Status)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Status (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Status: ", p), err)
		}
	}
	return err
}

func (p *LogInRes) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetLoginStatus() {
		if err := oprot.WriteFieldBegin("LoginStatus", thrift.I32, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:LoginStatus: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.LoginStatus)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.LoginStatus (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:LoginStatus: ", p), err)
		}
	}
	return err
}

func (p *LogInRes) writeField100(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrInfo() {
		if err := oprot.WriteFieldBegin("ErrInfo", thrift.STRUCT, 100); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 100:ErrInfo: ", p), err)
		}
		if err := p.ErrInfo.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ErrInfo), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 100:ErrInfo: ", p), err)
		}
	}
	return err
}

func (p *LogInRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LogInRes(%+v)", *p)
}

// Attributes:
//  - UID
type LogOutReq struct {
	UID int64 `thrift:"Uid,1" json:"Uid"`
}

func NewLogOutReq() *LogOutReq {
	return &LogOutReq{}
}

func (p *LogOutReq) GetUID() int64 {
	return p.UID
}
func (p *LogOutReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LogOutReq) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.UID = v
	}
	return nil
}

func (p *LogOutReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("LogOutReq"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LogOutReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Uid", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Uid: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.UID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Uid (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Uid: ", p), err)
	}
	return err
}

func (p *LogOutReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LogOutReq(%+v)", *p)
}

// Attributes:
//  - LoginStatus
//  - ErrInfo
type LogOutRes struct {
	LoginStatus *int32 `thrift:"LoginStatus,1" json:"LoginStatus,omitempty"`
	// unused fields # 2 to 99
	ErrInfo *base.ErrorInfo `thrift:"ErrInfo,100" json:"ErrInfo,omitempty"`
}

func NewLogOutRes() *LogOutRes {
	return &LogOutRes{}
}

var LogOutRes_LoginStatus_DEFAULT int32

func (p *LogOutRes) GetLoginStatus() int32 {
	if !p.IsSetLoginStatus() {
		return LogOutRes_LoginStatus_DEFAULT
	}
	return *p.LoginStatus
}

var LogOutRes_ErrInfo_DEFAULT *base.ErrorInfo

func (p *LogOutRes) GetErrInfo() *base.ErrorInfo {
	if !p.IsSetErrInfo() {
		return LogOutRes_ErrInfo_DEFAULT
	}
	return p.ErrInfo
}
func (p *LogOutRes) IsSetLoginStatus() bool {
	return p.LoginStatus != nil
}

func (p *LogOutRes) IsSetErrInfo() bool {
	return p.ErrInfo != nil
}

func (p *LogOutRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 100:
			if err := p.readField100(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LogOutRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.LoginStatus = &v
	}
	return nil
}

func (p *LogOutRes) readField100(iprot thrift.TProtocol) error {
	p.ErrInfo = &base.ErrorInfo{}
	if err := p.ErrInfo.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ErrInfo), err)
	}
	return nil
}

func (p *LogOutRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("LogOutRes"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField100(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LogOutRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetLoginStatus() {
		if err := oprot.WriteFieldBegin("LoginStatus", thrift.I32, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:LoginStatus: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.LoginStatus)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.LoginStatus (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:LoginStatus: ", p), err)
		}
	}
	return err
}

func (p *LogOutRes) writeField100(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrInfo() {
		if err := oprot.WriteFieldBegin("ErrInfo", thrift.STRUCT, 100); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 100:ErrInfo: ", p), err)
		}
		if err := p.ErrInfo.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ErrInfo), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 100:ErrInfo: ", p), err)
		}
	}
	return err
}

func (p *LogOutRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LogOutRes(%+v)", *p)
}

// Attributes:
//  - Account
//  - Password
//  - Email
//  - Addr
type RegisterReq struct {
	Account  string `thrift:"Account,1" json:"Account"`
	Password string `thrift:"Password,2" json:"Password"`
	Email    string `thrift:"Email,3" json:"Email"`
	Addr     string `thrift:"Addr,4" json:"Addr"`
}

func NewRegisterReq() *RegisterReq {
	return &RegisterReq{}
}

func (p *RegisterReq) GetAccount() string {
	return p.Account
}

func (p *RegisterReq) GetPassword() string {
	return p.Password
}

func (p *RegisterReq) GetEmail() string {
	return p.Email
}

func (p *RegisterReq) GetAddr() string {
	return p.Addr
}
func (p *RegisterReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RegisterReq) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Account = v
	}
	return nil
}

func (p *RegisterReq) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Password = v
	}
	return nil
}

func (p *RegisterReq) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Email = v
	}
	return nil
}

func (p *RegisterReq) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Addr = v
	}
	return nil
}

func (p *RegisterReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RegisterReq"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RegisterReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Account", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Account: ", p), err)
	}
	if err := oprot.WriteString(string(p.Account)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Account (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Account: ", p), err)
	}
	return err
}

func (p *RegisterReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Password", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:Password: ", p), err)
	}
	if err := oprot.WriteString(string(p.Password)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Password (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:Password: ", p), err)
	}
	return err
}

func (p *RegisterReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Email", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:Email: ", p), err)
	}
	if err := oprot.WriteString(string(p.Email)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Email (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:Email: ", p), err)
	}
	return err
}

func (p *RegisterReq) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Addr", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:Addr: ", p), err)
	}
	if err := oprot.WriteString(string(p.Addr)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Addr (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:Addr: ", p), err)
	}
	return err
}

func (p *RegisterReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RegisterReq(%+v)", *p)
}

// Attributes:
//  - Status
//  - ErrInfo
type RegisterRes struct {
	Status *int32 `thrift:"Status,1" json:"Status,omitempty"`
	// unused fields # 2 to 99
	ErrInfo *base.ErrorInfo `thrift:"ErrInfo,100" json:"ErrInfo,omitempty"`
}

func NewRegisterRes() *RegisterRes {
	return &RegisterRes{}
}

var RegisterRes_Status_DEFAULT int32

func (p *RegisterRes) GetStatus() int32 {
	if !p.IsSetStatus() {
		return RegisterRes_Status_DEFAULT
	}
	return *p.Status
}

var RegisterRes_ErrInfo_DEFAULT *base.ErrorInfo

func (p *RegisterRes) GetErrInfo() *base.ErrorInfo {
	if !p.IsSetErrInfo() {
		return RegisterRes_ErrInfo_DEFAULT
	}
	return p.ErrInfo
}
func (p *RegisterRes) IsSetStatus() bool {
	return p.Status != nil
}

func (p *RegisterRes) IsSetErrInfo() bool {
	return p.ErrInfo != nil
}

func (p *RegisterRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 100:
			if err := p.readField100(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RegisterRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = &v
	}
	return nil
}

func (p *RegisterRes) readField100(iprot thrift.TProtocol) error {
	p.ErrInfo = &base.ErrorInfo{}
	if err := p.ErrInfo.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ErrInfo), err)
	}
	return nil
}

func (p *RegisterRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RegisterRes"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField100(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RegisterRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetStatus() {
		if err := oprot.WriteFieldBegin("Status", thrift.I32, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Status: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Status)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Status (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Status: ", p), err)
		}
	}
	return err
}

func (p *RegisterRes) writeField100(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrInfo() {
		if err := oprot.WriteFieldBegin("ErrInfo", thrift.STRUCT, 100); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 100:ErrInfo: ", p), err)
		}
		if err := p.ErrInfo.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ErrInfo), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 100:ErrInfo: ", p), err)
		}
	}
	return err
}

func (p *RegisterRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RegisterRes(%+v)", *p)
}

// Attributes:
//  - UID
//  - OriginalPw
//  - NewPw_
type ChangePwReq struct {
	UID        int64  `thrift:"Uid,1" json:"Uid"`
	OriginalPw string `thrift:"OriginalPw,2" json:"OriginalPw"`
	NewPw_     string `thrift:"NewPw,3" json:"NewPw"`
}

func NewChangePwReq() *ChangePwReq {
	return &ChangePwReq{}
}

func (p *ChangePwReq) GetUID() int64 {
	return p.UID
}

func (p *ChangePwReq) GetOriginalPw() string {
	return p.OriginalPw
}

func (p *ChangePwReq) GetNewPw_() string {
	return p.NewPw_
}
func (p *ChangePwReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ChangePwReq) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.UID = v
	}
	return nil
}

func (p *ChangePwReq) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.OriginalPw = v
	}
	return nil
}

func (p *ChangePwReq) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.NewPw_ = v
	}
	return nil
}

func (p *ChangePwReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ChangePwReq"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ChangePwReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Uid", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Uid: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.UID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Uid (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Uid: ", p), err)
	}
	return err
}

func (p *ChangePwReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("OriginalPw", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:OriginalPw: ", p), err)
	}
	if err := oprot.WriteString(string(p.OriginalPw)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.OriginalPw (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:OriginalPw: ", p), err)
	}
	return err
}

func (p *ChangePwReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("NewPw", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:NewPw: ", p), err)
	}
	if err := oprot.WriteString(string(p.NewPw_)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.NewPw (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:NewPw: ", p), err)
	}
	return err
}

func (p *ChangePwReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ChangePwReq(%+v)", *p)
}

// Attributes:
//  - Status
//  - ErrInfo
type ChangePwRes struct {
	Status *int32 `thrift:"Status,1" json:"Status,omitempty"`
	// unused fields # 2 to 99
	ErrInfo *base.ErrorInfo `thrift:"ErrInfo,100" json:"ErrInfo,omitempty"`
}

func NewChangePwRes() *ChangePwRes {
	return &ChangePwRes{}
}

var ChangePwRes_Status_DEFAULT int32

func (p *ChangePwRes) GetStatus() int32 {
	if !p.IsSetStatus() {
		return ChangePwRes_Status_DEFAULT
	}
	return *p.Status
}

var ChangePwRes_ErrInfo_DEFAULT *base.ErrorInfo

func (p *ChangePwRes) GetErrInfo() *base.ErrorInfo {
	if !p.IsSetErrInfo() {
		return ChangePwRes_ErrInfo_DEFAULT
	}
	return p.ErrInfo
}
func (p *ChangePwRes) IsSetStatus() bool {
	return p.Status != nil
}

func (p *ChangePwRes) IsSetErrInfo() bool {
	return p.ErrInfo != nil
}

func (p *ChangePwRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 100:
			if err := p.readField100(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ChangePwRes) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = &v
	}
	return nil
}

func (p *ChangePwRes) readField100(iprot thrift.TProtocol) error {
	p.ErrInfo = &base.ErrorInfo{}
	if err := p.ErrInfo.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ErrInfo), err)
	}
	return nil
}

func (p *ChangePwRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ChangePwRes"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField100(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ChangePwRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetStatus() {
		if err := oprot.WriteFieldBegin("Status", thrift.I32, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Status: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Status)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.Status (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Status: ", p), err)
		}
	}
	return err
}

func (p *ChangePwRes) writeField100(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrInfo() {
		if err := oprot.WriteFieldBegin("ErrInfo", thrift.STRUCT, 100); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 100:ErrInfo: ", p), err)
		}
		if err := p.ErrInfo.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ErrInfo), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 100:ErrInfo: ", p), err)
		}
	}
	return err
}

func (p *ChangePwRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ChangePwRes(%+v)", *p)
}
