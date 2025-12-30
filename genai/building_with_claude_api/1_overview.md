https://anthropic.skilljar.com/claude-with-the-anthropic-api

Contents
- Anthropic Overview
- Accessing Claude via the API
- Prompt evaluations and engineering
- Tool use
- Retrieval Augmented Generation
- Model Context Protocol
- Claude code and computer use
- Workflows and Agents


### Claude Models
- Claude Opus: Highest level of intelligence. Supports reasoning
- Claude Sonnet: Intelligent model that balances quality, speed, cost. Supports reasoning.
- Claude Haiku: Most cost-efficient and latency optimized model. Supports reasoning. Fastest model.

Text generation process includes
- Tokenization
- Embedding: Number representation of tokens
- Contextualization: to refine each embedding down to a single precise definition, Contextualization process is used. In Contextualization each embedding is adjusted based upon other embeddings around it.
- Generation


- When Claude APIs are called, claude doesn't store any of the conversation history. Each request
is completely independent, with no memory of previous exchanges.
- For multi-turn conversations we need to handle the conversation state

### System Prompts
- System Prompts are a powerful way to customize how Claude responds to user input.
- We can shape Claude's tone, style and approach to match your specific use case.
- eg: giving the Claude a math tutor behavior. A tutor doesn't give out the answer, initially
gives hints, helps student to walk through step by step.

```python
messages = []

def add_user_message(messages, text):
    msg = {"role": "user", "content": text}
    messages.append(msg)

def add_assistant_message(messages, text):
    msg = {"role": "assistant", "content": text}
    messages.append(msg)

def chat(messages, system=None):
    params = {
        model=model,
        max_tokens=1000,
        messages=messages
    }
    if system:
        params["system"] = system
    message = client.messages.create(**params)
    return message.content[0].text
```

### Temperature
- Temperature is a decimal value between 0 and 1, is going to influence the distribution of
probabilities.
- Low Temperature: more deterministic output, less creative (0.0 - 0.3)
- Medium temperature: summarization, problem solving (0.4 - 0.7)
- High Temperature: more random output, more creative. (0.8 - 1.0)

### Response streaming
When streaming is enabled, Claude sends back several types of events:
1. Message Start
2. ContentBlockStart: start of a new block of containing text, tool use or other content
3. ContentBlockDelta: Chunks of the actual generated text
4. ContentBlockStop: The current content block has been completed
5. MessageDelta: The current message is complete
6. MessageStop: End of information aobut the current message

### Controlling model output
1. Prefilled assistant Messages
  - Provide a last message of assistant. Claude will use it as the starting point for its response.
2. Stop sequences
  - Forces Claude to stop generating text when it creates a specific series of characters

### Structured Data
```
messages = []

add_user_message(messages, "Generate a very short event bridge rule as json")
add_assistant_message(messages, "```json""")
chat(messages, stop_sequences=["```"])
```
