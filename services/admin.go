package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/structs"
	"aisecurity/structs/entities"
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

func (service *AdminService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Admin)
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	saved, err := service.entClient.Create().
		SetUsername(e.Username).
		SetNickname(e.Nickname).
		SetRealName(e.RealName).
		SetPassword(string(hashedPassword)).
		AddAdminRoleIDs(e.AdminRoleIDs...).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Admin: %w", err)
	}
	return saved, nil
}

func (service *AdminService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Admin)
	q := service.entClient.UpdateOneID(e.ID)
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
		Save(service.Ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed updating Admin")
	}
	return saved, nil
}

func (service *AdminService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := service.entClient.Query().WithCreator().All(service.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = v
	}
	return ents, nil
}

func (service *AdminService) GetByUserName(username string) (*dao.Admin, error) {
	one, err := service.entClient.Query().Where(admin.UsernameEQ(username)).Only(service.Ctx)
	if err != nil {
		return nil, err
	}
	return one, nil
}
