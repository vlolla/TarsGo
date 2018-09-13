//
// This file war generated by FastTars2go 1.0
// Generated from NotifyF.tars
// Tencent.

package notifyf

import (
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
)

type Notify struct {
	s m.Servant
}

func (_obj *Notify) ReportServer(SServerName string, SThreadId string, SMessage string, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(SServerName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(SThreadId, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(SMessage, 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "reportServer", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func (_obj *Notify) NotifyServer(SServerName string, Level NOTIFYLEVEL, SMessage string, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(SServerName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(Level), 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(SMessage, 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "notifyServer", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func (_obj *Notify) GetNotifyInfo(StKey *NotifyKey, StInfo *NotifyInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = StKey.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "getNotifyInfo", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(_resp.SBuffer)
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*StInfo).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}
func (_obj *Notify) ReportNotifyInfo(Info *ReportInfo, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = Info.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "reportNotifyInfo", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

func (_obj *Notify) SetServant(s m.Servant) {
	_obj.s = s
}

func (_obj *Notify) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

type _impNotify interface {
	ReportServer(SServerName string, SThreadId string, SMessage string) (err error)
	NotifyServer(SServerName string, Level NOTIFYLEVEL, SMessage string) (err error)
	GetNotifyInfo(StKey *NotifyKey, StInfo *NotifyInfo) (ret int32, err error)
	ReportNotifyInfo(Info *ReportInfo) (err error)
}

func (_obj *Notify) Dispatch(_val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(req.SBuffer)
	_os := codec.NewBuffer()
	_imp := _val.(_impNotify)
	switch req.SFuncName {
	case "reportServer":
		var SServerName string
		err = _is.Read_string(&SServerName, 1, true)
		if err != nil {
			return err
		}
		var SThreadId string
		err = _is.Read_string(&SThreadId, 2, true)
		if err != nil {
			return err
		}
		var SMessage string
		err = _is.Read_string(&SMessage, 3, true)
		if err != nil {
			return err
		}
		err := _imp.ReportServer(SServerName, SThreadId, SMessage)
		if err != nil {
			return err
		}
	case "notifyServer":
		var SServerName string
		err = _is.Read_string(&SServerName, 1, true)
		if err != nil {
			return err
		}
		var Level NOTIFYLEVEL
		err = _is.Read_int32((*int32)(&Level), 2, true)
		if err != nil {
			return err
		}
		var SMessage string
		err = _is.Read_string(&SMessage, 3, true)
		if err != nil {
			return err
		}
		err := _imp.NotifyServer(SServerName, Level, SMessage)
		if err != nil {
			return err
		}
	case "getNotifyInfo":
		var StKey NotifyKey
		err = StKey.ReadBlock(_is, 1, true)
		if err != nil {
			return err
		}
		var StInfo NotifyInfo
		ret, err := _imp.GetNotifyInfo(&StKey, &StInfo)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}

		err = StInfo.WriteBlock(_os, 2)
		if err != nil {
			return err
		}
	case "reportNotifyInfo":
		var Info ReportInfo
		err = Info.ReadBlock(_is, 1, true)
		if err != nil {
			return err
		}
		err := _imp.ReportNotifyInfo(&Info)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var status map[string]string
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      _os.ToBytes(),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}
	_ = length
	_ = have
	_ = ty
	return nil
}
