import os

from dotenv import load_dotenv
from mistralai.client import Mistral

load_dotenv()


def main():
    print("Hello from evals!")


client = Mistral(api_key=os.environ["MISTRAL_API_KEY"])
model = "mistral-large-latest"


TRANSACTION_DATA_FILE_PATH = "Store_Sales_Price_Elasticity_Promotions_Data.parquet"

SQL_GENERATION_PROMPT = """
Generate an SQL query based on a prompt. Do not reply with anything besides the SQL query.
The prompt is: {prompt}

The available columns are: {columns}
The table name is: {table_name}
"""


def generate_sql_query(prompt: str, columns: list, table_name: str) -> str:
    formatted_prompt = SQL_GENERATION_PROMPT.format(prompt, columns, table_name)
    # response = client.c
    return ""


# response = mistral_client.messages.create(
#     model=model,
#     max_tokens=1024,
#     messages=[{"role": "user", "content": prompt}],
# )
if __name__ == "__main__":
    main()
