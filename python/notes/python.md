Organize python code

**project_name**/
```run.py / main.py ```<-  entry point of our project
processing/
untils/ _ _ init _ _ .py

#modules
any file that has .py like run.py chai.py

#packages
any file that has __ init __ .py

PEP 8 *Guidline to write python code*

**Data Types**
- Objects
	- objects have identity, type and value
	- objects can be immutable(not changeable) or mutable(changeable)
- Numbers
	- when we divide a number and round it off to integer we do //
	  like  7/4 = 1.75 
	  but 7//4 = 1
	  - 2^3 is calculated as 2 ** 3
- Strings
	-  [0:8:2] 0 starting index, included, 8 ending index excluded every second character is skipped
	- [::-1] will reverse the entire siting
- Tuples
	-  reprented with **( )** ;
	- destructureing is possible
	  - nums = (1,2,3,4,5)
	  - (num1,num2,num3,num4,num5) = nums;
	- memebership
	  1 in nums -> returns True (case sensitive in terms of strings)
- List(also called as Array)
	-  nums = [1,2,3,4,5]
	- nums.append(9) //inserts at the last index
	- nums.remove(3) //removes the parameter from the list/array
	  
	  nums2= [11,111,11111,1111];
	- nums.extend(nums2) // [1,2,3,4,5,11,111 ...  ..  . . .]
	- nums.insert(2,69) //inserts 69 at the second pos
	- .pop(removes the last element)
	-   .reverse() the list
	- .sort() in ascending
- Set
	- set of all unique elements
	- represneted with **{ }**
	- intersection is represned by **&**
	- union is represned by **|**
	- present in A but not in B is *A **-** B*
- Dictionary
	- it is like map or obejcts
	- dicttionary_example = dict(type="cgpa", score=10.1, college="pesce")
	- example2 = {}
	  example2["cgpa"] = 9
	- **del** example["cgpa"]
	- keys() -> returns all the keys
	- .values() returns all the values
	- .items returns the set of key value pairs
	  .popitem will remove the item(key and value both)
	- if we have lets say another dictionary we can just
	  example.update(newExample)
	  
- conditions 
  if elif
  swtich case -> match variable_name
  
  functions
  using def
  def print_order(name, number):
	  print(f"{name} is {number}")
	  
	scopes
	nonlocal keyname -> refres to the varibale just outside the scope
	example
		def numb():
			n = 11
			def nuuuum():
				nonlocal n
				
