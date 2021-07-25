package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sa-back/graph/generated"
	"sa-back/graph/model"
)

func (r *queryResolver) listAnalysisTemplate() []model.Analysis {
	Analyzes := []model.Analysis{
		model.Analysis{
			Name:        "ABC",
			Description: "Разделение товаров на три категории\nпо степени их значимости. Поможет определить рентабельные продукты,\nважных клиентов и поставщиков.",
			Image:       "",
			Disabled:    false,
		},
		model.Analysis{
			Name:        "XYZ",
			Description: "Определение стабильности или устойчивости спроса на товары. Поможет определить какие товары обязательно должны быть на складе или прилавке.",
			Image:       "",
			Disabled:    true,
		},
		model.Analysis{
			Name:        "ABC-XYZ",
			Description: "Комбинация анализов которая поможет спланировать закупки. Результатом будет 6 групп которые покажут самые прибыльные и востребованные товары.",
			Image:       "",
			Disabled:    true,
		},
		model.Analysis{
			Name:        "FMR",
			Description: "Анализ выявит насколько товары востребованы. Поможет определить оптимальное место хранения товаров, снизить риск возникновения просроченности, и убрать из ассортимента непопулярную\nпродукцию.",
			Image:       "",
			Disabled:    true,
		},
		model.Analysis{
			Name:        "RFM",
			Description: "Разделение клиентов по частоте и сумме покупок. Поможет выявить тех, кто больше и чаще платит деньги, а кто давно ничего не покупал.",
			Image:       "",
			Disabled:    true,
		},
		model.Analysis{
			Name:        "VEN",
			Description: "Разделение товаров или ресурсов по степени их важности или необходимости. Поможет определить самые важные запчасти необходимые для производства\nили приоритетные лекарственные\nсредства для аптеки.",
			Image:       "",
			Disabled:    true,
		},
	}
	return Analyzes
}

func (r *queryResolver) ListAnalysis(ctx context.Context) ([]*model.Analysis, error) {
	var Analyzes []*model.Analysis
	for i := 0; i < len(r.listAnalysisTemplate()); i++ {
		Analyzes = append(Analyzes, &r.listAnalysisTemplate()[i])
	}
	return Analyzes, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
