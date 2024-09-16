package bid

import (
	"context"
	"sync"

	"github.com/andremelinski/auction/config/logger"
	"github.com/andremelinski/auction/internal/entity"
	"github.com/andremelinski/auction/internal/infra/database/auction"
	internalerror "github.com/andremelinski/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"userId"`
	AuctionId string  `bson:"auctionId"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func NewBidRepository(db *mongo.Database, auctionRepo *auction.AuctionRepository) *BidRepository {
	return &BidRepository{
		db.Collection("bids"),
		auctionRepo,
	}
}

func (bdr *BidRepository) CreateAuction(ctx context.Context, bidsArr []entity.Bid) *internalerror.InternalError {
	// salvando em batch
	wg := sync.WaitGroup{}
	for _, bid := range bidsArr {
		wg.Add(1)
		go func(bid entity.Bid) {
			auction, err := bdr.AuctionRepository.FindAuctionById(ctx, bid.AuctionId)

			if err != nil {
				logger.Error("error trying to find auction by id", err)
				return
			}

			if auction.Status != entity.Active {
				return
			}

			payload := &BidEntityMongo{
				Id:        bid.Id,
				UserId:    bid.UserId,
				AuctionId: bid.AuctionId,
				Amount:    bid.Amount,
				Timestamp: bid.Timestamp.Unix(),
			}

			if _, err := bdr.Collection.InsertOne(ctx, payload); err != nil {
				logger.Error("error trying to InsertOne bid", err)
				return
			}
		}(bid)

	}
	defer wg.Wait()
	return nil
}
