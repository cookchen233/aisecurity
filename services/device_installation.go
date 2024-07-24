package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/deviceinstallation"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
)

type DeviceInstallationService struct {
	Service
	entClient *dao.DeviceInstallationClient
}

func NewDeviceInstallationService(ctx context.Context) *DeviceInstallationService {
	s := &DeviceInstallationService{entClient: db.EntClient.DeviceInstallation}
	s.Ctx = ctx
	return s
}

func (s *DeviceInstallationService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.DeviceInstallation)
	c := s.entClient.Create().
		SetDeviceID(e.DeviceID).
		SetAliasName(e.AliasName).
		SetInstaller(e.Installer).
		SetInstallTime(e.InstallTime).
		SetLongitude(e.Longitude).
		SetLatitude(e.Latitude).
		SetLocation(e.Location).
		SetLocationData(e.LocationData)
	if e.AreaID > 0 {
		c.SetAreaID(e.AreaID)
	}
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating DeviceInstallationService")
	}
	// ConvertTo now returns a non-pointer type that implements IEntity
	return structs.ConvertTo[*dao.DeviceInstallation, entities.DeviceInstallation](saved), nil
}

func (s *DeviceInstallationService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.DeviceInstallation)
	u := s.entClient.UpdateOneID(e.ID).
		SetDeviceID(e.DeviceID).
		SetAliasName(e.AliasName).
		SetInstaller(e.Installer).
		SetInstallTime(e.InstallTime).
		SetLongitude(e.Longitude).
		SetLatitude(e.Latitude).
		SetLocation(e.Location).
		SetLocationData(e.LocationData)
	if e.AreaID > 0 {
		u.SetAreaID(e.AreaID)
	}
	save, err := u.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating DeviceInstallationService")
	}
	return structs.ConvertTo[*dao.DeviceInstallation, entities.DeviceInstallation](save), nil
}

func (s *DeviceInstallationService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *DeviceInstallationService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		WithArea().
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(deviceinstallation.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.DeviceInstallation)
		v2.AreaName = v.Edges.Area.Name
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

func (s *DeviceInstallationService) query(fit structs.IFilter) *dao.DeviceInstallationQuery {
	f := fit.(*filters.DeviceInstallation)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(deviceinstallation.IDEQ(f.ID))
	}
	return q.Clone()
}

func (s *DeviceInstallationService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *DeviceInstallationService) GetByDeviceID(deviceID int) (structs.IEntity, error) {
	first, err := s.entClient.Query().
		Where(deviceinstallation.DeviceID(deviceID)).
		WithArea().
		Order(deviceinstallation.ByID(sql.OrderDesc())).
		First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.DeviceInstallation, entities.DeviceInstallation](first), nil
}