- args and kwargs
  def special_chai(*ingredients, **extras):
	- print("Ingredients", ingredients)
	  print("extras", extras)
	special_chai("Cinnamon", "cardmom", sweetener="Honey", foam="yes")
	
	output:
	Ingredients ('cinnamon', 'cardmo') //tuple
	Extras {'sweetner': 'honey', 'foam': 'yes'} //dictionary
	
	lambdas and anonymous functions are same
	
	chai_types = ["light", "kadak", "ginger", "kadak"]
	
	 strong-chai = list(filter(lambda chai: chai == "kadak", char_types))
	 
	 _ _ dunder _ _ 
	 
	 - comprehensions
	   they are a concise way to create lists, sets, dictionaries or generators in python using a single line of code
	   where are they used
		 - filter items
		 - transform item
		 - create a new collection
		 - flatten nested structure
		- list comprehensions
			code: 
			```python.py
			   menu = ["Masal chai", 
						"Iced Lemon tea",
						"green Tea",
						"Iced peach tea",
						"Ginger chai"
						]
						
						iced_tea = [tea for tea in menu if "Iced" in tea]
			```
			
		- Generators
		  why 
		  - save memory
		  - don't want results immediately
		  - lazy evaluations
		    
		    ``` python.py
		    def get_chai_gen():
			    yield "Cup 1"
			    yield "Cup 2"
			    yield "Cup 3"
			chai = get_chai_gen # this will have refrence of the function
			
			print(next(chai)) # prints "Cup1"
			print(next(chai)) # prints "Cup2"
			
		    ```
		    - infinite generators
		      ```python
		      def infinite_chai():
			      count=0
			      while True:
				      yield f"Refill Count {count}"
				      count+=1
				
			  refill= infinite_chai()
			  
			  for _ in range(3)
				  print(next(refill))
			 #refill will be printed only 3 times
				      
		      
		      ```
		    send yield to generator
		    ```python.py
		    def chai_consumer():
			    print("Welcome! what chai would you like ?")
			    order = yield
			    while True:
				    print(f"Preparing : {order}")
				    order = yield # this will technically stops your program
				    
				stall = chai_customer()
				next(stall) #start the generator and this will print Welcome! what chai would you like ?
				stall.send("Masala Chai") #Preparing: masala Chai
				stall.send("Lemon chai")#preparing: lemon chai
		    ```  
		    yeild from and close the generators
		    
		    ```python.py
		    def local_chai():
			    yield "Masala Chai"
			    yield "Ginger Chai"
			    
			def imported_chai()
				yield "machha"
				yield "Ooolong"
				
			def full_menu():
				yield from local_chai()
				yield from imported_chai()
				
			for chai in full_menu():
				print(chai) #prints masala chai,Ginger Chai, machha, Ooolong	
		    ```
		    Object oriented programming in python
		    - it is a paradigm of coding
		    - it means a way to write the code
		    ```Python.py
		    class Chai:
			    origin = "India"
			print(Chai.origin)
			Chai.is_hot = True //adding properties to the class
			print(Chai.is_hot)
		    ```
		    Attribute shadowing..
		    ```python
		    class Chai:
			    temperature = "hot"
			    strength = "Strong"
			cutting = Chai()
			print(cutting.temperature)
			
			cutting.temperature = "Mild"
			print("After changing", cutting.temperature) #mild
			print("Direct look into the class", Chai.temperature) #hot
			
			del cutting.temperature
			print(cutting.temperature) #hot
		    ```
		    - self arguments
		      ```python
		      class Chaicup:
			      size = 150
			      
			      def describe(self)
				      return f"A {self.size}ml cup of chai"
			  
			  cup = Chaicup()
			  print(cup.describe()) # A 150ml cup of chai
			  print(Cahicup.describe()) # error
		      ```
		      - constructors (_____init __)


			```python
			class ChaiOrder:
				def __init__(self, type_, size):
					self.type = type_ #varaible value
					self.size = size # variable value
					
				def summery(self):
					return f"{self.size}ml of {self.type} chai"
					
					order = ChaiOrder("Masala", 200)
					print(order.summery())
					
					order_two = ChaiOrder("Ginger", 220)
					print(order_two.summery())
			```
		        
		        inheritance
		        
		        mro - method resultion order
		        class A:
					    label = "A: base Class"
				class B(A):
					label = "B: Masala Blend"
				class C(A):
					label = "c: Herbal blend"
				class D(B,C):
					pass
					
				cup = D()
				print(cup.label) /#prints B: Masal Blend


```Python

class ChaiUtils:
	@staicmethod
	def clean_ingredients(text):
		return [item.strip() for item in text.split(",")]
		
		raw = "water , milk, giner  ,  honey"
		
		cleaned = ChaiUtils.clean_ingredients(raw) #using class instead of making an object refrence
		print(cleaned) # ['water', 'milk', 'ginger', 'honey']
```

classmethod:
```python

class ChaiOrder:
	def __init__(self, tea_type, sweetness, size):
		self.tea_type = tea_type
		self.sweetness = sweetness
		self.size = size
		
	@classmethod
	def from_dict(cls, order_data):
		return cls(
			order_data["tea_type"],
			order_data["sweetness"],
			order_data["size"]
		) #cls here is the reference of the class (here ChaiOrder)
		
	@classmethod
	def from_string(cls, order_string)
		tea_type, sweetness, size = order_string.split("-")
		return cls(tea_type, sweetness, size)
		
order1 = ChaiOrder.from_dict({"tea_type": "masala", "sweetness": "medium", "size": "Large"})

order2 = ChaiOrder.from_string("Ginger-Low-Small")

print(order1) # <__main__.ChaiOrder object at 0x6172617618>
print(order1.__dict__) #this will print the dictionary
```
property decorator

