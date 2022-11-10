package postgres

import (
	"fmt"
	pb "monorepo/src/idl/auth_service"
	"time"

	"monorepo/src/libs/etc"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type authRepo struct {
	db *sqlx.DB
}

type staff struct {
	Id             uuid.UUID
	BranchId       uuid.UUID
	Role           string
	HashedPassword string
}

func New(db *sqlx.DB) *authRepo {
	return &authRepo{db: db}
}

// check staff by username or phonenumber if exists return role and id
func (r *authRepo) StaffLogin(req pb.StaffLoginRequest) (pb.AuthResponse, error) {

	var s staff
	//Select a staff if it is in db with active status
	err := r.db.QueryRow(`
		SELECT id, branch_id,  password, role FROM staff_auth WHERE is_active=true AND ( phone_number = $1 OR username = $2) `,
		req.PhoneNumber, req.Username).Scan(
		&s.Id,
		&s.BranchId,
		&s.HashedPassword,
		&s.Role)

	if err != nil {
		return pb.AuthResponse{}, err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(s.HashedPassword), []byte(req.Password))
	if err != nil {
		return pb.AuthResponse{}, err
	}

	return pb.AuthResponse{
		Id:       s.Id.String(),
		BranchId: s.BranchId.String(),
		Role:     s.Role,
	}, nil
}

// Sign up staff with incoming username and default password
func (r *authRepo) StaffSignUp(req pb.StaffSignUpRequest) (pb.AuthResponse, error) {

	if req.PhoneNumber != "" {
		s, err := r.signUpWithPhoneNumber(req)
		if err != nil {
			return pb.AuthResponse{}, err
		}
		return pb.AuthResponse{
			Id:       s.Id.String(),
			BranchId: s.BranchId.String(),
			Role:     s.Role,
		}, nil

	}
	s, err := r.signUpWithUsername(req)
	if err != nil {
		return pb.AuthResponse{}, err
	}
	return pb.AuthResponse{
		Id:       s.Id.String(),
		BranchId: s.BranchId.String(),
		Role:     s.Role,
	}, nil

}

func (r *authRepo) signUpWithPhoneNumber(req pb.StaffSignUpRequest) (*staff, error) {
	var s staff
	exists := r.check(req.PhoneNumber, "")
	if exists {
		return &staff{}, fmt.Errorf("user with the phone number %s is already exist", req.PhoneNumber)
	}

	//Default password with name + the last 4 digits of phoneNumber
	defaultPassHash, err := etc.GeneratePasswordHash(req.Name + req.PhoneNumber[len(req.PhoneNumber)-4:])
	if err != nil {
		return &staff{}, fmt.Errorf("password hash error: %w", bcrypt.ErrHashTooShort)
	}

	err = r.db.QueryRow(`
		INSERT INTO staff_auth(id, phone_number, name, role, password, branch_id, create_date, update_date)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, branch_id, role`,
		uuid.New(),
		req.PhoneNumber,
		req.Name,
		req.Role,
		defaultPassHash,
		req.BranchId,
		time.Now(),
		time.Now()).Scan(
		&s.Id,
		&s.BranchId,
		&s.Role)
	if err != nil {
		return &staff{}, err
	}

	return &s, nil
}

func (r *authRepo) signUpWithUsername(req pb.StaffSignUpRequest) (*staff, error) {
	var s staff
	exists := r.check("", req.Username)
	if exists {
		return &staff{}, fmt.Errorf("user with the username: %s is already exist", req.Username)
	}

	//Default password with name + the last 4 digits of username
	defaultPassHash, err := etc.GeneratePasswordHash(req.Name + req.Username[len(req.PhoneNumber)-4:])
	if err != nil {
		return &staff{}, fmt.Errorf("error when generating password hash")
	}

	err = r.db.QueryRow(`
		INSERT INTO staff_auth(id, phone_number, name, role, password, branch_id, create_date, update_date)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, branch_id, role`,
		uuid.New(),
		req.PhoneNumber,
		req.Name,
		req.Role,
		defaultPassHash,
		req.BranchId,
		time.Now(),
		time.Now()).Scan(
		&s.Id,
		&s.BranchId,
		&s.Role,
	)
	if err != nil {
		return &staff{}, err
	}

	return &s, nil
}

// checks if active user exists or not in db by username or password
func (r *authRepo) check(phoneNumber string, username string) bool {

	var exists bool
	if phoneNumber != "" {
		if err := r.db.QueryRow(`
				SELECT EXISTS(SELECT 1 FROM staff_auth WHERE phone_number = $1 AND is_active = true)
			`, phoneNumber).Scan(&exists); err != nil {
			return false
		}
	} else if username != "" {
		if err := r.db.QueryRow(`
	SELECT EXISTS(SELECT 1 FROM staff_auth WHERE username = $1 AND is_active = true)
	`, username).Scan(&exists); err != nil {
			return false
		}
	}

	return true
}
