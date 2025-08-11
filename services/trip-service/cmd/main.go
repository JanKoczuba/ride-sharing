package main

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	fare := &domain.RideFareModel{
		ID:                primitive.NewObjectID(),
		UserID:            "42",
		PackageSlug:       "",
		TotalPriceInCents: 0,
	}
	svc.CreateTrip(ctx, fare)

	for {
		time.Sleep(time.Second)
	}
}
