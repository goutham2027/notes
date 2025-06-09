import base64
import os

from opentelemetry.sdk.trace import TracerProvider

from openinference.instrumentation.smolagents import SmolagentsInstrumentor
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace.export import SimpleSpanProcessor

from huggingface_hub import InferenceClient
from smolagents import (
    CodeAgent,
    tool,
    DuckDuckGoSearchTool,
    InferenceClientModel,
    LiteLLMModel,
)

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

LANGFUSE_PUBLIC_KEY = os.environ["LANGFUSE_PUBLIC_KEY"]
LANGFUSE_SECRET_KEY = os.environ["LANGFUSE_SECRET_KEY"]
LANGFUSE_AUTH = base64.b64encode(
    f"{LANGFUSE_PUBLIC_KEY}:{LANGFUSE_SECRET_KEY}".encode()
).decode()

os.environ["OTEL_EXPORTER_OTLP_ENDPOINT"] = (
    "https://cloud.langfuse.com/api/public/otel"  # EU data region
)
os.environ["OTEL_EXPORTER_OTLP_HEADERS"] = f"Authorization=Basic {LANGFUSE_AUTH}"

trace_provider = TracerProvider()
trace_provider.add_span_processor(SimpleSpanProcessor(OTLPSpanExporter()))

SmolagentsInstrumentor().instrument(tracer_provider=trace_provider)

# agent = CodeAgent(tools=[suggest_menu], model=InferenceClientModel())
model = LiteLLMModel(
    model_id="anthropic/claude-3-5-sonnet-20240620",
    temperature=0.2,
    api_key=os.environ["ANTHROPIC_API_KEY"],
)
agent = CodeAgent(tools=[suggest_menu, DuckDuckGoSearchTool()], model=model)
agent.run("Prepare a formal menu for the party and also suggest menu items")
