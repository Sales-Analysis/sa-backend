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
	for i := 0; i < len(r.listAnalysisTemplate()); i++ {
		Analyzes = append(Analyzes, &r.listAnalysisTemplate()[i])
	}
	return Analyzes, nil
}

func (r *queryResolver) ListHiw(ctx context.Context) ([]*model.Hiw, error) {
	var Hiw []*model.Hiw
	for i := 0; i < len(r.listHiwTemplate()); i++ {
		Hiw = append(Hiw, &r.listHiwTemplate()[i])
	}
	return Hiw, nil
}

func (r *queryResolver) ListFaq(ctx context.Context) ([]*model.Faq, error) {
	var Faq []*model.Faq
	for i := 0; i < len(r.listFAQTemplate()); i++ {
		Faq = append(Faq, &r.listFAQTemplate()[i])
	}
	return Faq, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) ListFAQ(ctx context.Context) ([]*model.Faq, error) {
	var Faq []*model.Faq
	for i := 0; i < len(r.listFAQTemplate()); i++ {
		Faq = append(Faq, &r.listFAQTemplate()[i])
	}
	return Faq, nil
}
func (r *queryResolver) listAnalysisTemplate() []model.Analysis {
	Analyzes := []model.Analysis{
		model.Analysis{
			ID:          0,
			Name:        "ABC",
			Description: "Разделение товаров на три категории\nпо степени их значимости. Поможет определить рентабельные продукты,\nважных клиентов и поставщиков.",
			Image:       "data/analyses/abc.png",
			Disabled:    false,
		},
		model.Analysis{
			ID:          1,
			Name:        "XYZ",
			Description: "Определение стабильности или устойчивости спроса на товары. Поможет определить какие товары обязательно должны быть на складе или прилавке.",
			Image:       "data/analyses/xyz.png",
			Disabled:    true,
		},
		model.Analysis{
			ID:          2,
			Name:        "ABC-XYZ",
			Description: "Комбинация анализов которая поможет спланировать закупки. Результатом будет 6 групп которые покажут самые прибыльные и востребованные товары.",
			Image:       "data/analyses/abc-xyz.png",
			Disabled:    true,
		},
		model.Analysis{
			ID:          3,
			Name:        "FMR",
			Description: "Анализ выявит насколько товары востребованы. Поможет определить оптимальное место хранения товаров, снизить риск возникновения просроченности, и убрать из ассортимента непопулярную\nпродукцию.",
			Image:       "data/analyses/fmr.png",
			Disabled:    true,
		},
		model.Analysis{
			ID:          4,
			Name:        "RFM",
			Description: "Разделение клиентов по частоте и сумме покупок. Поможет выявить тех, кто больше и чаще платит деньги, а кто давно ничего не покупал.",
			Image:       "data/analyses/rfm.png",
			Disabled:    true,
		},
		model.Analysis{
			ID:          5,
			Name:        "VEN",
			Description: "Разделение товаров или ресурсов по степени их важности или необходимости. Поможет определить самые важные запчасти необходимые для производства\nили приоритетные лекарственные\nсредства для аптеки.",
			Image:       "data/analyses/ven.png",
			Disabled:    true,
		},
	}
	return Analyzes
}
func (r *queryResolver) listHiwTemplate() []model.Hiw {
	Hiw := []model.Hiw{
		model.Hiw{
			ID:          0,
			Title:       "Загрузите файл с данными",
			Description: "Подготовьте файл в котором будут категории товаров или товары, информация по продажам \nи прибыли",
			Image:       "data/hiw/step_1.png",
		},
		model.Hiw{
			ID:          1,
			Title:       "Настройте расчет",
			Description: "Выберите столбцы из вашего файла, по которым нужно произвести расчет и определить группы товаров",
			Image:       "data/hiw/step_2.png/",
		},
		model.Hiw{
			ID:          2,
			Title:       "Получите готовый отчет",
			Description: "Отчет рассчитает и распределит все товары \nи категории по подходящим группам. Вы сможете увидеть какие товары самые рентабельные, \nа какие нет.",
			Image:       "data/hiw/step_3.png",
		},
	}
	return Hiw
}
func (r *queryResolver) listFAQTemplate() []model.Faq {
	Faq := []model.Faq{
		model.Faq{
			ID:          0,
			Title:       "Какие знания нужны для проведения \nанализа?",
			Description: "Никаких. Сервис подскажет, что нужно сделат чтобы провести анализ и правильно разобраться в его результатах.",
		},
	}
	return Faq
}
