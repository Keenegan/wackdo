```mermaid
classDiagram

    class Employee {
        id : int
        name: String
    }

    class Role {
        id : int
        name: String
    }

    class Order {
        id : int
        createdAt : Date
        status : OrderStatus
        +getTotalPrice(): float
    }

    class OrderStatus {
        <<enumeration>>
        CREATED
        PAID
        CANCELLED
        DELIVERED
    }

    class OrderLine {
        id : int
        quantity : int
        unitPrice : float
    }

    class SellableItem {
        <<abstract>>
        id : int
        name : String
        basePrice : float
    }

    class Product {
        category: ProductCategory
    }

    class ProductCategory {
        <<enumeration>>
        DRINK
        FOOD
    }

    class Menu {
        description: String
    }

    class MenuOption {
        id: int
        price:  float
        name: String
    }

    SellableItem <|-- Product
    SellableItem <|-- Menu

    Order "1" *-- "1..*" OrderLine : contains
    OrderLine "1" --> "1" SellableItem : refersTo
    Employee "1..*" --> "1..*" Role : has
    Order "1" --> "1" Employee : has
    Menu "1" --> "1..*" Product : composedOf
    Menu "1" --> "0..*" MenuOption: has
    ```
    ```