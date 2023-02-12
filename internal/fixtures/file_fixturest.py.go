package fixtures

import "crypto/rand"

func PathFixture() string {
	path := make([]byte, 20)
	_, err := rand.Read(path)
	if err != nil {
		panic(err)
	}

	generator := []func(int) byte{
		func(i int) byte { return 'a' + path[i]%26 },
		func(i int) byte { return '0' + path[i]%10 },
	}
	for i := range path {
		path[i] = generator[path[i]%2](i)
	}

	return string(path)
}
