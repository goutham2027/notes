pip install jupyterlab

jupyter lab

pip install openai

nano llm - https://github.com/harishhirthi/Nano-LLM

fast ai

### Questions

**On what basis we select a GPU accelerator when we want to deploy an LLM model?**

- Understand computational requirements of the model
- Workload type: training vs inference

**How important the host VM when we are selecting a GPU accelerator?**

- CPUs handle data preprocessing and feeding data to the GPU.
  A slow CPU can become a bottleneck, unable to keep up with
  the GPU's processing speed, leading to underutilization of the GPU.
- The CPU manages the overall orchestration of tasks, including setting up
  kernels for execution on the GPU, handling I/O operations, and managing memory.
  Insufficient CPU resources can lead to increased latency and lower throughput.
  **How to deploy an LLM in a non-vertex AI way and expose an api?**
  **How important the communication between CPU and GPU in inference and training?**
  **How to quickly deploy an LLM in vertex AI?**
