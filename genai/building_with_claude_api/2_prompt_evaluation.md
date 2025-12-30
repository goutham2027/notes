# Prompt Evaluation

Prompt Engineering: Best practices and guidance to improve prompts.
    - Multishot prompting
    - Structuring with XML tags

Prompt Evaluation: Automated testing to measure how well your prompts work
    - Test against expected answers
    - Compare different versions of the same prompt.
    - Review outputs for errors

### Drafting a prompt
Option-1:
- Test the prompt once and decide it's good enough
- Might break in production

Option-2:
- Test the prompt a few times and tweak it to handle a corner case or two
- Unexpected outputs

Option-3:
- Run the prompt through an evaluation pipeline to score it, then iterate on the prompt.
- more work but more confidence

### Eval flow
- Draft a prompt
- Create an Eval dataset
- Feed through Claude
- Feed through a grader
- Change prompt and repeat

```python
def run_prompt(test_case):
    """Merges the prompt and test case input, then returns the result"""
    prompt = f"""
    Please solve the following task:
    {test_case["task"]}
    """
    messages = []
    add_user_message(messages, prompt)
    output = chat(messages)
    return output

def run_test_case(test_case):
    """Calls run_prompt, then grades the result"""
    output = run_prompt(test_case)
    # Grading
    score = 10
    return {
        "output": output,
        "test_case": test_case,
        "score": score
    }


def run_eval(dataset):
    """Loads the dataset and calls run_test_case with each case"""
    results = []
    for test_case in tasks:
        result = run_test_case(test_case)
        results.append(result)
    return results

tasks = [
    {'task': 'Write a Python function that takes an AWS S3 bucket name and returns True if it follows AWS naming conventions (lowercase, 3-63 characters, no underscores), False otherwise.'},
    {'task': "Create a JSON object that represents an AWS IAM policy allowing read-only access to a specific S3 bucket named 'my-data-bucket'."},
    {'task': "Write a regex pattern that matches valid AWS EC2 instance IDs, which start with 'i-' followed by exactly 17 hexadecimal characters."}
]
results = run_eval(tasks)
print(results)
```

### Model based Grading
- Grading systems provide objective signals about output quality.
- A grader takes model output and returns some kind of measurable feedback - typically a number
between 1 (poor quality) and 10 (high quality).

Types of graders
1. Code: Programmatically evaluate the result.
  - Checking output length
  - Verifying output does/doesn't have certain words
  - Syntax validation
  - Readability scores
2. Model: Ask a model to assign a score to the output, or compare two versions
3. Human: Ask a human to assign a score to the output, or compare two versions

Evaluation criteria
- How will we know if our prompt is producing good outputs?

eg: Format, valid syntax, task following (resposne should directly and clearly address the user's task)
Format and Valid syntax need Code Grader
Task following need Model Grader
