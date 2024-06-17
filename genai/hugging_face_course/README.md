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
