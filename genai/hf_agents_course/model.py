import os

from smolagents import LiteLLMModel

haiku_model3 = "claude-3-haiku-20240307"
sonnet_model3_5 = "claude-3-5-sonnet-20240620"

use_model = haiku_model3
# use_model = sonnet_model3_5

model = LiteLLMModel(
    model_id=f"anthropic/{use_model}",
    temperature=0.2,
    api_key=os.environ["ANTHROPIC_API_KEY"],
)
