package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/device"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
)

type Device struct {
	Service
	entClient *dao.DeviceClient
}

func NewDevice(ctx context.Context) *Device {
	s := &Device{entClient: db.EntClient.Device}
	s.Ctx = ctx
	return s
}

func (s *Device) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Device)
	c := s.entClient.
		Create().
		SetName(e.Name).
		SetBrand(e.Brand).
		SetDeviceType(e.DeviceType).
		SetModel(e.Model).
		SetSn(e.Sn)
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating Device")
	}
	// ConvertTo now returns a non-pointer type that implements IEntity
	return structs.ConvertTo[*dao.Device, entities.Device](saved), nil
}

func (s *Device) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Device)
	u := s.entClient.UpdateOneID(e.ID).
		SetName(e.Name).
		SetBrand(e.Brand).
		SetDeviceType(e.DeviceType).
		SetModel(e.Model).
		SetSn(e.Sn)
	save, err := u.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating Device")
	}
	return structs.ConvertTo[*dao.Device, entities.Device](save), nil
}

func (s *Device) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
	fit.SetPage(1)
	fit.SetLimit(1)
	list, err := s.GetList(fit)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	if len(list) == 0 {
		return nil, utils.ErrorWithStack(expects.NewDataNotFound())
	}
	return list[0], nil
}

func (s *Device) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.Device)
		ents[i] = v
		err := gconv.Struct(v, v2)
		if err != nil {
			utils.Logger.Warn("convert error", zap.Error(err))
		} else {
			ents[i] = v2
		}
	}
	return ents, nil
}

func (s *Device) query(fit structs.IFilter) *dao.DeviceQuery {
	f := fit.(*filters.Device)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(device.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(device.NameEQ(f.Name))
	}
	if f.SN != "" {
		q = q.Where(device.Sn(f.SN))
	}
	return q.Clone()
}

func (s *Device) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *Device) GetByName(name string) (structs.IEntity, error) {
	first, err := s.entClient.Query().Where(device.NameEQ(name)).First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.Device, entities.Device](first), nil
}

func (s *Device) GetBySn(name string) (structs.IEntity, error) {
	first, err := s.entClient.Query().Where(device.SnEQ(name)).First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.Device, entities.Device](first), nil
}

func (s *Device) GetOrCreateBySn(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Device)
	exDevice, err := s.GetBySn(e.Sn)
	if err == nil {
		return exDevice, nil
	}
	return s.Create(e)
}
