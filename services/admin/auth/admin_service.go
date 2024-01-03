package auth

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/adminrole"
	"aisecurity/services"
	"aisecurity/structs/service"
	"aisecurity/utils/db"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	services.Service
	entClient *dao.AdminClient
}

func NewAdminService() *AdminService {
	return &AdminService{
		entClient: db.EntClient.Admin,
	}
}

func (service *AdminService) Create(data service.Admin, role *dao.AdminRole) (*dao.Admin, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	saved, err := service.entClient.Create().
		AddAdminRoles(role).
		SetUsername(data.Username).
		SetName(data.Name).
		SetPassword(string(hashedPassword)).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Admin: %w", err)
	}
	return saved, nil
}

func (service *AdminService) Update(data service.Admin) (*dao.Admin, error) {
	q := service.entClient.UpdateOneID(data.ID)
	if data.Password != "" {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}
		q = q.SetPassword(string(hashedPassword))
	}
	saved, err := q.
		SetUsername(data.Username).
		SetName(data.Name).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating Admin: %w", err)
	}
	return saved, nil
}

func (service *AdminService) GetList() (dao.Admins, error) {
	all, err := service.entClient.Query().WithCreator().All(service.Ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (service *AdminService) GetSuperAdminList() (dao.Admins, error) {
	return db.EntClient.Admin.Query().Where(admin.HasAdminRolesWith(adminrole.IDEQ(1))).All(service.Ctx)
}

func (service *AdminService) GetByUserName(username string) (*dao.Admin, error) {
	one, err := service.entClient.Query().Where(admin.UsernameEQ(username)).Only(service.Ctx)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (service *AdminService) GetAdminRoles(data *dao.Admin) (dao.AdminRoles, error) {
	all, err := service.entClient.QueryAdminRoles(data).All(service.Ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (service *AdminService) RemoveAdminRoles(data *dao.Admin, adminRoles dao.AdminRoles) (*dao.Admin, error) {
	saved, err := service.entClient.UpdateOne(data).
		RemoveAdminRoles(adminRoles...).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed removing admin roles: %w", err)
	}
	return saved, nil
}

func (service *AdminService) AddAdminRole(data *dao.Admin, adminRoles []*dao.AdminRole) (*dao.Admin, error) {
	saved, err := service.entClient.UpdateOne(data).
		AddAdminRoles(adminRoles...).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed adding admin roles: %w", err)
	}
	return saved, nil
}
