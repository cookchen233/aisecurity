package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/department"
	employee2 "aisecurity/ent/dao/employee"
	"aisecurity/ent/dao/occupation"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type AdminService struct {
	Service
	entClient *dao.AdminClient
}

func NewAdminService(ctx context.Context) *AdminService {
	s := &AdminService{
		entClient: db.EntClient.Admin,
	}
	s.Ctx = ctx
	return s
}

func (s *AdminService) Create(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Admin)

	ex, err := s.GetByUserName(e.Username)
	if ex != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(expects.NewAlreadyExistedUsername()))
	}

	hashedPassword, err := utils.HashPassword(e.SetPassword)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}

	c := tx.Admin.Create().
		SetUsername(e.Username).
		SetNickname(e.Nickname).
		SetRealName(e.RealName).
		SetPassword(hashedPassword).
		SetMobile(e.Mobile).
		SetAdminStatus(enums.ENS1).
		SetAvatar(e.Avatar)

	if e.CreatorID > 0 {
		c.SetCreatorID(e.CreatorID)
	}

	if len(e.Permissions) > 0 {
		c.AddPermissions(e.Permissions...)
	}
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating Admin")
	}

	create := tx.Employee.Create().
		SetAdmin(saved)
	// set occupation
	if e.Occupation != nil {
		create.SetOccupation(e.Occupation)
	}
	_, err = create.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Employee"))
	}

	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	utils.Logger.Debug("hehe")

	return saved, nil
}

func (s *AdminService) Update(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Admin)

	ex, err := s.GetByUserName(e.Username)
	if err == nil && ex.(*entities.Admin).ID != e.ID {
		return nil, s.rollback(tx, utils.ErrorWithStack(expects.NewAlreadyExistedUsername()))
	}

	u := tx.Admin.UpdateOneID(e.ID)
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
		SetAvatar(e.Avatar).
		SetMobile(e.Mobile).
		SetAdminStatus(enums.ENS1).
		ClearPermissions()

	if len(e.Permissions) > 0 {
		u.AddPermissions(e.Permissions...)
	}

	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed updating Admin")
	}

	// create employee if not already
	employee, err := db.EntClient.Employee.Query().Where(employee2.AdminID(e.ID)).First(s.Ctx)
	if err != nil {
		if !dao.IsNotFound(err) {
			return nil, s.rollback(tx, err)
		} else {
			create := tx.Employee.Create().
				SetAdmin(saved)
			// set occupation
			if e.Occupation != nil {
				create.SetOccupation(e.Occupation)
			}
			_, err = create.Save(s.Ctx)
			if err != nil {
				return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Employee"))
			}
		}
	} else {
		update := tx.Employee.UpdateOne(employee)
		if e.Occupation != nil {
			update.SetOccupation(e.Occupation)
		}
		_, err = update.Save(s.Ctx)
		if err != nil {
			return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating Employee"))
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}

	return saved, nil
}

func (s *AdminService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *AdminService) GetByID(ID int) (structs.IEntity, error) {
	if ID == 0 {
		return nil, utils.ErrorWithStack(expects.NewInValidID())
	}
	var fit = &filters.Admin{
		StandardFilter: structs.StandardFilter{
			ID: ID,
		},
	}
	return s.GetDetails(fit)
}

func (s *AdminService) query(fit structs.IFilter) *dao.AdminQuery {
	f := fit.(*filters.Admin)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(admin.IDEQ(f.ID))
	}
	if f.Keywords != "" {
		q = q.Where(
			admin.Or(
				admin.NicknameContainsFold(f.Keywords),
				admin.UsernameContainsFold(f.Keywords),
				admin.RealNameContainsFold(f.Keywords),
			),
		)
	}
	if f.AdminStatus != 0 {
		q = q.Where(admin.AdminStatusEQ(f.AdminStatus))
	}
	if f.DepartmentID != 0 {
		q = q.Where(admin.HasEmployeeWith(
			employee2.HasDepartmentWith(
				department.IDEQ(f.DepartmentID),
			),
		))
	}
	if f.OccupationID != 0 {
		q = q.Where(admin.HasEmployeeWith(
			employee2.HasOccupationWith(
				occupation.IDEQ(f.OccupationID),
			),
		))
	}
	return q.Clone()
}

func (s *AdminService) GetTotal(fit structs.IFilter) (int, error) {
	// total
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *AdminService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).WithCreator().
		WithPermissions().
		WithEmployee(func(query *dao.EmployeeQuery) {
			query.WithDepartment().WithOccupation()
		}).
		Offset(fit.GetOffset()).
		Limit(fit.GetLimit()).
		Order(dao.Desc(admin.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	employeeService := NewEmployeeService(s.Ctx)
	for i, v := range list {
		v2 := structs.ConvertTo[*dao.Admin, entities.Admin](v)
		if v.Edges.Employee != nil && v.Edges.Employee.Edges.Department != nil {
			top, err := employeeService.GetTopDepartment(v.Edges.Employee.Edges.Department)
			if err != nil {
				return nil, utils.ErrorWithStack(err)
			}
			v2.Edges.Employee.Edges.TopDepartment = top
		}
		ents[i] = v2
	}
	return ents, nil
}

func (s *AdminService) CreateSuperAdmin() (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	_, err = tx.ExecContext(s.Ctx, `INSERT INTO "admins" (
                    "create_time",
                    "update_time",
                    "username",
                    "password",
                    "nickname",
                    "real_name",
                    "admin_status") VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		time.Now(), time.Now(), "root", hashedPassword, "超超", "超管", 1)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Super Admin"))
	}
	list, err := tx.Admin.Query().All(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed querying Super Admin"))
	}
	if len(list) > 1 {
		return nil, s.rollback(tx, utils.ErrorWithStack(fmt.Errorf("the count of Admin is not correct")))
	}
	if list[0].ID != 1 {
		return nil, s.rollback(tx, utils.ErrorWithStack(fmt.Errorf("the ID of Super Admin is not correct")))
	}
	if list[0].Username != "root" {
		return nil, s.rollback(tx, utils.ErrorWithStack(fmt.Errorf("the Username of Super Admin is not correct")))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return list[0], nil
}

func (s *AdminService) GetByUserName(username string) (structs.IEntity, error) {
	one, err := s.entClient.Query().Where(admin.UsernameEQ(username)).WithPermissions().Only(s.Ctx)
	if err != nil {
		return nil, err
	}
	o := structs.ConvertTo[*dao.Admin, entities.Admin](one)

	o.Password = one.Password
	return o, nil
}

//func (s *AdminService) GetWechatOpenid(admin structs.IEntity) (string, error) {
//	adm := admin.(*entities.Admin)
//	if adm.Mobile == "" {
//		return "", utils.ErrorWithStack(fmt.Errorf("the admin has not set mobile, username: %s", adm.Username))
//	}
//	userService := NewUserService(s.Ctx)
//	user, err := userService.GetByMobile(adm.Mobile)
//	if err != nil {
//		return "", utils.ErrorWithStack(err)
//	}
//	return user.(*entities.User).WechatOpenid, nil
//}

func (s *AdminService) GetByWechatOpenid(openid string) (structs.IEntity, error) {
	first, err := s.entClient.Query().Where(admin.WechatOpenidEQ(openid)).First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.Admin, entities.Admin](first), nil
}
