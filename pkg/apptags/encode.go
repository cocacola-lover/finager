package apptags

import (
	"encoding/binary"
	"hash/crc32"
	"io"
	"money_app/pkg/apperrors"
	"money_app/pkg/encodingutils"
	"os"
)

const tagsFile = "tags.bin"

func AddTag(tag string) error {
	file, err := os.OpenFile(tagsFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	taghash := crc32.ChecksumIEEE([]byte(tag))

	_, err = file.Write(encodingutils.Uint32ToBytes(taghash))
	if err != nil {
		return err
	}

	_, err = file.Write(encodingutils.StringToBytes(tag))
	if err != nil {
		return err
	}

	return nil
}

func readTagInto(r io.Reader, tags map[uint32]string) error {
	data := make([]byte, 8)
	n, err := r.Read(data)

	if err != nil {
		if err == io.EOF {
			if n == 0 {
				return err
			}
			return apperrors.ErrCorruptedData
		}
		return err
	}
	key := binary.BigEndian.Uint32(data[:4])
	length := int(binary.BigEndian.Uint32(data[4:]))

	data = make([]byte, length)
	n, err = r.Read(data)
	if err != nil {
		if err == io.EOF {
			if n != int(length) {
				return apperrors.ErrCorruptedData
			}
		} else {
			return err
		}
	}

	tags[key] = string(data)
	return nil
}

func ReadTags() (map[uint32]string, error) {

	tags := make(map[uint32]string)

	file, err := os.Open(tagsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	for {
		err := readTagInto(file, tags)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}

	return tags, nil
}
