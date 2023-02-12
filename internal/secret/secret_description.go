package secret

import (
	"fmt"
)

func (s *SecretDescription) ToStdout() {
	fmt.Println("  {")
	fmt.Printf("    id: %s\n", s.Id)
	fmt.Printf("    created at: %s\n", s.CreatedAt)
	fmt.Printf("    modified at: %s\n", s.ModifiedAt)
	fmt.Printf("    type: %s\n", SecretType_name[int32(s.SecretType)])
	fmt.Println("  }")
}
