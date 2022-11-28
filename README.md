# Shopping cart checkout technical challenge 

### Description

This program emulates a supermarket that scans items in and calculates total prices correctly for any combination of the products and offers.

### Set up and how to run

- To set up this project, clone this repository and run `go mod download` to download the dependencies.
- To run an example solution, run `go run main.go` in the root of the project. 
- To modify the example solution/run your own checkout, go to `main.go` and amend the variables to set it to the products, prices, discount etc... that you'd like to test.

### Testing
- To run the unit tests, navigate to the internal directory `cd internal` and run `go test`

## Planning

### Initial plan

I approached this test by reading through the instructions and making note of any questions I had.
I then created tables to decide how many classes, instance variables and methods I would require.
I also made a note of what tests I would expect to write, which in this case was just unit testing.

** Note these tables include setter functions, however I did not actually require it due to my choice to use Golang, as the fields of a struct can be set when the entity is initialised. 
I also would not require changing them once set for this exercise.


| Class        | Methods                    | Instance variables |
| :---         |     :---:                  |        :---:       |
| Product      | **setProductCode           | ID                 |
|              | setName                    | Product code       |
|              | setPrice                   | Name               |
|              |                            | Price              |

| Class        | Methods                    | Instance variables |
| :---         |     :---:                  |        :---:       |
| Checkout     | Scan                       | Discounts          |
|              | calculateTotal             | Basket             |
|              | calculateDiscount          | Price              |
|              | calculateTotalWithDiscount |                    |

| Class        | Methods                    | Instance variables |
| :---         |     :---:                  |        :---:       |
| Discount     | getDiscount                | Discount Type      |
|              | addProduct                 | Products           |
|              | buyOneGetOneFree           | ID                 |
|              | bulkDiscount               | Percentage         |
|              |                            | Quantity           |

### Changes to initial plan

When I started to work on the discount entity, I realised that I would need an easy way to gain information about how many quantities of a product where scanned. 
This is when I decided to create `Basket` and `Item` entities that would hold information about how many quantities of a product had been scanned. 
It would also mean a product would not be represented more than once in a list of products scanned making it more performant to iterate over and easier to find information on quantities to apply discounts.

| Class           |
| :---            |
| Basket (a list of items)|

| Class        | Methods                    | Instance variables |
| :---         |     :---:                  |        :---:       |
| Item         | increaseQuantity           | Product            |
|              |                            | Quantity           |

With the opportunity, I would've liked to clarity the following:
- The behaviour when the total is calculated. For now, I have printed a message with the final total, but perhaps it would be useful to print the overall total, discount received and then the final total
- Whether a supermarket entity was required. I did consider it when planning classes, however as the instructions focused on the behaviour of the checkout, I felt it was not required at this stage and for the test overall.


## Reflections

Overall, I am happy with the solution I have come up with. I think the classes are well encapsulated, and I have adhered to the single responsibility principle as much as possible, keeping methods short and focused on a single purpose.

Nevertheless, there were things I would like to improve with more time

- **Improve precision when calculating total and truncation and rounding.** Currently, I do not have any code in place to take into account what to do if a value needs to be rounded up or down. I think this is sufficient for the solution required for this technical test. However, if more complex division operations were performed, this means that the final total presented could be incorrect. I would consider using the math package in the future and including more tests to check that more complex operations are correctly handled.
- **Validate product codes/discount types.** I used constants to represent product codes and discount types to ensure they could be reused and minimise the risk of creating a product/discount with an incorrect code/type. However, as I did not employ setter methods which could confirm the correct type was being used, my current code is vulnerable to user error. 
- **Improve success/error handling.** Currently, there is not much error handling in the program but there are some success messages. I would've liked to clarify what success/ error handling would be required and the desired behaviour e.g. logging error messages. I think error handling could particularly be improved for mathematical operations and incorrect user behaviour e.g. trying to check out with an empty basket or scan a product that does not exist.
- **Improve discount handling.** I would like to have taken more time to flesh out discount behaviour for both the BOGOF discount and the bulk discounting. For example, with BOGOF, there is an assumption that a user has the correct number of items in the basket, and it will calculate the discount for the nearest even number (rounding down). But perhaps, it would be best to display a message to let the user know they can get another item.


## Future considerations

Other potentially useful features to include could be:
- A front end checkout user interface where users can scan products, view their basket and checkout.
- Allowing users to scan more than one quantity of a product at a time. This could be achieved expanding the scan method to also pass in quantity with product
- A supermarket/store interface where users can view products. Currently, there is no entity that holds information about all the products available to the user. A storefront/supermarket would be a good place to store this information.