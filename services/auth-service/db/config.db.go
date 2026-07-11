package db

import (
	"auth-service/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dbConfig struct {
	Db       string
	Host     string
	Port     int
	Username string
	Password string
	Max_Conn int
	Min_Conn int
}

var PGPool *pgxpool.Pool

func GetPGPool() (*pgxpool.Pool, error) {
	if PGPool == nil {
		return nil, fmt.Errorf("[DB ERROR] DB is not connected yet")
	}

	return PGPool, nil
}

func setConfigFromEnv() dbConfig {
	db := utils.GetenvWithDefaultValue("PG_DB", "central_db")
	port := utils.GetenvWithDefaultValue("PG_PORT", "5432")
	host := utils.GetenvWithDefaultValue("PG_HOST", "central-db")
	username := utils.GetenvWithDefaultValue("PG_USERNAME", "postgres")
	password := utils.GetenvWithDefaultValue("PG_PASSWORD", "postgres")
	maxConn := utils.GetenvWithDefaultValue("PG_MAX_CONN", "5")
	minConn := utils.GetenvWithDefaultValue("PG_MIN_CONN", "1")

	return dbConfig{
		Db:       db,
		Host:     host,
		Port:     utils.Conversion.StrToInt(port),
		Username: username,
		Password: password,
		Max_Conn: utils.Conversion.StrToInt(maxConn),
		Min_Conn: utils.Conversion.StrToInt(minConn),
	}
}

func DbConnect(ctx context.Context) (*pgxpool.Pool, error) {

	cfg := setConfigFromEnv()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Db,
	)

	pgConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("[ERROR DB] Error parsing DSN: %w", err)
	}

	pgConfig.MaxConns = int32(cfg.Max_Conn)
	pgConfig.MinConns = int32(cfg.Min_Conn)

	pool, err := pgxpool.NewWithConfig(ctx, pgConfig)
	if err != nil {
		return nil, fmt.Errorf("[ERROR DB] Failed to connect with db: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	PGPool = pool
	return pool, nil
}
