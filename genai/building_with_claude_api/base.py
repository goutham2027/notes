from dotenv import load_dotenv

load_dotenv()


from anthropic import Anthropic

client = Anthropic()
model = "claude-haiku-4-5"


messages = []


def add_user_message(messages, text):
    msg = {"role": "user", "content": text}
    messages.append(msg)


def add_assistant_message(messages, text):
    msg = {"role": "assistant", "content": text}
    messages.append(msg)


def chat(messages, system=None, temperature=1.0, stop_sequences=[]):
    params = {
        "model": model,
        "max_tokens": 1000,
        "messages": messages,
        "temperature": temperature,
        "stop_sequences": stop_sequences,
    }
    if system:
        params["system"] = system
    message = client.messages.create(**params)
    return message.content[0].text
