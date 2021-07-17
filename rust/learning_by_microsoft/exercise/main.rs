// Declare Car struct to describe vehicle with four named fields
struct Car {
    color: String,
    motor: Transmission,
    roof: bool,
    age: (String, u32),
}

#[derive(PartialEq, Debug)]
// Declare enum for Car transmission type
// TO DO: Fix enum definition so code compiles
enum Transmission {
    Manual,
    SemiAuto,
    Automatic,
}

// Build a new "Car" using the values of four input arguments
// - color (String)
// - motor (Transmission enum)
// - roof (boolean, true if the car has closed roof)
// - miles (u32)
// Call the car_quality(miles) function to get the car age
// Return an instance of a "Car" struct with the arrow `->` syntax
fn car_factory(color: String, motor: Transmission, roof: bool, miles: u32) -> Car {

    // Create a new "Car" instance as requested
    // - Bind first three fields to values of input arguments
    // TO DO: Replace the "mileage" field from the previous exercise with an "age" field
    // TO DO" The "age" field calls the "car_quality" function with the "miles" input argument
    let car = Car {
        color: color,
        motor: motor,
        roof: roof,
        age: car_quality(miles),
    };

    // Return new instance of "Car" struct
    return car
}

// Get the car quality by testing the value of the input argument
// - miles (u32)
// Create a tuple for the car quality with the age ("New" or "Used") and miles
// Return a tuple with the arrow `->` syntax
fn car_quality (miles: u32) -> (String, u32) {

    // Declare and initialize the return tuple value
    // For a new car, set the miles to 0
    // TO DO: Correct the quality declaration so we can change the values later
    let mut quality: (String, u32) = ("New".to_string(), 0);

    // Use a conditional expression to check the miles
    // If the car has accumulated miles, then the car is used
    if miles > 0 {
        // TO DO: Add the code to set the quality value for a used car
        quality.0 = "Used".to_string();
        quality.1 = miles;
    }

    // TO DO: Return the completed tuple
    return quality;
}


fn main() {
    // Initialize a hash map for car orders
    // - Keys: New or Used, Values: integer
    // - Keys: Manual or Automatic, Values: integer
    // TO DO: Fix syntax to create the "orders" hash map
    use std::collections::HashMap;
    let mut orders: HashMap<String, u32> = HashMap::new();
    let (mut new_cars, mut used_cars) = (1, 1);
    let (mut manual, mut auto) = (1, 1);

    // Create car color array
    // TO DO: Set the values: 0 = Blue, 1 = Green, 2 = Red, 3 = Silver
    let colors = ["Blue", "Green", "Red", "Silver"];

    // Declare the car type and initial values
    // TO DO: Create "car" as a "Car" struct
    // TO DO: Create "engine" as a "Transmission" enum
    // TO DO: Initialize "roof" to the value when the car has a hard top
    let mut miles = 1000; // Start used cars with 1,000 miles
    let mut roof = true;  // convertible = false | hard top = true
    let mut engine = Transmission::Manual;

    let mut car = Car{
        color: colors[1].to_string(),
        motor: engine,
        roof: roof,
        age: car_quality(miles),
    };

    // Initialize variables
    let (mut index, mut order) = (1, 1);

    // Order 11 cars
    // TO DO: Replace "loop expression" - loop 11 times, use "order" variable
    while order < 12 {

        // Set car transmission type, make some roofs convertible
        // TO DO: Add conditional expression
        // TO DO: Check order number, set engine type, fix syntax
        // TO DO: If order % 3 equals 0, engine is "Automatic"
        // TO DO: If order % 2 equals 0, engine is "SemiAuto" | else, engine is "Manual"
        // When order % 3, swap roof type for fun!
        if order % 3 == 0 {
            engine = Transmission::Automatic;
            // ADD <K, V> pair to hash map
            orders.insert("Automatic".to_string(), auto);
            auto = auto + 1;
            roof = !roof;
        } else if order % 2 == 0 {
            engine = Transmission::SemiAuto;
        } else {
            engine = Transmission::Manual;
            // ADD <K, V> pair to hash map
            orders.insert("Manual".to_string(), manual);
            manual = manual + 1;
        }

        // ADD hash map functionality
        // Order the cars, New are even numbers, Used are odd numbers
        // Corrected code: Index into `colors` array, vary color for the orders
        // TO DO: Fix syntax to add car age to "orders" hash map
        if index % 2 != 0 {
            car = car_factory(colors[index-1].to_string(), engine, roof, miles);
            // ADD <K, V> pair to hash map
            orders.insert("Used".to_string(), used_cars);
            used_cars = used_cars + 1;
        } else {
            car = car_factory(colors[index-1].to_string(), engine, roof, 0);
            // ADD <K, V> pair to hash map
            orders.insert("New".to_string(), new_cars);
            new_cars = new_cars + 1;
        }

        // Display car order details by roof type and age of car
        // TO DO: Add conditional expressions
        // TO DO: Print output based on four conditions, correct the syntax
        // TO DO: Used & closed roof, New & closed roof, Used convertible, New convertible

        //if used cars with closed roofs {
        if car.age.1 > 0 && !car.roof {
            println!("{}: {}, {:?}, Closed roof, {}, {} miles", order, car.age.0, car.motor, car.color, car.age.1);
        } else if car.age.1 == 0 && !car.roof {
        //} if new cars with closed roofs {
            println!("{}: {}, {:?}, Closed roof, {}", order, car.age.0, car.motor, car.color);
        } else if car.age.1 > 0 {
            //} if convertible used cars {
            println!("{}: {}, {:?}, Convertible, {}, {} miles", order, car.age.0, car.motor, car.color, car.age.1);
        //} if convertible new cars {
        } else if car.age.1 == 0 {
            println!("{}: {}, {:?}, Convertible, {}", order, car.age.0, car.motor, car.color);
        }

        // Change values for next loop
        // TO DO: Increment "order" by 1, and "miles" by 1,000
        order += 1;
        miles += 1000;

        // Adjust the index for the car details
        // Order 11 cars, use index range of 0 -- 4, then repeat from 0
        if index < 4 {
            index = index + 1;
        } else {
            index = 1;
        }
    }

    // TO DO: Display output from hash map, fix the syntax
     // Display the hash map of car orders, show <K, V> pairs
     println!("\nCar orders: {:?} {:?}", orders.keys(), orders.values());
}