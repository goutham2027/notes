from smolagents import ToolCallingAgent, DuckDuckGoSearchTool

from model import model

import trace_setup

agent = ToolCallingAgent(tools=[DuckDuckGoSearchTool()], model=model)
agent.visualize()

# agent.run(
#     "Before even coming up with list of books think usually what experienced programmers lack and need upskilling in and then suggest books"
# )
