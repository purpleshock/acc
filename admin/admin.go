package admin

import (
	"github.com/purpleshock/acc/admin/repository"
)

type AdminService struct {
	adminRepository repository.AdminRepository
	encrypt         repository.Encrypt
}
