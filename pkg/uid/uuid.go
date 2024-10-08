package uid

import (
	"github.com/google/uuid"
	"os"
	"regexp"
)

func init() {
	// k8s: inject pod ip
	uuid.SetNodeInterface(os.Getenv("NODE_IDENTIFIER"))
}

func UuidV1() uuid.UUID {
	// TODO when do this loc failed?
	return uuid.Must(uuid.NewUUID())
}

func IsValidUUID(uuid string) bool {
	regex := regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")
	return regex.MatchString(uuid)
}
