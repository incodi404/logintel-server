package db

import (
	"auth-service/db/queries"
	"auth-service/utils"
	"context"
	"fmt"
)

func InitializeDBSeed(ctx context.Context) error {
	err := AdminSeed(ctx)
	return err
}

func AdminSeed(ctx context.Context) error {
	adminUser := utils.GetenvWithDefaultValue("ADMIN_USER", "")
	adminPass := utils.GetenvWithDefaultValue("ADMIN_PASSWORD", "")
	if adminUser == "" || adminPass == "" {
		return fmt.Errorf("[ERROR] Admin username or password not found")
	}

	// hashing password
	hashedPass, err := utils.HashStr(adminPass)
	if err != nil {
		return fmt.Errorf("[ERROR] Hashing admin password has been failed: %w", err)
	}

	// getting pool
	db, err := GetPGPool()
	if err != nil {
		return err
	}

	// checking if admin exists or not
	var isAdminExists bool
	err = db.QueryRow(ctx, queries.CHECK_IF_ADMIN_EXISTS, adminUser).Scan(&isAdminExists)
	if err != nil {
		return fmt.Errorf("[ERROR] Error occured while checking existed admin: %w", err)
	}

	if isAdminExists {
		fmt.Println("[ERROR] Admin already exists, skipping creation")
		return nil
	}

	// adding admin
	var adminId int
	err = db.QueryRow(ctx, queries.ADD_NEW_ADMIN, adminUser, hashedPass).Scan(&adminId)
	if err != nil {
		return fmt.Errorf("[ERROR] Error occured while creating admin: %w", err)
	}

	fmt.Println("[INFO] Admin has been created")
	fmt.Printf("[INFO] Username: %s\nAdmin ID: %d\n", adminUser, adminId)

	return nil
}
