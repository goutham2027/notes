import json
import os

import duckdb
import pandas as pd
from dotenv import load_dotenv
from mistralai.client import Mistral
from pydantic import BaseModel, Field

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
        prompt=prompt, columns=columns, table_name=table_name
    )

    response = client.chat.complete(
        model=MODEL, messages=[{"role": "user", "content": formatted_prompt}]
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


DATA_ANALYSIS_PROMPT = """
Analyze the following data: {data}
Your job is to answer the following question: {prompt}
"""


def analyze_sales_data(prompt: str, data: str) -> str:
    formatted_prompt = DATA_ANALYSIS_PROMPT.format(data=data, prompt=prompt)

    response = client.chat.complete(
        model=MODEL, messages=[{"role": "user", "content": formatted_prompt}]
    )

    analysis = response.choices[0].message.content

    return analysis if analysis else "No analysis could be generated"


CHART_CONFIGURATION_PROMPT = """
Generate a chart configuration based on this data: {data}
The goal is to show: {visualization_goal}
"""


class VisualizationConfig(BaseModel):
    chart_type: str = Field(..., description="Type of chart to generate")
    x_axis: str = Field(..., description="Name of the x-axis column")
    y_axis: str = Field(..., description="Name of the y-axis column")
    title: str = Field(..., description="Title of the chart")


def extract_chart_config(data: str, visualization_goal: str) -> dict:
    formatted_prompt = CHART_CONFIGURATION_PROMPT.format(
        data=data, visualization_goal=visualization_goal
    )

    response = client.chat.parse(
        model=MODEL,
        messages=[{"role": "user", "content": formatted_prompt}],
        response_format=VisualizationConfig,
    )

    try:
        content = response.choices[0].message.parsed
        return {
            "chart_type": content.chart_type,
            "x_axis": content.x_axis,
            "y_axis": content.y_axis,
            "title": content.title,
            "data": data,
        }
    except Exception:
        return {
            "chart_type": "line",
            "x_axis": "date",
            "y_axis": "value",
            "title": visualization_goal,
            "data": data,
        }


CREATE_CHART_PROMPT = """
Write python code to create a chart based on the following configuration.
Only return the code, no other text.
config: {config}
"""


def create_chart(config: dict) -> str:
    formatted_prompt = CREATE_CHART_PROMPT.format(config=config)

    response = client.chat.complete(
        model=MODEL, messages=[{"role": "user", "content": formatted_prompt}]
    )

    code = response.choices[0].message.content
    code = code.replace("```python", "").replace("```", "")
    code = code.strip()

    return code


def generate_visualization(data: str, visualization_goal: str) -> str:
    config = extract_chart_config(data, visualization_goal)
    code = create_chart(config)
    return code


tools = [
    {
        "type": "function",
        "function": {
            "name": "lookup_sales_data",
            "description": "Look up data from Store Sales Price Elasticity Promotions dataset",
            "parameters": {
                "type": "object",
                "properties": {
                    "prompt": {
                        "type": "string",
                        "description": "The unchanged prompt that the user provided.",
                    }
                },
                "required": ["prompt"],
            },
        },
    },
    {
        "type": "function",
        "function": {
            "name": "analyze_sales_data",
            "description": "Analyze sales data to extract insights",
            "parameters": {
                "type": "object",
                "properties": {
                    "data": {
                        "type": "string",
                        "description": "The lookup_sales_data tool's output.",
                    },
                    "prompt": {
                        "type": "string",
                        "description": "The unchanged prompt that the user provided.",
                    },
                },
                "required": ["data", "prompt"],
            },
        },
    },
    {
        "type": "function",
        "function": {
            "name": "generate_visualization",
            "description": "Generate Python code to create data visualizations",
            "parameters": {
                "type": "object",
                "properties": {
                    "data": {
                        "type": "string",
                        "description": "The lookup_sales_data tool's output.",
                    },
                    "visualization_goal": {
                        "type": "string",
                        "description": "The goal of the visualization.",
                    },
                },
                "required": ["data", "visualization_goal"],
            },
        },
    },
]

tool_implementations = {
    "lookup_sales_data": lookup_sales_data,
    "analyze_sales_data": analyze_sales_data,
    "generate_visualization": generate_visualization,
}


SYSTEM_PROMPT = """
You are a helpful assistant that can answer questions about the Store Sales Price Elasticity Promptions dataset.
Use available tools to answer the user.
"""


def handle_tool_calls(tool_calls, messages):
    for tool_call in tool_calls:
        function = tool_implementations[tool_call.function.name]
        function_args = json.loads(tool_call.function.arguments)
        result = function(**function_args)
        messages.append(
            {"role": "tool", "content": result, "tool_call_id": tool_call.id}
        )
    return messages


def run_agent(messages):
    print("Running agent with messages:", messages)

    if isinstance(messages, str):
        messages = [{"role": "user", "content": messages}]

    if not any(
        isinstance(message, dict) and message.get("role") == "system"
        for message in messages
    ):
        system_prompt = {"role": "system", "content": SYSTEM_PROMPT}
        messages.append(system_prompt)

    while True:
        print("Making router call to Mistral")
        response = client.chat.complete(
            model=MODEL, messages=messages, tools=tools, tool_choice="auto"
        )
        messages.append(response.choices[0].message)
        # print(messages)
        tool_calls = response.choices[0].message.tool_calls
        print("Received response with tool calls:", bool(tool_calls))

        if tool_calls:
            print("Processing tool calls")
            messages = handle_tool_calls(tool_calls, messages)
        else:
            print("No tool calls, returning final response")
            return response.choices[0].message.content


if __name__ == "__main__":
    # example_data = lookup_sales_data(
    #     "Show me all the sales for store 1320 on November 1st, 2021"
    # )
    # print(example_data)
    # print(
    #     analyze_sales_data(
    #         prompt="what trends do you see in this data", data=example_data
    #     )
    # )

    # code = generate_visualization(
    #     example_data,
    #     "A bar chart of sales by product SKU. Put the product SKU on the x-axis and the sales on the y-axis.",
    # )
    # print(code)

    result = run_agent(
        "Show me the code for graph of sales by store in Nov 2021, and tell me what trends you see."
    )
    print(result)
