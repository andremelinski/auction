package auction

import (
	"context"

	"github.com/andremelinski/auction/internal/entity"
	internalerror "github.com/andremelinski/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                  `bson:"_id"`
	ProductName string                  `bson:"product_name"`
	Category    string                  `bson:"category"`
	Description string                  `bson:"description"`
	Condition   entity.ProductCondition `bson:"condition"`
	Status      entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                   `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(db *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		db.Collection("actions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionInfo entity.Auction) *internalerror.InternalError {
	// fazendo o bson para salvar
	mongoPlayload := &AuctionEntityMongo{
		Id:          auctionInfo.Id,
		ProductName: auctionInfo.ProductName,
		Category:    auctionInfo.Category,
		Description: auctionInfo.Description,
		Condition:   auctionInfo.Condition,
		Status:      auctionInfo.Status,
		Timestamp:   auctionInfo.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, mongoPlayload)
	if err != nil {
		return internalerror.NewNotFoundError(
			"error trying to create auction",
		)
	}
	return nil
}
