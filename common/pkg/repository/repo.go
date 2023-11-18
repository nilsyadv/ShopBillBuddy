package repository

import (
	"fmt"
	"net/http"
	"strconv"

	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Repository represents generic interface for interacting with DB
type Repository interface {
	Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) *wraperror.WrappedError
	GetFirst(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError
	GetForTenant(uow *UnitOfWork, out interface{}, id string, tenantID uuid.UUID, preloadAssociations []string) *wraperror.WrappedError
	GetAll(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError
	GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, queryProcessors []QueryProcessor) *wraperror.WrappedError
	GetAllUnscoped(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError
	GetAllUnscopedForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, queryProcessors []QueryProcessor) *wraperror.WrappedError
	GetCount(uow *UnitOfWork, out *int64, entity interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError
	GetCountForTenant(uow *UnitOfWork, out *int64, tenantID uuid.UUID, entity interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError

	Add(uow *UnitOfWork, out interface{}) *wraperror.WrappedError
	AddWithOmit(uow *UnitOfWork, out interface{}, omitFields []string) *wraperror.WrappedError
	Update(uow *UnitOfWork, out interface{}) *wraperror.WrappedError
	UpdateWithOmit(uow *UnitOfWork, out interface{}, omitFields []string) *wraperror.WrappedError
	Upsert(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError
	Delete(uow *UnitOfWork, out interface{}, where ...interface{}) *wraperror.WrappedError
	DeleteForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID) *wraperror.WrappedError
	DeletePermanent(uow *UnitOfWork, out interface{}, where ...interface{}) *wraperror.WrappedError

	AddAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) *wraperror.WrappedError
	RemoveAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) *wraperror.WrappedError
	ReplaceAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) *wraperror.WrappedError
}

// GormRepository implements Repository
type GormRepository struct {
}

// NewRepository returns a new repository object
func NewRepository() Repository {
	return &GormRepository{}
}

// QueryProcessor allows to modify the query before it is executed
type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, *wraperror.WrappedError)

// PreloadAssociations specified associations to be preloaded
func PreloadAssociations(preloadAssociations []string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, *wraperror.WrappedError) {
		for _, association := range preloadAssociations {
			db = db.Preload(association)
		}
		return db, nil
	}
}

// Paginate will restrict the output of query
func Paginate(limit int, offset int, count *int64) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, *wraperror.WrappedError) {
		if out != nil && count != nil {
			if err := db.Model(out).Count(count).Error; err != nil {
				return db, wraperror.NewDBErrorWrap(err, "encounter during record count", http.StatusInternalServerError)
			}
		}
		if limit != -1 {
			db = db.Limit(limit)
		}
		if offset > 0 {
			db = db.Offset(offset)
		}
		return db, nil
	}
}

// PaginateForWeb will take limit and offset parameters from URL and  will set X-Total-Count header in response
func PaginateForWeb(w http.ResponseWriter, r *http.Request) QueryProcessor {
	queryParams := r.URL.Query()
	limitParam := queryParams["limit"]
	offsetParam := queryParams["offset"]

	var err error
	limit := -1
	if len(limitParam) > 0 {
		limit, err = strconv.Atoi(limitParam[0])
		if err != nil {
			limit = -1
		}
	}
	offset := 0
	if len(offsetParam) > 0 {
		offset, err = strconv.Atoi(offsetParam[0])
		if err != nil {
			offset = 0
		}
	}

	return func(db *gorm.DB, out interface{}) (*gorm.DB, *wraperror.WrappedError) {

		if out != nil {
			var totalRecords int64
			if err := db.Model(out).Count(&totalRecords).Error; err != nil {
				return db, wraperror.NewDBErrorWrap(err, "encounter during record count", http.StatusInternalServerError)
			}
			w.Header().Add("Access-Control-Expose-Headers", "X-Total-Count")
			w.Header().Set("X-Total-Count", strconv.Itoa(int(totalRecords)))
		}

		if limit != -1 {
			db = db.Limit(limit)
		}
		if offset > 0 {
			db = db.Offset(offset)
		}

		return db, nil
	}
}