```python
class teaLeaf:
	def __init__(self, age):
		self._age = age 
		
		@property
		def age(self);
			return self._age+2
			
		@age.setter
		def age(self, age):
			if 1<= age<=5:
				self._age = age
			else:
				raise ValueError("tea leaf age must be between 1 and 5")
				
leaf = Tealeaf(2)
print(leaf.age)
```
- excpetion handling
  
  custom error
  
  ```python.py
  class OutOfIngredientsError(Exception):
	  pass
	def make_chai(milk, sugar):
		if milk == 0 or sugar == 0:
			raise OutOfIngredientsError("Missing milk or sugar")
		print("chai is ready...")
  ```
  mini project with exception handling
  ```python
  class InvalidChaiError(Exception): pass
  def bill(flavor, cups):
	  menu = {"masala": 20, "ginger": 40}
	  try:
		  if flavor not in menu:
			  raise InvalidChaiError("that chai is not avaliable")
		  if not isinstance(cups, int):
			  raise TypeError("Number of cups must be an integer")
		  total = menu[flavor]*cups
		  print(f"Your bill for {cups} of {flavor} chai rupees {total}" )
	   except Exception as e:
		   print("Error: ", e)
		finally:
			print("Thank you for visiting chaicode!")
			
bill
  ```
**Threads and concurrency in python**
- concurrency and parallelism
**concurrency**
```python
import threading
import time

def take_orders():
	for i in range(1,4):
		print(f"Taking order for #{i}")
		time.sleep(1)
def brew_chai():
	for i in range(1,4):
		print(f"Brewing chai for #{i}")
		time.sleep(2)

#create threads

order_thread=threading.Thread(target=take_oders)
brew_thread = threading.Thread(target=brew_chai)

order_thread.start()
brew_thread.start()

#wait for both to finish
order_thread.join()
brew_thread.join()

print(f"All orders taken and chai brewed")
```

**parallelism**
```python
from multiprocessing import Process
import time

def brew_chai(name):
	print(f"start {name} chai brewing")
	time.sleep(3)
	print(f"end of {name} chai brewing")
	
if __name__ == "__main__":
	chai_makers = [
	Process(target=brew_chai, args=(f"Chai Maker #{i+1}"))
	for i in range(3)
	]
	
	#Start all process
	for p in chai_makers:
		p.start()
	
	#wait for all to complete
	for p in chai_makers:
		p.join()	
	
	print("all chai served")

```
**GIL**(global)
```python
import threading
import time
def brew_chai():
	print(f"{threading.current_thread().name} started brewing...")
	count = 0
	for _ in range(100_000_000):
		count+=1
	print(f"{threading.current_thread().name} started brewing...")

thread1 = threading.Thread(target=brew_chai, name="Barista-1")
thread2 = threading.Thread(target=brew_chai, name="barista-2")

start = time.time()
thread1.start()
thread2.start()
thread1.join()
thread2.join()
end=time.time()

print(f"total time taken: {end-start:.2f} seconds")
```

Async, Event loop, coroutines and await in python

Process Pool executor with aysnc io
```python
import aysncio
import concurrent.futures import ProcessPoolExecutor

def encrypt(data):
	return f"🔒 {data[::-1]}"
async def main():
	loo=asyncio.get_running_loop()
	with ProcessPoolexecutor() as pool
		result = await loop.run_in_executor(pool, encrypt, "credit_card_1234")
		print(f"{result}")
		
if __name__ = "__main__":
	asyncio.run(main())
```
background worker
```python
import asyncio
import threading
import time

def background_worker():
		while True:
			time.sleep(1)
			print(f"Logging the system health")
aysnc def fetch_orders():
	await asyncio.sleep(3)
	print("order fetched")

threading.thread(target=background_worker, daemon=True).start()

asyncio.run(fetch_orders())
```

pydantic
used for
- data validation
- setting management
```python
from pydantic import BaseModel

class User(BaseModel):
	id: int
	name: str
	is_active: bool
	
input_data = {'id':101, 'name': 'Chaicode', 'is_active': True}

user = User(**input_data)
print(user)
```

pydantic default conversions:
```python
from pydantic import BaseModel

class Product(BaseModel):
	id: int
	name: str
	price: float
	in_stock: bool = True
	
product_one = Product(id=1, name="gajendra", price=24.2)

```

Mixing pydantic and typing

```python
from pydantic import BaseModel
from typing import List, Dict, Optional

class Cart(BaseModel):
	user_id: int
	items: List[str]
	quantities: Dict[str, int]
	
class BlogPost(BaseModel):
	title: str
	content: str
	image: Optional[str] = None
```

