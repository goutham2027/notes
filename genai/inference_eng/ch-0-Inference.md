https://inferenceengineering.tech/

Training is the process of learning model weights from data. Inference is serving generative AI models in production.

Inference for generative AI models is complex and not straigtforward as class ML models.

### Chapter-0 Inference

3 layers of inference
1. Runtime: optimizing the performance of a single model on a single GPU backed instance
2. Infrastructure: scaling across clusters, regions and clouds without creating silos while maintaining uptime
3. Tooling: providing engineers working on inference with the right level of abstraction to balance control with productivity.


#### Runtime - Model Performance
- Responsible to ensure model is running as performantly and efficiently as possible
- Software stack
    - CUDA - Low level GPU programming framework
    - PyTorch - Deep learning framework for model execution
    - Inference engines (vLLM, SGLang, TensorRT-LLM) for optimized serving
- Optimization
    - Batching: run incoming requests in parallel, to increase throughput
    - Caching: reuse the kv cache, the cached results of the attention algorithm
    - Quantization: lower the precision of select pieces of the model to access more compute and reduce memory burden
    - Speculation: generate and validate draft tokens to produce more than one token per forward pass during decode
    - Parallelism: leverage more than one GPU to accelerate large models
    - Disaggregation: separate the two phases of LLM inference, prefill and decode onto independently scaling workers
- Kernels
    - FlashAttention - Optimized attention kernels
    - Optimized CUDA kernels for maximum GPU utilization


#### Infrastructure - Scale & Reliability
No matter how performant a single instance of a model server is, it will outgrow the single server. This is a systems problem (not CUDA or PyTorch) that needs to be solved at
the infrastructure layer.

- Autoscaling
- Multi-cloud
- Reliability

Global unification: treating all available resources as a single unified pool of compute, eliminating silos


#### Tooling - Developer Experience
- Abstraction level
    - Balance control with productivity - from blackbox API to raw compute primitives
- Model deployment
    - Give a platform model weights, get back a serving endpoint with monitoring
- Observability
    - Metrics, logging, tracing, and alerting for inference workloads
    