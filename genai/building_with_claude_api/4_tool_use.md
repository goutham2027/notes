### Tool Use

Tool functions
- Plain python function that will be executed when Claude decided it needs some additional information to help the user.
Best practices
  - Use well named descriptive arguments
  - Validate the inputs, raising an error if they fail validation
  - Return meaningful errors - Claude will try to call to use your function a second time

Tool Schemas
- JSON schema spec to describe the function
Best Practices
  - Explain what the tool does, when to use it and what it returns
  - Provide super detailed descriptions
```json
{
  "name": "get_weather",
  "description": "Retrieves current weather for the given location",
  "input_schema": {
    "type": "object",
    "properties": {
      "location": {
        "type": "string",
        "description": "The location for which weather is required""
      }
    },
    required: [
      "location"
    ]
    }
  }
}
```

Write a valid json schema spec for the purposes of tool calling for this function. Follow the best practices listed in the attached documentation
https://platform.claude.com/docs/en/agents-and-tools/tool-use/overview

User message - TextBlock
Assistant message - ToolUseBlock
User Message - ToolResult Block

ToolResultBlock contains
  - tool_use_id: must match the `id` of the ToolUse block that this ToolResult corresponds to
  - content: output from running the tool, serialized as a string
  - is_error: True if an error occurred.
  - type: tool_result


### Multi-turn conversations with tools

stop reason: tells us why Claude stopped generating text

tool_use - claude has decided that it needs to call a tool
end_turn - claude has finished generating it's assistant message
max_tokens - claude has hit the token output limit and can't generate any more output
stop_sequence - claude has encountered one of your provided stop sequences


### Batch tool use
- multiple tool call requests in a single claude message instead of multiple rounds.
- Batch tool: additional tool that accepts a list of calls to other tools
- Tricking claude to call multiple tools in parallel

### Prompt based structured output
- Tool for structured data
- Controlling tool use
  - we can force claude to call specific tools
  ```python
  tool_choice = {"type": "auto"} # model decides if it needs to use a tool (default)
  
  tool_choice = {"type": "any"} # model must use a tool - it can decide which to use
  
  tool_choice = {"type": "tool", "name": "TOOL_NAME"} # model must use the listed tool
  ```
  
  Get the JSON schema using claude with this prompt
  ```
  write a tool schema for a function named article_summary. It should be called with a title (string), author (string) and list of key insights (list of strings).
  
  documentation
  https://platform.claude.com/docs/en/agents-and-tools/tool-use/implement-tool-use
  ```
  
### Finegrained tool calling
- Disables JSON validation step.
- Finegrained tool calling sends groups of chunks without wiating for a full top level key to be created
- `fine_grained=True`

### The text edit tool
- Tool built directly into claude
- Allows claude to create, read and edit both directories and files
- Only the JSON schema is build into Claude - you have to provide an implementation of the actual tool

### The Web search tool
- Search the web
- Implementation exists
CitationWebSearchResultLocation
