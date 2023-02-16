package secret

import (
	"fmt"
	"time"
)

func (s *SecretDescription) ToStdout() {
	fmt.Println("  {")
	fmt.Printf("    id          : %s\n", s.Id)
	fmt.Printf("    created at  : %s\n", s.CreatedAt)
	if t, err := time.Parse(time.RFC3339, s.ModifiedAt); err == nil && !t.IsZero() {
		fmt.Printf("    modified at : %s\n", s.ModifiedAt)
	}
	fmt.Printf("    type        : %s\n", SecretType_name[int32(s.SecretType)])
	fmt.Println("  }")
}
