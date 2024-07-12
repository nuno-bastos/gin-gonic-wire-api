package repo

import (
	"context"

	security_model "github.com/nuno-bastos/gin-gonic-wire-api/model/security"
	interfaces "github.com/nuno-bastos/gin-gonic-wire-api/repo/interface"

	"gorm.io/gorm"
)

type securityMatrixRuleRepository struct {
	DB *gorm.DB
}

func NewSecurityMatrixRuleRepository(DB *gorm.DB) (interfaces.SecurityMatrixRuleRepository, error) {
	repo := &securityMatrixRuleRepository{DB: DB}

	// Perform the migration
	if err := repo.Migrate(); err != nil {
		return nil, err
	}

	return repo, nil
}

// Migrate uses GORM's AutoMigrate to handle the table creation for GoMatrixRule.
//
// Returns:
// - error: an error object if the migration fails, nil otherwise.
func (p *securityMatrixRuleRepository) Migrate() error {
	return p.DB.AutoMigrate(&security_model.GoMatrixRule{})
}

// WriteRules writes a batch of GoMatrixRule entries to the database. It deletes
// existing records before inserting new ones.
//
// Parameters:
// - ctx: context for managing request lifecycle.
// - rules: a slice of GoMatrixRule to be written to the database.
//
// Returns:
// - error: an error object if the operation fails, nil otherwise.
func (p *securityMatrixRuleRepository) WriteRules(ctx context.Context, rules []security_model.GoMatrixRule) error {
	tx := p.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Delete existing records
	if err := tx.Exec("DELETE FROM dbo.GoMatrixRule").Error; err != nil {
		tx.Rollback()
		return err
	}

	batchSize := 156 /*
	* The batchSize variable determines the number of rule entries in each transaction batch creation iteration.
	* Performance varies with different values (e.g., sample values for my machine):
	*
	* - Low values (e.g., 1-50) result in slower processing speed
	*
	* - Moderate values (e.g., 100-200) achieve optimal performance, processing
	*   at approximately twice the speed or half the time compared to lower values.
	*
	* - Higher values (e.g., 400+, max 700 due to local SQL server limit) result
	*   in longer processing times, around the same as higher low values.
	*
	*
	*   Adjust according to performance :)
	*	In this case, 156 was determined the optimal batch size from a Bayesian Optimization analysis.
	*
	 */

	// Process the rules in batches to optimize performance and reduce memory usage.
	// The batch size is determined by the batchSize variable.
	for i := 0; i < len(rules); i += batchSize {
		end := i + batchSize // Calculate the end index for the current batch.
		if end > len(rules) {
			end = len(rules) // Ensure the end index does not exceed the total number of rules.
		}
		batch := rules[i:end] // Extract the current batch from the rules slice.

		// Insert the current batch into the database using CreateInBatches.
		if err := tx.CreateInBatches(batch, len(batch)).Error; err != nil {
			tx.Rollback() // Rollback the transaction if an error occurs.
			return err
		}
	}

	return tx.Commit().Error
}
