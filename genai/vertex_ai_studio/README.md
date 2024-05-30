Source: https://www.cloudskillsboost.google/course_templates/552/

### Vertex AI studio course

- tool to access Google's generative AI models
- it facilitates
  - testing
  - tuning
  - augmenting
  - deploying

### course overview

- Generative AI workflow
- Gemini multimodal
- Prompt design
- Model tuning

### Intro

**How does AI generate new content?**

- learns from a massive amount of existing content inclues: text, image and video. This process is called training
- results in the creation of a foundational model
- a foundation model is usually a large model interms of number of parameters, size of training data and high requirements
  computational power
- an llm like PaLM - Pathways Language Model is an example for foundational model
- models from Google:
  - PaLM for language generation
  - Gemini for multimodal processing
  - Codey for code generation
  - Imagen for image processing

How can we use the foundation models to power our applications?
and
How can we further train or tune the foundation model to solve a problem in a specific field?

### Vertex AI

- comprehensive machine learning platform by gcp
- supports end-to-end ML processes:
  - model creation
  - deployment
  - management
- vertex ai provides two primary capabilities:
  1. Predictive AI
  - we can build ML models for forecasting
  2. Generative AI
  - we can use and tune gen AI models to produce content

**Genreative AI workflow on Vertex AI**
![Vertex AI](./generative-ai-workflow-on%20-vertex-ai.png)

### Gemini Multimodal model

**What is a multimodal model?**

- Large foundational model that is capable of processing information from multiple modalities including
  text, image and video.
- The generated content can also be in multiple modalities

**What is a prompt?**

- A prompt is a natural language request to a model to receive a response

**Anatomy of a prompt**

- input (required): represents a request to get a response from the model
  eg: question input, task input, partial input, completion input
- context (optional):
  - instructions that specify how the model should behave
  - information that the model references
    eg: you are the IT help desk. Please consistently advise users to restart
    their computers, regardless of the nature of their inquiries.
- examples (optional): input-output pairs to illustrate the model of an ideal response
  - incorporating examples in the prompt is an effective technique for tailoring the response format

### Prompt design

- The process of designing the input text to get the desired response back from the model

The Temperature setting controls the degree of randomness in the response, with 0 being the most expected answer
and 1 being the most creative

The safety setting allows to adjust the likelihood of receiving a response that could contain harmful content.
Content is blocked based on the probability that it's harmful.

There are 3 methods we can use to shape the model's response in a way that you desire

1. Zero-shot prompting

- Writing a single command so that the model can adopt a certain behavior is called zero-shot prompting
- Prompting without any example of the task to the LLM

2. One-shot prompting

- Where the LLM is given a single example of the task
  eg: if we want the LLM to write a poem, you might provide a single example poem

3. Few-shot prompting

- Where the model is given a small number of examples
  eg: if we want the model to write a news article we might give it a few news articles to read

Structured prompt

**Vertex AI Studio: Language**

- Design a prompt
- Start a conversation
- Tune a model