// Order will order the results
func Order(value interface{}, reorder bool) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, *wraperror.WrappedError) {
		db = db.Order(value)
		return db, nil
	}
}

// Filter will filter the results
func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, *wraperror.WrappedError) {
		db = db.Where(condition, args...)
		return db, nil
	}
}

// FilterWithOR will filter the results with an 'OR'
func FilterWithOR(columnName []string, condition []string, filterValues []interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, *wraperror.WrappedError) {
		if len(condition) != len(columnName) && len(condition) != len(filterValues) {
			return db, nil
		}
		if len(condition) == 1 {
			db = db.Where(fmt.Sprintf("%v %v ?", columnName[0], condition[0]), filterValues[0])
			return db, nil
		}
		str := ""
		for i := 0; i < len(columnName); i++ {
			if i == len(columnName)-1 {
				str = fmt.Sprintf("%v%v %v ?", str, columnName[i], condition[i])
			} else {
				str = fmt.Sprintf("%v%v %v ? OR ", str, columnName[i], condition[i])
			}
		}
		db = db.Where(str, filterValues...)
		return db, nil
	}
}

// GetFirst gets first record matching the given criteria
func (repository *GormRepository) GetFirst(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
			}
		}
	}
	if err := db.First(out).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during 1st record extraction", http.StatusInternalServerError)
	}
	return nil
}

// Get a record for specified entity with specific id
func (repository *GormRepository) Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) *wraperror.WrappedError {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	if err := db.First(out, "id = ?", id).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during 1st record extraction with id", http.StatusInternalServerError)
	}
	return nil
}

// GetForTenant a record for specified entity with specific id and for specified tenant
func (repository *GormRepository) GetForTenant(uow *UnitOfWork, out interface{}, id string, tenantID uuid.UUID, preloadAssociations []string) *wraperror.WrappedError {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	if err := db.First(out, "id = ? AND tenantid = ?", id, tenantID).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during 1st record extraction with id and tenant id", http.StatusInternalServerError)
	}
	return nil
}

// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
			}
		}
	}
	if err := db.Find(out).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during records extraction", http.StatusInternalServerError)
	}
	return nil
}

// GetAllForTenant returns all objects of specifeid tenantID
func (repository *GormRepository) GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, queryProcessors []QueryProcessor) *wraperror.WrappedError {
	queryProcessors = append([]QueryProcessor{Filter("tenantID = ?", tenantID)}, queryProcessors...)
	return repository.GetAll(uow, out, queryProcessors)
}

// GetAllUnscoped retrieves all the records (including deleted) for a specified entity and returns it
func (repository *GormRepository) GetAllUnscoped(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
			}
		}
	}
	if err := db.Unscoped().Find(out).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
	}
	return nil
}

// GetAllUnscopedForTenant returns all objects (including deleted) of specifeid tenantID
func (repository *GormRepository) GetAllUnscopedForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, queryProcessors []QueryProcessor) *wraperror.WrappedError {
	queryProcessors = append([]QueryProcessor{Filter("tenantID = ?", tenantID)}, queryProcessors...)
	return repository.GetAllUnscoped(uow, out, queryProcessors)
}

// GetCount gets count of the given entity type
func (repository *GormRepository) GetCount(uow *UnitOfWork, count *int64, entity interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
			}
		}
	}
	if err := db.Model(entity).Count(count).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during count extraction", http.StatusInternalServerError)
	}
	return nil
}

// GetCountForTenant gets count of the given entity type for specified tenant
func (repository *GormRepository) GetCountForTenant(uow *UnitOfWork, count *int64, tenantID uuid.UUID, entity interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError {

	db := uow.DB.Where("tenantID = ?", tenantID)

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
			}
		}
	}
	if err := db.Model(entity).Count(count).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
	}
	return nil
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) *wraperror.WrappedError {
	if err := uow.DB.Create(entity).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during add", http.StatusInternalServerError)
	}
	return nil
}

