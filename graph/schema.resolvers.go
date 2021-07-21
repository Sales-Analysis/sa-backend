package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sa-back/graph/generated"
	"sa-back/graph/model"
)

func (r *queryResolver) ListAnalysis(ctx context.Context) ([]*model.Analysis, error) {
	var Analyzes []*model.Analysis
	a := model.Analysis{
		Name:        "ABC",
		Description: "Разделение товаров на три категории по степени их значимости. Поможет определить рентабельные продукты, важных клиентов и поставщиков",
		Image:       "Тут будет путь до картинки",
		Disabled:    false,
	}
	Analyzes = append(Analyzes, &a)
	return Analyzes, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
