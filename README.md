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