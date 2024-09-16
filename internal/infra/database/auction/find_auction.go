package auction

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/andremelinski/auction/config/logger"
	"github.com/andremelinski/auction/internal/entity"
	internalerror "github.com/andremelinski/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*entity.Auction, *internalerror.InternalError) {
	auctionEntityMongo := &AuctionEntityMongo{}

	filter := bson.M{"_id": id}

	err := ar.Collection.FindOne(ctx, filter).Decode(auctionEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("auction not found", err)
			return nil, internalerror.NewNotFoundError(
				fmt.Sprintf("auction not found with this id = %s", id),
			)
		}
		return nil, internalerror.NewNotFoundError(
			"error trying to find auction by id",
		)
	}

	auctionEntity := &entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}

	return auctionEntity, nil
}

func (ar *AuctionRepository) FindAuctions(ctx context.Context,
	status entity.AuctionStatus,
	category, productName string) ([]entity.Auction, *internalerror.InternalError) {

	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}
	if category != "" {
		filter["category"] = category
	}
	if productName != "" {
		// case insensitive
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)

	if err != nil {
		logger.Error("auction not found", err)
		return nil, internalerror.NewNotFoundError(
			"error trying to find auctions",
		)
	}

	defer cursor.Close(ctx)

	var (
		auctionEntityMongoArr []AuctionEntityMongo
		auctionEntity         []entity.Auction
	)
	if err := cursor.All(ctx, &auctionEntityMongoArr); err != nil {
		logger.Error("error trying to find auctions", err)
		return nil, internalerror.NewInternalserverError("error trying to find auctions")
	}

	for _, auctionMongo := range auctionEntityMongoArr {
		auctionEntity = append(auctionEntity,
			entity.Auction{
				Id:          auctionMongo.Id,
				ProductName: auctionMongo.ProductName,
				Category:    auctionMongo.Category,
				Description: auctionMongo.Description,
				Condition:   auctionMongo.Condition,
				Status:      auctionMongo.Status,
				Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
			},
		)
	}

	return auctionEntity, nil
}
