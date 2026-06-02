import json

from pydantic import BaseModel, ValidationError, EmailStr, Field
from typing import Optional
from datetime import date


class UserInput(BaseModel):
    name: str
    email: EmailStr
    query: str
    order_id: Optional[int] = Field(
        None,  # default value
        description="5-digit order number (cannot start with 0)",
        ge=10000,
        le=99999
    )
    purchase_date: Optional[date] = None

# try:
#     user_input = UserInput(
#         name="Joe",
#         email="jo@eeample.com",
#         query="I forgot my password"
#     )
#     print(user_input)
# except ValidationError as e:
#     print(e)


def validate_user_input(input_data):
    try:
        user_input = UserInput(**input_data)
        print("Valid user input data")
        print(f"{user_input.model_dump_json(indent=2)}")
        return user_input
    except ValidationError as e:
        print("Validation error occurred")
        for error in e.errors():
            print(f"{error['loc'][0]}: {error['msg']}")
        return None

# input_data = {
#     "name": "Joe",
#     "email": "joe@example.com",
#     "query": "forgot my password",
#     "order_id": 12345,
#     "purchase_date": "2026-04-17",
#     "system_message": "logging status", # additional field
#     "iteration": 1 # additional field
# }
# response = validate_user_input(input_data)
# print(response)

json_data = '''
{
    "name": "Joe",
    "email": "joe@example.com",
    "query": "forgot my password",
    "order_id": "012345",
    "system_message": "logging status"
}
'''
try:
    user_input = UserInput.model_validate_json(json_data)
    print(user_input.model_dump_json(indent=2))
except ValidationError as e:
    print("Validation error occurred")
    for error in e.errors():
        print(f"{error['loc'][0]}: {error['msg']}")
