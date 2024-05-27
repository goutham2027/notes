https://learn.deeplearning.ai/courses/llmops/lesson/1/introduction

Extension to ML Ops

Lifecycle

- data preparation
- model tuning
- deployment
- maintenance
- monitoring

Prompt management

Evaluating and comparing the performance of different prompts

### Fundamentals

ML Ops is an ML eng culture and practice that aims at unifying
ML development(dev) and ML operation (ops)

Automation and monitoring at all steps of ML system construction including:

- Integration
- Testing
- Releasing
- Deployment
- infra management

**ML Ops framework**

- Data ingestion
- Data validation
- Data trasformation
- Model
- Model Analysis
- Serving
- Logging

We focuss on

- Data transformation
- Model
- Serving

Per usecase, we build a workflow
One model for one use case in ML Ops

**LLM Ops**
ML Ops for LLMs (LLMOps)

- Focus on the llm development and managing the model in production

Examples

- experiment on foundation models
- prompt design and managment
- supervised tuning
- monitoring
- evaluate generative output

LLM system design

- Broader design of the entire end-to-end application - FE, BE, Data engineering etc

Examples

- Chain multiple LLM together
- Grounding: additional information
- Track history

**Example of a LLM driven application**
![Example](./example-llm-driven-app.png)

**LLMOps pipeline**

- Data preparation and versioning
- Pipeline design (Supervised tuning)
- Artifact (configuration and workflow)
- Pipeline exection
- Deploy LLM
- Prompting and predictions
- Responsible AI

![LLMOps Simplified Pipeline](./llmops-pipeline.png)

Check

- Prompt design and prompt management
- Model evaluation
- Model monitoring
- Testing
