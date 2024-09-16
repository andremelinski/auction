package bid

import (
	"context"
	"time"

	"github.com/andremelinski/auction/internal/entity"
	internalerror "github.com/andremelinski/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
)

func (bdr *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]entity.Bid, *internalerror.InternalError) {
	var (
		bidEntityMongoArr []BidEntityMongo
		bidEntityArr      []entity.Bid
	)

	filter := bson.M{"auctionId": auctionId}
	cursor, err := bdr.Collection.Find(ctx, filter)
	if err != nil {
		return nil, internalerror.NewNotFoundError(
			"error trying to find auction by id",
		)
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &bidEntityMongoArr); err != nil {
		return nil, internalerror.NewNotFoundError(
			"error trying to bid cursor.All",
		)
	}

	for _, bid := range bidEntityMongoArr {
		bidEntityArr = append(bidEntityArr, entity.Bid{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: time.Unix(bid.Timestamp, 0),
		})

	}
	return bidEntityArr, nil
}
