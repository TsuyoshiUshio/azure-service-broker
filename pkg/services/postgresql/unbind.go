package postgresql

import (
	"fmt"

	"github.com/Azure/azure-service-broker/pkg/service"
)

func (m *module) Unbind(
	provisioningContext service.ProvisioningContext, // nolint: unparam
	bindingContext service.BindingContext,
) error {
	pc, ok := provisioningContext.(*postgresqlProvisioningContext)
	if !ok {
		return fmt.Errorf(
			"error casting provisioningContext as *postgresqlProvisioningContext",
		)
	}
	bc, ok := bindingContext.(*postgresqlBindingContext)
	if !ok {
		return fmt.Errorf(
			"error casting bindingContext as *postgresqlBindingContext",
		)
	}

	db, err := getDBConnection(pc, primaryDB)
	if err != nil {
		return err
	}
	defer db.Close() // nolint: errcheck

	_, err = db.Exec(
		fmt.Sprintf("drop role %s", bc.LoginName),
	)
	if err != nil {
		return fmt.Errorf(`error dropping role "%s": %s`, bc.LoginName, err)
	}

	return nil
}
