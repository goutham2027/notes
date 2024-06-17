# main source: https://huggingface.co/EleutherAI/pythia-70m

from transformers import GPTNeoXForCausalLM, AutoTokenizer

model = GPTNeoXForCausalLM.from_pretrained(
    "EleutherAI/pythia-70m-deduped",
    revision="step3000",
    cache_dir="./pythia-70m-deduped/step3000",
)

tokenizer = AutoTokenizer.from_pretrained(
    "EleutherAI/pythia-70m-deduped",
    revision="step3000",
    cache_dir="./pythia-70m-deduped/step3000",
)

inputs = tokenizer("Today is", return_tensors="pt")
tokens = model.generate(**inputs, max_length=50, num_return_sequences=5, num_beams=5)
for token in tokens:
    print(tokenizer.decode(token))

from llama import BasicModelRunner
model = BasicModelRunner("EleutherAI/pythia-70m-deduped")
model.