// AddWithOmit add specified Entity by omitting passed fields
func (repository *GormRepository) AddWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) *wraperror.WrappedError {
	if err := uow.DB.Omit(omitFields...).Create(entity).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during add with omit", http.StatusInternalServerError)
	}
	return nil
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) *wraperror.WrappedError {
	if err := uow.DB.Model(entity).Updates(entity).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during update", http.StatusInternalServerError)
	}
	return nil
}

// Update or insert if not found
func (repository *GormRepository) Upsert(uow *UnitOfWork, entity interface{}, queryProcessors []QueryProcessor) *wraperror.WrappedError {
	db := uow.DB
	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, entity)
			if err != nil {
				return wraperror.NewDBErrorWrap(err, "encounter during query processing", http.StatusInternalServerError)
			}
		}
	}
	result := db.Model(entity).Updates(entity)
	if result.Error != nil {
		return wraperror.NewDBErrorWrap(result.Error, "encounter during update", http.StatusInternalServerError)
	}

	if result.RowsAffected == 0 {
		if err := uow.DB.Create(entity).Error; err != nil {
			return wraperror.NewDBErrorWrap(err, "encounter during insertion", http.StatusInternalServerError)
		}
	}

	return nil
}

// UpdateWithOmit updates specified Entity by omitting passed fields
func (repository *GormRepository) UpdateWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) *wraperror.WrappedError {
	if err := uow.DB.Model(entity).Omit(omitFields...).Updates(entity).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during update with omit", http.StatusInternalServerError)
	}
	return nil
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}, where ...interface{}) *wraperror.WrappedError {
	if err := uow.DB.Delete(entity, where...).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during delete records", http.StatusInternalServerError)
	}
	return nil
}

// DeleteForTenant all recrod(s) of specified entity / entity type for given tenant
func (repository *GormRepository) DeleteForTenant(uow *UnitOfWork, entity interface{}, tenantID uuid.UUID) *wraperror.WrappedError {
	if err := uow.DB.Delete(entity, "tenantid = ?", tenantID).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during delete records for tenant", http.StatusInternalServerError)
	}
	return nil
}

// DeletePermanent deletes record permanently specified Entity
func (repository *GormRepository) DeletePermanent(uow *UnitOfWork, entity interface{}, where ...interface{}) *wraperror.WrappedError {
	if err := uow.DB.Unscoped().Delete(entity, where...).Error; err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during permanent delete", http.StatusInternalServerError)
	}
	return nil
}

// AddAssociations adds associations to the given out entity
func (repository *GormRepository) AddAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) *wraperror.WrappedError {
	if err := uow.DB.Model(out).Association(associationName).Append(associations...); err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during adding association", http.StatusInternalServerError)
	}
	return nil
}

// RemoveAssociations removes associations from the given out entity
func (repository *GormRepository) RemoveAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) *wraperror.WrappedError {
	if err := uow.DB.Model(out).Association(associationName).Delete(associations...); err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during remove association", http.StatusInternalServerError)
	}
	return nil
}

// ReplaceAssociations removes associations from the given out entity
func (repository *GormRepository) ReplaceAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) *wraperror.WrappedError {
	if err := uow.DB.Model(out).Association(associationName).Replace(associations...); err != nil {
		return wraperror.NewDBErrorWrap(err, "encounter during replace association", http.StatusInternalServerError)
	}
	return nil
}

// Contains checks if value present in array
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// ContainsKey checks if key present in map
func ContainsKey(keyValuePair map[string][]string, keyToCheck string) bool {
	if _, keyFound := keyValuePair[keyToCheck]; keyFound {
		return true
	}
	return false
}

// DoesColumnExistInTable returns bool if the column exist in table
func DoesColumnExistInTable(uow *UnitOfWork, tableName string, ColumnName string) bool {
	//tableName := uow.DB.NewScope(rules).TableName() // rules --> model, need to send from client controller
	return uow.DB.Migrator().HasColumn(tableName, ColumnName)
}
