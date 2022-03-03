module github.com/clivewalkden/m2-db-sync

go 1.16

require (
	github.com/fatih/color v1.13.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
)

replace github.com/clivewalkden/m2-db-sync/common => ../common/

replace github.com/clivewalkden/m2-db-sync/validation => ../validation/
