// Copyright 2019 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tiny

import (
	"bytes"

	"github.com/andybalholm/brotli"
)

// BrotliEncode brotli encode
func BrotliEncode(buf []byte, quality int) ([]byte, error) {
	if quality <= 0 || quality > maxBrotliQuality {
		quality = defauttBrotliQuality
	}
	buffer := new(bytes.Buffer)
	w := brotli.NewWriterLevel(buffer, quality)
	_, err := w.Write(buf)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
