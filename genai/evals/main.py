import os

import pandas as pd
import duckdb
from dotenv import load_dotenv
from mistralai.client import Mistral

load_dotenv()


client = Mistral(api_key=os.environ["MISTRAL_API_KEY"])
MODEL = "mistral-large-latest"


TRANSACTION_DATA_FILE_PATH = "Store_Sales_Price_Elasticity_Promotions_Data.parquet"

SQL_GENERATION_PROMPT = """
Generate an SQL query based on a prompt. Do not reply with anything besides the SQL query.
The prompt is: {prompt}

The available columns are: {columns}
The table name is: {table_name}
"""


def generate_sql_query(prompt: str, columns: list, table_name: str) -> str:
    formatted_prompt = SQL_GENERATION_PROMPT.format(
        prompt=prompt,
        columns=columns,
        table_name=table_name
    )

    response = client.chat.complete(
        model=MODEL,
        messages=[
            {"role": "user", "content": formatted_prompt}
        ]
    )

    return response.choices[0].message.content


def lookup_sales_data(prompt: str) -> str:
    try:
        table_name = "sales"

        df = pd.read_parquet(TRANSACTION_DATA_FILE_PATH)
        duckdb.sql(f"CREATE TABLE IF NOT EXISTS {table_name} AS SELECT * FROM df")

        sql_query = generate_sql_query(prompt, df.columns, table_name)
        sql_query = sql_query.strip()
        sql_query = sql_query.replace("```sql", "").replace("```", "")

        result = duckdb.sql(sql_query)

        return result.to_csv()
    except Exception as e:
        return f"Error accessing data {str(e)}"

# response = mistral_client.messages.create(
#     model=model,
#     max_tokens=1024,
#     messages=[{"role": "user", "content": prompt}],
# )
if __name__ == "__main__":
    example_data = lookup_sales_data("Show me all the sales for store 1320 on November 1st, 2021")
    print(example_data)
