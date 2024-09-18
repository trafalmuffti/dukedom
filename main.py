import random

class Dukedom:
    def __init__(self):
        self.year = 1
        self.population = 100  # Starting population
        self.grain_stores = 1000  # Bushels of grain
        self.land = 500  # Acres of land
        self.treasury = 500  # Amount of money
        self.army = 20  # Number of soldiers
        self.grain_per_acre = 3  # Yield per acre

    def display_status(self):
        print(f"\nYear: {self.year}")
        print(f"Population: {self.population}")
        print(f"Grain Stores: {self.grain_stores} bushels")
        print(f"Land: {self.land} acres")
        print(f"Treasury: {self.treasury} gold coins")
        print(f"Army: {self.army} soldiers\n")

    def plant_grain(self):
        while True:
            try:
                acres = int(input(f"How many acres will you plant? (0-{self.land}): "))
                if 0 <= acres <= self.land:
                    required_grain = acres // 2
                    if required_grain <= self.grain_stores:
                        self.grain_stores -= required_grain
                        yield_grain = acres * self.grain_per_acre
                        self.grain_stores += yield_grain
                        print(f"You planted {acres} acres, using {required_grain} bushels of grain.")
                        print(f"You harvested {yield_grain} bushels of grain.")
                    else:
                        print("You don't have enough grain to plant that many acres.")
                    break
                else:
                    print("Invalid number of acres.")
            except ValueError:
                print("Please enter a valid number.")

    def feed_population(self):
        while True:
            try:
                grain = int(input(f"How much grain will you distribute to your people? (0-{self.grain_stores}): "))
                if 0 <= grain <= self.grain_stores:
                    self.grain_stores -= grain
                    food_needed = self.population * 2
                    if grain < food_needed:
                        deaths = (food_needed - grain) // 2
                        self.population -= deaths
                        print(f"{deaths} people starved.")
                    else:
                        print("Everyone has been fed.")
                    break
                else:
                    print("Invalid amount of grain.")
            except ValueError:
                print("Please enter a valid number.")

    def recruit_army(self):
        while True:
            try:
                recruits = int(input("How many soldiers will you recruit? (0-50): "))
                cost = recruits * 10
                if 0 <= recruits <= 50 and cost <= self.treasury:
                    self.army += recruits
                    self.treasury -= cost
                    print(f"You recruited {recruits} soldiers, costing {cost} gold coins.")
                    break
                else:
                    print("You can't afford that many recruits.")
            except ValueError:
                print("Please enter a valid number.")

    def collect_taxes(self):
        tax_income = self.population * 3
        self.treasury += tax_income
        print(f"You collected {tax_income} gold coins in taxes.")

    def random_events(self):
        event = random.choice(["good_harvest", "bad_harvest", "enemy_attack", "plague", None])
        if event == "good_harvest":
            bonus = self.land // 5
            self.grain_stores += bonus
            print(f"A bumper crop! You gain {bonus} bushels of grain.")
        elif event == "bad_harvest":
            loss = self.land // 5
            self.grain_stores -= loss
            print(f"A drought hit your fields. You lost {loss} bushels of grain.")
        elif event == "enemy_attack":
            soldiers_lost = random.randint(1, self.army // 2)
            if soldiers_lost < self.army:
                self.army -= soldiers_lost
                print(f"Your dukedom was attacked! {soldiers_lost} soldiers were lost defending the realm.")
            else:
                print("The enemy overran your army! Your dukedom has fallen.")
                self.population = 0  # Game Over condition
        elif event == "plague":
            plague_deaths = random.randint(1, self.population // 3)
            self.population -= plague_deaths
            print(f"A plague swept through your dukedom, killing {plague_deaths} people.")

    def next_year(self):
        self.year += 1
        population_growth = random.randint(0, 10)
        self.population += population_growth
        print(f"{population_growth} people were born this year.")

    def run(self):
        print("Welcome to Dukedom! Lead your people wisely.")
        while self.population > 0 and self.grain_stores > 0:
            self.display_status()
            self.collect_taxes()
            self.plant_grain()
            self.feed_population()
            self.recruit_army()
            self.random_events()
            self.next_year()

        print("Your reign has ended. The dukedom has failed.")

# Start the game
dukedom = Dukedom()
dukedom.run()
