package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-request-response-test-funcionales/internal/domain"
	"io"
	"os"
)

func NewJsonStorage(filename string) (jsonStorage *JsonStorage, err error) {
	file1, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		err = ErrJsonStorageCantOpenFile
		return
	}
	file2, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		err = ErrJsonStorageCantOpenFile
		return
	}

	//Si no hay ningun producto inserto uno de prueba
	fileInfo, _ := file1.Stat()

	if fileInfo.Size() == 0 {
		productos := []domain.Producto{
			{
				ID:        1,
				Name:      "Producto de ejemplo",
				CodeValue: "1A",
				Price:     9.99,
			},
		}
		data, err := json.Marshal(productos)
		if err != nil {
			panic(err)
		}

		_, err = file1.Write(data)
		if err != nil {
			panic(err)
		}
	}

	// Consigo el array para calcular el ultimo id insertado
	var productos []domain.Producto
	err = json.NewDecoder(file2).Decode(&productos)
	if err != nil {
		return
	}
	lastInsertId := productos[len(productos)-1].ID
	fmt.Printf("LAST_INSERT_ID: %v", lastInsertId)
	jsonStorage = &JsonStorage{filename, lastInsertId}
	return
}

type JsonStorage struct {
	filename     string
	lastInsertID int
}

func (jsonS *JsonStorage) GetAll() (products []*domain.Producto, err error) {
	file, err := os.OpenFile(jsonS.filename, os.O_RDWR, 0644)
	if err != nil {
		err = ErrJsonStorageCantOpenFile
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		return
	}
	fmt.Println(products)
	return
}

func (jsonS *JsonStorage) GetProductById(id int) (product *domain.Producto, err error) {
	var productos []domain.Producto

	file, err := os.OpenFile(jsonS.filename, os.O_RDWR, 0644)
	if err != nil {
		err = ErrJsonStorageCantOpenFile
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&productos)
	if err != nil {
		return
	}

	for i, p := range productos {
		if p.ID == id {
			product = &productos[i]
			return
		}
	}
	err = ErrJsonStorageProductNotFound
	return
}

