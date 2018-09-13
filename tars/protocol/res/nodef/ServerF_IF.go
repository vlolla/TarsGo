//
// This file war generated by FastTars2go 1.0
// Generated from NodeF.tars
// Tencent.

package nodef

import (
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
)

type ServerF struct {
	s m.Servant
}

func (_obj *ServerF) KeepAlive(ServerInfo *ServerInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = ServerInfo.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "keepAlive", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(_resp.SBuffer)
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}
func (_obj *ServerF) ReportVersion(App string, ServerName string, Version string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(App, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(ServerName, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(Version, 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "reportVersion", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(_resp.SBuffer)
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}

func (_obj *ServerF) SetServant(s m.Servant) {
	_obj.s = s
}

func (_obj *ServerF) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

type _impServerF interface {
	KeepAlive(ServerInfo *ServerInfo) (ret int32, err error)
	ReportVersion(App string, ServerName string, Version string) (ret int32, err error)
}

func (_obj *ServerF) Dispatch(_val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(req.SBuffer)
	_os := codec.NewBuffer()
	_imp := _val.(_impServerF)
	switch req.SFuncName {
	case "keepAlive":
		var ServerInfo ServerInfo
		err = ServerInfo.ReadBlock(_is, 1, true)
		if err != nil {
			return err
		}
		ret, err := _imp.KeepAlive(&ServerInfo)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	case "reportVersion":
		var App string
		err = _is.Read_string(&App, 1, true)
		if err != nil {
			return err
		}
		var ServerName string
		err = _is.Read_string(&ServerName, 2, true)
		if err != nil {
			return err
		}
		var Version string
		err = _is.Read_string(&Version, 3, true)
		if err != nil {
			return err
		}
		ret, err := _imp.ReportVersion(App, ServerName, Version)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
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
