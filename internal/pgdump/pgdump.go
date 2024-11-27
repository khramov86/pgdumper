package pgdump

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/JCoupalK/go-pgdump"
)

func BackupPostgreSQL(username, password, hostname, dbname string, port int, outputDir string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostname, port, username, password, dbname)

	currentTime := time.Now()
	dumpFilename := filepath.Join(outputDir, fmt.Sprintf("%s-%s.sql", dbname, currentTime.Format("20060102T150405")))

	dumper := pgdump.NewDumper(psqlInfo)

	if err := dumper.DumpDatabase(dumpFilename); err != nil {
		fmt.Printf("Error dumping database: %v", err)
		os.Remove(dumpFilename)
		return
	}

	fmt.Println("Backup successfully saved to", dumpFilename)
}