adding validations using field

```python
from typing import Optional
from pydantic import BaseModel, Field

class Employee(BaseModel):
	id: int
	name: str = Field(
		..., #required
		min_length=3,
		max_length=50,
		description="Employee Name",
		example="Hitesh Choudhary"
	)
	department: Optional[str] = 'General'
	salary: float = Field(
		...,
		ge=10000
	)
```

field and model validator
```python
from pydantic import baseModel, field_validator

class user(BaseModel):
	username: str
	
	@field_validator('username')
	def username_length(cls, v):
		if len(v) < 4:
			raise ValueError("username must be at least 4 characters")
		return v
	

class signupdata(BaseModel):
	password: str,
	confirm_password: str	
	@model_validator(mode='after')
	def passwordmatch(cls, values):
		if values.password != values.confirm_password:
			raise ValueError("Password do not match")
		return values
```

computed property
```python
from pydantic import BaseModel, computer_field

class Product(BaseModel):
	price: float
	quantity: int
	
	@computed_field
	@property
	def total_price(self) -> float:
		return self.price * self.quantity

class Booking(BaseModel):
	user_id: int
	room_id: int
	nights: int = Field(.., ge=1)
	rate_per_night: float
	
	@computed_field
	@property
	def total_amount(self)->float:
		return self.nights*self.rate_per_night
```

filed validator
```python

class Product(BaseModel);
	price: str
	
	@field_validator('price', mode='before')
	def
```

nested model in pydantic
```python
from typing import List, Optional
from pydantic import BaseModel

class Address(BaseModel):
	street: str
	city: str
	postal_code: str

class User(BaseModel):
	id: int
	name: str
	address: Address
	
address = Address(
	street="123 something",
	city="Jaipur",
	postal_code="10001"
)

user = User(
	id=1
	name="Hitesh"
	address=address	
)
```

self-referencing/ recursive model

```python
from typing import List, Optional
from pydantic import BaseModel

class Comment(BaseModel):
	id: int
	content: str
	replies: Optional[List['Comment']] = None
	
Comment.model_rebuild() #required if we use recursive/self-referensing model

comment = Comment(
	id: 1
	content="First comment"
	replies=[
		Comment(id=2, content="reply1"),
		Comment(id=3, content="reply2"),
	]
)
```

best practices 

model organization
- Define leaf nodes first - Models with no dependencies
- Build Upwards - Gradually compose more complex models
- user clear naming - Make relationships obvious
- group related models - keep models in logical modules
performance consideration
- Deep nesting impacts performance - keep reasonable depth
- large lists of nested models - Consider pagination
- Circular references - Use carefully, can cause memory issues
- Lazy loading - Consider for expensive nested computations
data modelling tips
- Model real-world relationship - Mirror your domain structure
- use optional appropriately - Not all relationships are required
- consider union types - for polymorphic relationships
- validate business rules - use model validators for cross model logic


Pydantic seralization
```python
from pydantic import baseModel, ConfigDict
from typing import List
from datetime import datetime

class Address(BaseModel):
	street: str
	city: str
	zip_code: str
	
class User(BaseModel):
	id: int
	name: str
	email: str
	is_active: bool = True
	createdAt = datetime
	address: Address
	tags: List[str] = []
	
	model_config = ConfigDict(
		json_encoders = {datetime: lambda v: v.strftime('%d-%m-%Y %H:%M:%S')}
	)
	
user = User(
	id=1,
	name="Hitesh",
	email="h@hitesh.ai",
	createdAt=datetime(2024,3,15,14,30)
	address=Address(
		street="Something",
		city="Jaipur",
		zip="52631"
	),
	is_active=False,
	tags=["premium", "subscriber"]	
)

python_dict = user.model_dump() #converts everything to dictionary not nested shit
print(python_dict)
```

decoraators in python

using decorators you can add some featuers in the existing functions
```python
def div(a,b):
	print(a/b)
#you cannot modify the above function but you need to write some logic so that numerator should always be greater than denominator 
	
	
def smart_div(func):
	
	
	def inner(a,b):
		if a<b:
			a,b = b,a
		return func(a,b)
	return inner
	
#we are calling div function but in a smarter way..

 dev = smart_div(div)
 
 div(2,4)
```