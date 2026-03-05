package main

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/presenter"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/storage"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// TODO запуск логирования

	// TODO запуск конфигов

	// TODO получаем значение переменных окружения и/или через конфиг

	// подключение к БД
	db, err := storage.InitDB()
	if err != nil {
		fmt.Println("Подключиться к БД указателей с яндекс диска не удалось - ", err)
		//TODO логирование
	}
	storageImpl := storage.NewStorage(db)

	presenterImpl := presenter.New()

	repositoryImpl := repository.NewRepository(storageImpl)

	// TODO usacase

	// TODO подключение к ВК

	// TODO запуск шедулеров

	// TODO автозапуск при сбоях
}
