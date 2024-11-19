# Function signature

`func (studentHandler *StudentHandler) GetStudents(requestContext *fiber.Ctx)`

**(studentHandler \*StudentHandler)**

- The receiver of the method
- Meaning - This function is a method that belongs to any instance of the StudentHandler struct
- Inside this method you can access the fields and other methods of StudentHandler via the studentHandler variable

  **(requestContext \*fiber.Ctx)**

- A pointer to an instance of the fiber.Ctx struct
- Which represents the context of the HTTP request
- This context includes information about the request and response, and provides methods for interacting with them
