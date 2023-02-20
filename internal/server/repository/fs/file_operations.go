package fs

import (
	"context"
	"os"
)

func SaveFile(ctx context.Context, filePath string, data []byte) error {
	buffer := make([]byte, 1024)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	for idx := 0; idx < len(data); idx = idx + len(buffer) {
		select {
		default:
			buffer = data[idx:min(len(data), idx+len(buffer))]
			_, err = f.Write(buffer)
			if err != nil {
				f.Close()
				err = os.RemoveAll(filePath)
				if err != nil {
				}
				return err
			}
		case <-ctx.Done():
			f.Close()
			os.RemoveAll(filePath)
			return ctx.Err()
		}
	}
	f.Close()
	return nil
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func ReadFile(ctx context.Context, storageFilePath string) ([]byte, error) {
	buffer := make([]byte, 1024)
	f, err := os.OpenFile(storageFilePath, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	res := []byte{}
outerLoop:
	for {
		select {
		default:
			count, err := f.Read(buffer)
			if count == 0 {
				f.Close()
				break outerLoop
			}
			if err != nil {
				f.Close()
				return nil, err
			}
			res = append(res, buffer[:count]...)
		case <-ctx.Done():
			f.Close()
			return res, ctx.Err()
		}
	}

	return res, nil
}

func DeleteFile(ctx context.Context, filePath string) error {
	select {
	default:
		os.RemoveAll(filePath)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
