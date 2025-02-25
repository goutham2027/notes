# Hugging face course

https://huggingface.co/learn/nlp-course/en/chapter1/1

### NLP - Natural Language Processing

- A field of lingustics and ML focused on understanding everything related to human language
- Aim is to understand the context of the words

Use cases

- Classifying whole sentences: review sentiment, email spam detection
- Classifying each word in a sentence: identifying the grammatical compnents of a sentence
- Generating text content: completing a prompt with auto generated text
- Extracting an answer from a text
- Generating a new sentence from an input text: translation

NLP not limited to text, can handle speech recognition and computer vision

### Transformers

- Used to solve all kinds of NLP tasks
- transformers library provides the functionality to create and use the shared models
- model hub contains thousands of pretrained models that anyone can download and use

**pipeline function**

- the pipeline function returns an end-to-end object that performs an NLP task on one or several tasks

pre processing -> model -> post processing

```python
from transformers import pipeline

# by default, this pipeline selects a pretrained model that has been fine-tuned for sentiment
# analysis in English

# model is downloaded and cached when classifier object is created
classifier = pipeline("sentiment-analysis")
classifier("I've been waiting for a hugging face course my whole life")

# output
# [{'label': 'POSITIVE', 'score': 0.9598047137260437}]
```

3 main steps when we pass some text to a pipeline

1. The text is preprocessed into a format the model can understand
2. The preprocessed inputs are passed to the model
3. The predictions of the model are post-processed, so that we can make sense of them

There are many pipelines:

- feature extraction (get the vector representation of a text)
- question answering
- sentiment analysis
- summarization
- text generation
- translation
- zero shot classification

**zero shot classification**

- eg: classify texts that haven't been labelled
- allows us to specify which labels to use for the classification, so we don't have to rely on the labels of the pretrained model

- This pipeline is called zero-shot because you don't need to fine tune the model on your data to use it. It can directly return probability scores for any list of labels we want.

```python
from transformed import pipeline

classifier = pipeline("zero-shot-classification")
classifier("This is a course about the Transformers library", candidate_labels=["education", "politics", "business"])

# Output
# {'sequence': 'This is a course about the Transformers library',
#  'labels': ['education', 'business', 'politics'],
#  'scores': [0.8445963859558105, 0.111976258456707, 0.043427448719739914]}
```

```python
# to change the models
from transformers import pipeline

generator = pipeline("text-generation", model="distilgpt2")
```

### How do Transformers work?

- The Transformer architecture was introduced in June 2017
- [Paper: Attention is all you need](https://arxiv.org/pdf/1706.03762)
- Evolution and links to papers can be found from here - https://huggingface.co/learn/nlp-course/en/chapter1/4?fw=pt

**Transformers are language models**

- All the Transformer models (GPT, BERT, BART, T5) have been trained as language models
- They have been trained on large amounts of raw text in a self-supervised fashion
- Self supervised learning is a type of training in which the objective is automatically computed from the inputs of the model. Humans are not needed to label the data
- Though the model develops a statistical understanding of the language, but not useful for specific tasks.
- General pretrained model then goes through a process called Transfer Learning. The model is fine-tuned in a supervised way (human annotated labels)
- Eg: predicting the next word in a sentence reading the `n` previous words. This is called _causal language modeling_ because the output depends on the past and present inputs, but not the future ones.
- Masked models can predict the masked/missing word in the sentence.

- Pretraining is computational heavy and everytime someone pretrains from scratch would add unnecessary global costs.
- That's why sharing of language models is important. Sharing the trained weights and building on top of already trained weights reduces the overall compute cost and carbon footprint.

- https://mlco2.github.io/impact/
- https://codecarbon.io/

**Transfer Learning**
