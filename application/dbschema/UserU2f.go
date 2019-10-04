// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
	
	"time"
)

type Slice_UserU2f []*UserU2f

func (s Slice_UserU2f) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_UserU2f) RangeRaw(fn func(m *UserU2f) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// UserU2f 两步验证
type UserU2f struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*UserU2f
	namer   func(string) string
	connID  int
	context echo.Context
	
	Id     	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid    	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Token  	string  	`db:"token" bson:"token" comment:"签名" json:"token" xml:"token"`
	Type   	string  	`db:"type" bson:"type" comment:"类型" json:"type" xml:"type"`
	Extra  	string  	`db:"extra" bson:"extra" comment:"扩展设置" json:"extra" xml:"extra"`
	Created	uint    	`db:"created" bson:"created" comment:"绑定时间" json:"created" xml:"created"`
}

func (this *UserU2f) Trans() *factory.Transaction {
	return this.trans
}

func (this *UserU2f) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *UserU2f) SetContext(ctx echo.Context) factory.Model {
	this.context = ctx
	return this
}

func (this *UserU2f) Context() echo.Context {
	return this.context
}

func (this *UserU2f) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *UserU2f) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *UserU2f) Objects() []*UserU2f {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *UserU2f) NewObjects() factory.Ranger {
	return &Slice_UserU2f{}
}

func (this *UserU2f) InitObjects() *[]*UserU2f {
	this.objects = []*UserU2f{}
	return &this.objects
}

func (this *UserU2f) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *UserU2f) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *UserU2f) Short_() string {
	return "user_u2f"
}

func (this *UserU2f) Struct_() string {
	return "UserU2f"
}

func (this *UserU2f) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *UserU2f) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *UserU2f) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *UserU2f) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *UserU2f) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *UserU2f) GroupBy(keyField string, inputRows ...[]*UserU2f) map[string][]*UserU2f {
	var rows []*UserU2f
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*UserU2f{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*UserU2f{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *UserU2f) KeyBy(keyField string, inputRows ...[]*UserU2f) map[string]*UserU2f {
	var rows []*UserU2f
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*UserU2f{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *UserU2f) AsKV(keyField string, valueField string, inputRows ...[]*UserU2f) map[string]interface{} {
	var rows []*UserU2f
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *UserU2f) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *UserU2f) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	err = DBI.Fire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", this, nil)
	}
	return
}

func (this *UserU2f) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	
	if err = DBI.Fire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", this, mw, args...)
}

func (this *UserU2f) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *UserU2f) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *UserU2f) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	m := *this
	m.FromMap(kvset)
	if err = DBI.Fire("updating", &m, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", &m, mw, args...)
}

func (this *UserU2f) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { 
		return DBI.Fire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
		return DBI.Fire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", this, mw, args...)
		} else {
			err = DBI.Fire("created", this, nil)
		}
	} 
	return 
}

func (this *UserU2f) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.Fire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", this, mw, args...)
}

func (this *UserU2f) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *UserU2f) Reset() *UserU2f {
	this.Id = 0
	this.Uid = 0
	this.Token = ``
	this.Type = ``
	this.Extra = ``
	this.Created = 0
	return this
}

func (this *UserU2f) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["Token"] = this.Token
	r["Type"] = this.Type
	r["Extra"] = this.Extra
	r["Created"] = this.Created
	return r
}

func (this *UserU2f) FromMap(rows map[string]interface{}) {
	for key, value := range rows {
		switch key {
			case "id": this.Id = param.AsUint64(value)
			case "uid": this.Uid = param.AsUint(value)
			case "token": this.Token = param.AsString(value)
			case "type": this.Type = param.AsString(value)
			case "extra": this.Extra = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
		}
	}
}

func (this *UserU2f) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint64(vv)
				case "Uid": this.Uid = param.AsUint(vv)
				case "Token": this.Token = param.AsString(vv)
				case "Type": this.Type = param.AsString(vv)
				case "Extra": this.Extra = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
			}
	}
}

func (this *UserU2f) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["uid"] = this.Uid
	r["token"] = this.Token
	r["type"] = this.Type
	r["extra"] = this.Extra
	r["created"] = this.Created
	return r
}

func (this *UserU2f) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *UserU2f) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