func (jsonS *JsonStorage) CreateProduct(producto *domain.Producto) (lastInsertId int, err error) {
	var productos []*domain.Producto

	file, err := os.OpenFile(jsonS.filename, os.O_RDWR, 0644)
	if err != nil {
		err = ErrJsonStorageCantOpenFile
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		err = ErrJsonStorageFileRead
		return
	}

	err = json.Unmarshal(content, &productos)
	if err != nil {
		err = ErrJsonStorageCantParseJsonToProduct
		return
	}

	producto.ID = jsonS.lastInsertID + 1
	productos = append(productos, producto)

	var newJson []byte

	newJson, err = json.Marshal(productos)
	if err != nil {
		err = ErrJsonStorageCantParseProductToJson
		return
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteAt(newJson, 0)
	if err != nil {
		err = ErrJsonStorageWritingFile
		return
	}

	jsonS.lastInsertID++
	lastInsertId = jsonS.lastInsertID
	return
}

func (jsonS *JsonStorage) UpdateProduct(id int, producto *domain.Producto) (err error) {
	var productos []*domain.Producto

	file, err := os.OpenFile(jsonS.filename, os.O_RDWR, 0644)
	if err != nil {
		err = ErrJsonStorageCantOpenFile
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		err = ErrJsonStorageFileRead
		return
	}

	err = json.Unmarshal(content, &productos)
	if err != nil {
		err = ErrJsonStorageCantParseJsonToProduct
		return
	}

	for i, p := range productos {
		if p.ID == id {
			productos[i] = producto

			var newJson []byte

			newJson, err = json.Marshal(productos)
			if err != nil {
				err = ErrJsonStorageCantParseProductToJson
				return
			}

			err = file.Truncate(0)
			if err != nil {
				panic(err)
			}

			_, err = file.WriteAt(newJson, 0)
			if err != nil {
				err = ErrJsonStorageWritingFile
				return
			}

			return
		}
	}
	err = ErrJsonStorageProductNotFound
	return
}

func (jsonS *JsonStorage) DeleteProduct(id int) (err error) {
	var productos []*domain.Producto

	file, err := os.OpenFile(jsonS.filename, os.O_RDWR, 0644)
	if err != nil {
		err = ErrJsonStorageCantOpenFile
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		err = ErrJsonStorageFileRead
		return
	}

	err = json.Unmarshal(content, &productos)
	if err != nil {
		err = ErrJsonStorageCantParseJsonToProduct
		return
	}

	for i, p := range productos {
		if p.ID == id {
			productos = append(productos[:i], productos[i+1:]...)

			var newJson []byte

			newJson, err = json.Marshal(productos)
			if err != nil {
				err = ErrJsonStorageCantParseProductToJson
				return
			}

			err = file.Truncate(0)
			if err != nil {
				panic(err)
			}

			_, err = file.WriteAt(newJson, 0)
			if err != nil {
				err = ErrJsonStorageWritingFile
				return
			}

			return
		}
	}
	err = ErrJsonStorageProductNotFound
	return
}

var (
	ErrJsonStorageCantOpenFile           = errors.New("could not open file, check path or permissions")
	ErrJsonStorageFileRead               = errors.New("could not read file")
	ErrJsonStorageCantParseJsonToProduct = errors.New("could not parse json to product")
	ErrJsonStorageCantParseProductToJson = errors.New("could not parse product to json")
	ErrJsonStorageProductNotFound        = errors.New("product not found")
	ErrJsonStorageWritingFile            = errors.New("could not write json")
)

//func (jsonS *JsonStorage) validCodeValue(actualCode string) error {
//	for _, p := range jsonS.Productos {
//		if p.CodeValue == actualCode {
//			return errors.New("cannot create code_value, unique violation")
//		}
//	}
//	return nil
//}

//func (jsonS *JsonStorage) GetByID(id int) (producto *domain.Producto, err error) {
//	var productos []*domain.Producto
//
//	content, err := io.ReadAll(jsonS.file)
//	if err != nil {
//		err = ErrJsonStorageFileRead
//		return
//	}
//
//	err = json.Unmarshal(content, &productos)
//	if err != nil {
//		err = ErrJsonStorageCantParseJsonToProduct
//		return
//	}
//
//	for i, p := range productos {
//		if p.ID == id {
//			producto = productos[i]
//			return
//		}
//	}
//	err = ErrJsonStorageProductNotFound
//	return
//}
//
//func (jsonS *JsonStorage) Save(producto *domain.Producto) (int, error) {
//	return 0, nil
//}
//
//func (jsonS *JsonStorage) Update(id int, producto *domain.Producto) (err error) {
//	var productos []*domain.Producto
//
//	content, err := io.ReadAll(jsonS.file)
//	if err != nil {
//		err = ErrJsonStorageFileRead
//		return
//	}
//
//	err = json.Unmarshal(content, &productos)
//	if err != nil {
//		err = ErrJsonStorageCantParseJsonToProduct
//		return
//	}
//
//	for i, p := range productos {
//		if p.ID == id {
//			productos[i] = producto
//
//			var newJson []byte
//
//			newJson, err = json.Marshal(productos)
//			if err != nil {
//				err = ErrJsonStorageCantParseProductToJson
//				return
//			}
//
//			err = os.WriteFile(jsonS.filename, newJson, 0644)
//			if err != nil {
//				err = ErrJsonStorageWritingFile
//				return
//			}
//		}
//	}
//	err = ErrJsonStorageProductNotFound
//	return
//}
//
//func (jsonS *JsonStorage) Delete(id int) (err error) {
//	var productos []*domain.Producto
//
//	content, err := io.ReadAll(jsonS.file)
//	if err != nil {
//		err = ErrJsonStorageFileRead
//		return
//	}
//
//	err = json.Unmarshal(content, &productos)
//	if err != nil {
//		err = ErrJsonStorageCantParseJsonToProduct
//		return
//	}
//
//	for i, p := range productos {
//		if p.ID == id {
//			productos = append(productos[:i], productos[i+1:]...)
//
//			var newJson []byte
//
//			newJson, err = json.Marshal(productos)
//			if err != nil {
//				err = ErrJsonStorageCantParseProductToJson
//				return
//			}
//
//			err = os.WriteFile(jsonS.filename, newJson, 0644)
//			if err != nil {
//				err = ErrJsonStorageWritingFile
//				return
//			}
//		}
//	}
//	err = ErrJsonStorageProductNotFound
//	return
//}
