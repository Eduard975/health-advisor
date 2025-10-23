def make_exercise_sentence(row):
    return (
        f"Performing {row['Activity, Exercise or Sport (1 hour)'].lower()} for one hour burns approximately "
        f"{row['130 lb']} calories for a person weighing 130 pounds, "
        f"{row['155 lb']} calories for someone weighing 155 pounds, "
        f"{row['180 lb']} calories for a person at 180 pounds, and "
        f"{row['205 lb']} calories for someone weighing 205 pounds. "
        f"This activity has an estimated energy cost of {row['Calories per kg']:.2f} calories per kilogram of body weight."
    )

def make_food_sentence(row):
    return (
        f"{row['food'].capitalize()} provides {row['Caloric Value']} kcal per 100 grams. "
        f"It contains {row['Fat']} g of total fat "
        f"({row['Saturated Fats']} g saturated, {row['Monounsaturated Fats']} g monounsaturated, "
        f"and {row['Polyunsaturated Fats']} g polyunsaturated), "
        f"{row['Carbohydrates']} g of carbohydrates with {row['Sugars']} g of sugars, "
        f"{row['Protein']} g of protein, and {row['Dietary Fiber']} g of dietary fiber. "
        f"It has {row['Cholesterol']} mg of cholesterol, {row['Sodium']} g of sodium, and {row['Water']} g of water. "
        f"Vitamin content per 100g: A = {row['Vitamin A']} mg, B1 (Thiamine) = {row['Vitamin B1']} mg, "
        f"B2 (Riboflavin) = {row['Vitamin B2']} mg, B3 (Niacin) = {row['Vitamin B3']} mg, "
        f"B5 (Pantothenic Acid) = {row['Vitamin B5']} mg, B6 = {row['Vitamin B6']} mg, "
        f"B11 (Folic Acid) = {row['Vitamin B11']} mg, B12 = {row['Vitamin B12']} mg, "
        f"Vitamin C = {row['Vitamin C']} mg, Vitamin D = {row['Vitamin D']} mg, "
        f"Vitamin E = {row['Vitamin E']} mg, Vitamin K = {row['Vitamin K']} mg. "
        f"Mineral content per 100g: Calcium = {row['Calcium']} mg, Copper = {row['Copper']} mg, "
        f"Iron = {row['Iron']} mg, Magnesium = {row['Magnesium']} mg, "
        f"Manganese = {row['Manganese']} mg, Phosphorus = {row['Phosphorus']} mg, "
        f"Potassium = {row['Potassium']} mg, Selenium = {row['Selenium']} mg, "
        f"and Zinc = {row['Zinc']} mg. "
        f"The overall nutrition density score is {row['Nutrition Density']}. "
    )