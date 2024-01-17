package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/video"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"time"
)

type Video struct {
	Service
	entClient *dao.VideoClient
}

func NewVideo(ctx context.Context) *Video {
	s := &Video{entClient: db.EntClient.Video}
	s.Ctx = ctx
	return s
}

func (s *Video) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Video)
	c := s.entClient.
		Create().
		SetName(e.Name).
		SetURL(e.URL).
		SetDuration(e.Duration).
		SetSize(e.Size)
	if e.UploadedAt != nil {
		c = c.SetUploadedAt(*e.UploadedAt)
	}
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating Video")
	}
	// ConvertTo now returns a non-pointer type that implements IEntity
	return structs.ConvertTo[*dao.Video, entities.Video](saved), nil
}

func (s *Video) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Video)
	u := s.entClient.UpdateOneID(e.ID).
		SetName(e.Name).
		SetURL(e.URL).
		SetDuration(e.Duration).
		SetSize(e.Size)
	if e.UploadedAt != nil {
		u = u.SetUploadedAt(*e.UploadedAt)
	}
	save, err := u.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating Video")
	}
	return structs.ConvertTo[*dao.Video, entities.Video](save), nil
}

func (s *Video) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *Video) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
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
		v2 := new(entities.Video)
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

func (s *Video) query(fit structs.IFilter) *dao.VideoQuery {
	f := fit.(*filters.Video)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(video.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(video.NameEQ(f.Name))
	}
	if f.URL != "" {
		q = q.Where(video.URLEQ(f.URL))
	}
	return q.Clone()
}

func (s *Video) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *Video) GetByName(name string) (structs.IEntity, error) {
	first, err := s.entClient.Query().Where(video.NameEQ(name)).First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.Video, entities.Video](first), nil
}

func (s *Video) CreateOrUpdateByName(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Video)
	exVideo, err := s.GetByName(e.Name)
	var now = time.Now()
	if err == nil {
		ex := exVideo.(*entities.Video)
		if ex.Size > 0 && ex.URL != "" {
			return ex, nil
		}
		if e.URL != "" {
			ex.URL = e.URL
		}
		if e.Size > 0 {
			ex.Size = e.Size
		}
		if ex.UploadedAt == nil {
			if e.Size > 0 || e.URL != "" {
				ex.UploadedAt = &now
			}
		}
		saved, err := s.Update(ex)
		if err != nil {
			return nil, utils.ErrorWithStack(err)
		}
		return saved, nil
	}
	e = &entities.Video{
		Video: dao.Video{
			Name: e.Name,
			URL:  e.URL,
			Size: e.Size,
		},
	}
	if e.Size > 0 || e.URL != "" {
		e.UploadedAt = &now
	}
	saved, err := s.Create(e)
	return saved, nil
}
