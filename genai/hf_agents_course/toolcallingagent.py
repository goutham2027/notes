from smolagents import ToolCallingAgent, DuckDuckGoSearchTool

from model import model
import trace_setup

agent = ToolCallingAgent(tools=[DuckDuckGoSearchTool()], model=model)
agent.run("Search for the best programming books for experienced programmers")
