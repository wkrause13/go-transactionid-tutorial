package repositories

import "go.uber.org/zap"

type GenericDataRepo interface {
	CloneWithTransID(transID int) GenericDataRepo
	ListUsers() []string
}

type DataRepo struct {
	logger *zap.Logger
	db     *MockDB
}

func NewDataRepo() DataRepo {
	db := connectToDb()
	logger, _ := zap.NewProduction()
	return DataRepo{logger: logger, db: &db}
}

func (d DataRepo) CloneWithTransID(transID int) GenericDataRepo {
	newLogger := d.logger.With(zap.Int("transID", transID))
	return DataRepo{logger: newLogger, db: d.db}
}

func (d DataRepo) ListUsers() []string {
	d.logger.Info("getting users")
	return d.db.GetUsers()
}

type MockDB struct{}

func (m MockDB) GetUsers() []string {
	return []string{"amy", "bob", "carl"}
}

func connectToDb() MockDB {
	return MockDB{}
}
