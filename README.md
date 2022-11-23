### S.O.L.I.D. Design Principles
The SOLID principles were introduced by *Robert C. Martin* (also known as Uncle Bob) in his 2000 paper *“Design Principles and Design Patterns*.” These concepts were later built upon by Michael Feathers, who introduced us to the SOLID acronym. And in the last 20 years, these five principles have revolutionized the world of object-oriented programming, changing the way that we write software.

SOLID principle is an acronym that stands for 5 key principles in software development: 

S – Single Responsibility Principle 

O – Open Closed Principle 

L – Liskov Substitution Principle

I – Interface Segregation Principle

D – Dependency Inversion Principle

#### S – Single Responsibility Principle 
The single responsibility principle states that a class should only have one responsibility. Furthermore, it should only have one reason to change.

**Benefits:**
1. **Testing** – A class with one responsibility will have far fewer test cases.
2. **Lower coupling** – Less functionality in a single class will have fewer dependencies.
3. **Organization** – Smaller, well-organized classes are easier to search than monolithic ones.

#### O – Open Closed Principle (OCP)
The Open Closed Principle(OCP) states that classes should be open for extension but closed for modification. “*Open to extension*” means that you should design your classes so that new functionality can be added as new requirements are generated. “*Closed for modification*” means that once you have developed a class you should never modify it, except to correct bugs.
In doing so, we stop ourselves from modifying existing code and causing potential new bugs in an otherwise happy application.
Of course, the one exception to the rule is when fixing bugs in existing code.

**Benefits:**
1. **Easier extensibility** - Offers a better extensibility to your code.
2. **Easier to maintain** - It suggests using interfaces. Interfaces in code offers an additional level of abstraction which in turn enables loose coupling. The implementations of an interface are independent of each other and don’t need to share any code.
3. **Flexibility** - if any change request arises in future, your code will be more flexible to extend.

#### L – Liskov Substitution Principle
Introduced by Barbara Liskov, the Liskov Substitution Principle (LSP) states that an object of a superclass should be replaceable by objects of its subclasses without causing issues in the code. Therefore, a child class should never change the characteristics of its parent class (such as the argument list and return types). Basically, derived classes should never do less than their base class.

**“Derived types must be completely substitutable for their base types”**

All subclasses must, therefore, operate in the same manner as their base classes. The specific functionality of the subclass may be different but must conform to the expected behavior of the base class. To be a true behavioral subtype, the subclass must not only implement the base class’s methods and properties, but also conform to its implied behavior.

**Benefits:**
1. Code reusability
2. Easier to maintain
3. Lower coupling

#### I – Interface Segregation Principle
ISP guides us to create multiple, smaller, cohesive interfaces. Larger interfaces should be split into smaller ones. By doing so, we can ensure that implementing classes only needs to be concerned about the methods that are of interest to them. So basically, the interface segregation principles as you prefer the interfaces, which are small but client specific instead of monolithic and bigger interface.

**“Don’t depend on things you don’t need”**

Single Responsibility Principle is concerned with classes, while the Interface Segregation Principle is concerned with interfaces. 

**Benefits:**
1. Increased code readability
2. Easy to implement and maintain
3. Better organization of code
4. Don’t need to throw exceptions unnecessarily

#### References
1. [https://www.baeldung.com/solid-principles](https://www.baeldung.com/solid-principles)
2. [https://www.javatpoint.com/solid-principles-java](https://www.javatpoint.com/solid-principles-java)
3. [https://www.jrebel.com/blog/solid-principles-in-java](https://www.jrebel.com/blog/solid-principles-in-java)
4. [https://javatechonline.com/solid-principles-the-liskov-substitution-principle/](https://javatechonline.com/solid-principles-the-liskov-substitution-principle/)