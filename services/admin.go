package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/risklocation"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	Service
	entClient *dao.AdminClient
}

func NewAdminService() *AdminService {
	return &AdminService{
		entClient: db.EntClient.Admin,
	}
}

func (s *AdminService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Admin)
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	saved, err := s.entClient.Create().
		SetUsername(e.Username).
		SetNickname(e.Nickname).
		SetRealName(e.RealName).
		SetPassword(string(hashedPassword)).
		AddAdminRoleIDs(e.AdminRoleIDs...).
		Save(s.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Admin: %w", err)
	}
	return saved, nil
}

func (s *AdminService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Admin)
	q := s.entClient.UpdateOneID(e.ID)
	if e.Password != "" {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}
		q = q.SetPassword(string(hashedPassword))
	}
	saved, err := q.
		SetUsername(e.Username).
		SetNickname(e.Nickname).
		SetRealName(e.RealName).
		Save(s.Ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed updating Admin")
	}
	return saved, nil
}

func (s *AdminService) query(fit structs.IFilter) *dao.RiskLocationQuery {
	f := fit.(*filters.Admin)
	q := db.EntClient.RiskLocation.Query()
	if f.ID != 0 {
		q = q.Where(risklocation.IDEQ(f.ID))
	}
	if f.Nickname != "" {
		q = q.Where(risklocation.NameContainsFold(f.Nickname))
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
		Offset(fit.GetOffset()).
		Limit(fit.GetLimit()).
		All(s.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = v
	}
	return ents, nil
}

func (s *AdminService) GetByUserName(username string) (*dao.Admin, error) {
	one, err := s.entClient.Query().Where(admin.UsernameEQ(username)).Only(s.Ctx)
	if err != nil {
		return nil, err
	}
	return one, nil
}
