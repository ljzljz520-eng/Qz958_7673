package model

type IngredientRule struct {
	Name        string
	Category    string
	DefaultUnit string
	Min         float64
	Max         float64
	Allergens   []string
	Seasonal    bool
}

var catalog = []IngredientRule{
	{"all-purpose flour", "flour", "kg", 0.1, 100, nil, false}, {"bread flour", "flour", "kg", 0.1, 100, nil, false}, {"cake flour", "flour", "kg", 0.1, 80, nil, false}, {"whole wheat flour", "flour", "kg", 0.1, 80, nil, false}, {"rye flour", "flour", "kg", 0.1, 60, nil, false}, {"almond flour", "flour", "kg", 0.05, 40, []string{"nuts"}, false}, {"coconut flour", "flour", "kg", 0.05, 30, []string{"coconut"}, false}, {"oat flour", "flour", "kg", 0.1, 60, nil, false}, {"rice flour", "flour", "kg", 0.1, 60, nil, false}, {"cornmeal", "flour", "kg", 0.1, 60, nil, false},
	{"granulated sugar", "sweetener", "kg", 0.1, 100, nil, false}, {"brown sugar", "sweetener", "kg", 0.1, 90, nil, false}, {"powdered sugar", "sweetener", "kg", 0.1, 70, nil, false}, {"molasses", "sweetener", "ml", 10, 5000, nil, false}, {"honey", "sweetener", "ml", 10, 5000, nil, false}, {"maple syrup", "sweetener", "ml", 10, 5000, nil, false}, {"agave", "sweetener", "ml", 10, 3000, nil, false}, {"date syrup", "sweetener", "ml", 10, 3000, nil, false},
	{"unsalted butter", "fat", "kg", 0.05, 80, []string{"milk"}, false}, {"salted butter", "fat", "kg", 0.05, 80, []string{"milk"}, false}, {"shortening", "fat", "kg", 0.05, 60, nil, false}, {"coconut oil", "fat", "ml", 10, 5000, []string{"coconut"}, false}, {"olive oil", "fat", "ml", 10, 5000, nil, false}, {"vegetable oil", "fat", "ml", 10, 5000, nil, false},
	{"whole milk", "dairy", "ml", 20, 10000, []string{"milk"}, false}, {"cream", "dairy", "ml", 20, 8000, []string{"milk"}, false}, {"cream cheese", "dairy", "kg", 0.05, 50, []string{"milk"}, false}, {"yogurt", "dairy", "kg", 0.05, 50, []string{"milk"}, false}, {"buttermilk", "dairy", "ml", 20, 8000, []string{"milk"}, false}, {"condensed milk", "dairy", "ml", 20, 5000, []string{"milk"}, false},
	{"cocoa powder", "cocoa", "kg", 0.05, 40, nil, false}, {"dark chocolate", "cocoa", "kg", 0.05, 50, []string{"soy"}, false}, {"milk chocolate", "cocoa", "kg", 0.05, 50, []string{"milk", "soy"}, false}, {"white chocolate", "cocoa", "kg", 0.05, 50, []string{"milk", "soy"}, false}, {"cacao nibs", "cocoa", "kg", 0.02, 20, nil, false},
	{"baking powder", "leavener", "kg", 0.01, 20, nil, false}, {"baking soda", "leavener", "kg", 0.01, 20, nil, false}, {"instant yeast", "leavener", "kg", 0.01, 20, nil, false}, {"active yeast", "leavener", "kg", 0.01, 20, nil, false}, {"cream of tartar", "leavener", "kg", 0.01, 10, nil, false},
	{"vanilla extract", "flavor", "ml", 1, 1000, nil, false}, {"almond extract", "flavor", "ml", 1, 1000, []string{"nuts"}, false}, {"lemon zest", "flavor", "kg", 0.01, 10, nil, true}, {"orange zest", "flavor", "kg", 0.01, 10, nil, true}, {"cinnamon", "spice", "kg", 0.01, 20, nil, false}, {"nutmeg", "spice", "kg", 0.01, 10, []string{"nuts"}, false}, {"ginger", "spice", "kg", 0.01, 20, nil, false},
	{"pistachio", "nut", "kg", 0.01, 30, []string{"nuts"}, false}, {"walnut", "nut", "kg", 0.01, 30, []string{"nuts"}, false}, {"pecan", "nut", "kg", 0.01, 30, []string{"nuts"}, false}, {"hazelnut", "nut", "kg", 0.01, 30, []string{"nuts"}, false}, {"cashew", "nut", "kg", 0.01, 30, []string{"nuts"}, false}, {"peanut", "nut", "kg", 0.01, 30, []string{"peanuts"}, false}, {"sesame", "seed", "kg", 0.01, 20, []string{"sesame"}, false}, {"poppy seed", "seed", "kg", 0.01, 20, nil, false}, {"sunflower seed", "seed", "kg", 0.01, 20, nil, false}, {"pumpkin seed", "seed", "kg", 0.01, 20, nil, false},
	{"strawberry", "fruit", "kg", 0.05, 100, nil, true}, {"blueberry", "fruit", "kg", 0.05, 100, nil, true}, {"raspberry", "fruit", "kg", 0.05, 100, nil, true}, {"cherry", "fruit", "kg", 0.05, 100, nil, true}, {"apple", "fruit", "kg", 0.05, 100, nil, false}, {"pear", "fruit", "kg", 0.05, 100, nil, false}, {"peach", "fruit", "kg", 0.05, 100, nil, true}, {"apricot", "fruit", "kg", 0.05, 80, nil, true}, {"cranberry", "fruit", "kg", 0.05, 80, nil, false}, {"raisin", "fruit", "kg", 0.05, 80, nil, false},
	{"fig", "fruit", "kg", 0.05, 60, nil, false}, {"date", "fruit", "kg", 0.05, 60, nil, false}, {"banana", "fruit", "kg", 0.05, 100, nil, false}, {"pineapple", "fruit", "kg", 0.05, 100, nil, false}, {"mango", "fruit", "kg", 0.05, 100, nil, true}, {"lime", "fruit", "kg", 0.02, 60, nil, true}, {"lemon", "fruit", "kg", 0.02, 60, nil, false}, {"orange", "fruit", "kg", 0.02, 60, nil, false}, {"grapefruit", "fruit", "kg", 0.02, 40, nil, false}, {"passionfruit", "fruit", "kg", 0.02, 40, nil, true},
	{"sea salt", "seasoning", "kg", 0.01, 20, nil, false}, {"kosher salt", "seasoning", "kg", 0.01, 20, nil, false}, {"vanilla bean", "flavor", "kg", 0.001, 5, nil, false}, {"cardamom", "spice", "kg", 0.001, 10, nil, false}, {"clove", "spice", "kg", 0.001, 10, nil, false}, {"allspice", "spice", "kg", 0.001, 10, nil, false}, {"anise", "spice", "kg", 0.001, 10, nil, false}, {"fennel", "spice", "kg", 0.001, 10, nil, false}, {"saffron", "spice", "g", 0.1, 100, nil, false}, {"matcha", "tea", "kg", 0.001, 10, nil, false},
	{"espresso powder", "coffee", "kg", 0.001, 20, nil, false}, {"instant coffee", "coffee", "kg", 0.001, 20, nil, false}, {"black tea", "tea", "kg", 0.001, 20, nil, false}, {"earl grey", "tea", "kg", 0.001, 20, nil, false}, {"chai blend", "tea", "kg", 0.001, 20, nil, false}, {"rose water", "flavor", "ml", 1, 1000, nil, false}, {"orange blossom water", "flavor", "ml", 1, 1000, nil, false}, {"rum extract", "flavor", "ml", 1, 1000, nil, false}, {"brandy extract", "flavor", "ml", 1, 1000, nil, false}, {"mint extract", "flavor", "ml", 1, 1000, nil, false},
}

func Catalog() []IngredientRule { return append([]IngredientRule(nil), catalog...) }
func FindIngredient(name string) (IngredientRule, bool) {
	for _, r := range catalog {
		if r.Name == name {
			return r, true
		}
	}
	return IngredientRule{}, false
}
func ValidateAgainstCatalog(r Record) error {
	rule, ok := FindIngredient(r.Ingredient)
	if !ok {
		return nil
	}
	if r.Quantity < rule.Min || r.Quantity > rule.Max {
		return ErrInvalid("quantity outside ingredient range")
	}
	return nil
}
func CategoryNames() []string {
	seen := map[string]bool{}
	out := []string{}
	for _, r := range catalog {
		if !seen[r.Category] {
			seen[r.Category] = true
			out = append(out, r.Category)
		}
	}
	return out
}
