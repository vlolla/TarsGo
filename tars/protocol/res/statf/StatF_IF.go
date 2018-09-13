//
// This file war generated by FastTars2go 1.0
// Generated from StatF.tars
// Tencent.

package statf

import (
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
)

type StatF struct {
	s m.Servant
}

func (_obj *StatF) ReportMicMsg(Msg map[StatMicMsgHead]StatMicMsgBody, BFromClient bool, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.WriteHead(codec.MAP, 1)
	if err != nil {
		return ret, err
	}
	err = _os.Write_int32(int32(len(Msg)), 0)
	if err != nil {
		return ret, err
	}
	for k0, v0 := range Msg {

		err = k0.WriteBlock(_os, 0)
		if err != nil {
			return ret, err
		}

		err = v0.WriteBlock(_os, 1)
		if err != nil {
			return ret, err
		}
	}

	err = _os.Write_bool(BFromClient, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "reportMicMsg", _os.ToBytes(), _status, _context, _resp)
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
func (_obj *StatF) ReportSampleMsg(Msg []StatSampleMsg, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.WriteHead(codec.LIST, 1)
	if err != nil {
		return ret, err
	}
	err = _os.Write_int32(int32(len(Msg)), 0)
	if err != nil {
		return ret, err
	}
	for _, v := range Msg {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return ret, err
		}
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(0, "reportSampleMsg", _os.ToBytes(), _status, _context, _resp)
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

func (_obj *StatF) SetServant(s m.Servant) {
	_obj.s = s
}

func (_obj *StatF) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

type _impStatF interface {
	ReportMicMsg(Msg map[StatMicMsgHead]StatMicMsgBody, BFromClient bool) (ret int32, err error)
	ReportSampleMsg(Msg []StatSampleMsg) (ret int32, err error)
}

func (_obj *StatF) Dispatch(_val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(req.SBuffer)
	_os := codec.NewBuffer()
	_imp := _val.(_impStatF)
	switch req.SFuncName {
	case "reportMicMsg":
		var Msg map[StatMicMsgHead]StatMicMsgBody
		err, have = _is.SkipTo(codec.MAP, 1, true)
		if err != nil {
			return err
		}

		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		Msg = make(map[StatMicMsgHead]StatMicMsgBody)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {
			var k1 StatMicMsgHead
			var v1 StatMicMsgBody

			err = k1.ReadBlock(_is, 0, false)
			if err != nil {
				return err
			}

			err = v1.ReadBlock(_is, 1, false)
			if err != nil {
				return err
			}

			Msg[k1] = v1
		}
		var BFromClient bool
		err = _is.Read_bool(&BFromClient, 2, true)
		if err != nil {
			return err
		}
		ret, err := _imp.ReportMicMsg(Msg, BFromClient)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	case "reportSampleMsg":
		var Msg []StatSampleMsg
		err, have, ty = _is.SkipToNoCheck(1, true)
		if err != nil {
			return err
		}

		if ty == codec.LIST {
			err = _is.Read_int32(&length, 0, true)
			if err != nil {
				return err
			}
			Msg = make([]StatSampleMsg, length, length)
			for i2, e2 := int32(0), length; i2 < e2; i2++ {

				err = Msg[i2].ReadBlock(_is, 0, false)
				if err != nil {
					return err
				}
			}
		} else if ty == codec.SIMPLE_LIST {
			err = fmt.Errorf("type not support SIMPLE_LIST.")
			if err != nil {
				return err
			}
		} else {
			err = fmt.Errorf("require vector, but not.")
			if err != nil {
				return err
			}
		}
		ret, err := _imp.ReportSampleMsg(Msg)
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
