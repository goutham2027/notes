import base64
import os


from huggingface_hub import InferenceClient
from smolagents import (
    CodeAgent,
    tool,
    DuckDuckGoSearchTool,
    InferenceClientModel,
)
from model import model
import trace_setup  # noqa

# set HF_TOKEN
api_key = os.environ["HF_TOKEN"]

# client = InferenceClient(
#     model="meta-llama/Meta-Llama-3-8B-Instruct", provider="novita", api_key=api_key
# )

# output = client.chat.completions.create(
#     messages=[
#         {"role": "user", "content": "The capital of France is"},
#     ],
#     stream=False,
#     max_tokens=1024,
# )

# print(output.choices[0].message.content)


@tool
def suggest_menu(occassion: str) -> str:
    """
    Suggests a menu based on the occassion.
    Args:
      occassion (str): The type of occassion for the party. Allowed values are:
                       - "casual": Menu for casual party
                       - "formal": Menu for formal party
                       - "superhero": Menu for superhero party
                       - "custom": Custom menu.
    """
    menu = {
        "casual": "Pizza, snacks, and drinks",
        "formal": "3 course dinner with wine and dessert",
        "superhero": "Buffet with high-energy and healthy food.",
    }
    default_menu = "Custom menu for the butler."

    return menu.get(occassion, default_menu)


# agent = CodeAgent(tools=[DuckDuckGoSearchTool()], model=InferenceClientModel())
# agent = CodeAgent(tools=[DuckDuckGoSearchTool()], model=InferenceClientModel())
# agent.run(
#     "Search for the best music recommendations for a party at the Wayne's mansion."
# )


# agent = CodeAgent(tools=[suggest_menu], model=InferenceClientModel())

agent = CodeAgent(tools=[suggest_menu, DuckDuckGoSearchTool()], model=model)
agent.run("Prepare a formal menu for the party and also suggest menu items")
