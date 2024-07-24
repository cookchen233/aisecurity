package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/user"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"github.com/pkg/errors"
)

type UserService struct {
	Service
	entClient *dao.UserClient
}

func NewUserService(ctx context.Context) *UserService {
	s := &UserService{
		entClient: db.EntClient.User,
	}
	s.Ctx = ctx
	return s
}

func (s *UserService) Create(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.User)

	ex, err := s.GetByUserName(e.Username)
	if ex != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(expects.NewAlreadyExistedUsername()))
	}

	hashedPassword, err := utils.HashPassword(e.SetPassword)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}

	c := tx.User.Create().
		SetUsername(e.Username).
		SetNickname(e.Nickname).
		SetRealName(e.RealName).
		SetMobile(e.Mobile).
		SetPassword(hashedPassword).
		SetUserStatus(enums.ENS1).
		SetAvatar(e.Avatar)

	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating User")
	}

	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}

	return saved, nil
}

func (s *UserService) Update(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.User)

	ex, err := s.GetByUserName(e.Username)
	if err == nil && ex.(*entities.User).ID != e.ID {
		return nil, s.rollback(tx, utils.ErrorWithStack(expects.NewAlreadyExistedUsername()))
	}

	u := tx.User.UpdateOneID(e.ID)
	if e.SetPassword != "" {
		hashedPassword, err := utils.HashPassword(e.SetPassword)
		if err != nil {
			return nil, utils.ErrorWithStack(err)
		}
		u = u.SetPassword(hashedPassword)
	}
	u = u.SetUsername(e.Username).
		SetNickname(e.Nickname).
		SetRealName(e.RealName).
		SetMobile(e.Mobile).
		SetUserStatus(enums.ENS1).
		SetAvatar(e.Avatar)

	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed updating User")
	}

	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}

	return saved, nil
}

func (s *UserService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *UserService) GetByID(ID int) (structs.IEntity, error) {
	if ID == 0 {
		return nil, utils.ErrorWithStack(expects.NewInValidID())
	}
	var fit = &filters.User{
		StandardFilter: structs.StandardFilter{
			ID: ID,
		},
	}
	return s.GetDetails(fit)
}

func (s *UserService) query(fit structs.IFilter) *dao.UserQuery {
	f := fit.(*filters.User)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(user.IDEQ(f.ID))
	}
	if f.Keywords != "" {
		q = q.Where(
			user.Or(
				user.NicknameContainsFold(f.Keywords),
				user.UsernameContainsFold(f.Keywords),
				user.RealNameContainsFold(f.Keywords),
			),
		)
	}
	if f.UserStatus != 0 {
		q = q.Where(user.UserStatusEQ(f.UserStatus))
	}

	return q.Clone()
}

func (s *UserService) GetTotal(fit structs.IFilter) (int, error) {
	// total
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *UserService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		Offset(fit.GetOffset()).
		Limit(fit.GetLimit()).
		Order(dao.Desc(user.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := structs.ConvertTo[*dao.User, entities.User](v)
		ents[i] = v2
	}
	return ents, nil
}

func (s *UserService) GetByUserName(username string) (structs.IEntity, error) {
	one, err := s.entClient.Query().Where(user.UsernameEQ(username)).Only(s.Ctx)
	if err != nil {
		return nil, err
	}
	o := structs.ConvertTo[*dao.User, entities.User](one)

	return o, nil
}

func (s *UserService) GetByMobile(mobile string) (structs.IEntity, error) {
	one, err := s.entClient.Query().Where(user.MobileEQ(mobile)).Only(s.Ctx)
	if err != nil {
		return nil, err
	}
	o := structs.ConvertTo[*dao.User, entities.User](one)

	return o, nil
}

func (s *UserService) GetByWechatOpenid(openid string) (structs.IEntity, error) {
	first, err := s.entClient.Query().Where(user.WechatOpenidEQ(openid)).First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.User, entities.User](first), nil
}

func (s *UserService) CreateByWechatOpenidIfNotExists(openid string) (structs.IEntity, error) {
	ex, err := s.GetByWechatOpenid(openid)
	if err == nil {
		return ex, nil
	}

	create, err := s.Create(&entities.User{
		User: dao.User{
			Username:     string([]rune(openid)[:10]),
			WechatOpenid: openid,
		},
	})
	if err != nil {
		return nil, err
	}

	return create, nil
}
